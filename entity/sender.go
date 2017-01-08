package entity

// Sender is struct of SMTP/POP3 config
type Sender struct {
	Type      string `json:"type"`
	Server    string `json:"server"`
	Password  string `json:"password"`
	EnableSSL bool   `json:"enable_ssl"`
}
