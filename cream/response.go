package cream

import (
	"net/http"
	"io"
	"github.com/rexlv/fungo/codec"
)

type Response interface {
	Reset(http.ResponseWriter)
	SetError(error)
	SetStatus(int)
	SetData(interface{})
	SetEncoder(codec.Encoder)
}

type BaseResponse struct {
	http.ResponseWriter
	writer io.Writer

	err error
	data interface{}
	status int
	encoder codec.Encoder

	commited bool
}

func (rep *BaseResponse) Reset(w http.ResponseWriter) {
	rep.ResponseWriter = w
}

func (rep *BaseResponse) SetError(err error) {
	rep.err = err
}

func (rep *BaseResponse) SetData(i interface{}) {
	rep.data = i
}

func (rep *BaseResponse) SetEncoder(encoder codec.Encoder) {
	rep.encoder = encoder
}

func (rep *BaseResponse) SetStatus(code int) {
	rep.status = code
}

func (rep *BaseResponse)  Status() int {
	return rep.status
}

func (rep *BaseResponse) Flush() {
	rep.writer.Write()
}
