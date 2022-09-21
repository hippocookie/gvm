package classfile

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantInfo) ConstantInfo {
	switch tag {
		case CONSTANT_Integer: return &ConstantIntegerInfo{}
		case CONSTANT_Float: return &ConstantFloatInfo{}
		case CONSTANT_Long: return &ConstantLongInfo{}
		case CONSTANT_Double: return &ConstantDoubleInfo{}
		case CONSTANT_Utf8: return &ConstantUtf8Info{}
		case CONSTANT_String: return &ConstantStringInfo{cp: cp}
		case CONSTANT_Class: return &ConstantClassInfo{cp: cp}
		case CONSTANT_Fieldref: 
			return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
		case CONSTANT_Methodref: 
			return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
		case CONSTANT_InterfaceMethodref: 
			return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
		case CONSTANT_NameAndType: return &ConstantNameAndTypeInfo{}
		case CONSTANT_MethodType: return &ConstantMethodTypeInfo{}
		case CONSTANT_MethodHandle: return &ConstantMethodHandleInfo{}
		case CONSTANT_InvokeDynamic: return &ConstantInvokeDynamicInfo{}
		default: panic("java.lang.ClassFormatError: constant pool tag!")
	}
}

