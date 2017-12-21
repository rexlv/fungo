// build test
package base_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rexlv/fungo/cream"
	"github.com/rexlv/fungo/cream/middleware"
	. "github.com/smartystreets/goconvey/convey"
)

type MyServer struct {
	*cream.Server
}

func (ms *MyServer) Mux() {
	ms.GET("/ping", pong, middleware.DefaultRecovery)
}

func pong(c cream.Context) error {
	return c.JSON(200, map[string]int64{
		"ping": time.Now().Unix(),
	})
}

func Test_MyServerGET(t *testing.T) {
	server := &MyServer{
		Server: cream.New(),
	}
	server.GET("/ping", pong, middleware.DefaultRecovery)

	Convey("注入一个路由", t, func() {
		r, _ := http.NewRequest("GET", "/ping", nil)
		w := httptest.NewRecorder()
		server.ServeHTTP(w, r)
		So(w.Result().StatusCode, ShouldEqual, 200)
	})
}
