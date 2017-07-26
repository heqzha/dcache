package model

import (
	"bytes"
	"encoding/gob"
)

//TODO
type Container struct {
	data []byte
}

func (c *Container) Set(val interface{}) error {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	if err := enc.Encode(val); err != nil {
		return err
	}
	c.data = buffer.Bytes()
	return nil
}

func (c *Container) Get(val interface{}) error {
	buffer := bytes.NewBuffer(c.data)
	dec := gob.NewDecoder(buffer)
	if err := dec.Decode(val); err != nil {
		return err
	}
	return nil
}
