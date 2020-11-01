package eventHandler

import (
	"context"
	"encoding/json"
	"log"

	fileServiceProto "github.com/Condition17/fleet-services/file-service/proto/file-service"
	"github.com/Condition17/fleet-services/lib"
	baseservice "github.com/Condition17/fleet-services/lib/base-service"
	"github.com/Condition17/fleet-services/run-controller-service/errors"
	"github.com/Condition17/fleet-services/run-controller-service/events"
	proto "github.com/Condition17/fleet-services/run-controller-service/proto/run-controller-service"
	testRunServiceProto "github.com/Condition17/fleet-services/test-run-service/proto/test-run-service"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
)

type handlerFunc = func(broker.Event) error

type eventHandler struct {
	baseservice.BaseHandler
	FileService    fileServiceProto.FileService
	TestRunService testRunServiceProto.TestRunService
}

func (h eventHandler) buildCallContext(event *proto.Event) context.Context {
	return metadata.Set(context.Background(), "Token", string(event.Meta.Token))
}

func (h eventHandler) sendErrorToWssQueue(err error) {
	h.SendEventToWssQueue(context.Background(), events.WSS_ERROR, []byte(err.Error()))
}

func (h eventHandler) HandleTestRunCreated(event *proto.Event) {
	// unmarshal event speciffic data
	var eventData *proto.TestRunCreatedEventData
	if err := json.Unmarshal(event.Data, &eventData); err != nil {
		log.Println(errors.EventUnmarshalError(event))
		return
	}

	// create a new file for the received test run
	var fileSpec fileServiceProto.File = fileServiceProto.File{
		Name:         eventData.FileSpec.Name,
		Size:         eventData.FileSpec.Size,
		MaxChunkSize: eventData.FileSpec.MaxChunkSize,
	}
	createFileResp, err := h.FileService.CreateFile(h.buildCallContext(event), &fileSpec)
	if err != nil {
		h.sendErrorToWssQueue(errors.FileCreationError(eventData.TestRunSpec))
		return
	}

	// assign the created file to the current test run
	var assignmentDetails testRunServiceProto.AssignRequest = testRunServiceProto.AssignRequest{
		TestRunId: eventData.TestRunSpec.Id,
		FileId:    createFileResp.File.Id,
	}
	if _, err := h.TestRunService.AssignFile(h.buildCallContext(event), &assignmentDetails); err != nil {
		h.sendErrorToWssQueue(errors.FileCreationError(eventData.TestRunSpec))
		return
	}

	// send informations to the client through WS service
	wssEventData, _ := json.Marshal(&proto.FileEntityCreatedEventData{
		TestRunId: assignmentDetails.TestRunId,
		FileId:    assignmentDetails.FileId,
	})
	h.SendEventToWssQueue(context.Background(), events.WSS_FILE_ENTITY_CREATED, wssEventData)
}

func handleEvent(handler eventHandler, event *proto.Event) {
	switch event.Type {
	case events.TEST_RUN_CREATED:
		handler.HandleTestRunCreated(event)
	default:
		log.Printf("The event with type '%s' is not a recognized fleet test run pipeline event", event.Type)
	}
}

func newEventHandler(service micro.Service) eventHandler {
	return eventHandler{
		BaseHandler:    baseservice.NewBaseHandler(service),
		FileService:    fileServiceProto.NewFileService(lib.GetFullExternalServiceName("file-service"), client.DefaultClient),
		TestRunService: testRunServiceProto.NewTestRunService(lib.GetFullExternalServiceName("test-run-service"), client.DefaultClient),
	}
}

func NewHandler(service micro.Service) handlerFunc {
	var handler eventHandler = newEventHandler(service)

	return func(e broker.Event) error {
		var event *proto.Event

		if err := json.Unmarshal(e.Message().Body, &event); err != nil {
			return err
		}
		handleEvent(handler, event)

		return nil
	}
}
