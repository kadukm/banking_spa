package handling

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kadukm/banking_spa/server/utils"
)

func GetPaymentViaBank(c *gin.Context) {
	payment := utils.PaymentViaBankDTO{}
	if err := c.ShouldBindQuery(&payment); err != nil {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: err.Error()})
		return
	}
	if !paymentViaBankIsRIght(payment) {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: "wrong data"})
		return
	}
	//TODO: send real pdf
	c.JSON(http.StatusOK, utils.ServerResponse{Ok: true, Result: "All is ok c:"})
}

func paymentViaBankIsRIght(payment utils.PaymentViaBankDTO) bool {
	return utils.INNIsRight(payment.INN) &&
		utils.BIKIsRight(payment.BIK) &&
		utils.AccountNumberIsRight(payment.AccountNumber) &&
		utils.ForWhatIsRight(payment.ForWhat) &&
		utils.AmountIsRight(payment.Amount)
}
