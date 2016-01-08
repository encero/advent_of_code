package main

import (
  "fmt"
  "regexp"
  "bufio"
  "os"
  "io"
  "strings"
)

func main()  {
  in := bufio.NewReader(os.Stdin)

  reg1 := regexp.MustCompile(`"`)
  reg2 := regexp.MustCompile(`\\`)

  code := 0;
  memory := 0;

  for  {
    line, err := in.ReadString('\n')

    line = strings.Trim(line, "\n")

    c := len(line)

    line2 := reg2.ReplaceAllString(line, "##")
    line2 = reg1.ReplaceAllString(line2, "##")

    fmt.Printf("\n%s\n%s\n%d|%d",line, line2, len(line2) + 2, c)

    code += c
    memory += len(line2) + 2

    if (err == io.EOF) {
      break
    }
  }

  fmt.Printf("code: %d memory:%d\n", code, memory)
  fmt.Printf("%d\n", memory - code - 2)
  fmt.Printf("end\n");
}
