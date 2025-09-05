package middleware

import (
	MessageTemplate "authentication/pkg/templates"
	"authentication/utils"
	"authentication/utils/logger"
	"github.com/gin-gonic/gin"
)

func ErrorHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				var pm utils.PanicMessage
				switch x := r.(type) {
				case utils.PanicMessage:
					pm = x
				default:
					pm = utils.PanicMessage{MessageKey: 0}
				}

				if pm.Error != nil {
					eventData := map[string]interface{}{
						"error":   *pm.Error,
						"depth":   4,
						"message": "An Error Occurred",
					}
					logger.LogErrorWithDepth(eventData)
				}

				// Fetch the message template using the key
				template, exists := MessageTemplate.MessageTemplates[pm.MessageKey]
				if !exists {
					template = MessageTemplate.MessageTemplates[0] // Default message
				}
				c.JSON(template.Status, template.Message)
				c.Abort()
			}
		}()

		c.Next()
	}
}
