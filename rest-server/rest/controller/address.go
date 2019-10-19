package controller

import (
	"github.com/gin-gonic/gin"
	"mdu/explorer/common/model"
	"mdu/explorer/common/util/response"
	"mdu/explorer/rest-server/rest/server"
)

func init() {
	server.DefaultServer.RegisterRoute("get", "/address/:hash", defaultAddressController.Info)
}

type AddressController struct {
}

var (
	defaultAddressController = &AddressController{}
)

func (val *AddressController) Info(ctx *gin.Context) {
	address := ctx.Param("hash")
	addressInfo, err := model.NewAddressesModel().Find(address)
	if err != nil {
		ctx.JSON(200, response.InternalServerError(err.Error()))
	}
	if addressInfo == nil {
		ctx.JSON(200, response.NotFound(nil))
	}
	ctx.JSON(200, response.Ok(addressInfo))
}
