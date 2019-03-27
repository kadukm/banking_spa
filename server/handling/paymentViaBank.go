package handling

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"github.com/kadukm/banking_spa/server/utils"
)

func GetPaymentViaBank(c *gin.Context) {
	payment := utils.PaymentViaBankDTO{}
	if err := c.ShouldBindQuery(&payment); err != nil {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: err.Error()})
	} else if !paymentViaBankIsRight(payment) {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: "wrong data"})
	} else if err := buildPaymentViaBankPDF(payment, c.Writer); err != nil {
		c.JSON(http.StatusBadRequest, utils.ServerResponse{Ok: false, Result: err.Error()})
	}
}

func paymentViaBankIsRight(payment utils.PaymentViaBankDTO) bool {
	return utils.INNIsRight(payment.INN) &&
		utils.BIKIsRight(payment.BIK) &&
		utils.AccountNumberIsRight(payment.AccountNumber) &&
		utils.ForWhatIsRight(payment.ForWhat) &&
		utils.AmountIsRight(payment.Amount)
}

func buildPaymentViaBankPDF(payment utils.PaymentViaBankDTO, writer io.Writer) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1251")
	pdf.AddPage()
	pdf.AddFont("Helvetica", "", "helvetica_1251.json")
	pdf.SetFont("Helvetica", "", 16)
	_, lineHt := pdf.GetFontSize()
	text := buildTextForPaymentViaBank(payment)
	pdf.Write(lineHt, tr(text))
	return pdf.Output(writer)
}

func buildTextForPaymentViaBank(payment utils.PaymentViaBankDTO) string {
	return fmt.Sprintf(`ИНН: %s
БИК: %s
Номер счёта: %s
За что: %s
Сумма: %v`,
		payment.INN, payment.BIK, payment.AccountNumber, payment.ForWhat, payment.Amount)
}
