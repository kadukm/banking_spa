package handling

import (
	"net/http"

	"github.com/kadukm/banking_spa/server/utils"

	"github.com/gin-gonic/gin"
)

func GetCompany(c *gin.Context) {
	company := utils.CompanyDTO{
		Status:         "Индивидуальный предприниматель",
		Name:           "Гофер",
		PhotoPath:      "/assets/images/gopher.jpg",
		Phone:          "+79998887766",
		Site:           "golang.org",
		Email:          "gopher@golang.org",
		FullInfoPath:   "#",
		RequisitesPath: "#",
	}
	c.JSON(http.StatusOK, company)
}
