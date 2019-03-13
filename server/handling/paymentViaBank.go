package handling

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kadukm/banking_spa/server/utils"
)

func PostPaymentViaBank(c *gin.Context) {
	if !utils.MIMEContentTypeIsJSON(c.Request) {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: "Wrong Content-Type"})
		return
	}
	payment := utils.PaymentViaBankDTO{}
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: err.Error()})
		return
	}
	if !paymentViaBankIsRIght(payment) {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: "wrong data"})
		return
	}
	//TODO: send read pdf
	c.JSON(http.StatusOK, utils.ServerResponse{Ok: true, Result: "All is ok c:"})
}

func paymentViaBankIsRIght(payment utils.PaymentViaBankDTO) bool {
	return utils.INNIsRight(payment.INN) &&
		utils.BIKIsRight(payment.BIK) &&
		utils.AccountNumberIsRight(payment.AccountNumber) &&
		utils.ForWhatIsRight(payment.ForWhat) &&
		utils.AmountIsRight(payment.Amount)
}
