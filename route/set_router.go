package route

import (
	"SmartHouseAPI/handler"
	"github.com/gin-gonic/gin"
)

func SetRouter(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	// Пример группировки роутов
	api := r.Group("/api")
	{
		api.POST("/light/:id", h.ToggleLight)
		api.POST("/heater", h.SetHeater)
		api.POST("/cooler", h.SetCooler)
		api.POST("/login", h.Login)

	}

	return r
}
