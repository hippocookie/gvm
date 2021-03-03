package heap

import "math"

type Slot struct {
	num int32
	ref *Object
}
type Slots []Slot

func newLocalVars(maxLocals uint) Slot {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (self Slot) SetInt(index uint, val int32) {
	self[index].num = val
}

func (self Slot) GetInt(index uint) int32 {
	return self[index].num
}

func (self Slot) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}

func (self Slot) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

func (self Slot) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}

func (self Slot) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}

func (self Slot) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

func (self Slot) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

func (self Slot) SetRef(index uint, ref *Object) {
	self[index].ref = ref
}

func (self Slot) GetRef(index uint) *Object {
	return self[index].ref
}
