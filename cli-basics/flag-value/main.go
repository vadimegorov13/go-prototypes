package main

import (
	"flag"
	"fmt"
	"strings"
)

type idsFlag []string

type person struct {
	name string
}

func (ids idsFlag) String() string {
	return strings.Join(ids, ", ")
}

func (ids *idsFlag) Set(id string) error {
	*ids = append(*ids, id)
	return nil
}

func (p person) String() string {
	return fmt.Sprintf("My name is %s", p.name)
}

func (p *person) Set(name string) error {
	p.name = name
	return nil
}

func main() {
	var ids idsFlag
	var p person

	flag.Var(&ids, "id", "The id to be appended to the list")
	flag.Var(&p, "name", "The name of the person")
	flag.Parse()
	fmt.Println(ids)
	fmt.Println(p)
}
