package procedure

import (
	"log"

	"github.com/YarikRevich/FVm/internal/core/memory"
	"github.com/YarikRevich/FVm/internal/core/operators"
	"github.com/YarikRevich/FVm/internal/core/registers"
	"github.com/YarikRevich/FVm/internal/core/traps"
)

type Procedure struct {
	instr uint16
}

func (p *Procedure) UpdateInstr() {
	registers.RegStore[registers.R_PC]++
	p.instr = memory.MemRead(registers.RegStore[registers.R_PC])
}

func (p Procedure) GetOpCode() uint16 {
	return p.instr >> 12
}

func (p Procedure) ProcessOpCode() {
	op := p.GetOpCode()
	switch op {
	case registers.OP_ADD:
		operators.ADD(p.instr)
	case registers.OP_AND:
		operators.AND(p.instr)
	case registers.OP_NOT:
		operators.NOT(p.instr)
	case registers.OP_BR:
		operators.BRANCH(p.instr)
	case registers.OP_JMP:
		operators.JUMP(p.instr)
	case registers.OP_JSR:
		operators.JUMB_REGISTER(p.instr)
	case registers.OP_LD:
		operators.LOAD(p.instr)
	case registers.OP_LDI:
		operators.LOAD_INDIRECT(p.instr)
	case registers.OP_LDR:
		operators.LOAD_REGISTERS(p.instr)
	case registers.OP_LEA:
		operators.EFFECTIV_LOAD_ADDRESS(p.instr)
	case registers.OP_ST:
		operators.STORE(p.instr)
	case registers.OP_STI:
		operators.STORE_INDIRECT(p.instr)
	case registers.OP_STR:
		operators.STORE_REGISTER(p.instr)
	case registers.OP_TRAP:

	case registers.OP_RES:
		operators.RES()
	case registers.OP_RTI:
		operators.RTI()
	default:
		log.Fatalln("Not correct op code")
	}
}

func (p Procedure) ProcessTrapCode() {
	switch p.instr & 0xFF {
	case traps.TRAP_GETC:
		break
	case traps.TRAP_OUT:
		break
	case traps.TRAP_PUTS:
		break
	case traps.TRAP_IN:
		break
	case traps.TRAP_PUTSB:
		break
	case traps.TRAP_HALT:
	}
}

func (p Procedure) Run() {
	registers.RegStore[registers.R_PC] = 0x3000

	running := 1
	for running != 0 {
		p.UpdateInstr()
		p.ProcessOpCode()
		p.ProcessTrapCode()
	}
}

func New() *Procedure {
	return new(Procedure)
}
