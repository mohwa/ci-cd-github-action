package protocol

import (
	"github.com/mohwa/ci-cd-github-action/api/rest/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohwa/ci-cd-github-action/internal/handler/rest/settings"
	"github.com/mohwa/ci-cd-github-action/internal/response"
)

func InitRouterGroupRest(restRouterGroup *gin.RouterGroup) {
	restRouterGroup.GET("/settings", GetSettings)
	restRouterGroup.POST("/settings", SetSettings)
}

func GetSettings(c *gin.Context) {
	result, err := settings.GetSettings()
	if err != nil {
		c.JSON(http.StatusOK, response.JsonResponse{
			Result: nil,
		})

		return
	}

	c.JSON(http.StatusOK, response.JsonResponse{
		Result: model.Settings{
			BgColor: result.BgColor,
		},
	})
}

func SetSettings(c *gin.Context) {
	var json model.Settings

	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.JsonResponse{
			Result: gin.H{
				"code":    -1,
				"message": err.Error(),
			},
		})
		return
	}

	err = settings.SetSettings(model.Settings{BgColor: json.BgColor})
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.JsonResponse{
			Result: gin.H{
				"code":    -1,
				"message": err.Error(),
			},
		})

		return
	}

	c.JSON(http.StatusOK, response.JsonResponse{
		Result: "OK",
	})
}
