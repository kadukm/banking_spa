package handling

import (
	"net/http"

	"github.com/kadukm/banking_spa/server/db"

	"github.com/gin-gonic/gin"
	"github.com/kadukm/banking_spa/server/utils"
)

func PostPaymentRequest(c *gin.Context) {
	request := utils.PaymentRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Message: err.Error()})
		return
	}
	if !paymentRequestIsRight(request) {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Message: "Wrong data"})
		return
	}
	if err := db.AddNewPaymentRequest(request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.ServerResponse{Ok: true, Message: "All is ok c:"})
}

func paymentRequestIsRight(request utils.PaymentRequestDTO) bool {
	return utils.IDIsRight(request.ID) &&
		utils.InnIsRight(request.Inn) &&
		utils.BikIsRight(request.Bik) &&
		utils.AccountNumberIsRight(request.AccountNumber) &&
		utils.ForWhatIsRight(request.ForWhat) &&
		utils.ValueIsRight(request.Value) &&
		utils.PhoneIsRight(request.Phone) &&
		utils.EmailIsRight(request.Email)
}
