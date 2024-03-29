package handling

import (
	"net/http"

	"github.com/kadukm/banking_spa/server/db"

	"github.com/gin-gonic/gin"
	"github.com/kadukm/banking_spa/server/utils"
)

func GetPaymentRequestsSorted(c *gin.Context) {
	sortOptions := utils.MongoSortDTO{}
	if err := c.ShouldBindQuery(&sortOptions); err != nil {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: "Запрос некорректен"})
		return
	}
	if payments, err := db.GetPaymentRequestsSorted(sortOptions); err == nil {
		c.JSON(http.StatusOK, utils.ServerResponse{Ok: true, Result: payments})
	} else {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: "Невозможно получить платеж"})
	}
}

func GetPaymentRequests(c *gin.Context) {
	filter := utils.MongoFilterDTO{}
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: "Запрос некорректен"})
		return
	}
	if payments, err := db.GetPaymentRequests(filter); err == nil {
		c.JSON(http.StatusOK, utils.ServerResponse{Ok: true, Result: payments})
	} else {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: "Невозможно получить платеж"})
	}
}

func PostPaymentRequest(c *gin.Context) {
	if !utils.MIMEContentTypeIsJSON(c.Request) {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: "Wrong Content-Type"})
		return
	}
	request := utils.PaymentRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: "Запрос некорректен"})
		return
	}
	if !paymentRequestIsRight(request) {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: "Wrong data"})
		return
	}
	if err := db.AddNewPaymentRequest(request); err != nil {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: "Невозможно добавить платеж"})
		return
	}

	c.JSON(http.StatusOK, utils.ServerResponse{Ok: true, Result: "Запрос прошел успешно"})
}

func paymentRequestIsRight(request utils.PaymentRequestDTO) bool {
	return utils.IDIsRight(request.ID) &&
		utils.INNIsRight(request.INN) &&
		utils.BIKIsRight(request.BIK) &&
		utils.AccountNumberIsRight(request.AccountNumber) &&
		utils.ForWhatIsRight(request.ForWhat) &&
		utils.AmountIsRight(request.Amount) &&
		utils.PhoneIsRight(request.Phone) &&
		utils.EmailIsRight(request.Email)
}
