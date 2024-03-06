package request

type RequestID struct {
	ID uint `json:"id" validate:"required"`
}

type WxLogin struct {
	Code string `json:"code" validate:"required"`
}
