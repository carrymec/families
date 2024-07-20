package person

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/carrymec/families/common"
	"go.uber.org/zap"
	"net/http"
)

type Controller struct {
	lg      *zap.Logger
	service *Service
}

func NewPersonController(lg *zap.Logger, service *Service) *Controller {
	return &Controller{
		lg:      lg,
		service: service,
	}
}

func (p *Controller) Create(ctx *gin.Context) {
	var per Person
	if err := ctx.ShouldBind(&per); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := p.service.CreatePerson(ctx, per)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Resp{
			Code: -1,
			Msg:  fmt.Sprintf("创建用户失败: %s", err.Error()),
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Resp{
		Code: 0,
		Msg:  "",
		Data: id,
	})
	return

}
