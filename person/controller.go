package person

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PersonController struct {
	lg      *zap.Logger
	service *Service
}

func NewPersonController(lg *zap.Logger, service *Service) *PersonController {
	return &PersonController{
		lg:      lg,
		service: service,
	}
}

// Create
func (p *PersonController) Create(ctx *gin.Context) {

}
