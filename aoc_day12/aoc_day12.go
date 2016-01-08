package main

import (
  "fmt"
//  "regexp"
  "io/ioutil"
  "os"
//  "strconv"
  "encoding/json"
)

func load(v interface{}) (float64, bool) {
  var sum float64 = 0

  switch vv := v.(type) {
  case map[string]interface{}:
    for _, item := range vv {
      value, ok := load(item)
      if !ok {
        return 0, true
      }

      sum += value
    }
  case []interface{}:
    for _, item := range vv {
      value, _ := load(item)
      sum += value
    }
  case float64:
    return vv, true
  case string:
    return 0, vv != "red"
  }

  return sum, true
}

func main() {
  bytes, _ := ioutil.ReadAll(os.Stdin)

  var v interface{}

  json.Unmarshal(bytes, &v)

  sum, _ := load(v)

  //fmt.Printf("%v", v)


  fmt.Printf("sum: %d\n", sum);
}
