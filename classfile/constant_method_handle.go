package classfile

type ConstantMethodHandleInfo struct {
	cp                ConstantPool
	classIndex        uint16
	methodHandleIndex uint16
}

func (self *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.methodHandleIndex = reader.readUint16()
}
