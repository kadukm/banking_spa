package handling

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kadukm/banking_spa/server/db"
	"github.com/kadukm/banking_spa/server/utils"
)

func PostPaymentFromCard(c *gin.Context) {
	payment := utils.PaymentFromCard{}
	if err := c.ShouldBindJSON(&payment); err == nil {
		//TODO: check all fields
		db.AddNewPaymentFromCard(payment)
		c.JSON(http.StatusOK, utils.ServerResponse{Ok: true, Message: "All is ok c:"})
	} else {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Message: "I don't feel so good..."})
	}
}

func PatchPaymentFromCard(c *gin.Context) {
	paymentIDstr := c.Param("paymentId")
	paymentID, _ := strconv.ParseUint(paymentIDstr, 10, 64)
	patch := utils.PaymentFromCardPatch{}
	if err := c.ShouldBindJSON(&patch); err == nil {
		db.PatchPaymentFromCard(patch, paymentID)
		c.JSON(http.StatusOK, utils.ServerResponse{Ok: true, Message: "All is ok c:"})
	} else {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Message: "I don't feel so good..."})
	}
}
