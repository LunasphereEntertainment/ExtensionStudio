package templates

import (
	"fmt"
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet"
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet/nodes"
	"io"
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
		"Mission.goxml",
		"ActionSet.goxml",
	))
	templateNames = map[reflect.Type]string{
		reflect.TypeOf(nodes.Computer{}):               "Computer.goxml",
		reflect.TypeOf(hacknet.ExtensionInfo{}):        "ExtensionInfo.goxml",
		reflect.TypeOf(hacknet.Mission{}):              "Mission.goxml",
		reflect.TypeOf(hacknet.ConditionalActionSet{}): "ActionSet.goxml",
	}
}

func ExecuteTemplate[T interface{}](data T, output io.Writer) error {
	tmplName, ok := templateNames[reflect.TypeOf(data)]
	if !ok {
		return fmt.Errorf("no template loaded for type '%s'", reflect.TypeOf(data))
	}

	err := tmpl.ExecuteTemplate(output, tmplName, data)
	return err
}
