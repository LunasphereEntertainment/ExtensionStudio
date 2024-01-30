package templates

import (
	"encoding/xml"
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
		"C:\\Users\\admir\\CodingProjects\\ExtensionStudio\\hacknet\\templates\\Computer.goxml",
		"C:\\Users\\admir\\CodingProjects\\ExtensionStudio\\hacknet\\templates\\ActionSet.goxml",
		"C:\\Users\\admir\\CodingProjects\\ExtensionStudio\\hacknet\\templates\\ExtensionInfo.goxml",
		"C:\\Users\\admir\\CodingProjects\\ExtensionStudio\\hacknet\\templates\\Mission.goxml",
		//"hacknet/templates/Computer.goxml",
		//"hacknet/templates/ExtensionInfo.goxml",
		//"hacknet/templates/Mission.goxml",
		//"hacknet/templates/ActionSet.goxml",
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
		enc := xml.NewEncoder(output)
		enc.Indent("", "    ")
		return enc.Encode(data)
	}

	err := tmpl.ExecuteTemplate(output, tmplName, data)
	return err
}
