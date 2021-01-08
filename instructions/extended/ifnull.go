package extended

import (
	"gvm/instructions/base"
	"gvm/rtda"
)

type IFNULL struct {
	base.BranchInstruction
}

type IFNONNULL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}
