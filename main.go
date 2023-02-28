package gotest

import (
	"fmt"
)

type Parent struct {
	id int
	name string
}

func (p *Parent) Show() {
	fmt.Println("Parent")
	fmt.Printf("ID:%d, Name: %s\n", p.id, p.name)
}

type Children struct {
	id int
	Parent
}

func (c *Children) Show() {
	fmt.Println("Children")
	fmt.Printf("ID:%d, Name: %s\n", c.id, c.Parent.name)
}

func main() {
	a := &Children{
		0,
		Parent{
			1,"kek",
		},
	}
	a.Show()

}
