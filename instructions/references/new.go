package references

type NEW struct {
	base.Index16Instruction
}

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func newObject(class *Class) *Object {
	return &Object {
		class: class,
		fields: newSlots(class.instanceSlotCount),
	}
}