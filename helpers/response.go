package helpers

import "github.com/gin-gonic/gin"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponseHTTP(c *gin.Context, code int, status bool, message string, data interface{}) {
	resp := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	c.JSON(code, resp)
}
