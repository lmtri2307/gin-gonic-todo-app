package auth

import (
	"go-todo-app/utils/tokenutil"
	"strings"

	"github.com/gin-gonic/gin"
)

func toPayload(rawPayload any) (*TokenPayload, error) {
	payloadMap, ok := rawPayload.(map[string]interface{})
	if !ok {
		return nil, &Errors.InvalidToken
	}

	var payload TokenPayload

	if id, ok := payloadMap["Id"].(float64); ok {
		payload.Id = int(id)
	} else {
		return nil, &Errors.InvalidToken
	}

	return &payload, nil
}
func JwtAuthMiddleware(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	splitedStrings := strings.Split(authHeader, " ")

	if len(splitedStrings) != 2 {
		c.Error(&Errors.UnAuthorized)
		c.Abort()
		return
	}

	authToken := splitedStrings[1]
	rawPayload, err := tokenutil.VerifyToken(authToken, "secret")
	if err != nil {
		c.Error(&Errors.InvalidToken)
		c.Abort()
		return
	}

	payload, err := toPayload(rawPayload)
	if err != nil {
		c.Error(&Errors.InvalidToken)
		c.Abort()
		return
	}

	c.Set("x-user-id", payload.Id)
	c.Next()
}
