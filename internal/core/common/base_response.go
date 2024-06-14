package common

type BaseResponse struct {
	Code    int    `json:"code"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message"`
	Errors  any    `json:"errors,omitempty"`
}

func (b *BaseResponse) WithCode(code int) {
	b.Code = code
}

func (b *BaseResponse) WithData(data any) {
	b.Data = data
}

func (b *BaseResponse) WithMessage(msg string) {
	b.Message = msg
}
func (b *BaseResponse) WithErrors(err error) {
	b.Errors = err
}
