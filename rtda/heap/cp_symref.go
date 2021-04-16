package heap

import "gvm/classfile"

// symbolic reference
type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func newSymRef(cp *ConstantPool,
	classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &SymRef{}
	ref.cp = cp
	ref.className = classInfo.Name()
	return ref
}

func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}

func (self *Class) isAccessableTo(other *Class) bool {
	return self.IsPublic() || self.getPackageName() == other.getPackageName()
}

func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

