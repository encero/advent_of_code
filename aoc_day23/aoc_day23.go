package main

import (
  "fmt"
  "regexp"
  "bufio"
  "os"
  "io"
  "strconv"
)

func read() CPU {
  in := bufio.NewReader(os.Stdin)

  i_hlf := regexp.MustCompile("^hlf ([ab])\n$")
  i_tpl := regexp.MustCompile("^tpl ([ab])\n$")
  i_inc := regexp.MustCompile("^inc ([ab])\n$")
  i_jmp := regexp.MustCompile("^jmp ([+-][0-9]+)\n$")
  i_jie := regexp.MustCompile("^jie ([ab]), ([+-][0-9]+)\n$")
  i_jio := regexp.MustCompile("^jio ([ab]), ([+-][0-9]+)\n$")

  cpu := CPU{0, make([]Instruction, 0), 1, 0}

  for  {
    line, err := in.ReadString('\n')

    if (err == io.EOF || line == "") { break }

    match := i_hlf.FindStringSubmatch(line)
    if len(match) > 0 {
      cpu.add(hlf{match[1]})
      continue
    }

    match = i_tpl.FindStringSubmatch(line)
    if len(match) > 0 {
      cpu.add(tpl{match[1]})
      continue
    }

    match = i_inc.FindStringSubmatch(line)
    if len(match) > 0 {
      cpu.add(inc{match[1]})
      continue
    }

    match = i_jmp.FindStringSubmatch(line)
    if len(match) > 0 {
      cpu.add(jmp{toInt(match[1])})
      continue
    }

    match = i_jie.FindStringSubmatch(line)
    if len(match) > 0 {
      cpu.add(jie{match[1], toInt(match[2])})
      continue
    }

    match = i_jio.FindStringSubmatch(line)
    if len(match) > 0 {
      cpu.add(jio{match[1], toInt(match[2])})
      continue
    }

    fmt.Println("not recognized", line)
  }

  return cpu
}

func toInt(number string) int {
  v, _ := strconv.Atoi(number)

  return v
}

func main()  {
  cpu := read()

  cpu.run()

  fmt.Println("b ->", cpu.b, "a ->", cpu.a)

  fmt.Printf("end\n");
}
