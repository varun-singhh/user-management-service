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

type Register struct {
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Password   string `json:"password,omitempty"`
	Permission string `json:"permission,omitempty"`
}

type Data struct {
	Data struct {
		User struct {
			ID         int    `json:"id"`
			Email      string `json:"email,omitempty"`
			Phone      string `json:"phone,omitempty"`
			Password   string `json:"password"`
			Permission string `json:"permission"`
			Status     string `json:"status"`
			CreatedAt  string `json:"created_at"`
		} `json:"user"`
	} `json:"data"`
}
