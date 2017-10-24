package entity

import (
	"encoding/json"
	"fmt"
	"os"

	errors "github.com/go-errors/errors"
)

type storage struct {
	path string
}

func (s *storage) load(ptr interface{}) {
	file, err := os.Open(s.path)
	defer file.Close()
	if os.IsNotExist(err) {
		logger.Printf("[storage] storage file '%s' does not exist, so use an empty one\n", s.path)
		return
	}
	if err != nil {
		panic(errors.Wrap(err, 0))
	}

	decodeErr := json.NewDecoder(file).Decode(ptr)
	if decodeErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to decode storage file '%s'. You might want to remove it.", s.path)
		panic(errors.Wrap(fmt.Errorf("[storage] failed to decode storage file '%s': %v", s.path, decodeErr), 0))
	}
	logger.Printf("[storage] storage file '%s' decoded successfully\n", s.path)
}

func (s *storage) dump(ptr interface{}) {
	file, err := os.OpenFile(s.path, os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		panic(errors.Wrap(err, 0))
	}

	encodeErr := json.NewEncoder(file).Encode(ptr)
	if encodeErr != nil {
		panic(errors.Wrap(encodeErr, 0))
	}
	logger.Printf("[storage] storage file '%s' written successfully\n", s.path)
}
