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
	processor.process()
	//fmt.Println(processor)
	strOutputs := make([]string, len(processor.outputs))
	for i, output := range processor.outputs {
		strOutputs[i] = strconv.Itoa(output)
	}
	outStr := strings.Join(strOutputs, ",")
	return outStr
}

func (Solution) Part2(input []byte) string {
	processor := buildProcessor(input)
	processor.requiresMatching = true
	i := 0
	//defer func() { fmt.Println(i) }()
	//fmt.Println(math.MaxInt)
	for {
		processor.outputs = make([]int, 0, len(processor.program))
		processor.a = i
		processor.pointerIndex = 0
		processor.quit = false
		finished := processor.process()
		if finished && sliceEqual(processor.outputs, processor.program) {
			return strconv.Itoa(i)
		}
		i++
	}
}

func (Solution) GetExample(part int) []byte {
	if part == 1 {
		return []byte("Register A: 729\nRegister B: 0\nRegister C: 0\n\nProgram: 0,1,5,4,3,0")
	} else {
		return []byte("Register A: 2024\nRegister B: 0\nRegister C: 0\n\nProgram: 0,3,5,4,3,0")
	}
}

func (Solution) ExampleAnswer1() string {
	return "4,6,3,5,6,3,5,2,1,0"
}
func (Solution) ExampleAnswer2() string {
	return "117440"
}

type Processor struct {
	program               []int
	a, b, c, pointerIndex int
	outputs               []int
	requiresMatching      bool
	quit                  bool
}

func (p *Processor) comboOperand(v int) int {
	switch v {
	case 4:
		return p.a
	case 5:
		return p.b
	case 6:
		return p.c
	default:
		return v
	}
}

type Instruction func(p *Processor, v int)

func (p *Processor) dv(v int) int {
	return p.a >> p.comboOperand(v)
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
		if p.requiresMatching && (len(p.program) <= len(p.outputs) || value != p.program[len(p.outputs)]) {
			//fmt.Println("quitting early", p.program, p.outputs, value)
			p.quit = true
			return
		}
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

func (p *Processor) process() bool {
	done := false
	for !done {
		//fmt.Println(processor)
		done = p.doInstruction()
		if p.quit {
			return false
		}
	}
	return true
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
