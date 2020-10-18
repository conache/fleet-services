package model

import (
	proto "github.com/Condition17/fleet-services/test-run-service/proto/test-run-service"
	userModels "github.com/Condition17/fleet-services/user-service/model"
	"gorm.io/gorm"
)

type TestRun struct {
	gorm.Model
	Name   string `gorm:"not null;type:varchar(100);default:null"`
	UserID uint32
	User   userModels.User
}

func MarshalTestRun(testRun *proto.TestRun) *TestRun {
	return &TestRun{
		Model: gorm.Model{ID: uint(testRun.Id)},
		Name:  testRun.Name,
	}
}

func UnmarshalTestRun(testRun *TestRun) *proto.TestRun {
	return &proto.TestRun{
		Id:   uint32(testRun.ID),
		Name: testRun.Name,
	}
}

func UnmarshalTestRunsCollection(testRuns []*TestRun) []*proto.TestRun {
	collection := make([]*proto.TestRun, 0)
	for _, run := range testRuns {
		collection = append(collection, UnmarshalTestRun(run))
	}
	return collection
}
