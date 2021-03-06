package model

import (
	proto "github.com/Condition17/fleet-services/resource-manager-service/proto/resource-manager-service"
	testRunModels "github.com/Condition17/fleet-services/test-run-service/model"
	"gorm.io/gorm"
)

type FileSystem struct {
	gorm.Model

	IP                  string `gorm:"not null;type:varchar(100);default:null"`
	Name                string `gorm:"not null;type:varchar(100);default:null"`
	FileShareCapacityGb int64  `gorm:"not nulll"`
	FileShareName       string `gorm:"not null;type:varchar(100);default:null"`
	TestRunID           uint32 `gorm:"unique;constaint:OnDelete:CASCADE;"`
	TestRun             testRunModels.TestRun
}

func UnmarshalFileSystem(fileSystem *FileSystem) *proto.FileSystem {
	// TODO: find a better way to do this
	testRunData := testRunModels.UnmarshalTestRun(&fileSystem.TestRun)
	return &proto.FileSystem{
		Id:                  uint32(fileSystem.ID),
		IP: 				 fileSystem.IP,
		Name:                fileSystem.Name,
		FileShareCapacityGb: fileSystem.FileShareCapacityGb,
		FileShareName:       fileSystem.FileShareName,
		TestRunId:           fileSystem.TestRunID,
		TestRun: &proto.TestRun{
			Id:     testRunData.Id,
			Name:   testRunData.Name,
			FileId: testRunData.FileId,
			UserId: testRunData.UserId,
			User: &proto.User{
				Id:      testRunData.User.Id,
				Name:    testRunData.User.Name,
				Company: testRunData.User.Company,
				Email:   testRunData.User.Email,
			},
		},
	}
}
