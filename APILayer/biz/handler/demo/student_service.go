// Code generated by hertz generator.

package demo

import (
	"context"
	"fmt"
	demo "github.com/Raccoon-njuse/httpsvr/biz/model/demo"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
)

// Register .
// @router /add-student-info [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demo.Student
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	cli := initGenericClient()
	httpReq, err := adaptor.GetCompatRequest(c.GetRequest())
	if err != nil {
		panic("get http request failed")
	}
	customReq, err := generic.FromHTTPRequest(httpReq)
	if err != nil {
		panic("get customReq failed")
	}
	resp, err := cli.GenericCall(ctx, "Register", customReq)
	realResp := resp.(*generic.HTTPResponse)
	fmt.Println(realResp)
	c.JSON(consts.StatusOK, realResp.Body)
}

// Query .
// @router /query [GET]
func Query(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demo.QueryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	cli := initGenericClient()
	httpReq, err := adaptor.GetCompatRequest(c.GetRequest())
	if err != nil {
		panic("get http request failed")
	}
	customReq, err := generic.FromHTTPRequest(httpReq)
	if err != nil {
		panic("get custom req failed")
	}
	resp, err := cli.GenericCall(ctx, "Query", customReq)
	//fmt.Println(resp)
	realResp := resp.(*generic.HTTPResponse)
	//fmt.Println(realResp)
	//fmt.Println(realResp.Body)
	c.JSON(consts.StatusOK, realResp.Body)
}

func initGenericClient() genericclient.Client {
	// 本地文件 idl 解析
	// YOUR_IDL_PATH thrift 文件路径: 举例 ./idl/example.thrift
	// includeDirs: 指定 include 路径，默认用当前文件的相对路径寻找 include
	p, err := generic.NewThriftFileProvider("../StudentService/student.thrift")
	if err != nil {
		panic(err)
	}
	// 构造 http 类型的泛化调用
	g, err := generic.HTTPThriftGeneric(p)
	if err != nil {
		panic(err)
	}
	cli, err := genericclient.NewClient("destServiceName", g, client.WithHostPorts("127.0.0.1:8889"))
	if err != nil {
		panic(err)
	}
	return cli
}