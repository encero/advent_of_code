package main

import (
  "fmt"
)

type CPU struct {
  pointer int
  instructions []Instruction
  a, b uint
}

func (cpu *CPU) run() {
  for cpu.pointer >= 0 && cpu.pointer < len(cpu.instructions) {
    instruction := cpu.instructions[cpu.pointer]
    instruction.apply(cpu)
    fmt.Printf("%T %v -> (%v)\n", instruction, instruction, cpu)
  }
}

func (cpu *CPU) add(i Instruction) {
  cpu.instructions = append(cpu.instructions, i)
}

func (cpu *CPU) String() string {
  return fmt.Sprintf("p->%d a->%d b->%d", cpu.pointer, cpu.a, cpu.b)
}

type Instruction interface {
  apply(*CPU)
}

type hlf struct {
  register string
}

func (instruction hlf) apply(cpu *CPU) {
  switch instruction.register {
    case "a":
      cpu.a = cpu.a / 2
    case "b":
      cpu.b = cpu.b / 2
  }

  cpu.pointer ++
}

type tpl struct {
  register string
}

func (instruction tpl) apply(cpu *CPU) {
  switch instruction.register {
    case "a":
      cpu.a = cpu.a * 3
    case "b":
      cpu.b = cpu.b * 3
  }

  cpu.pointer ++
}

type inc struct {
  register string
}

func (instruction inc) apply(cpu *CPU) {
  switch instruction.register {
    case "a":
      cpu.a ++
    case "b":
      cpu.b ++
  }

  cpu.pointer ++
}

type jmp struct {
  offset int
}

func (instruction jmp) apply(cpu *CPU) {
  cpu.pointer += instruction.offset
}

type jie struct {
  register string
  offset int
}

func (instuction jie) apply(cpu *CPU) {
  if (instuction.register == "a" && cpu.a % 2 == 0) || (instuction.register == "b" && cpu.b % 2 == 0) {
    cpu.pointer += instuction.offset
  } else {
    cpu.pointer ++
  }
}

type jio struct {
  register string
  offset int
}

func (instuction jio) apply(cpu *CPU) {
  if (instuction.register == "a" && cpu.a  == 1) || (instuction.register == "b" && cpu.b == 1) {
    cpu.pointer += instuction.offset
  } else {
    cpu.pointer ++
  }
}
