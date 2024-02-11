package main

import (
	"encoding/xml"
	"github.com/LunasphereEntertainment/ExtensionStudio/hacknet/templates"
	"os"
)

func LoadXML[T interface{}](path string) (out *T, err error) {
	out = new(T)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	err = xml.NewDecoder(f).Decode(out)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	return out, err
}

func SaveXML(path string, model interface{}) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC, 644)
	if err != nil {
		return err
	}

	err = templates.ExecuteTemplate(model, f)
	if err != nil {
		return err
	}
	return f.Close()
}
