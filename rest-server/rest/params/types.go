package params

import (
	"github.com/gin-gonic/gin"
	"mdu/explorer/common/util"
	"strconv"
)

type HeightOrHash string

func (h HeightOrHash) IsHeight() bool {
	return false
}

func (h HeightOrHash) IsHash() bool {
	return false
}

func (h HeightOrHash) ToHash() string {
	return ""
}

func (h HeightOrHash) ToHeight() int64 {
	return -1
}

func ExtractPage(ctx *gin.Context) util.Page {
	index, err := strconv.Atoi(ctx.Param("index"))
	if err != nil {

	}
	pageSize, err := strconv.Atoi(ctx.Param("pageSize"))
	return util.Page{int32(index), int32(pageSize)}
}
