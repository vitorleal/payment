package cielo

type SaleStatus string

var saleStatus = [...]SaleStatus{
	"NotFinished",
	"Authorized",
	"PaymentConfirmed",
	"Denied",
	"Voided",
	"Refunded",
	"Pending",
	"Aborted",
	"Schedueled",
}

// Get the value for the sale status
func (ss SaleStatus) String() string {
	return ""
}
