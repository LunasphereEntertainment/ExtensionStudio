package templates

import (
	"bytes"
	"fmt"
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet"
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet/nodes"
	"reflect"
	"text/template"
)

var (
	tmpl          *template.Template
	templateNames map[reflect.Type]string
)

func init() {
	tmpl = template.Must(template.ParseFiles(
		"Computer.goxml",
		"ExtensionInfo.goxml",
	))
	templateNames = map[reflect.Type]string{
		reflect.TypeOf(nodes.Computer{}):        "Computer.goxml",
		reflect.TypeOf(hacknet.ExtensionInfo{}): "ExtensionInfo.goxml",
	}
}

func ExecuteTemplate[T interface{}](data T) ([]byte, error) {
	tmplName, ok := templateNames[reflect.TypeOf(data)]
	if !ok {
		return nil, fmt.Errorf("no template loaded for type '%s'", reflect.TypeOf(data))
	}

	var buff bytes.Buffer
	err := tmpl.ExecuteTemplate(&buff, tmplName, data)
	return buff.Bytes(), err
}
