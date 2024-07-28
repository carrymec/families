package relationship

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Controller struct {
	lg      *zap.Logger
	service *Service
}

func NewRelationController(lg *zap.Logger, service *Service) *Controller {
	return &Controller{
		lg:      lg,
		service: service,
	}
}

func (p *Controller) Create(ctx *gin.Context) {

}

func (p *Controller) Query(ctx *gin.Context) {

}

func (p *Controller) FindById(ctx *gin.Context) {

}

func (p *Controller) Update(ctx *gin.Context) {

}

func (p *Controller) Delete(ctx *gin.Context) {

}
