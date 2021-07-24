package registers

import (
	"github.com/YarikRevich/FVm/internal/core/signs"
)

const (
	R_R0 = iota
	R_R1
	R_R2
	R_R3
	R_R4
	R_R5
	R_R6
	R_R7
	R_PC
	R_COND
	R_COUNT
)

const (
	OP_BR = iota
	OP_ADD
	OP_LD
	OP_ST
	OP_JSR
	OP_AND
	OP_LDR
	OP_STR
	OP_RTI
	OP_NOT
	OP_LDI
	OP_STI
	OP_JMP
	OP_RES
	OP_LEA
	OP_TRAP
)

const (
	FL_POS = 1 << 0
	FL_ZRO = 1 << 1
	FL_NEG = 1 << 2
)

var (
	RegStore [R_COUNT]uint16
)


func UpdateConditionReg(r uint16){
	switch {
		case RegStore[r] == 0:
			RegStore[R_COND] = FL_ZRO
		case (signs.LShift(RegStore[r], 15)) != 0:
			RegStore[R_COND] = FL_NEG
		default:
			RegStore[R_COND] = FL_POS
	}
}

