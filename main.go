package main

import (
	"bytes"
	"fmt"
	"os"

	"emudev.com/disassembly"
)

func main() {
	file, err := os.ReadFile("invaders.h")
	if err != nil {
		panic(err)
	}
	pc := 0
	buf := bytes.NewBuffer(file)
	// for true {
	// 	b, err := buf.ReadByte()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	ok := Disassemble8080(b, pc)
	// 	if ok != 0 {
	// 		fmt.Println("Issue with disassembly")
	// 		break
	// 	}
	// 	pc++
	// }
	for pc < buf.Len() {
		pc = disassembly.Disassemble80800p(buf, pc)
	}

	fmt.Println("Success")
}

func Disassemble8080(instruction byte, pc int) int {
	switch instruction {
	case 0x00:
		fmt.Printf("%07x NOP\n", pc)
	}
	return 0
}
