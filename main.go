package main

import "fmt"

/*
func (c Car)  Print(){
	fmt.Println(c)
}
func swap(ml,m2 *int) {
	var temp int
	temp = *m2
	*m2 = *ml
	*ml = temp
}

type Car struct {
		Name string
		Miles int
		ModelNo int
}
func todo() {
	arr := []int{1,2,3,4}
	fmt.Println(arr)
}
func (c Car) GetName() string {
	return c.Name
}*/

type Car interface {
	Drive()
	Stop()
}
type Lambo struct {
	LamboModel string
}
type Chevy struct {
	ChevyModel string
}

func NewModel(arg string) Car {
	return &Lambo{arg}
}
func (l *Lambo) Drive() {
	fmt.Println("Lambo moving")
	fmt.Println(l.LamboModel)
}
func (l *Lambo) Stop() {
	fmt.Println("Stopping moving")

}
func (c *Chevy) Drive() {
	fmt.Println("Chevy moving")
	fmt.Println(c.ChevyModel)
}
func main() {
	l := Lambo{"gallardo"}
	c := Chevy{"C369"}
	l.Drive()
	l.Stop()
	c.Drive()
}
