package Models

type Transporter struct {
	IdTransporter int       `json:"IdTransporter,omitempty"`
	Orders        []Order   `json:"Orders,omitempty"`
	Products      []Product `json:"Products,omitempty"`
}
