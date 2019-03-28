package utils

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CSRFTokenName  = "X-CSRF-TOKEN"
	tokenLength    = 64
	allowedLetters = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM_1234567890"
)

func CSRFGeneration(c *gin.Context) {
	token := generateCSRFToken()
	c.SetCookie(CSRFTokenName, token, 0, "/", "localhost", false, false)
}

func CheckCSRFToken(c *gin.Context) {
	tokenFromHeader := c.GetHeader(CSRFTokenName)
	tokenFromCookie, err := c.Cookie(CSRFTokenName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ServerResponse{Ok: false, Result: err.Error()})
	} else if tokenFromHeader != tokenFromCookie {
		c.AbortWithStatusJSON(http.StatusBadRequest, ServerResponse{Ok: false, Result: "Wrong CSRF Token"})
	}
}

func generateCSRFToken() string {
	res := make([]byte, tokenLength)
	for i := range res {
		res[i] = allowedLetters[rand.Intn(len(allowedLetters))]
	}
	return string(res)
}
