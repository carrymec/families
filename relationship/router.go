package relationship

import "github.com/gin-gonic/gin"

func (p *Controller) Register(engine *gin.Engine) {
	engine.POST("/api/v1/relations", p.Create)
	engine.POST("/api/v1/query_relations", p.Query)
	engine.GET("/api/v1/relations/:id", p.FindById)
	engine.PUT("/api/v1/relations/:id", p.Update)
	engine.DELETE("/api/v1/relations/:id", p.Delete)
}
