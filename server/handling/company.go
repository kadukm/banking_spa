package handling

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCompanyInfo(c *gin.Context) {
	c.String(http.StatusOK, "not implemented yet")
}
