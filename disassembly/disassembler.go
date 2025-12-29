package disassembly

import (
	"bytes"
	"fmt"
)

func Disassemble80800p(ins *bytes.Buffer, pc int) int {
	insSet := ins.Bytes()
	instruction := insSet[pc]
	switch instruction {
	// 0x00 - 0x0f
	case 0x00:
		fmt.Printf("%07x NOP\n", pc)
	case 0x01:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("LXI instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x LXI B,#$%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2

	case 0x02:
		fmt.Printf("%07x STAX B\n", pc)
	case 0x03:
		fmt.Printf("%07x INX B\n", pc)
	case 0x04:
		fmt.Printf("%07x INR B\n", pc)
	case 0x05:
		fmt.Printf("%07x DCR B\n", pc)
	case 0x06:
		if len(insSet) <= pc+1 {
			panic("MVI instruction has no destination")
		}
		destination := insSet[pc+1]
		fmt.Printf("%07x MVI B,#$%02x\n", pc, destination)
		pc = pc + 1
	case 0x07:
		fmt.Printf("%07x RLC\n", pc)
	case 0x08:
		fmt.Printf("%07x NOP\n", pc)
	case 0x09:
		fmt.Printf("%07x DAD B\n", pc)
	case 0x0a:
		fmt.Printf("%07x LDAX B\n", pc)
	case 0x0b:
		fmt.Printf("%07x DCX B\n", pc)
	case 0x0c:
		fmt.Printf("%07x INR C\n", pc)
	case 0x0d:
		fmt.Printf("%07x DCR C\n", pc)
	case 0x0e:
		if len(insSet) <= pc+1 {
			panic("MVI instruction has no destination")
		}
		destination := insSet[pc+1]
		fmt.Printf("%07x MVI C,#$%02x\n", pc, destination)
		pc = pc + 1
	case 0x0f:
		fmt.Printf("%07x RCC\n", pc)

	// 0x10 - 0x1f
	case 0x10:
		fmt.Printf("%07x NOP\n", pc)
	case 0x11:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("LXI instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x LXI D,#$%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0x12:
		fmt.Printf("%07x STAX C\n", pc)
	case 0x13:
		fmt.Printf("%07x INX C\n", pc)
	case 0x14:
		fmt.Printf("%07x INR C\n", pc)
	case 0x15:
		fmt.Printf("%07x DCR C\n", pc)
	case 0x16:
		if len(insSet) <= pc+1 {
			panic("MVI instruction has no destination")
		}
		destination := insSet[pc+1]
		fmt.Printf("%07x MVI D,#$%02x\n", pc, destination)
		pc = pc + 1
	case 0x17:
		fmt.Printf("%07x RAL\n", pc)
	case 0x18:
		fmt.Printf("%07x NOP\n", pc)
	case 0x19:
		fmt.Printf("%07x DAD D\n", pc)
	case 0x1a:
		fmt.Printf("%07x LDAX D\n", pc)
	case 0x1b:
		fmt.Printf("%07x DCX D\n", pc)
	case 0x1c:
		fmt.Printf("%07x INR E\n", pc)
	case 0x1d:
		fmt.Printf("%07x DCR E\n", pc)
	case 0x1e:
		if len(insSet) <= pc+1 {
			panic("MVI instruction has no destination")
		}
		destination := insSet[pc+1]
		fmt.Printf("%07x MVI E,#$%02x\n", pc, destination)
		pc = pc + 1
	case 0x1f:
		fmt.Printf("%07x RAR\n", pc)

	//0x20 - 0x2f
	case 0x20:
		fmt.Printf("%07x NOP\n", pc)
	case 0x21:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("LXI instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x LXI H,#$%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0x22:
		fmt.Printf("%07x STAX D\n", pc)
	case 0x23:
		fmt.Printf("%07x INX D\n", pc)
	case 0x24:
		fmt.Printf("%07x INR D\n", pc)
	case 0x25:
		fmt.Printf("%07x DCR D\n", pc)
	case 0x26:
		if len(insSet) <= pc+1 {
			panic("MVI instruction has no destination")
		}
		destination := insSet[pc+1]
		fmt.Printf("%07x MVI H,#$%02x\n", pc, destination)
		pc = pc + 1
	case 0x27:
		fmt.Printf("%07x DAA\n", pc)
	case 0x28:
		fmt.Printf("%07x NOP\n", pc)
	case 0x29:
		fmt.Printf("%07x DAD H\n", pc)
	case 0x2a:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("JMLHLDP instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x LHLD $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0x2b:
		fmt.Printf("%07x DCX H\n", pc)
	case 0x2c:
		fmt.Printf("%07x INR L\n", pc)
	case 0x2d:
		fmt.Printf("%07x DCR L\n", pc)
	case 0x2e:
		if len(insSet) <= pc+1 {
			panic("MVI instruction has no destination")
		}
		destination := insSet[pc+1]
		fmt.Printf("%07x MVI L,#$%02x\n", pc, destination)
		pc = pc + 1
	case 0x2f:
		fmt.Printf("%07x CMA\n", pc)

	//0x30 - 0x3f
	case 0x30:
		fmt.Printf("%07x NOP\n", pc)
	case 0x31:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("LXI instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x LXI SP,#$%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0x32:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("STA instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x STA $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0x33:
		fmt.Printf("%07x INX SP\n", pc)
	case 0x34:
		fmt.Printf("%07x INR M\n", pc)
	case 0x35:
		fmt.Printf("%07x DCR M\n", pc)
	case 0x36:
		if len(insSet) <= pc+1 {
			panic("MVI instruction has no destination")
		}
		destination := insSet[pc+1]
		fmt.Printf("%07x MVI M,#$%02x\n", pc, destination)
		pc = pc + 1
	case 0x37:
		fmt.Printf("%07x STC\n", pc)
	case 0x38:
		fmt.Printf("%07x NOP\n", pc)
	case 0x39:
		fmt.Printf("%07x DAD SP\n", pc)
	case 0x3a:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("LDA instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x LDA $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0x3b:
		fmt.Printf("%07x DCX SP\n", pc)
	case 0x3c:
		fmt.Printf("%07x INR A\n", pc)
	case 0x3d:
		fmt.Printf("%07x DCR A\n", pc)
	case 0x3e:
		if len(insSet) <= pc+1 {
			panic("MVI instruction has no destination")
		}
		destination := insSet[pc+1]
		fmt.Printf("%07x MVI A,#$%02x\n", pc, destination)
		pc = pc + 1
	case 0x3f:
		fmt.Printf("%07x CMC\n", pc)

	// 0x40 - 0x4f
	case 0x40:
		fmt.Printf("%07x MOV B,B\n", pc)
	case 0x41:
		fmt.Printf("%07x MOV B,C\n", pc)
	case 0x42:
		fmt.Printf("%07x MOV B,D\n", pc)
	case 0x43:
		fmt.Printf("%07x MOV B,E\n", pc)
	case 0x44:
		fmt.Printf("%07x MOV B,H\n", pc)
	case 0x45:
		fmt.Printf("%07x MOV B,L\n", pc)
	case 0x46:
		fmt.Printf("%07x MOV B,M\n", pc)
	case 0x47:
		fmt.Printf("%07x MOV B,A\n", pc)
	case 0x48:
		fmt.Printf("%07x MOV C,B\n", pc)
	case 0x49:
		fmt.Printf("%07x MOV C,C\n", pc)
	case 0x4a:
		fmt.Printf("%07x MOV C,D\n", pc)
	case 0x4b:
		fmt.Printf("%07x MOV C,E\n", pc)
	case 0x4c:
		fmt.Printf("%07x MOV C,H\n", pc)
	case 0x4d:
		fmt.Printf("%07x MOV C,L\n", pc)
	case 0x4e:
		fmt.Printf("%07x MOV C,M\n", pc)
	case 0x4f:
		fmt.Printf("%07x MOV C,A\n", pc)

	// 0x50 - 0x5f
	case 0x50:
		fmt.Printf("%07x MOV D,B\n", pc)
	case 0x51:
		fmt.Printf("%07x MOV D,C\n", pc)
	case 0x52:
		fmt.Printf("%07x MOV D,D\n", pc)
	case 0x53:
		fmt.Printf("%07x MOV D,E\n", pc)
	case 0x54:
		fmt.Printf("%07x MOV D,H\n", pc)
	case 0x55:
		fmt.Printf("%07x MOV D,L\n", pc)
	case 0x56:
		fmt.Printf("%07x MOV D,M\n", pc)
	case 0x57:
		fmt.Printf("%07x MOV D,A\n", pc)
	case 0x58:
		fmt.Printf("%07x MOV E,B\n", pc)
	case 0x59:
		fmt.Printf("%07x MOV E,C\n", pc)
	case 0x5a:
		fmt.Printf("%07x MOV E,D\n", pc)
	case 0x5b:
		fmt.Printf("%07x MOV E,E\n", pc)
	case 0x5c:
		fmt.Printf("%07x MOV E,H\n", pc)
	case 0x5d:
		fmt.Printf("%07x MOV E,L\n", pc)
	case 0x5e:
		fmt.Printf("%07x MOV E,M\n", pc)
	case 0x5f:
		fmt.Printf("%07x MOV E,A\n", pc)

	// 0x60 - 0x6f
	case 0x60:
		fmt.Printf("%07x MOV H,B\n", pc)
	case 0x61:
		fmt.Printf("%07x MOV H,C\n", pc)
	case 0x62:
		fmt.Printf("%07x MOV H,D\n", pc)
	case 0x63:
		fmt.Printf("%07x MOV H,E\n", pc)
	case 0x64:
		fmt.Printf("%07x MOV H,H\n", pc)
	case 0x65:
		fmt.Printf("%07x MOV H,L\n", pc)
	case 0x66:
		fmt.Printf("%07x MOV H,M\n", pc)
	case 0x67:
		fmt.Printf("%07x MOV H,A\n", pc)
	case 0x68:
		fmt.Printf("%07x MOV L,B\n", pc)
	case 0x69:
		fmt.Printf("%07x MOV L,C\n", pc)
	case 0x6a:
		fmt.Printf("%07x MOV L,D\n", pc)
	case 0x6b:
		fmt.Printf("%07x MOV L,E\n", pc)
	case 0x6c:
		fmt.Printf("%07x MOV L,H\n", pc)
	case 0x6d:
		fmt.Printf("%07x MOV L,L\n", pc)
	case 0x6e:
		fmt.Printf("%07x MOV L,M\n", pc)
	case 0x6f:
		fmt.Printf("%07x MOV L,A\n", pc)

	// 0x70 - 0x7f
	case 0x70:
		fmt.Printf("%07x MOV M,B\n", pc)
	case 0x71:
		fmt.Printf("%07x MOV M,C\n", pc)
	case 0x72:
		fmt.Printf("%07x MOV M,D\n", pc)
	case 0x73:
		fmt.Printf("%07x MOV M,E\n", pc)
	case 0x74:
		fmt.Printf("%07x MOV M,H\n", pc)
	case 0x75:
		fmt.Printf("%07x MOV M,L\n", pc)
	case 0x76:
		fmt.Printf("%07x HLT\n", pc)
	case 0x77:
		fmt.Printf("%07x MOV M,A\n", pc)
	case 0x78:
		fmt.Printf("%07x MOV A,B\n", pc)
	case 0x79:
		fmt.Printf("%07x MOV A,C\n", pc)
	case 0x7a:
		fmt.Printf("%07x MOV A,D\n", pc)
	case 0x7b:
		fmt.Printf("%07x MOV A,E\n", pc)
	case 0x7c:
		fmt.Printf("%07x MOV A,H\n", pc)
	case 0x7d:
		fmt.Printf("%07x MOV A,L\n", pc)
	case 0x7e:
		fmt.Printf("%07x MOV A,M\n", pc)
	case 0x7f:
		fmt.Printf("%07x MOV A,A\n", pc)

	// 0x80 - 0x8f
	case 0x80:
		fmt.Printf("%07x ADD B\n", pc)
	case 0x81:
		fmt.Printf("%07x ADD C\n", pc)
	case 0x82:
		fmt.Printf("%07x ADD D\n", pc)
	case 0x83:
		fmt.Printf("%07x ADD E\n", pc)
	case 0x84:
		fmt.Printf("%07x ADD H\n", pc)
	case 0x85:
		fmt.Printf("%07x ADD L\n", pc)
	case 0x86:
		fmt.Printf("%07x ADD M\n", pc)
	case 0x87:
		fmt.Printf("%07x ADD A\n", pc)
	case 0x88:
		fmt.Printf("%07x ADC B\n", pc)
	case 0x89:
		fmt.Printf("%07x ADC C\n", pc)
	case 0x8a:
		fmt.Printf("%07x ADC D\n", pc)
	case 0x8b:
		fmt.Printf("%07x ADC E\n", pc)
	case 0x8c:
		fmt.Printf("%07x ADC H\n", pc)
	case 0x8d:
		fmt.Printf("%07x ADC L\n", pc)
	case 0x8e:
		fmt.Printf("%07x ADC M\n", pc)
	case 0x8f:
		fmt.Printf("%07x ADC A\n", pc)

	// 0x90 - 0x9f
	case 0x90:
		fmt.Printf("%07x SUB B\n", pc)
	case 0x91:
		fmt.Printf("%07x SUB C\n", pc)
	case 0x92:
		fmt.Printf("%07x SUB D\n", pc)
	case 0x93:
		fmt.Printf("%07x SUB E\n", pc)
	case 0x94:
		fmt.Printf("%07x SUB H\n", pc)
	case 0x95:
		fmt.Printf("%07x SUB L\n", pc)
	case 0x96:
		fmt.Printf("%07x SUB M\n", pc)
	case 0x97:
		fmt.Printf("%07x SUB A\n", pc)
	case 0x98:
		fmt.Printf("%07x SBB B\n", pc)
	case 0x99:
		fmt.Printf("%07x SBB C\n", pc)
	case 0x9a:
		fmt.Printf("%07x SBB D\n", pc)
	case 0x9b:
		fmt.Printf("%07x SBB E\n", pc)
	case 0x9c:
		fmt.Printf("%07x SBB H\n", pc)
	case 0x9d:
		fmt.Printf("%07x SBB L\n", pc)
	case 0x9e:
		fmt.Printf("%07x SBB M\n", pc)
	case 0x9f:
		fmt.Printf("%07x SBB A\n", pc)

	// 0xa0 - 0xaf
	case 0xa0:
		fmt.Printf("%07x ANA B\n", pc)
	case 0xa1:
		fmt.Printf("%07x ANA C\n", pc)
	case 0xa2:
		fmt.Printf("%07x ANA D\n", pc)
	case 0xa3:
		fmt.Printf("%07x ANA E\n", pc)
	case 0xa4:
		fmt.Printf("%07x ANA H\n", pc)
	case 0xa5:
		fmt.Printf("%07x ANA L\n", pc)
	case 0xa6:
		fmt.Printf("%07x ANA M\n", pc)
	case 0xa7:
		fmt.Printf("%07x ANA A\n", pc)
	case 0xa8:
		fmt.Printf("%07x XRA B\n", pc)
	case 0xa9:
		fmt.Printf("%07x XRA C\n", pc)
	case 0xaa:
		fmt.Printf("%07x XRA D\n", pc)
	case 0xab:
		fmt.Printf("%07x XRA E\n", pc)
	case 0xac:
		fmt.Printf("%07x XRA H\n", pc)
	case 0xad:
		fmt.Printf("%07x XRA L\n", pc)
	case 0xae:
		fmt.Printf("%07x XRA M\n", pc)
	case 0xaf:
		fmt.Printf("%07x XRA A\n", pc)

	// 0xb0 - 0xbf
	case 0xb0:
		fmt.Printf("%07x ORA B\n", pc)
	case 0xb1:
		fmt.Printf("%07x ORA C\n", pc)
	case 0xb2:
		fmt.Printf("%07x ORA D\n", pc)
	case 0xb3:
		fmt.Printf("%07x ORA E\n", pc)
	case 0xb4:
		fmt.Printf("%07x ORA H\n", pc)
	case 0xb5:
		fmt.Printf("%07x ORA L\n", pc)
	case 0xb6:
		fmt.Printf("%07x ORA M\n", pc)
	case 0xb7:
		fmt.Printf("%07x ORA A\n", pc)
	case 0xb8:
		fmt.Printf("%07x CMP B\n", pc)
	case 0xb9:
		fmt.Printf("%07x CMP C\n", pc)
	case 0xba:
		fmt.Printf("%07x CMP D\n", pc)
	case 0xbb:
		fmt.Printf("%07x CMP E\n", pc)
	case 0xbc:
		fmt.Printf("%07x CMP H\n", pc)
	case 0xbd:
		fmt.Printf("%07x CMP L\n", pc)
	case 0xbe:
		fmt.Printf("%07x CMP M\n", pc)
	case 0xbf:
		fmt.Printf("%07x CMP A\n", pc)

	// 0xc0 - 0xcf
	case 0xc0:
		fmt.Printf("%07x RNZ\n", pc)
	case 0xc1:
		fmt.Printf("%07x POP B\n", pc)
	case 0xc2:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("JNZ instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x JNZ $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xc3:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("JMP instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x JMP $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xc4:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("CNZ instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x CNZ $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xc5:
		fmt.Printf("%07x PUSH B\n", pc)
	case 0xc6:
		if len(insSet) <= pc+1 {
			panic("ADI instruction has no destination")
		}
		destination := insSet[pc+1]
		fmt.Printf("%07x ADI #$%02x\n", pc, destination)
		pc = pc + 1
	case 0xc7:
		fmt.Printf("%07x RST 0\n", pc)
	case 0xc8:
		fmt.Printf("%07x RZ\n", pc)
	case 0xc9:
		fmt.Printf("%07x RET\n", pc)
	case 0xca:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("JZ instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x JZ $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xcb:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("JMP instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x JMP $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xcc:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("CZ instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x CZ $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xcd:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("CALL instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x CALL $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xce:
		if len(insSet) <= pc+1 {
			panic("ACI instruction has no destination")
		}
		destination := insSet[pc+1]
		fmt.Printf("%07x ACI #$%02x\n", pc, destination)
		pc = pc + 1
	case 0xcf:
		fmt.Printf("%07x RST 1\n", pc)

	// 0xd0 - 0xdf
	case 0xd0:
		fmt.Printf("%07x RNC\n", pc)
	case 0xd1:
		fmt.Printf("%07x POP D\n", pc)
	case 0xd2:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("JNC instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x JNC $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xd3:
		if len(insSet) <= pc+1 {
			panic("OUT instruction has no destination")
		}
		destination := insSet[pc+1]
		fmt.Printf("%07x OUT #$%02x\n", pc, destination)
		pc = pc + 1
	case 0xd4:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("CNC instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x CNC $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xd5:
		fmt.Printf("%07x PUSH D\n", pc)
	case 0xd6:
		if len(insSet) <= pc+1 {
			panic("SUI instruction has no destination")
		}
		destination := insSet[pc+1]
		fmt.Printf("%07x SUI #$%02x\n", pc, destination)
		pc = pc + 1
	case 0xd7:
		fmt.Printf("%07x RST 2\n", pc)
	case 0xd8:
		fmt.Printf("%07x RC\n", pc)
	case 0xd9:
		fmt.Printf("%07x RET\n", pc)
	case 0xda:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("JC instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x JC $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xdb:
		if len(insSet) <= pc+1 {
			panic("IN instruction has no destination")
		}
		destination := insSet[pc+1]
		fmt.Printf("%07x IN #$%02x\n", pc, destination)
		pc = pc + 1
	case 0xdc:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("CC instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x CC $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xdd:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("CALL instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x CALL $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xde:
		if len(insSet) <= pc+1 {
			panic("SBI instruction has no destination")
		}
		destination := insSet[pc+1]
		fmt.Printf("%07x SBI #$%02x\n", pc, destination)
		pc = pc + 1
	case 0xdf:
		fmt.Printf("%07x RST 3\n", pc)

	// 0xe0 - 0xef
	case 0xe0:
		fmt.Printf("%07x RPO\n", pc)
	case 0xe1:
		fmt.Printf("%07x POP H\n", pc)
	case 0xe2:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("JPO instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x JPO $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xe3:
		fmt.Printf("%07x XTHL\n", pc)
	case 0xe4:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("CPO instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x CPO $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xe5:
		fmt.Printf("%07x PUSH H\n", pc)
	case 0xe6:
		if len(insSet) <= pc+1 {
			panic("ANI instruction has no destination")
		}
		destination := insSet[pc+1]
		fmt.Printf("%07x ANI #$%02x\n", pc, destination)
		pc = pc + 1
	case 0xe7:
		fmt.Printf("%07x RST 4\n", pc)
	case 0xe8:
		fmt.Printf("%07x RPE\n", pc)
	case 0xe9:
		fmt.Printf("%07x PCHL\n", pc)
	case 0xea:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("JPE instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x JPE $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xeb:
		fmt.Printf("%07x XCHG\n", pc)
	case 0xec:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("CPE instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x CPE $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xed:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("CALL instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x CALL $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xee:
		if len(insSet) <= pc+1 {
			panic("XRI instruction has no destination")
		}
		destination := insSet[pc+1]
		fmt.Printf("%07x XRI #$%02x\n", pc, destination)
		pc = pc + 1
	case 0xef:
		fmt.Printf("%07x RST 5\n", pc)

	// 0xf0 - 0xff
	case 0xf0:
		fmt.Printf("%07x RP\n", pc)
	case 0xf1:
		fmt.Printf("%07x POP PSW\n", pc)
	case 0xf2:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("JP instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x JP $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xf3:
		fmt.Printf("%07x DI\n", pc)
	case 0xf4:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("CP instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x CP $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xf5:
		fmt.Printf("%07x PUSH PSW\n", pc)
	case 0xf6:
		if len(insSet) <= pc+1 {
			panic("ORI instruction has no destination")
		}
		destination := insSet[pc+1]
		fmt.Printf("%07x ORI #$%02x\n", pc, destination)
		pc = pc + 1
	case 0xf7:
		fmt.Printf("%07x RST 6\n", pc)
	case 0xf8:
		fmt.Printf("%07x RM\n", pc)
	case 0xf9:
		fmt.Printf("%07x SPHL\n", pc)
	case 0xfa:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("JM instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x JM $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xfb:
		fmt.Printf("%07x EI\n", pc)
	case 0xfc:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("CM instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x CM $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xfd:
		if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
			panic("CALL instruction has no destination")
		}
		destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		fmt.Printf("%07x CALL $%02x%02x\n", pc, destination[1], destination[0])
		pc = pc + 2
	case 0xfe:
		if len(insSet) <= pc+1 {
			panic("CPI instruction has no destination")
		}
		destination := insSet[pc+1]
		fmt.Printf("%07x CPI #$%02x\n", pc, destination)
		pc = pc + 1
	case 0xff:
		fmt.Printf("%07x RST 7\n", pc)
	}
	pc++

	return pc
}

/*
{case 0x00: printf("NOP"); break;
		case 0x01: printf("LXI    B,#$%02x%02x", code[2], code[1]); opbytes=3; break;
		case 0x02: printf("STAX   B"); break;
		case 0x03: printf("INX    B"); break;
		case 0x04: printf("INR    B"); break;
		case 0x05: printf("DCR    B"); break;
		case 0x06: printf("MVI    B,#$%02x", code[1]); opbytes=2; break;
		case 0x07: printf("RLC"); break;
		case 0x08: printf("NOP"); break;
		case 0x09: printf("DAD    B"); break;
		case 0x0a: printf("LDAX   B"); break;
		case 0x0b: printf("DCX    B"); break;
		case 0x0c: printf("INR    C"); break;
		case 0x0d: printf("DCR    C"); break;
		case 0x0e: printf("MVI    C,#$%02x", code[1]); opbytes = 2;	break;
		case 0x0f: printf("RRC"); break;}
*/

/*
	case 0x10: printf("NOP"); break;
	case 0x11: printf("LXI    D,#$%02x%02x", code[2], code[1]); opbytes=3; break;
	case 0x12: printf("STAX   D"); break;
	case 0x13: printf("INX    D"); break;
	case 0x14: printf("INR    D"); break;
	case 0x15: printf("DCR    D"); break;
	case 0x16: printf("MVI    D,#$%02x", code[1]); opbytes=2; break;
	case 0x17: printf("RAL"); break;
	case 0x18: printf("NOP"); break;
	case 0x19: printf("DAD    D"); break;
	case 0x1a: printf("LDAX   D"); break;
	case 0x1b: printf("DCX    D"); break;
	case 0x1c: printf("INR    E"); break;
	case 0x1d: printf("DCR    E"); break;
	case 0x1e: printf("MVI    E,#$%02x", code[1]); opbytes = 2; break;
	case 0x1f: printf("RAR"); break;

*/

/*

	case 0x20: printf("NOP"); break;
	case 0x21: printf("LXI    H,#$%02x%02x", code[2], code[1]); opbytes=3; break;
	case 0x22: printf("SHLD   $%02x%02x", code[2], code[1]); opbytes=3; break;
	case 0x23: printf("INX    H"); break;
	case 0x24: printf("INR    H"); break;
	case 0x25: printf("DCR    H"); break;
	case 0x26: printf("MVI    H,#$%02x", code[1]); opbytes=2; break;
	case 0x27: printf("DAA"); break;
	case 0x28: printf("NOP"); break;
	case 0x29: printf("DAD    H"); break;
	case 0x2a: printf("LHLD   $%02x%02x", code[2], code[1]); opbytes=3; break;
	case 0x2b: printf("DCX    H"); break;
	case 0x2c: printf("INR    L"); break;
	case 0x2d: printf("DCR    L"); break;
	case 0x2e: printf("MVI    L,#$%02x", code[1]); opbytes = 2; break;
	case 0x2f: printf("CMA"); break;

*/

/*

	case 0x30: printf("NOP"); break;
	case 0x31: printf("LXI    SP,#$%02x%02x", code[2], code[1]); opbytes=3; break;
	case 0x32: printf("STA    $%02x%02x", code[2], code[1]); opbytes=3; break;
	case 0x33: printf("INX    SP"); break;
	case 0x34: printf("INR    M"); break;
	case 0x35: printf("DCR    M"); break;
	case 0x36: printf("MVI    M,#$%02x", code[1]); opbytes=2; break;
	case 0x37: printf("STC"); break;
	case 0x38: printf("NOP"); break;
	case 0x39: printf("DAD    SP"); break;
	case 0x3a: printf("LDA    $%02x%02x", code[2], code[1]); opbytes=3; break;
	case 0x3b: printf("DCX    SP"); break;
	case 0x3c: printf("INR    A"); break;
	case 0x3d: printf("DCR    A"); break;
	case 0x3e: printf("MVI    A,#$%02x", code[1]); opbytes = 2; break;
	case 0x3f: printf("CMC"); break;
*/

/*
	case 0x40: printf("MOV    B,B"); break;
	case 0x41: printf("MOV    B,C"); break;
	case 0x42: printf("MOV    B,D"); break;
	case 0x43: printf("MOV    B,E"); break;
	case 0x44: printf("MOV    B,H"); break;
	case 0x45: printf("MOV    B,L"); break;
	case 0x46: printf("MOV    B,M"); break;
	case 0x47: printf("MOV    B,A"); break;
	case 0x48: printf("MOV    C,B"); break;
	case 0x49: printf("MOV    C,C"); break;
	case 0x4a: printf("MOV    C,D"); break;
	case 0x4b: printf("MOV    C,E"); break;
	case 0x4c: printf("MOV    C,H"); break;
	case 0x4d: printf("MOV    C,L"); break;
	case 0x4e: printf("MOV    C,M"); break;
	case 0x4f: printf("MOV    C,A"); break;
*/

/*

	case 0x50: printf("MOV    D,B"); break;
	case 0x51: printf("MOV    D,C"); break;
	case 0x52: printf("MOV    D,D"); break;
	case 0x53: printf("MOV    D.E"); break;
	case 0x54: printf("MOV    D,H"); break;
	case 0x55: printf("MOV    D,L"); break;
	case 0x56: printf("MOV    D,M"); break;
	case 0x57: printf("MOV    D,A"); break;
	case 0x58: printf("MOV    E,B"); break;
	case 0x59: printf("MOV    E,C"); break;
	case 0x5a: printf("MOV    E,D"); break;
	case 0x5b: printf("MOV    E,E"); break;
	case 0x5c: printf("MOV    E,H"); break;
	case 0x5d: printf("MOV    E,L"); break;
	case 0x5e: printf("MOV    E,M"); break;
	case 0x5f: printf("MOV    E,A"); break;
*/

/*

	case 0x60: printf("MOV    H,B"); break;
	case 0x61: printf("MOV    H,C"); break;
	case 0x62: printf("MOV    H,D"); break;
	case 0x63: printf("MOV    H.E"); break;
	case 0x64: printf("MOV    H,H"MOV); break;
	case 0x65: printf("MOV    H,L"); break;
	case 0x66: printf("MOV    H,M"); break;
	case 0x67: printf("MOV    H,A"); break;
	case 0x68: printf("MOV    L,B"); break;
	case 0x69: printf("MOV    L,C"); break;
	case 0x6a: printf("MOV    L,D"); break;
	case 0x6b: printf("MOV    L,E"); break;
	case 0x6c: printf("MOV    L,H"); break;
	case 0x6d: printf("MOV    L,L"); break;
	case 0x6e: printf("MOV    L,M"); break;
	case 0x6f: printf("MOV    L,A"); break;

*/

/*

	case 0x70: printf("MOV    M,B"); break;
	case 0x71: printf("MOV    M,C"); break;
	case 0x72: printf("MOV    M,D"); break;
	case 0x73: printf("MOV    M.E"); break;
	case 0x74: printf("MOV    M,H"); break;
	case 0x75: printf("MOV    M,L"); break;
	case 0x76: printf("HLT");        break;
	case 0x77: printf("MOV    M,A"); break;
	case 0x78: printf("MOV    A,B"); break;
	case 0x79: printf("MOV    A,C"); break;
	case 0x7a: printf("MOV    A,D"); break;
	case 0x7b: printf("MOV    A,E"); break;
	case 0x7c: printf("MOV    A,H"); break;
	case 0x7d: printf("MOV    A,L"); break;
	case 0x7e: printf("MOV    A,M"); break;
	case 0x7f: printf("MOV    A,A"); break;

*/

/*

	case 0x80: printf("ADD    B"); break;
	case 0x81: printf("ADD    C"); break;
	case 0x82: printf("ADD    D"); break;
	case 0x83: printf("ADD    E"); break;
	case 0x84: printf("ADD    H"); break;
	case 0x85: printf("ADD    L"); break;
	case 0x86: printf("ADD    M"); break;
	case 0x87: printf("ADD    A"); break;
	case 0x88: printf("ADC    B"); break;
	case 0x89: printf("ADC    C"); break;
	case 0x8a: printf("ADC    D"); break;
	case 0x8b: printf("ADC    E"); break;
	case 0x8c: printf("ADC    H"); break;
	case 0x8d: printf("ADC    L"); break;
	case 0x8e: printf("ADC    M"); break;
	case 0x8f: printf("ADC    A"); break;

*/

/*

	case 0x90: printf("SUB    B"); break;
	case 0x91: printf("SUB    C"); break;
	case 0x92: printf("SUB    D"); break;
	case 0x93: printf("SUB    E"); break;
	case 0x94: printf("SUB    H"); break;
	case 0x95: printf("SUB    L"); break;
	case 0x96: printf("SUB    M"); break;
	case 0x97: printf("SUB    A"); break;
	case 0x98: printf("SBB    B"); break;
	case 0x99: printf("SBB    C"); break;
	case 0x9a: printf("SBB    D"); break;
	case 0x9b: printf("SBB    E"); break;
	case 0x9c: printf("SBB    H"); break;
	case 0x9d: printf("SBB    L"); break;
	case 0x9e: printf("SBB    M"); break;
	case 0x9f: printf("SBB    A"); break;
*/

/*

	case 0xa0: printf("ANA    B"); break;
	case 0xa1: printf("ANA    C"); break;
	case 0xa2: printf("ANA    D"); break;
	case 0xa3: printf("ANA    E"); break;
	case 0xa4: printf("ANA    H"); break;
	case 0xa5: printf("ANA    L"); break;
	case 0xa6: printf("ANA    M"); break;
	case 0xa7: printf("ANA    A"); break;
	case 0xa8: printf("XRA    B"); break;
	case 0xa9: printf("XRA    C"); break;
	case 0xaa: printf("XRA    D"); break;
	case 0xab: printf("XRA    E"); break;
	case 0xac: printf("XRA    H"); break;
	case 0xad: printf("XRA    L"); break;
	case 0xae: printf("XRA    M"); break;
	case 0xaf: printf("XRA    A"); break;
*/
/*

	case 0xb0: printf("ORA    B"); break;
	case 0xb1: printf("ORA    C"); break;
	case 0xb2: printf("ORA    D"); break;
	case 0xb3: printf("ORA    E"); break;
	case 0xb4: printf("ORA    H"); break;
	case 0xb5: printf("ORA    L"); break;
	case 0xb6: printf("ORA    M"); break;
	case 0xb7: printf("ORA    A"); break;
	case 0xb8: printf("CMP    B"); break;
	case 0xb9: printf("CMP    C"); break;
	case 0xba: printf("CMP    D"); break;
	case 0xbb: printf("CMP    E"); break;
	case 0xbc: printf("CMP    H"); break;
	case 0xbd: printf("CMP    L"); break;
	case 0xbe: printf("CMP    M"); break;
	case 0xbf: printf("CMP    A"); break;
*/

/*

	case 0xc0: printf("RNZ"); break;
	case 0xc1: printf("POP    B"); break;
	case 0xc2: printf("JNZ    $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xc3: printf("JMP    $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xc4: printf("CNZ    $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xc5: printf("PUSH   B"); break;
	case 0xc6: printf("ADI    #$%02x",code[1]); opbytes = 2; break;
	case 0xc7: printf("RST    0"); break;
	case 0xc8: printf("RZ"); break;
	case 0xc9: printf("RET"); break;
	case 0xca: printf("JZ     $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xcb: printf("JMP    $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xcc: printf("CZ     $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xcd: printf("CALL   $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xce: printf("ACI    #$%02x",code[1]); opbytes = 2; break;
	case 0xcf: printf("RST    1"); break;
*/

/*
	case 0xd0: printf("RNC"); break;
	case 0xd1: printf("POP    D"); break;
	case 0xd2: printf("JNC    $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xd3: printf("OUT    #$%02x",code[1]); opbytes = 2; break;
	case 0xd4: printf("CNC    $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xd5: printf("PUSH   D"); break;
	case 0xd6: printf("SUI    #$%02x",code[1]); opbytes = 2; break;
	case 0xd7: printf("RST    2"); break;
	case 0xd8: printf("RC");  break;
	case 0xd9: printf("RET"); break;
	case 0xda: printf("JC     $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xdb: printf("IN     #$%02x",code[1]); opbytes = 2; break;
	case 0xdc: printf("CC     $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xdd: printf("CALL   $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xde: printf("SBI    #$%02x",code[1]); opbytes = 2; break;
	case 0xdf: printf("RST    3"); break;
*/

/*
	case 0xe0: printf("RPO"); break;
	case 0xe1: printf("POP    H"); break;
	case 0xe2: printf("JPO    $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xe3: printf("XTHL");break;
	case 0xe4: printf("CPO    $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xe5: printf("PUSH   H"); break;
	case 0xe6: printf("ANI    #$%02x",code[1]); opbytes = 2; break;
	case 0xe7: printf("RST    4"); break;
	case 0xe8: printf("RPE"); break;
	case 0xe9: printf("PCHL");break;
	case 0xea: printf("JPE    $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xeb: printf("XCHG"); break;
	case 0xec: printf("CPE     $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xed: printf("CALL   $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xee: printf("XRI    #$%02x",code[1]); opbytes = 2; break;
	case 0xef: printf("RST    5"); break;
*/

/*

	case 0xf0: printf("RP");  break;
	case 0xf1: printf("POP    PSW"); break;
	case 0xf2: printf("JP     $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xf3: printf("DI");  break;
	case 0xf4: printf("CP     $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xf5: printf("PUSH   PSW"); break;
	case 0xf6: printf("ORI    #$%02x",code[1]); opbytes = 2; break;
	case 0xf7: printf("RST    6"); break;
	case 0xf8: printf("RM");  break;
	case 0xf9: printf("SPHL");break;
	case 0xfa: printf("JM     $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xfb: printf("EI");  break;
	case 0xfc: printf("CM     $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xfd: printf("CALL   $%02x%02x",code[2],code[1]); opbytes = 3; break;
	case 0xfe: printf("CPI    #$%02x",code[1]); opbytes = 2; break;
	case 0xff: printf("RST    7"); break;

*/
