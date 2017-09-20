package cream

import (
	"encoding/json"
	"encoding/xml"
)

func RegisterRender(mime MIME, render Renderer) {
	renders[mime] = render
}

func render(mime MIME) Renderer {
	if render, ok := renders[mime]; ok {
		return render
	}

	return renders[MIMEApplicationJSON]
}

var renders = map[MIME]Renderer{
	MIMEApplicationJSON:            defaultJSONRender,
	MIMEApplicationJSONCharsetUTF8: defaultJSONRender,
	MIMEApplicationXML:             defaultXMLRender,
	MIMEApplicationXMLCharsetUTF8:  defaultXMLRender,
	MIMEApplicationProtobuf:        defaultProtobufRender,
}

// Renderer
type Renderer interface {
	Render(interface{}) ([]byte, error)
}

type JSONRender struct {
	Pretty bool
}

var defaultJSONRender = JSONRender{Pretty: false}
var defaultXMLRender = XMLRender{}
var defaultProtobufRender = ProtobufRender{}

func (render JSONRender) Render(obj interface{}) (bs []byte, err error) {
	return json.Marshal(obj)
}

type XMLRender struct {
}

func (render XMLRender) Render(obj interface{}) (bs []byte, err error) {
	return xml.Marshal(obj)
}

type ProtobufRender struct {
}

func (render ProtobufRender) Render(obj interface{}) (bs []byte, err error) {
	return
}
