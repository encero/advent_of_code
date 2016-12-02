package main

import (
  "fmt"
  "regexp"
  "bufio"
  "os"
  "io"
  "strconv"
//"strings"
)

type Ingredient struct {
  name string;
  capacity, durability, flavor, texture, calories, spoons int;
}

func (i Ingredient) String()string {return fmt.Sprintf("c:%d d:%d f:%d t:%d c:%d s:%d\n", i.capacity, i.durability, i.flavor, i.texture, i.calories, i.spoons)}

func toInt(number string) int {
  value, _ := strconv.Atoi(number)
  return value
}

func CreateIngredient(spec []string) *Ingredient {
  return &Ingredient{spec[0], toInt(spec[1]), toInt(spec[2]), toInt(spec[3]), toInt(spec[4]), toInt(spec[5]), 0}
}

func read() []*Ingredient {
  in := bufio.NewReader(os.Stdin)

//Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
  reg := regexp.MustCompile(`^(\w+): capacity (-?[0-9]+), durability (-?[0-9]+), flavor (-?[0-9]+), texture (-?[0-9]+), calories (-?[0-9]+)\n$`)

  var ingredients []*Ingredient

  for  {
    line, err := in.ReadString('\n')
    if err == io.EOF || line == "" {
      break
    }

    match := reg.FindStringSubmatch(line)

    ingredients =append(ingredients, CreateIngredient(match[1:]))
  }

  return ingredients
}

func sumUp(ingredients []*Ingredient, f func(i *Ingredient) int) int {
  sum := 0

  for _, v := range ingredients {
    sum += f(v) * v.spoons
  }

  if sum < 0 {
    return 0
  } else {
    return sum
  }
}

func calc(max *int, ingredients []*Ingredient) {
  if sumUp(ingredients, func(i *Ingredient) int { return i.calories } ) != 500 {
    return
  }

  current :=
    sumUp(ingredients, func(i *Ingredient) int { return i.capacity }) *
    sumUp(ingredients, func(i *Ingredient) int { return i.durability }) *
    sumUp(ingredients, func(i *Ingredient) int { return i.flavor }) *
    sumUp(ingredients, func(i *Ingredient) int { return i.texture })

  // fmt.Println(ingredients, current, *max, sumUp(ingredients, func(i *Ingredient) int { return i.capacity }),
  // sumUp(ingredients, func(i *Ingredient) int { return i.durability }),
  // sumUp(ingredients, func(i *Ingredient) int { return i.flavor }),
  // sumUp(ingredients, func(i *Ingredient) int { return i.texture }))

  if current > *max {
    *max = current
  }
}

func combine(spoons, index int, max *int, ingredients []*Ingredient) {
  if (index >= len(ingredients) - 1) {
    ingredients[index].spoons = spoons
    calc(max, ingredients)
  } else {
    for i:= 1; i <= (spoons - (len(ingredients) - index - 1) ); i++ {
      ingredients[index].spoons = i
      combine(spoons - i, index + 1, max, ingredients)
    }
  }
}

func main()  {
  ingredients := read()

  max := 0

  combine(100, 0, &max, ingredients)

  fmt.Println(ingredients)

  fmt.Printf("max: %d end\n", max);
}
