package main

import "fmt"

type A interface {
	DoTask() bool
}

type a struct{}

func (sa *a) DoTask() bool {
	return true
}

type b struct {
	a A
}

func (sb *b) DoSomething() bool {
	//Do some logic
	fmt.Println(sb.a.DoTask())
	//Do some logic
	return true
}

func main() {

	m := &b{a: &a{}}

	m.DoSomething()
}
