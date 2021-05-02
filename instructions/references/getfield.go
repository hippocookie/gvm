package references

type GET_FIELD struct {
	base.Index16Instruction
}

func (self *GET_FIELD) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nul {
		panic("java.lang.NullPointerException")
	}
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := ref.Fields()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slot.GetFloat(slotId))
	case 'J':
		stack.PushLong(slot.GetLong(slotId))
	case 'D':
		stack.PushDouble(slot.GetDouble(slotId))
	case 'L', '[': stack.GetRef(slotId)
	}
}
