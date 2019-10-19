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
	server.DefaultServer.RegisterRoute("get", "/block/:heightOrHash", defaultBlockController.BlockInfo)
	server.DefaultServer.RegisterRoute("get", "/blocks/:index/:pageSize", defaultBlockController.QueryBlockPage)
}

type BlockController struct {
}

var (
	defaultBlockController = &BlockController{}
)

func (val *BlockController) BlockInfo(ctx *gin.Context) {
	heightOrHash := params.HeightOrHash(ctx.Param("heightOrHash"))
	var block *model.Block
	var err error
	if heightOrHash.IsHeight() {
		block, err = model.NewBlocksModel().FindBlockByHeight(heightOrHash.ToHeight())
	} else {
		block, err = model.NewBlocksModel().FindBlockByHash(heightOrHash.ToHash())
	}

	if err != nil {
		ctx.JSON(200, response.InternalServerError(err.Error()))
	}
	if block == nil {
		ctx.JSON(200, response.NotFound(nil))
	}

	ctx.JSON(200, response.Ok(block))
}

func (val *BlockController) QueryBlockPage(ctx *gin.Context) {
	page := params.ExtractPage(ctx)

	blocksModel := model.NewBlocksModel()
	query := model.BlockQuery{PageIndex: page.PageIndex, PageSize: page.PageSize}
	totalItems, err := blocksModel.CountBlocks(query)
	if nil != err {
		ctx.String(500, err.Error())
		return
	}

	items, err := blocksModel.List(query)
	if nil != err {
		ctx.String(500, err.Error())
		return
	}

	pageInfo := util.NewPageInfo(page, totalItems, items)
	ctx.JSON(200, response.Ok(pageInfo))
}
