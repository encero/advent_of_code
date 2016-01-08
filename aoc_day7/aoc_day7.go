package main

import (
  "fmt"
  "bufio"
  "os"
  "regexp"
  "io"
  "strings"
  "strconv"
)



type Operator func(left uint16, right uint16) uint16

func operatorTROUGH (left uint16, right uint16) uint16 {
  return left
}

func operatorNOT (left uint16, right uint16) uint16 {
  return ^left
}

func operatorAND (left uint16, right uint16) uint16 {
  return left & right
}

func operatorOR (left uint16, right uint16) uint16 {
  return left | right
}

func operatorSHIFTL (left uint16, right uint16) uint16 {
  return left << right
}

func operatorSHIFTR (left uint16, right uint16) uint16 {
  return left >> right
}

type Wire struct {
  name string
  op Operator
  left *Wire
  right *Wire
  hasValue bool
  value uint16
}

var reTrough, reUnary, reBinary *regexp.Regexp
var wires []*Wire
var wireMap map[string]*Wire

func init() {
  reTrough = regexp.MustCompile("^([a-z]+|[0-9]+) -> ([a-z]+)$")
  reUnary = regexp.MustCompile("^NOT ([a-z]+|[0-9]+) -> ([a-z]+)$")
  reBinary = regexp.MustCompile("^([a-z]+|[0-9]+) (RSHIFT|LSHIFT|AND|OR) ([a-z]+|[0-9]+) -> ([a-z]+)$")

  wireMap = make(map[string]*Wire)
}

func createWire(name string) *Wire {
  wire := Wire{name, nil, nil, nil, false, 0}

  wires = append(wires, &wire)
  wireMap[name] = &wire

  return &wire
}

func getWire(name string) *Wire {
  value, err := strconv.Atoi(name)

  if err != nil {
    value, ok := wireMap[name]

    if ok {
      return value
    } else {
      return createWire(name)
    }
  } else {
    return &Wire{"", nil, nil, nil, true, uint16(value)}
  }
}

func parse(line string) {
  var match []string

  // No op wire
  match = reTrough.FindStringSubmatch(line)
  if len(match) > 0 {
    wire := getWire(match[2])
    wire.op = operatorTROUGH

    wire.left = getWire(match[1])

    return
  }

  //not operator
  match = reUnary.FindStringSubmatch(line)
  if len(match) > 0 {
    wire := getWire(match[2])
    wire.op = operatorNOT

    wire.left = getWire(match[1])

    return
  }

  match = reBinary.FindStringSubmatch(line)
  if len(match) > 0 {
    wire := getWire(match[4])

    wire.left = getWire(match[1])
    wire.right = getWire(match[3])

    switch match[2] {
      case "OR":
        wire.op = operatorOR
      case "AND":
        wire.op = operatorAND
      case "LSHIFT":
        wire.op = operatorSHIFTL
      case "RSHIFT":
        wire.op = operatorSHIFTR
    }
    return
  }

}

func eval(wire *Wire) uint16 {
  if wire == nil {
    return 0
  }

  if wire.hasValue {
    return wire.value
  }

  wire.value = wire.op(eval(wire.left), eval(wire.right))
  wire.hasValue = true

  return wire.value
}

func main()  {
  in := bufio.NewReader(os.Stdin)

  for  {
    line, err := in.ReadString('\n')

    parse(strings.Trim(line, "\n "))

    if (err == io.EOF) {
      break
    }
  }

  a := getWire("a")
  result := eval(a)

  fmt.Printf("%d\n", result)
  fmt.Printf("end\n");
}
