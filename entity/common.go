package entity

// Resp is the common struct of http response body
type Resp struct {
	Success bool   `json:"success"`
	ErrNo   int    `json:"errno"`
	Errmsg  string `json:"errmsg"`
}
