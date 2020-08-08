package mp4

import (
	"io"
	"io/ioutil"
)

func UnknownDecode(r io.Reader) (Box, error) {

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	u := newUnknownBox()
	u.notDecoded = data

	return u, nil
}

type UnknownBox struct {
	boxType string
	boxSize int

	notDecoded []byte
}

func newUnknownBox() *UnknownBox {
	b := UnknownBox{}
	return &b
}

func (u *UnknownBox) setHeader(h BoxHeader) {
	u.boxType = h.Type
	u.boxSize = int(h.Size)
}

func (u *UnknownBox) Type() string {
	return u.boxType
}

func (u *UnknownBox) Size() int {
	return u.boxSize
}

func (u *UnknownBox) Encode(w io.Writer) error {
	err := EncodeHeader(u, w)
	if err != nil {
		return err
	}
	buf := makebuf(u)

	copy(buf, u.notDecoded)
	_, err = w.Write(buf)
	return err
}
