package constants

import (
	"gvm/instructions/base"
	"gvm/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {

}
