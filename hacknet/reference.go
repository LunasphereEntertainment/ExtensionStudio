package hacknet

import (
	"encoding/xml"
	"fmt"
	"github.com/google/uuid"
	"os"
	"path"
)

type ExternalReference[T interface{}] struct {
	Path      string
	ProjectID uuid.UUID
}

func (e *ExternalReference[T]) Load(basePath string) (*T, error) {
	if e.Path == "NONE" {
		return nil, fmt.Errorf("nothing to load here")
	}

	f, err := os.Open(path.Join(basePath, e.Path))
	if err != nil {
		return nil, err
	}

	out := new(T)

	err = xml.NewDecoder(f).Decode(out)

	return out, err
}
