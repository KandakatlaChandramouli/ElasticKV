package checkpoint

import (
	"encoding/gob"
	"os"
)

func Save(
	path string,
	state map[uint64][]byte,
	metadata Metadata,
) error {

	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := gob.NewEncoder(
		file,
	)

	err = encoder.Encode(metadata)

	if err != nil {
		return err
	}

	return encoder.Encode(state)
}

func Load(
	path string,
) (
	map[uint64][]byte,
	Metadata,
	error,
) {

	file, err := os.Open(path)

	if err != nil {

		return nil,
			Metadata{},
			err
	}

	defer file.Close()

	decoder := gob.NewDecoder(
		file,
	)

	var metadata Metadata

	err = decoder.Decode(
		&metadata,
	)

	if err != nil {

		return nil,
			Metadata{},
			err
	}

	state := make(
		map[uint64][]byte,
	)

	err = decoder.Decode(
		&state,
	)

	if err != nil {

		return nil,
			Metadata{},
			err
	}

	return state,
		metadata,
		nil
}
