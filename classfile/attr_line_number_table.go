package classfile

type LineNumberTableAttribute struct {
	lineNumberTable [] *LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumerTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumerTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableEntry{
			startPc: reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}