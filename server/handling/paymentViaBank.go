package handling

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kadukm/banking_spa/server/utils"
)

func PostPaymentViaBank(c *gin.Context) {
	request := utils.PaymentRequestDTO{}
	if err := c.ShouldBindJSON(&request); err == nil {
		//TODO: check all fields
		//TODO: send real pdf
		c.JSON(http.StatusOK, utils.ServerResponse{Ok: true, Message: "All is ok c:"})
	} else {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Message: "I don't feel so good..."})
	}
}
