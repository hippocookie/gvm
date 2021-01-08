package instructions

package instructions
  import "fmt"
  import "jvmgo/ch05/instructions/base"
  import . "jvmgo/ch05/instructions/comparisons"
  import . "jvmgo/ch05/instructions/constants"
  import . "jvmgo/ch05/instructions/control"
  import . "jvmgo/ch05/instructions/conversions"
  import . "jvmgo/ch05/instructions/extended"
  import . "jvmgo/ch05/instructions/loads"
  import . "jvmgo/ch05/instructions/math"
  import . "jvmgo/ch05/instructions/stack"
  import . "jvmgo/ch05/instructions/stores"
  func NewInstruction(opcode byte) base.Instruction {
  switch opcode {
    case 0x00: return &NOP{}
    case 0x01: return &ACONST_NULL{}
    ...
      default:
    panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
  }
}