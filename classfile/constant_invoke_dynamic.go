package classfile

type ConstantInvokeDynamicInfo struct {
	cp                 ConstantPool
	classIndex         uint16
	invokeDynamicIndex uint16
}

func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.invokeDynamicIndex = reader.readUint16()
}
