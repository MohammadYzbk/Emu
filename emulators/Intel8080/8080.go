package Intel8080

import (
	"errors"
)

var ErrUnimplementedInstruction error = errors.New("unimplemented instruction")

type ConditionCodes struct {
	Z   uint8
	S   uint8
	P   uint8
	CY  uint8
	AC  uint8
	Pad uint8
}

type State struct {
	A         uint8
	B         uint8
	C         uint8
	D         uint8
	E         uint8
	H         uint8
	L         uint8
	SP        uint16
	PC        uint16
	Memory    []uint8
	CC        ConditionCodes
	IntEnable uint8
}

func Emulate80800p(state *State) (uint16, error) {
	instruction := &state.Memory[state.PC]
	switch *instruction {
	// 0x00 - 0x0f
	case 0x00:
	case 0x01:
		state.C = state.Memory[state.PC+1]
		state.B = state.Memory[state.PC+2]
		state.PC += 2
		// case 0x02:
		// 	fmt.Printf("%07x STAX B\n", pc)
		// case 0x03:
		// 	fmt.Printf("%07x INX B\n", pc)
		// case 0x04:
		// 	fmt.Printf("%07x INR B\n", pc)
		// case 0x05:
		// 	fmt.Printf("%07x DCR B\n", pc)
		// case 0x06:
		// 	if len(insSet) <= pc+1 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := insSet[pc+1]
		// 	fmt.Printf("%07x MVI B,#$%02x\n", pc, destination)
		// 	pc = pc + 1
		// case 0x07:
		// 	fmt.Printf("%07x RLC\n", pc)
		// case 0x08:
		// 	fmt.Printf("%07x NOP\n", pc)
		// case 0x09:
		// 	fmt.Printf("%07x DAD B\n", pc)
		// case 0x0a:
		// 	fmt.Printf("%07x LDAX B\n", pc)
		// case 0x0b:
		// 	fmt.Printf("%07x DCX B\n", pc)
		// case 0x0c:
		// 	fmt.Printf("%07x INR C\n", pc)
		// case 0x0d:
		// 	fmt.Printf("%07x DCR C\n", pc)
		// case 0x0e:
		// 	if len(insSet) <= pc+1 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := insSet[pc+1]
		// 	fmt.Printf("%07x MVI C,#$%02x\n", pc, destination)
		// 	pc = pc + 1
		// case 0x0f:
		// 	fmt.Printf("%07x RCC\n", pc)

		// // 0x10 - 0x1f
		// case 0x10:
		// 	fmt.Printf("%07x NOP\n", pc)
		// case 0x11:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x LXI D,#$%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0x12:
		// 	fmt.Printf("%07x STAX C\n", pc)
		// case 0x13:
		// 	fmt.Printf("%07x INX C\n", pc)
		// case 0x14:
		// 	fmt.Printf("%07x INR C\n", pc)
		// case 0x15:
		// 	fmt.Printf("%07x DCR C\n", pc)
		// case 0x16:
		// 	if len(insSet) <= pc+1 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := insSet[pc+1]
		// 	fmt.Printf("%07x MVI D,#$%02x\n", pc, destination)
		// 	pc = pc + 1
		// case 0x17:
		// 	fmt.Printf("%07x RAL\n", pc)
		// case 0x18:
		// 	fmt.Printf("%07x NOP\n", pc)
		// case 0x19:
		// 	fmt.Printf("%07x DAD D\n", pc)
		// case 0x1a:
		// 	fmt.Printf("%07x LDAX D\n", pc)
		// case 0x1b:
		// 	fmt.Printf("%07x DCX D\n", pc)
		// case 0x1c:
		// 	fmt.Printf("%07x INR E\n", pc)
		// case 0x1d:
		// 	fmt.Printf("%07x DCR E\n", pc)
		// case 0x1e:
		// 	if len(insSet) <= pc+1 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := insSet[pc+1]
		// 	fmt.Printf("%07x MVI E,#$%02x\n", pc, destination)
		// 	pc = pc + 1
		// case 0x1f:
		// 	fmt.Printf("%07x RAR\n", pc)

		// //0x20 - 0x2f
		// case 0x20:
		// 	fmt.Printf("%07x NOP\n", pc)
		// case 0x21:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x LXI H,#$%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0x22:
		// 	fmt.Printf("%07x STAX D\n", pc)
		// case 0x23:
		// 	fmt.Printf("%07x INX D\n", pc)
		// case 0x24:
		// 	fmt.Printf("%07x INR D\n", pc)
		// case 0x25:
		// 	fmt.Printf("%07x DCR D\n", pc)
		// case 0x26:
		// 	if len(insSet) <= pc+1 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := insSet[pc+1]
		// 	fmt.Printf("%07x MVI H,#$%02x\n", pc, destination)
		// 	pc = pc + 1
		// case 0x27:
		// 	fmt.Printf("%07x DAA\n", pc)
		// case 0x28:
		// 	fmt.Printf("%07x NOP\n", pc)
		// case 0x29:
		// 	fmt.Printf("%07x DAD H\n", pc)
		// case 0x2a:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x LHLD $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0x2b:
		// 	fmt.Printf("%07x DCX H\n", pc)
		// case 0x2c:
		// 	fmt.Printf("%07x INR L\n", pc)
		// case 0x2d:
		// 	fmt.Printf("%07x DCR L\n", pc)
		// case 0x2e:
		// 	if len(insSet) <= pc+1 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := insSet[pc+1]
		// 	fmt.Printf("%07x MVI L,#$%02x\n", pc, destination)
		// 	pc = pc + 1
		// case 0x2f:
		// 	fmt.Printf("%07x CMA\n", pc)

		// //0x30 - 0x3f
		// case 0x30:
		// 	fmt.Printf("%07x NOP\n", pc)
		// case 0x31:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x LXI SP,#$%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0x32:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x STA $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0x33:
		// 	fmt.Printf("%07x INX SP\n", pc)
		// case 0x34:
		// 	fmt.Printf("%07x INR M\n", pc)
		// case 0x35:
		// 	fmt.Printf("%07x DCR M\n", pc)
		// case 0x36:
		// 	if len(insSet) <= pc+1 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := insSet[pc+1]
		// 	fmt.Printf("%07x MVI M,#$%02x\n", pc, destination)
		// 	pc = pc + 1
		// case 0x37:
		// 	fmt.Printf("%07x STC\n", pc)
		// case 0x38:
		// 	fmt.Printf("%07x NOP\n", pc)
		// case 0x39:
		// 	fmt.Printf("%07x DAD SP\n", pc)
		// case 0x3a:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x LDA $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0x3b:
		// 	fmt.Printf("%07x DCX SP\n", pc)
		// case 0x3c:
		// 	fmt.Printf("%07x INR A\n", pc)
		// case 0x3d:
		// 	fmt.Printf("%07x DCR A\n", pc)
		// case 0x3e:
		// 	if len(insSet) <= pc+1 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := insSet[pc+1]
		// 	fmt.Printf("%07x MVI A,#$%02x\n", pc, destination)
		// 	pc = pc + 1
		// case 0x3f:
		// 	fmt.Printf("%07x CMC\n", pc)

		// // 0x40 - 0x4f
		// case 0x40:
		// 	fmt.Printf("%07x MOV B,B\n", pc)
		// case 0x41:
		// 	fmt.Printf("%07x MOV B,C\n", pc)
		// case 0x42:
		// 	fmt.Printf("%07x MOV B,D\n", pc)
		// case 0x43:
		// 	fmt.Printf("%07x MOV B,E\n", pc)
		// case 0x44:
		// 	fmt.Printf("%07x MOV B,H\n", pc)
		// case 0x45:
		// 	fmt.Printf("%07x MOV B,L\n", pc)
		// case 0x46:
		// 	fmt.Printf("%07x MOV B,M\n", pc)
		// case 0x47:
		// 	fmt.Printf("%07x MOV B,A\n", pc)
		// case 0x48:
		// 	fmt.Printf("%07x MOV C,B\n", pc)
		// case 0x49:
		// 	fmt.Printf("%07x MOV C,C\n", pc)
		// case 0x4a:
		// 	fmt.Printf("%07x MOV C,D\n", pc)
		// case 0x4b:
		// 	fmt.Printf("%07x MOV C,E\n", pc)
		// case 0x4c:
		// 	fmt.Printf("%07x MOV C,H\n", pc)
		// case 0x4d:
		// 	fmt.Printf("%07x MOV C,L\n", pc)
		// case 0x4e:
		// 	fmt.Printf("%07x MOV C,M\n", pc)
		// case 0x4f:
		// 	fmt.Printf("%07x MOV C,A\n", pc)

		// // 0x50 - 0x5f
		// case 0x50:
		// 	fmt.Printf("%07x MOV D,B\n", pc)
		// case 0x51:
		// 	fmt.Printf("%07x MOV D,C\n", pc)
		// case 0x52:
		// 	fmt.Printf("%07x MOV D,D\n", pc)
		// case 0x53:
		// 	fmt.Printf("%07x MOV D,E\n", pc)
		// case 0x54:
		// 	fmt.Printf("%07x MOV D,H\n", pc)
		// case 0x55:
		// 	fmt.Printf("%07x MOV D,L\n", pc)
		// case 0x56:
		// 	fmt.Printf("%07x MOV D,M\n", pc)
		// case 0x57:
		// 	fmt.Printf("%07x MOV D,A\n", pc)
		// case 0x58:
		// 	fmt.Printf("%07x MOV E,B\n", pc)
		// case 0x59:
		// 	fmt.Printf("%07x MOV E,C\n", pc)
		// case 0x5a:
		// 	fmt.Printf("%07x MOV E,D\n", pc)
		// case 0x5b:
		// 	fmt.Printf("%07x MOV E,E\n", pc)
		// case 0x5c:
		// 	fmt.Printf("%07x MOV E,H\n", pc)
		// case 0x5d:
		// 	fmt.Printf("%07x MOV E,L\n", pc)
		// case 0x5e:
		// 	fmt.Printf("%07x MOV E,M\n", pc)
		// case 0x5f:
		// 	fmt.Printf("%07x MOV E,A\n", pc)

		// // 0x60 - 0x6f
		// case 0x60:
		// 	fmt.Printf("%07x MOV H,B\n", pc)
		// case 0x61:
		// 	fmt.Printf("%07x MOV H,C\n", pc)
		// case 0x62:
		// 	fmt.Printf("%07x MOV H,D\n", pc)
		// case 0x63:
		// 	fmt.Printf("%07x MOV H,E\n", pc)
		// case 0x64:
		// 	fmt.Printf("%07x MOV H,H\n", pc)
		// case 0x65:
		// 	fmt.Printf("%07x MOV H,L\n", pc)
		// case 0x66:
		// 	fmt.Printf("%07x MOV H,M\n", pc)
		// case 0x67:
		// 	fmt.Printf("%07x MOV H,A\n", pc)
		// case 0x68:
		// 	fmt.Printf("%07x MOV L,B\n", pc)
		// case 0x69:
		// 	fmt.Printf("%07x MOV L,C\n", pc)
		// case 0x6a:
		// 	fmt.Printf("%07x MOV L,D\n", pc)
		// case 0x6b:
		// 	fmt.Printf("%07x MOV L,E\n", pc)
		// case 0x6c:
		// 	fmt.Printf("%07x MOV L,H\n", pc)
		// case 0x6d:
		// 	fmt.Printf("%07x MOV L,L\n", pc)
		// case 0x6e:
		// 	fmt.Printf("%07x MOV L,M\n", pc)
		// case 0x6f:
		// 	fmt.Printf("%07x MOV L,A\n", pc)

		// // 0x70 - 0x7f
		// case 0x70:
		// 	fmt.Printf("%07x MOV M,B\n", pc)
		// case 0x71:
		// 	fmt.Printf("%07x MOV M,C\n", pc)
		// case 0x72:
		// 	fmt.Printf("%07x MOV M,D\n", pc)
		// case 0x73:
		// 	fmt.Printf("%07x MOV M,E\n", pc)
		// case 0x74:
		// 	fmt.Printf("%07x MOV M,H\n", pc)
		// case 0x75:
		// 	fmt.Printf("%07x MOV M,L\n", pc)
		// case 0x76:
		// 	fmt.Printf("%07x HLT\n", pc)
		// case 0x77:
		// 	fmt.Printf("%07x MOV M,A\n", pc)
		// case 0x78:
		// 	fmt.Printf("%07x MOV A,B\n", pc)
		// case 0x79:
		// 	fmt.Printf("%07x MOV A,C\n", pc)
		// case 0x7a:
		// 	fmt.Printf("%07x MOV A,D\n", pc)
		// case 0x7b:
		// 	fmt.Printf("%07x MOV A,E\n", pc)
		// case 0x7c:
		// 	fmt.Printf("%07x MOV A,H\n", pc)
		// case 0x7d:
		// 	fmt.Printf("%07x MOV A,L\n", pc)
		// case 0x7e:
		// 	fmt.Printf("%07x MOV A,M\n", pc)
		// case 0x7f:
		// 	fmt.Printf("%07x MOV A,A\n", pc)

		// // 0x80 - 0x8f
		// case 0x80:
		// 	fmt.Printf("%07x ADD B\n", pc)
		// case 0x81:
		// 	fmt.Printf("%07x ADD C\n", pc)
		// case 0x82:
		// 	fmt.Printf("%07x ADD D\n", pc)
		// case 0x83:
		// 	fmt.Printf("%07x ADD E\n", pc)
		// case 0x84:
		// 	fmt.Printf("%07x ADD H\n", pc)
		// case 0x85:
		// 	fmt.Printf("%07x ADD L\n", pc)
		// case 0x86:
		// 	fmt.Printf("%07x ADD M\n", pc)
		// case 0x87:
		// 	fmt.Printf("%07x ADD A\n", pc)
		// case 0x88:
		// 	fmt.Printf("%07x ADC B\n", pc)
		// case 0x89:
		// 	fmt.Printf("%07x ADC C\n", pc)
		// case 0x8a:
		// 	fmt.Printf("%07x ADC D\n", pc)
		// case 0x8b:
		// 	fmt.Printf("%07x ADC E\n", pc)
		// case 0x8c:
		// 	fmt.Printf("%07x ADC H\n", pc)
		// case 0x8d:
		// 	fmt.Printf("%07x ADC L\n", pc)
		// case 0x8e:
		// 	fmt.Printf("%07x ADC M\n", pc)
		// case 0x8f:
		// 	fmt.Printf("%07x ADC A\n", pc)

		// // 0x90 - 0x9f
		// case 0x90:
		// 	fmt.Printf("%07x SUB B\n", pc)
		// case 0x91:
		// 	fmt.Printf("%07x SUB C\n", pc)
		// case 0x92:
		// 	fmt.Printf("%07x SUB D\n", pc)
		// case 0x93:
		// 	fmt.Printf("%07x SUB E\n", pc)
		// case 0x94:
		// 	fmt.Printf("%07x SUB H\n", pc)
		// case 0x95:
		// 	fmt.Printf("%07x SUB L\n", pc)
		// case 0x96:
		// 	fmt.Printf("%07x SUB M\n", pc)
		// case 0x97:
		// 	fmt.Printf("%07x SUB A\n", pc)
		// case 0x98:
		// 	fmt.Printf("%07x SBB B\n", pc)
		// case 0x99:
		// 	fmt.Printf("%07x SBB C\n", pc)
		// case 0x9a:
		// 	fmt.Printf("%07x SBB D\n", pc)
		// case 0x9b:
		// 	fmt.Printf("%07x SBB E\n", pc)
		// case 0x9c:
		// 	fmt.Printf("%07x SBB H\n", pc)
		// case 0x9d:
		// 	fmt.Printf("%07x SBB L\n", pc)
		// case 0x9e:
		// 	fmt.Printf("%07x SBB M\n", pc)
		// case 0x9f:
		// 	fmt.Printf("%07x SBB A\n", pc)

		// // 0xa0 - 0xaf
		// case 0xa0:
		// 	fmt.Printf("%07x ANA B\n", pc)
		// case 0xa1:
		// 	fmt.Printf("%07x ANA C\n", pc)
		// case 0xa2:
		// 	fmt.Printf("%07x ANA D\n", pc)
		// case 0xa3:
		// 	fmt.Printf("%07x ANA E\n", pc)
		// case 0xa4:
		// 	fmt.Printf("%07x ANA H\n", pc)
		// case 0xa5:
		// 	fmt.Printf("%07x ANA L\n", pc)
		// case 0xa6:
		// 	fmt.Printf("%07x ANA M\n", pc)
		// case 0xa7:
		// 	fmt.Printf("%07x ANA A\n", pc)
		// case 0xa8:
		// 	fmt.Printf("%07x XRA B\n", pc)
		// case 0xa9:
		// 	fmt.Printf("%07x XRA C\n", pc)
		// case 0xaa:
		// 	fmt.Printf("%07x XRA D\n", pc)
		// case 0xab:
		// 	fmt.Printf("%07x XRA E\n", pc)
		// case 0xac:
		// 	fmt.Printf("%07x XRA H\n", pc)
		// case 0xad:
		// 	fmt.Printf("%07x XRA L\n", pc)
		// case 0xae:
		// 	fmt.Printf("%07x XRA M\n", pc)
		// case 0xaf:
		// 	fmt.Printf("%07x XRA A\n", pc)

		// // 0xb0 - 0xbf
		// case 0xb0:
		// 	fmt.Printf("%07x ORA B\n", pc)
		// case 0xb1:
		// 	fmt.Printf("%07x ORA C\n", pc)
		// case 0xb2:
		// 	fmt.Printf("%07x ORA D\n", pc)
		// case 0xb3:
		// 	fmt.Printf("%07x ORA E\n", pc)
		// case 0xb4:
		// 	fmt.Printf("%07x ORA H\n", pc)
		// case 0xb5:
		// 	fmt.Printf("%07x ORA L\n", pc)
		// case 0xb6:
		// 	fmt.Printf("%07x ORA M\n", pc)
		// case 0xb7:
		// 	fmt.Printf("%07x ORA A\n", pc)
		// case 0xb8:
		// 	fmt.Printf("%07x CMP B\n", pc)
		// case 0xb9:
		// 	fmt.Printf("%07x CMP C\n", pc)
		// case 0xba:
		// 	fmt.Printf("%07x CMP D\n", pc)
		// case 0xbb:
		// 	fmt.Printf("%07x CMP E\n", pc)
		// case 0xbc:
		// 	fmt.Printf("%07x CMP H\n", pc)
		// case 0xbd:
		// 	fmt.Printf("%07x CMP L\n", pc)
		// case 0xbe:
		// 	fmt.Printf("%07x CMP M\n", pc)
		// case 0xbf:
		// 	fmt.Printf("%07x CMP A\n", pc)

		// // 0xc0 - 0xcf
		// case 0xc0:
		// 	fmt.Printf("%07x RNZ\n", pc)
		// case 0xc1:
		// 	fmt.Printf("%07x POP B\n", pc)
		// case 0xc2:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x JNZ $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xc3:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x JMP $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xc4:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x CNZ $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xc5:
		// 	fmt.Printf("%07x PUSH B\n", pc)
		// case 0xc6:
		// 	if len(insSet) <= pc+1 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := insSet[pc+1]
		// 	fmt.Printf("%07x ADI #$%02x\n", pc, destination)
		// 	pc = pc + 1
		// case 0xc7:
		// 	fmt.Printf("%07x RST 0\n", pc)
		// case 0xc8:
		// 	fmt.Printf("%07x RZ\n", pc)
		// case 0xc9:
		// 	fmt.Printf("%07x RET\n", pc)
		// case 0xca:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x JZ $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xcb:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x JMP $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xcc:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x CZ $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xcd:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x CALL $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xce:
		// 	if len(insSet) <= pc+1 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := insSet[pc+1]
		// 	fmt.Printf("%07x ACI #$%02x\n", pc, destination)
		// 	pc = pc + 1
		// case 0xcf:
		// 	fmt.Printf("%07x RST 1\n", pc)

		// // 0xd0 - 0xdf
		// case 0xd0:
		// 	fmt.Printf("%07x RNC\n", pc)
		// case 0xd1:
		// 	fmt.Printf("%07x POP D\n", pc)
		// case 0xd2:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x JNC $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xd3:
		// 	if len(insSet) <= pc+1 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := insSet[pc+1]
		// 	fmt.Printf("%07x OUT #$%02x\n", pc, destination)
		// 	pc = pc + 1
		// case 0xd4:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x CNC $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xd5:
		// 	fmt.Printf("%07x PUSH D\n", pc)
		// case 0xd6:
		// 	if len(insSet) <= pc+1 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := insSet[pc+1]
		// 	fmt.Printf("%07x SUI #$%02x\n", pc, destination)
		// 	pc = pc + 1
		// case 0xd7:
		// 	fmt.Printf("%07x RST 2\n", pc)
		// case 0xd8:
		// 	fmt.Printf("%07x RC\n", pc)
		// case 0xd9:
		// 	fmt.Printf("%07x RET\n", pc)
		// case 0xda:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x JC $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xdb:
		// 	if len(insSet) <= pc+1 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := insSet[pc+1]
		// 	fmt.Printf("%07x IN #$%02x\n", pc, destination)
		// 	pc = pc + 1
		// case 0xdc:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x CC $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xdd:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x CALL $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xde:
		// 	if len(insSet) <= pc+1 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := insSet[pc+1]
		// 	fmt.Printf("%07x SBI #$%02x\n", pc, destination)
		// 	pc = pc + 1
		// case 0xdf:
		// 	fmt.Printf("%07x RST 3\n", pc)

		// // 0xe0 - 0xef
		// case 0xe0:
		// 	fmt.Printf("%07x RPO\n", pc)
		// case 0xe1:
		// 	fmt.Printf("%07x POP H\n", pc)
		// case 0xe2:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x JPO $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xe3:
		// 	fmt.Printf("%07x XTHL\n", pc)
		// case 0xe4:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x CPO $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xe5:
		// 	fmt.Printf("%07x PUSH H\n", pc)
		// case 0xe6:
		// 	if len(insSet) <= pc+1 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := insSet[pc+1]
		// 	fmt.Printf("%07x ANI #$%02x\n", pc, destination)
		// 	pc = pc + 1
		// case 0xe7:
		// 	fmt.Printf("%07x RST 4\n", pc)
		// case 0xe8:
		// 	fmt.Printf("%07x RPE\n", pc)
		// case 0xe9:
		// 	fmt.Printf("%07x PCHL\n", pc)
		// case 0xea:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x JPE $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xeb:
		// 	fmt.Printf("%07x XCHG\n", pc)
		// case 0xec:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x CPE $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xed:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x CALL $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xee:
		// 	if len(insSet) <= pc+1 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := insSet[pc+1]
		// 	fmt.Printf("%07x XRI #$%02x\n", pc, destination)
		// 	pc = pc + 1
		// case 0xef:
		// 	fmt.Printf("%07x RST 5\n", pc)

		// // 0xf0 - 0xff
		// case 0xf0:
		// 	fmt.Printf("%07x RP\n", pc)
		// case 0xf1:
		// 	fmt.Printf("%07x POP PSW\n", pc)
		// case 0xf2:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x JP $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xf3:
		// 	fmt.Printf("%07x DI\n", pc)
		// case 0xf4:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x CP $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xf5:
		// 	fmt.Printf("%07x PUSH PSW\n", pc)
		// case 0xf6:
		// 	if len(insSet) <= pc+1 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := insSet[pc+1]
		// 	fmt.Printf("%07x ORI #$%02x\n", pc, destination)
		// 	pc = pc + 1
		// case 0xf7:
		// 	fmt.Printf("%07x RST 6\n", pc)
		// case 0xf8:
		// 	fmt.Printf("%07x RM\n", pc)
		// case 0xf9:
		// 	fmt.Printf("%07x SPHL\n", pc)
		// case 0xfa:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x JM $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xfb:
		// 	fmt.Printf("%07x EI\n", pc)
		// case 0xfc:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x CM $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xfd:
		// 	if len(insSet) <= pc+1 || len(insSet) <= pc+2 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := append([]byte{insSet[pc+1]}, insSet[pc+2])
		// 	fmt.Printf("%07x CALL $%02x%02x\n", pc, destination[1], destination[0])
		// 	pc = pc + 2
		// case 0xfe:
		// 	if len(insSet) <= pc+1 {
		// 		return 0, ErrUnimplementedInstruction
		// 	}
		// 	destination := insSet[pc+1]
		// 	fmt.Printf("%07x CPI #$%02x\n", pc, destination)
		// 	pc = pc + 1
		// case 0xff:
		// 	fmt.Printf("%07x RST 7\n", pc)
	}
	state.PC += 1

	return state.PC, nil
}
