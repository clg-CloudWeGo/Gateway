// Code generated by hertz generator.

package service

import (
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	service "github.com/clg-CloudWeGo/idlmanagementservice/biz/model/hertzSvr/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type IDLMessage struct {
	ID      int
	SvcName string
	IDL     string
}

var DB, _ = gorm.Open(sqlite.Open("IDLMessage.db"), &gorm.Config{})

// AddIDL .
// @router /idl/add [POST]
func AddIDL(ctx context.Context, c *app.RequestContext) {
	var err error
	var req service.IDLInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(service.IDLResponse)
	var serviceMessage IDLMessage
	DB.First(&serviceMessage, "svc_name = ?", req.Name)
	if serviceMessage.ID != 0 {
		resp.Success = false
		resp.Message = "ERROR:SERVICE ALREADY EXISTED"
		c.JSON(consts.StatusOK, resp)
		return
	}

	DB.Create(&IDLMessage{
		SvcName: req.Name,
		IDL:     req.Idl,
	})
	resp.Success = true
	resp.Message = "SUCCESS ADD"
	c.JSON(consts.StatusOK, resp)
}

// DeleteIDL .
// @router /idl/delete [POST]
func DeleteIDL(ctx context.Context, c *app.RequestContext) {
	var err error
	var req service.IDLMessage
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(service.IDLResponse)
	var serviceMessage IDLMessage
	//检查服务是否存在
	DB.First(&serviceMessage, "svc_name = ?", req.Name)
	if serviceMessage.ID == 0 {
		resp.Success = false
		resp.Message = "ERROR:NO SUCH SERVICE"
		c.JSON(consts.StatusOK, resp)
		return
	}

	DB.Delete(&serviceMessage, serviceMessage.ID)
	resp.Success = true
	resp.Message = "SUCCESS DELETE"
	c.JSON(consts.StatusOK, resp)
}

// UpdateIDL .
// @router /idl/update [POST]
func UpdateIDL(ctx context.Context, c *app.RequestContext) {
	var err error
	var req service.IDLInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(service.IDLResponse)
	var serviceMessage IDLMessage
	//检查服务是否存在
	DB.First(&serviceMessage, "svc_name = ?", req.Name)
	if serviceMessage.ID == 0 {
		resp.Success = false
		resp.Message = "ERROR:NO SUCH SERVICE"
		c.JSON(consts.StatusOK, resp)
		return
	}
	DB.Model(&serviceMessage).Update("id_l", req.Idl)

	resp.Success = true
	resp.Message = "SUCCESS UPDATE"
	c.JSON(consts.StatusOK, resp)
}

// QueryIDL .
// @router /idl/query [GET]
func QueryIDL(ctx context.Context, c *app.RequestContext) {
	var err error
	var req service.IDLQueryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(service.IDLInfo)

	var serviceMessage IDLMessage
	DB.First(&serviceMessage, "svc_name = ?", req.Name)
	if serviceMessage.ID == 0 {
		resp.Idl = "ERROR:NO SUCH SERVICE"
		resp.Name = req.Name
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp = &service.IDLInfo{
		Name: req.Name,
		Idl:  serviceMessage.IDL,
	}
	c.JSON(consts.StatusOK, resp)
}
