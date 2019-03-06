package handling

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kadukm/banking_spa/server/utils"
)

func PostPaymentViaBank(c *gin.Context) {
	payment := utils.PaymentViaBankDTO{}
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Message: err.Error()})
		return
	}
	if !paymentViaBankIsRIght(payment) {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Message: "wrong data"})
		return
	}
	//TODO: send read pdf
	c.JSON(http.StatusOK, utils.ServerResponse{Ok: true, Message: "All is ok c:"})
}

func paymentViaBankIsRIght(payment utils.PaymentViaBankDTO) bool {
	return utils.INNIsRight(payment.INN) &&
		utils.BIKIsRight(payment.BIK) &&
		utils.AccountNumberIsRight(payment.AccountNumber) &&
		utils.ForWhatIsRight(payment.ForWhat) &&
		utils.ValueIsRight(payment.Value)
}
