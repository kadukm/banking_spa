package handling

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kadukm/banking_spa/server/db"
	"github.com/kadukm/banking_spa/server/utils"
)

const (
	sidName           = "SID"
	sidLength         = 64
	allowedSIDLetters = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM_1234567890"
)

var sid2login map[string]string
var login2sid map[string]string

func init() {
	sid2login = make(map[string]string)
	login2sid = make(map[string]string)
}

func Login(c *gin.Context) {
	userInfoDTO := utils.UserInfoDTO{}
	if err := c.BindJSON(&userInfoDTO); err != nil {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: err.Error()})
		return
	}
	user, err := db.GetUser(userInfoDTO.Login)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: err.Error()})
		return
	}
	calculatedHash := calculateHash(user.Salt, userInfoDTO.Password)
	if calculatedHash != user.Hash {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: "Неверный логин или пароль"})
		return
	}

	if oldSid, ok := login2sid[userInfoDTO.Login]; ok {
		delete(sid2login, oldSid)
	}
	sid := generateSID()
	sid2login[sid] = userInfoDTO.Login
	login2sid[userInfoDTO.Login] = sid
	c.SetCookie(sidName, sid, 0, "/", "localhost", false, true)
	c.JSON(http.StatusOK, utils.ServerResponse{Ok: true, Result: "Авторизация прошла успешно"})
}

func CheckSession(c *gin.Context) {
	if sid, err := c.Cookie(sidName); err == nil {
		login := sid2login[sid]
		user, err := db.GetUser(login)
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: err.Error()})
		}
		if user.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: "Access denie"})
		}
	} else {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		c.Abort()
	}
}

func generateSID() string {
	res := make([]byte, sidLength)
	for i := range res {
		res[i] = allowedSIDLetters[rand.Intn(len(allowedSIDLetters))]
	}
	return string(res)
}

func calculateHash(salt, password string) string {
	h := sha1.New()
	io.WriteString(h, salt)
	io.WriteString(h, password)
	return fmt.Sprintf("%x", h.Sum(nil))
}
