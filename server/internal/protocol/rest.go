package protocol

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohwa/ci-cd-github-action/internal/response"
)

func InitRouterGroupRest(rest *gin.RouterGroup) {
	rest.GET("/appInfo", GetAppInfo)
	rest.POST("/appInfo", SetAppInfo)
}

func GetAppInfo(c *gin.Context) {
	c.JSON(http.StatusOK, response.JsonResponse{
		Result: map[string]any{
			"bgColor": "red",
		},
	})
}

func SetAppInfo(c *gin.Context) {
	c.JSON(http.StatusOK, response.JsonResponse{
		Result: "OK",
	})
}
