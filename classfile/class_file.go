package classfile

import "fmt"

type ClassFile struct {
	// magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags uint16
	thisClass uint16
	superClass uint16
	interfaces []uint16
	fields []*MemberInfo
	methods []*MemberInfo
	attributes []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = *ClassFile{}
	cf.read(cr)
}

func (self *ClassFile) read(reader *ClassReader) {

}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {

}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {

}

func (self *ClassFile) MinorVersion() uint16 {

}

func (self *ClassFile) MajorVersion() uint16 {

}

func (self *ClassFile) ConstantPool() ConstantPool {

}

func (self *ClassFile) AccessFlags() uint16 {

}

func (self *ClassFile) Fields() []*MemberInfo {

}

func (self *ClassFile) Methods() []*MemberInfo {

}

func (self *ClassFile) ClassName() string {

}

func (self *ClassFile) SuperClassName() string {

}

func (self *ClassFile) InterfaceNames() []string {

}