package classfile

type AttributeInfo struct {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {

}

func readAttribute(reader *ClassReader, cp ConstantPoll) AttributeInfo {

}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	
}