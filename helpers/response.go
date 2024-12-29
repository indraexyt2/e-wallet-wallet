package helpers

import "github.com/gin-gonic/gin"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponseHTTP(c *gin.Context, code int, message string, data interface{}, err error) {
	if err != nil {
		c.JSON(code, Response{
			Status:  false,
			Message: err.Error(),
		})
	}
	c.JSON(code, Response{
		Status:  true,
		Message: message,
		Data:    data,
	})
}
