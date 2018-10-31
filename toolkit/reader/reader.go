package reader

import (
	"bytes"
	"io"
	"io/ioutil"
)

type multipleReader interface {
	Reader() io.ReadCloser
}

type myMultipleReader struct {
	data []byte
}

func NewMultipleReader(reader io.Reader) (multipleReader, error) {
	var data []byte
	var err error
	if reader != nil {
		data, err = ioutil.ReadAll(reader)
	} else {
		data = []byte{}
	}
	return &myMultipleReader{
		data: data,
	}, err
}

func (mr *myMultipleReader) Reader() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader(mr.data))
}
