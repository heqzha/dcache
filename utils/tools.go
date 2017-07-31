package utils

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func DCacheEncode(val interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	if err := enc.Encode(val); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func DCacheDecode(val interface{}, data []byte) error {
	buffer := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buffer)
	if err := dec.Decode(val); err != nil {
		fmt.Println("dec.Decode", err)
		return err
	}
	return nil
}
