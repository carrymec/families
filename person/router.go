package person

import "github.com/gin-gonic/gin"

func (p *Controller) Register(engine *gin.Engine) {
	engine.POST("/api/v1/create_person", p.Create)
}
