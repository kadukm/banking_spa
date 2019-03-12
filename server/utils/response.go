package utils

type ServerResponse struct {
	Ok     bool        `json:"ok"`
	Result interface{} `json:"result"`
}
