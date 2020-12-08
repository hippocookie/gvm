package classfile

type ConstantMethodTypeInfo struct {
	cp              ConstantPool
	classIndex      uint16
	methodTypeIndex uint16
}

func (self *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.methodTypeIndex = reader.readUint16()
}
