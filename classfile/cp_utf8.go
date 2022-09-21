package classfile

import (
	"fmt"
	"unicode/utf16"
)

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUft8Info) readInfo(reader *ClassReader) {
	length := reader.readUint16()
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}

