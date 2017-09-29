package stone

type SaleError struct {
	ErrorReport *ErrorReport `json:",omitempty"`
}

type ErrorReport struct {
	Category            string       `json:",omitempty"`
	ErrorItemCollection []*ErrorItem `json:",omitempty"`
}

type ErrorItem struct {
	Description  string `json:",omitempty"`
	ErrorCode    uint32 `json:",omitempty"`
	ErrorField   string `json:",omitempty"`
	SeverityCode string `json:",omitempty"`
}
