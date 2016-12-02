package main

import (
  "fmt"
  "regexp"
  "bufio"
  "os"
  "io"
  "strconv"
)

func toInt(number string) uint {
  value, _ := strconv.Atoi(number)
  return uint(value)
}

type Aunt struct {
  number uint
  traits map[string]uint
}

func read() []*Aunt{
  in := bufio.NewReader(os.Stdin)

  aunts := make([]*Aunt, 500)

  //                          Sue 1: children: 1, cars: 8, vizslas: 7
  reg := regexp.MustCompile("^Sue ([0-9]+): ([a-z]+): ([0-9]+), ([a-z]+): ([0-9]+), ([a-z]+): ([0-9]+)\n$")

  index := 0

  for  {
    line, err := in.ReadString('\n')

    if (err == io.EOF) {
      break
    }

    subMatch := reg.FindStringSubmatch(line)

    if len(subMatch) > 0 {

      aunts[index] = & Aunt{toInt(subMatch[1]), map[string]uint{
        subMatch[2]: toInt(subMatch[3]),
        subMatch[4]: toInt(subMatch[5]),
        subMatch[6]: toInt(subMatch[7]),
        }}
        index ++
    }

  }

  return aunts
}

func main()  {
  aunts := read()

  sue := Aunt{
    0, map[string]uint{
      "children": 3,
      "cats": 7,
      "samoyeds": 2,
      "pomeranians": 3,
      "akitas": 0,
      "vizslas": 0,
      "goldfish": 5,
      "trees": 3,
      "cars": 2,
      "perfumes": 1,
    }}

  for _, aunt := range aunts {
    ok := true

    for trait, value := range aunt.traits {
      if trait == "cats" || trait == "trees" {
        if sue.traits[trait] >= value {
          ok = false
          break
        }
      } else if trait == "pomeranians" || trait == "goldfish" {
        if sue.traits[trait] <= value {
          ok = false
          break
        }
      } else {
        if sue.traits[trait] != value {
          ok = false
          break
        }
      }
    }

    if ok {
      fmt.Println(aunt)
    }
  }

  fmt.Printf("end\n");
}
