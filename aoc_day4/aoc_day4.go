package main

import (
  "fmt"
  "crypto/md5"
  "encoding/hex"
  "strconv"
  "strings"
  "io"
)

func hash(key string, i int) (string) {
  h := md5.New()

  io.WriteString(h, key + strconv.Itoa(i))

  return hex.EncodeToString(h.Sum(nil))
}

func main() {
  i := 0;

  for {
    if strings.HasPrefix(hash("yzbqklnj", i), "000000") {
      break
    }

    i = i + 1
  }

  fmt.Printf("index: %d\n", i)
}
