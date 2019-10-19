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
	server.DefaultServer.RegisterRoute("get", "/validator/:operaddress", defaultValidator.ValidatorInfo)
	server.DefaultServer.RegisterRoute("get", "/validators/:index/:pageSize", defaultValidator.QueryValidatorPage)
}

var (
	defaultValidator = &Validator{}
)

type Validator struct {
}

func (val *Validator) ValidatorInfo(ctx *gin.Context) {
	operatorAddress := ctx.Param("operaddress")
	validator, err := model.NewValidatorModel().Info(operatorAddress)
	if err != nil {
		ctx.JSON(200, response.InternalServerError(err.Error()))
	}
	if validator == nil {
		ctx.JSON(200, response.NotFound(nil))
	}
	ctx.JSON(200, validator)
}

func (val *Validator) QueryValidatorPage(ctx *gin.Context) {
	page := params.ExtractPage(ctx)
	validatorModel := model.NewValidatorModel()
	query := model.ValidatorQuery{PageIndex: page.PageIndex, PageSize: page.PageSize}
	totalItems, err := validatorModel.Count(query)
	if nil != err {
		ctx.JSON(200, response.InternalServerError(err.Error()))
		return
	}

	items, err := validatorModel.List(query)
	if nil != err {
		ctx.JSON(200, response.InternalServerError(err.Error()))
		return
	}

	pageInfo := util.NewPageInfo(page, totalItems, items)
	ctx.JSON(200, response.Ok(pageInfo))
}
