package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	storageuploaderservice "storage-uploader-service/proto/storage-uploader-service"
)

type StorageUploaderService struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *StorageUploaderService) Call(ctx context.Context, req *storageuploaderservice.Request, rsp *storageuploaderservice.Response) error {
	log.Info("Received StorageUploaderService.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *StorageUploaderService) Stream(ctx context.Context, req *storageuploaderservice.StreamingRequest, stream storageuploaderservice.StorageUploaderService_StreamStream) error {
	log.Infof("Received StorageUploaderService.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&storageuploaderservice.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *StorageUploaderService) PingPong(ctx context.Context, stream storageuploaderservice.StorageUploaderService_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&storageuploaderservice.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
