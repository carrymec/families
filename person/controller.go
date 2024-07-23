package person

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github/carrymec/families/common"
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

func (p *Controller) Query(ctx *gin.Context) {
	var per Query
	if err := ctx.ShouldBind(&per); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	persons, err := p.service.Query(ctx, per)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Resp{
			Code: -1,
			Msg:  fmt.Sprintf("查询用户失败: %s", err.Error()),
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Resp{
		Code: 0,
		Msg:  "",
		Data: persons,
	})
	return

}

func (p *Controller) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusInternalServerError, common.Resp{
			Code: -1,
			Msg:  "参数id为空",
			Data: nil,
		})
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Resp{
			Code: -1,
			Msg:  fmt.Sprintf("参数id错误 %s", err.Error()),
			Data: nil,
		})
		return
	}
	person, err := p.service.FindById(ctx, int64(intId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Resp{
			Code: -1,
			Msg:  fmt.Sprintf("查询用户失败 %s", err.Error()),
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Resp{
		Code: 0,
		Data: person,
	})
	return
}

func (p *Controller) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusInternalServerError, common.Resp{
			Code: -1,
			Msg:  "参数id为空",
			Data: nil,
		})
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Resp{
			Code: -1,
			Msg:  fmt.Sprintf("参数id错误 %s", err.Error()),
			Data: nil,
		})
		return
	}
	var per Person
	if err := ctx.ShouldBind(&per); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = p.service.Update(ctx, int64(intId), per)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Resp{
			Code: -1,
			Msg:  fmt.Sprintf("更新用户失败 %s", err.Error()),
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Resp{
		Code: 0,
		Data: intId,
	})
	return
}

func (p *Controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusInternalServerError, common.Resp{
			Code: -1,
			Msg:  "参数id为空",
			Data: nil,
		})
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Resp{
			Code: -1,
			Msg:  fmt.Sprintf("参数id错误 %s", err.Error()),
			Data: nil,
		})
		return
	}
	err = p.service.Delete(ctx, int64(intId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Resp{
			Code: -1,
			Msg:  fmt.Sprintf("删除用户失败 %s", err.Error()),
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Resp{
		Code: 0,
		Data: nil,
	})
	return
}
