package person

import "github.com/gin-gonic/gin"

func (p *Controller) Register(engine *gin.Engine) {
	engine.POST("/api/v1/create_person", p.Create)
	engine.POST("/api/v1/query_persons", p.Query)
	engine.GET("/api/v1/persons/:id", p.FindById)
}
