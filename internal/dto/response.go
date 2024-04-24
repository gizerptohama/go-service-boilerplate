package dto

type Response struct {
	PageCount int     `json:"page_count,omitempty"`
	ItemCount int     `json:"item_count,omitempty"`
	Page      int     `json:"page,omitempty"`
	Message   string  `json:"message,omitempty"`
	Data      any     `json:"data,omitempty"`
	Errors    []error `json:"errors,omitempty"`
}

func NewMessageResponse(message string) (resp *Response) {
	return &Response{Message: message}
}

func NewDataResponse(data any) (resp *Response) {
	return &Response{Data: data}
}

func NewDataPaginationResponse(data any, pageCount int, itemCount int, page int) (resp *Response) {
	return &Response{PageCount: pageCount, ItemCount: itemCount, Data: data, Page: page}
}

func NewErrorResponse(err error) (resp *Response) {
	return &Response{Message: err.Error()}
}

func NewMultiErrorResponse(err ...error) (resp *Response) {
	errs := make([]error, 0)
	errs = append(errs, err...)
	return &Response{Errors: errs}
}

type TokenResponse struct {
	Token string `json:"token"`
}

func NewTokenResponse(token string) (resp *TokenResponse) {
	return &TokenResponse{Token: token}
}
