package stone

const Name = "stone"

// BankingBillet information
const (
	SantanderBankNumber = "033" // Santander
)

// CreditCard operation
const (
	Authorize           = "AuthOnly"
	AuthorizeAndCapture = "AuthAndCapture"
)

// Autorization status
const (
	Voided                   = "Voided"
	PartialCanceled          = "PartialVoid"
	Refunded                 = "Refunded"
	PartialRefund            = "PartialRefund"
	Captured                 = "Captured"
	NotAuthorized            = "NotAuthorized"
	AuthorizedPendingCapture = "AuthorizedPendingCapture"
	WithError                = "WithError"
	PendingVoid              = "PendingVoid"
	PendingRefund            = "PendingRefund"
)

// Currency
const BRL = "BRL"

// Document Type
const (
	CPF  = "CPF"
	CNPJ = "CNPJ"
)

// Recurrent Payment
const (
	RecurrentDaily   = "Daily"
	RecurrentWeekly  = "Weekly"
	RecurrentMonthly = "Monthly"
	RecurrentYearly  = "Yearly"
)
