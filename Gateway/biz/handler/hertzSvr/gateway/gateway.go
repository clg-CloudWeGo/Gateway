package gateway

import (
	"context"
	"fmt"
	"github.com/clg-CloudWeGo/gateway/biz/handler/hertzSvr"
	"github.com/clg-CloudWeGo/gateway/biz/handler/hertzSvr/idlManager"
	"github.com/clg-CloudWeGo/gateway/biz/handler/hertzSvr/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Gateway(ctx context.Context, c *app.RequestContext) {
	serviceName := c.Param("service")
	methodName := c.Param("method")

	// 获取服务对应的clientInfo
	var clientInfo hertzSvr.ClientInfo

	// 查找缓存中是否有对应service的client，若无则进行创建并放入缓存
	if _, isOk := hertzSvr.Clients[serviceName]; isOk {
		clientInfo = hertzSvr.Clients[serviceName]
	} else {
		// 调用idl管理平台API，查找对应service的idl
		idlContent := idlManager.GetIDLContent(serviceName)
		provider, err := utils.NewProvider(idlContent)
		if err != nil {
			c.JSON(consts.StatusBadRequest, &hertzSvr.Response{
				Success: false,
				Message: "Error: fail to load idl for service " + serviceName + "." + err.Error(),
			})
			return
		}
		clientInfo.Provider = provider
		clientInfo.Cli, err = utils.NewClient(serviceName, provider, hertzSvr.Resolver)
		if err != nil {
			c.JSON(consts.StatusBadRequest, &hertzSvr.Response{
				Success: false,
				Message: "Error: fail to make new client for service " + serviceName + "." + err.Error(),
			})
		}

		hertzSvr.Clients[serviceName] = clientInfo
	}

	// 进行HTTP泛化调用
	resp, err := utils.GetHTTPGenericResponse(ctx, c, methodName, clientInfo.Cli)
	fmt.Println("this is generic call resp.Body")
	fmt.Println(resp.Body)
	if err != nil {
		c.JSON(consts.StatusBadGateway, hertzSvr.Response{
			Success: false,
			Message: "Error: fail to generic call " + serviceName + "." + err.Error(),
		})
		return
	}

	c.JSON(consts.StatusOK, resp.Body)
}
