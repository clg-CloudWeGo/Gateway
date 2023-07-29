package main

import (
	"context"
	"fmt"
	demo "github.com/Raccoon-njuse/rpcsvr/kitex_gen/demo"
	model2 "github.com/Raccoon-njuse/rpcsvr/model"
	_ "github.com/golang/mock/mockgen/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct {
	db *gorm.DB
}

// InitDB eg: 初始化db，注意服务启动时初始化
func (s *StudentServiceImpl) InitDB() {
	db, err := gorm.Open(sqlite.Open("foo.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// drop table
	db.Migrator().DropTable(model2.Student{})
	// create table
	err = db.Migrator().CreateTable(model2.Student{})
	if err != nil {
		panic(err)
	}
	s.db = db
}

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *demo.Student) (resp *demo.RegisterResp, err error) {
	var stuRes *model2.Student
	tempresult := s.db.Table("students").First(&stuRes, student.Id)
	if tempresult.Error == nil {
		//标明查询成功，无法插入，返回查询到的数据
		tempReq := &demo.QueryReq{Id: student.Id}
		tempStudent, _ := s.Query(ctx, tempReq)
		fmt.Println(tempStudent)
		resp = &demo.RegisterResp{
			Message: "Already exists!" + "\n" +
				"name:" + tempStudent.Name + "\n" +
				"Id:" + string(tempStudent.Id) + "\n" +
				"Email:" + strings.Join(tempStudent.Email, ",") + "\n" +
				"CollegeName:" + tempStudent.College.Name + "\n" +
				"CollegeAddress:" + tempStudent.College.Address,
			Success: true,
		}
		return
	}
	modelStudent := student2Model(student)
	result := s.db.Table("students").Create(modelStudent)
	if result.Error != nil {
		panic("insert data fail")
	}
	resp = &demo.RegisterResp{
		Message: "Register success!" + "\n" +
			"name:" + modelStudent.Name + "\n" +
			"Id:" + string(modelStudent.Id) + "\n" +
			"CollegeName:" + modelStudent.CollegeName + "\n" +
			"CollegeAddress:" + modelStudent.CollegeAddress + "\n" +
			"Email:" + modelStudent.Email,
		Success: true,
	}
	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *demo.QueryReq) (resp *demo.Student, err error) {
	// TODO: Your code here...

	var stuRes *model2.Student
	result := s.db.Table("students").First(&stuRes, req.Id)
	if result.Error != nil {
		panic("query data failed")
	}
	resp = model2Student(stuRes)
	return
}

func student2Model(student *demo.Student) *model2.Student {
	return &model2.Student{
		Id:             student.Id,
		Name:           student.Name,
		Email:          strings.Join(student.Email, ","),
		CollegeName:    student.College.Name,
		CollegeAddress: student.College.Address,
	}
}

func model2Student(student *model2.Student) *demo.Student {
	return &demo.Student{
		Id:      student.Id,
		Name:    student.Name,
		Email:   strings.Split(student.Email, ","),
		College: &demo.College{Name: student.CollegeName, Address: student.CollegeAddress},
	}
}
