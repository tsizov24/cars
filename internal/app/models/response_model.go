package models

type Response struct {
	Description string `json:"description:"`
}

func GetResponse(description string) *Response {
	return &Response{Description: description}
}
