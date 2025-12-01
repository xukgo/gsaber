package cio

import "github.com/xukgo/gsaber/utils/stringUtil"

type RoBytes struct {
	buffer  []byte
	strVal  string
	isOwner bool
}

func NewOwnerRoBytes(buffer []byte) *RoBytes {
	s := &RoBytes{
		buffer:  buffer,
		strVal:  stringUtil.NoCopyBytes2String(buffer),
		isOwner: true,
	}
	return s
}

func NewBorrowRoBytes(buffer []byte) *RoBytes {
	s := &RoBytes{
		buffer:  buffer,
		strVal:  stringUtil.NoCopyBytes2String(buffer),
		isOwner: false,
	}
	return s
}

func (c *RoBytes) Reset(buffer []byte, isOwner bool) {
	c.buffer = buffer
	c.strVal = stringUtil.NoCopyBytes2String(buffer)
	c.isOwner = isOwner
}

func (c *RoBytes) IsOwner() bool {
	return c.isOwner
}

func (c *RoBytes) Bytes() []byte {
	return c.buffer
}

func (c *RoBytes) String() string {
	return c.strVal
}
