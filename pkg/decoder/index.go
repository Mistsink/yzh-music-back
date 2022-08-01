package decoder

import (
	"compress/flate"
	"compress/gzip"
	"errors"
	"io"
)

func NewDecoder(encodingType string, rd io.Reader) (reader io.Reader, err error) {
	switch encodingType {
	case "gzip":
		reader, err = gzip.NewReader(rd)
		return reader, err
	case "deflate":
		reader = flate.NewReader(rd)
		return
	case "br":
		return nil, errors.New("not support br decompress")
	case "identity":
		return rd, nil
	default:
		return nil, errors.New("not found appropriate decoder")
	}
}
