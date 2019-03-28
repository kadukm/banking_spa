package handling

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kadukm/banking_spa/server/utils"
)

const (
	CSRFTokenName      = "X-CSRF-TOKEN"
	tokenLength        = 64
	allowedCSRFLetters = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM_1234567890"
)

func CSRFGeneration(c *gin.Context) {
	token := generateCSRFToken()
	c.SetCookie(CSRFTokenName, token, 0, "/", "localhost", false, false)
}

func CheckCSRFToken(c *gin.Context) {
	tokenFromHeader := c.GetHeader(CSRFTokenName)
	tokenFromCookie, err := c.Cookie(CSRFTokenName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: "Wrong CSRF Token"})
	} else if tokenFromHeader != tokenFromCookie {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: "Wrong CSRF Token"})
	}
}

func generateCSRFToken() string {
	res := make([]byte, tokenLength)
	for i := range res {
		res[i] = allowedCSRFLetters[rand.Intn(len(allowedCSRFLetters))]
	}
	return string(res)
}
