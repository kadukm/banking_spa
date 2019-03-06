package handling

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCompany(c *gin.Context) {
	c.String(http.StatusOK, "not implemented yet")
}
