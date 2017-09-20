package cream

import (
	"io"
	"net/http"
)

type Response interface {
	Reset(http.ResponseWriter)
	SetError(error)
	SetStatus(int)
	SetData(interface{})
	Writer() io.Writer
	Header() http.Header
}

type BaseResponse struct {
	http.ResponseWriter
	writer io.Writer

	err    error
	data   interface{}
	status int

	commited bool
}

func (rep *BaseResponse) Header() http.Header {
	return rep.ResponseWriter.Header()
}

func (rep *BaseResponse) Writer() io.Writer {
	return rep.ResponseWriter
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

func (rep *BaseResponse) SetStatus(code int) {
	rep.status = code
}

func (rep *BaseResponse) Status() int {
	return rep.status
}

func (rep *BaseResponse) Flush() {
}
