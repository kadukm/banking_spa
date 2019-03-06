package handling

import (
	"net/http"

	"github.com/kadukm/banking_spa/server/utils"

	"github.com/gin-gonic/gin"
)

func GetCompany(c *gin.Context) {
	company := utils.CompanyDTO{
		Status:         "Индивидуальный предприниматель",
		Name:           "Швецова Мария Валерьевна",
		PhotoPath:      "/static/images/test_photo.jpg",
		Phone:          "+79998887766",
		Site:           "www.mary.com",
		Email:          "mary@tochka.ru",
		FullInfoPath:   "#",
		RequisitesPath: "#",
	}
	c.JSON(http.StatusOK, company)
}
