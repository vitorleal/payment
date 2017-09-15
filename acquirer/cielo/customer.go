package cielo

type Customer struct {
	Name            string
	Email           string   `json:",omitempty"`
	Birthdate       string   `json:",omitempty"`
	Address         *Address `json:",omitempty"`
	DeliveryAddress *Address `json:",omitempty"`
}

type Address struct {
	Street     string `json:",omitempty"`
	Number     string `json:",omitempty"`
	Complement string `json:",omitempty"`
	ZipCode    string `json:",omitempty"`
	City       string `json:",omitempty"`
	State      string `json:",omitempty"`
	Country    string `json:",omitempty"`
}
