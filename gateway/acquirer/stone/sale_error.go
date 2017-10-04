package stone

type (
	// SaleError represents an Stone api error
	SaleError struct {
		ErrorReport *ErrorReport `json:",omitempty"`
	}

	// Error report represents the Stone api error report
	ErrorReport struct {
		Category            string       `json:",omitempty"`
		ErrorItemCollection []*ErrorItem `json:",omitempty"`
	}

	// Error item represnet the Stone api error report item
	ErrorItem struct {
		Description  string `json:",omitempty"`
		ErrorCode    uint32 `json:",omitempty"`
		ErrorField   string `json:",omitempty"`
		SeverityCode string `json:",omitempty"`
	}
)
