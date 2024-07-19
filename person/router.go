package person

import "github.com/gin-gonic/gin"

func (s *PersonController) Register(engine *gin.Engine) {
	engine.POST("/api/v1/person", s.Create)
}
