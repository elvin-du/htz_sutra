package controller

import (
	"github.com/gin-gonic/gin"
	"mdu/explorer/common/model"
	"mdu/explorer/common/util"
	"mdu/explorer/common/util/response"
	"mdu/explorer/rest-server/rest/params"
	"mdu/explorer/rest-server/rest/server"
)

func init() {
	server.DefaultServer.RegisterRoute("get", "/tx/:hash", defaultTxController.TxInfo)
	server.DefaultServer.RegisterRoute("get", "/txs/:index/:pageSize", defaultTxController.QueryTxPage)
}

type TxController struct {
}

var (
	defaultTxController = &TxController{}
)

func (val *TxController) TxInfo(ctx *gin.Context) {
	hash := ctx.Param("hash")
	tx, err := model.NewTxModel().FindTxByHash(hash)

	if err != nil {
		ctx.JSON(200, response.InternalServerError(err.Error()))
	}
	if tx == nil {
		ctx.JSON(200, response.NotFound(nil))
	}
	ctx.JSON(200, response.Ok(tx))
}

func (val *TxController) QueryTxPage(ctx *gin.Context) {
	page := params.ExtractPage(ctx)

	txModel := model.NewTxModel()
	query := model.TxQuery{PageIndex: page.PageIndex, PageSize: page.PageSize}
	totalItems, err := txModel.CountTx(query)
	if nil != err {
		ctx.JSON(200, response.InternalServerError(err.Error()))
		return
	}

	items, err := txModel.List(query)
	if nil != err {
		ctx.JSON(200, response.InternalServerError(err.Error()))
		return
	}

	pageInfo := util.NewPageInfo(page, totalItems, items)
	ctx.JSON(200, response.Ok(pageInfo))
}
