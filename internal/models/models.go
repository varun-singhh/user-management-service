package models

type Page struct {
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}

type AllDataResponse struct {
	Count  string      `json:"count"`
	Offset string      `json:"offset"`
	Limit  string      `json:"limit"`
	Data   interface{} `json:"data"`
}
