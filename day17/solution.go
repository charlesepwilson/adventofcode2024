package day17

import (
	"bytes"
	"strconv"
	"strings"
)

type Solution struct{}

func (Solution) Day() int { return 17 }

func (Solution) Part1(input []byte) string {
	processor := buildProcessor(input)
	done := false
	for !done {
		//fmt.Println(processor)
		done = processor.doInstruction()
	}
	//fmt.Println(processor)
	strOutputs := make([]string, len(processor.outputs))
	for i, output := range processor.outputs {
		strOutputs[i] = strconv.Itoa(output)
	}
	outStr := strings.Join(strOutputs, ",")
	return outStr
}

func (Solution) Part2(input []byte) string {
	return ""
}

func (Solution) GetExample(part int) []byte {
	return []byte("Register A: 729\nRegister B: 0\nRegister C: 0\n\nProgram: 0,1,5,4,3,0")
}

func (Solution) ExampleAnswer1() string {
	return "4,6,3,5,6,3,5,2,1,0"
}
func (Solution) ExampleAnswer2() string {
	return ""
}

type Processor struct {
	program               []int
	a, b, c, pointerIndex int
	outputs               []int
}

func (p *Processor) comboOperand(v int) int {
	if v == 4 {
		return p.a
	} else if v == 5 {
		return p.b
	} else if v == 6 {
		return p.c
	} else if v == 7 {
		panic("invalid operand")
	} else {
		return v
	}
}

type Instruction func(p *Processor, v int)

func (p *Processor) dv(v int) int {
	return p.a >> p.comboOperand(v)
	//denominator := 1 << p.comboOperand(v)
	//if denominator == 0 {
	//	return 0
	//}
	//return numerator / denominator
}

var instructions = []Instruction{
	func(p *Processor, v int) { // adv
		p.a = p.dv(v)
	},
	func(p *Processor, v int) { // bxl
		p.b = p.b ^ v
	},
	func(p *Processor, v int) { // bst
		p.b = p.comboOperand(v) % 8
	},
	func(p *Processor, v int) { // jnz
		if p.a == 0 {
			return
		}
		p.pointerIndex = v - 2
	},
	func(p *Processor, v int) { // bxc
		p.b = p.b ^ p.c
	},
	func(p *Processor, v int) { // out
		value := p.comboOperand(v) % 8
		p.outputs = append(p.outputs, value)
	},
	func(p *Processor, v int) { // bdv
		p.b = p.dv(v)
	},
	func(p *Processor, v int) { // cdv
		p.c = p.dv(v)
	},
}

func (p *Processor) doInstruction() (done bool) {
	if p.pointerIndex >= len(p.program) {
		return true
	}
	opCode := p.program[p.pointerIndex]
	operand := p.program[p.pointerIndex+1]
	instructions[opCode](p, operand)
	p.pointerIndex += 2
	return false
}

func buildProcessor(input []byte) Processor {
	registerText, programText, _ := bytes.Cut(input, []byte("\n\n"))
	registerLines := bytes.Split(registerText, []byte("\n"))
	startVals := [3]int{}
	for i, line := range registerLines {
		_, s, _ := bytes.Cut(line, []byte(": "))
		n, _ := strconv.Atoi(string(s))
		startVals[i] = n
	}
	_, programSec, _ := bytes.Cut(programText, []byte(": "))
	programNumsStrs := bytes.Split(programSec, []byte(","))
	programNums := make([]int, len(programNumsStrs))
	for i, str := range programNumsStrs {
		num, _ := strconv.Atoi(string(str))
		programNums[i] = num
	}
	processor := Processor{
		program:      programNums,
		a:            startVals[0],
		b:            startVals[1],
		c:            startVals[2],
		pointerIndex: 0,
	}
	return processor
}
