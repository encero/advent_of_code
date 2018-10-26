package main

import (
	"fmt"
	"github.com/encero/advent_of_code/tools"
	"regexp"
	"strconv"
	"strings"
)

//                                       yqjtia (148) -> xxikrr, jxcxf, bqgxg
var fullEx = regexp.MustCompile(`^([a-z]+) \(([0-9]+)\)(?: -> ([a-z, ]+))?$`)

type Leaf struct {
	Name        string
	Weight      int64
	childsTotal int64
	NodeNames   []string
	Childs      []*Leaf
	Parent      *Leaf
}

func (l *Leaf) getTotal() int64 {
	if l.childsTotal == 0 {
		subTotal := int64(0)

		for _, ch := range l.Childs {
			subTotal += ch.getTotal()
		}

		l.childsTotal = subTotal + l.Weight
	}

	return l.childsTotal
}

func main() {
	leafs := make(map[string]*Leaf)

	lines := tools.ReadLines("./day_7.txt")

	// parse all lines to Leaf structs
	for _, l := range lines {
		match := fullEx.FindStringSubmatch(l) // find name weight and connection part if is there any

		weight, _ := strconv.ParseInt(match[2], 10, 64) // parse weight to int

		names := []string{}

		if match[3] != "" { // some names in third part
			names = strings.Split(match[3], ",") // simple split instead of regexp

			names = StringMap(names, func(s string) string {
				return strings.Trim(s, " ")
			})
		}

		leaf := &Leaf{
			Name:      match[1],
			Weight:    weight,
			NodeNames: names,
		}

		leafs[leaf.Name] = leaf // save leaf to map with its name for futher reference
	}

	// map all leafs to its childrens and parents
	for _, leaf := range leafs {
		for _, name := range leaf.NodeNames {
			leaf.Childs = append(leaf.Childs, leafs[name])
			leafs[name].Parent = leaf
		}
	}

	// find leaf with no parent ( solution 1 )
	// precompute all sum weights
	for _, leaf := range leafs {
		leaf.getTotal()

		if leaf.Parent == nil {
			fmt.Println("bottom leaf:", leaf.Name)
		}
	}

	for _, leaf := range leafs {
		balanced := true
		for i, ch := range leaf.Childs {
			for j := i + 1; j < len(leaf.Childs); j++ {

				if ch.getTotal() != leaf.Childs[j].getTotal() {
					balanced = false
				}
			}
		}

		if ! balanced {
			fmt.Println("Unbalanced leaf", leaf.Name , "weight:", leaf.Weight)
			for _, ch := range leaf.Childs {
				fmt.Print(ch.Name, " ", ch.getTotal(), "|")
			}
			fmt.Println()
			fmt.Println()
		}
	}
}

func StringMap(arr []string, f func(string) string) []string {
	tmp := make([]string, len(arr))

	for i, s := range arr {
		tmp[i] = f(s)
	}

	return tmp
}
