package gateway

type PaymentStatus uint32

const (
	Approved PaymentStatus = 1 + iota
	Confirmed
	Denied
	Canceled
	Refunded
	Pending
	Aborted
	Schedueled
)

var paymentStatus = [...]string{
	"approved",
	"confirmed",
	"denied",
	"canceled",
	"refunded",
	"pending",
	"aborted",
	"schedueled",
}

// Get the value of the payment status
func (ps PaymentStatus) String() string {
	return ""
}
