package handling

import (
	"net/http"
	"strconv"

	"github.com/kadukm/banking_spa/server/db"
	"github.com/kadukm/banking_spa/server/utils"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	companyID := c.Param("companyID")
	strMaxCount := c.DefaultQuery("maxcount", "3")
	maxCount, err := strconv.ParseInt(strMaxCount, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: "Запрос некорректен"})
		return
	}

	if products, err := db.GetProducts(companyID, maxCount); err == nil {
		c.JSON(http.StatusOK, utils.ServerResponse{Ok: true, Result: products})
	} else {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: "Невозможно получить платеж"})
	}
}
