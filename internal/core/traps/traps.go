package traps

import (
	"github.com/YarikRevich/FVm/internal/core/memory"
	"github.com/YarikRevich/FVm/internal/core/registers"
)

const (
	TRAP_GETC = 0x20
	TRAP_OUT = 0x21
	TRAP_PUTS = 0x22
	TRAP_IN = 0x23
	TRAP_PUTSB = 0x24
	TRAP_HALT = 0x25
)

func Putc(){
	ch := append(memory.Memory, registers.RegStore[registers.R_R0])
	
};