package main

import (
	"encoding/xml"
	"os"
)

func LoadXML[T interface{}](path string) (*T, error) {
	out := new(T)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	err = xml.NewDecoder(f).Decode(out)

	return out, err
}

func SaveXML(path string, model interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	return xml.NewEncoder(f).Encode(model)
}
