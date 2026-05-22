package manifest

import (
	"encoding/json"
	"os"
)

type Segment struct {
	ID      uint64 `json:"id"`
	Path    string `json:"path"`
	MinKey  uint64 `json:"min_key"`
	MaxKey  uint64 `json:"max_key"`
	Entries uint64 `json:"entries"`
	Level   uint64 `json:"level"`
}

type Manifest struct {
	Segments []Segment `json:"segments"`
}

func Open(path string) (*Manifest, error) {

	file, err := os.Open(path)

	if err != nil {

		if os.IsNotExist(err) {

			return &Manifest{
				Segments: []Segment{},
			}, nil
		}

		return nil, err
	}

	defer file.Close()

	var manifest Manifest

	err = json.NewDecoder(file).Decode(&manifest)

	if err != nil {
		return nil, err
	}

	return &manifest, nil
}

func (m *Manifest) Save(path string) error {

	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	encoder.SetIndent("", "  ")

	return encoder.Encode(m)
}
