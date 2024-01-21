package hacknet

import (
	"encoding/xml"
	"fmt"
	"os"
	"path"
)

type ExternalReference[T interface{}] string

func (e *ExternalReference[T]) Load(basePath string) (*T, error) {
	if *e == "NONE" {
		return nil, fmt.Errorf("nothing to load here")
	}

	f, err := os.Open(path.Join(basePath, string(*e)))
	if err != nil {
		return nil, err
	}

	out := new(T)

	err = xml.NewDecoder(f).Decode(out)

	return out, err
}
