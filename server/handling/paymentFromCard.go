package handling

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kadukm/banking_spa/server/db"
	"github.com/kadukm/banking_spa/server/utils"
)

func PostPaymentFromCard(c *gin.Context) {
	if !utils.MIMEContentTypeIsJSON(c.Request) {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Message: "Wrong Content-Type"})
		return
	}
	request := utils.PaymentFromCardDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Message: err.Error()})
		return
	}
	if !paymentFromCardIsRight(request) {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Message: "Wrong data"})
		return
	}
	if err := db.AddNewPaymentFromCard(request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.ServerResponse{Ok: true, Message: "All is ok c:"})
}

func PatchPaymentFromCard(c *gin.Context) {
	paymentID := c.Param("paymentID")
	patch := utils.PatchPaymentFromCardDTO{}
	if err := c.ShouldBindJSON(&patch); err != nil {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Message: err.Error()})
		return
	}
	if err := db.PatchPaymentFromCard(patch, paymentID); err != nil {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.ServerResponse{Ok: true, Message: "All is ok c:"})
}

func paymentFromCardIsRight(payment utils.PaymentFromCardDTO) bool {
	return utils.IDIsRight(payment.ID) &&
		utils.CardNumberIsRight(payment.CardNumber) &&
		utils.CardExpiresIsRight(payment.CardExpires) &&
		utils.CardCvcIsRight(payment.CardCVC) &&
		utils.ValueIsRight(payment.Value) &&
		utils.CommentIsRight(payment.Comment) &&
		utils.EmailIsRight(payment.Email)
}
