package operators

import (
	"log"

	"github.com/YarikRevich/FVm/internal/core/memory"
	"github.com/YarikRevich/FVm/internal/core/registers"
	"github.com/YarikRevich/FVm/internal/core/signs"
)

func ADD(instr uint16) {
	dr := (instr >> 9) & 0x7
	r1 := (instr >> 6) & 0x7
	if (signs.LShift(instr, 15) & 0x1) != 0 {
		registers.RegStore[dr] = registers.RegStore[r1] + signs.GetExtendedSign(instr&0x1F, 5)
	} else {
		registers.RegStore[dr] = registers.RegStore[r1] + (instr & 0x7)
	}

	registers.UpdateConditionReg(dr)
}

func RTI() {
	log.Println("RTI op code is not available in LC-3")
}

func RES() {
	log.Println("RTI op code is not available in LC-3")
}

func AND(instr uint16) {
	dr := (instr >> 9) & 0x7
	r1 := (instr >> 6) & 0x7
	if (signs.LShift(instr, 15) & 0x1) != 0 {
		registers.RegStore[dr] = registers.RegStore[r1] & signs.GetExtendedSign(instr&0x1F, 5)
	} else {
		registers.RegStore[dr] = registers.RegStore[r1] & (instr & 0x7)
	}

	registers.UpdateConditionReg(dr)
}

func NOT(instr uint16) {
	dr := (instr >> 9) & 0x7
	r1 := (instr >> 6) & 0x7
	registers.RegStore[dr] = ^registers.RegStore[r1]

	registers.UpdateConditionReg(dr)
}

func BRANCH(instr uint16) {
	pc_offset := signs.GetExtendedSign((instr)&0x1ff, 9)
	conf_flag := (instr >> 9) & 0x7
	if conf_flag&registers.RegStore[registers.R_COND] != 0 {
		registers.RegStore[registers.R_PC] += pc_offset
	}
}

func JUMP(instr uint16) {
	r1 := (instr >> 6) & 0x7
	registers.RegStore[registers.R_PC] = registers.RegStore[r1]
}

func JUMB_REGISTER(instr uint16) {
	r1 := (instr >> 6) & 0x7
	long_pc_offset := signs.GetExtendedSign(instr&0x7ff, 11)
	long_flag := (instr >> 11) & 0x7
	registers.RegStore[registers.R_R7] = registers.RegStore[registers.R_PC]
	if long_flag != 0 {
		registers.RegStore[registers.R_PC] += long_pc_offset
	} else {
		registers.RegStore[registers.R_PC] = r1
	}
}

func LOAD(instr uint16) {
	dr := (instr >> 9) & 0x7
	pc_offset := signs.GetExtendedSign(instr&0x1ff, 9)
	registers.RegStore[dr] = memory.MemRead(registers.RegStore[registers.R_PC] + pc_offset)
	registers.UpdateConditionReg(dr)
}

func LOAD_INDIRECT(instr uint16) {
	dr := (instr >> 9) & 0x7
	pc_offset := signs.GetExtendedSign(instr&0x1FF, 9)
	registers.RegStore[dr] = memory.MemRead(memory.MemRead(registers.RegStore[registers.R_PC] + pc_offset))

	registers.UpdateConditionReg(dr)
}

func LOAD_REGISTERS(instr uint16) {
	dr := (instr >> 9) & 0x7
	r1 := (instr >> 6) & 0x7
	offset := signs.GetExtendedSign(instr&0x3F, 6)
	registers.RegStore[dr] = memory.MemRead(registers.RegStore[r1] + offset)
	registers.UpdateConditionReg(dr)
}

func EFFECTIV_LOAD_ADDRESS(instr uint16) {
	dr := (instr >> 9) & 0x7
	pc_offset := signs.GetExtendedSign(instr&0x1ff, 9)
	registers.RegStore[dr] = registers.RegStore[registers.R_PC] + pc_offset
	registers.UpdateConditionReg(dr)
}

func STORE(instr uint16) {
	dr := (instr >> 9) & 0x7
	pc_offset := signs.GetExtendedSign(instr&0x1ff, 9)
	memory.MemWrite(registers.RegStore[registers.R_PC]+pc_offset, registers.RegStore[dr])
}

func STORE_INDIRECT(instr uint16) {
	dr := (instr >> 9) & 0x7
	pc_offset := signs.GetExtendedSign(instr&0x1ff, 9)
	memory.MemWrite(memory.MemRead(registers.RegStore[registers.R_PC]+pc_offset), registers.RegStore[dr])
}

func STORE_REGISTER(instr uint16) {
	dr := (instr >> 9) & 0x7
	r1 := (instr >> 6) & 0x7
	offset := signs.GetExtendedSign(instr&0x3F, 6)
	memory.MemWrite(registers.RegStore[r1]+offset, registers.RegStore[dr])
}
