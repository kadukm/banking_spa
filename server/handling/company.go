package handling

import (
	"net/http"

	"github.com/kadukm/banking_spa/server/db"
	"github.com/kadukm/banking_spa/server/utils"

	"github.com/gin-gonic/gin"
)

func GetCompany(c *gin.Context) {
	companyID := c.Param("companyID")
	if company, err := db.GetCompany(companyID); err == nil {
		c.JSON(http.StatusOK, company)
	} else {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Message: err.Error()})
	}
}
