package main

import "fmt"


type Item struct {
	name string
}

type Player struct {
	name string
	backpack []Item
}

func main() {
    steve := Player{name: "Steve", backpack: []Item{}}
	apple := Item{name: "Apple"}
	steve.pickItem(apple)
	steve.printBackpack()
	steve.useItem("Apple")
	steve.printBackpack()
	steve.useItem("Apple")
	steve.useItem("Orange")
}

func (p *Player)pickItem(it Item) {
	p.backpack = append(p.backpack, it)
	fmt.Printf("%s picked: %s\n", p.name, it.name)
}

func (p *Player)useItem(name string) {
	index := -1
	for i, it := range p.backpack {
		if it.name == name {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Printf("%s tried to use but %s doesn't have it.\n", p.name, name)
		return
	}

	usedItem := p.backpack[index] 
	p.backpack = append(p.backpack[:index], p.backpack[index+1:]...)
	fmt.Printf("%s used: %s\n", p.name, usedItem.name)
}


func (p *Player) printBackpack() {
	if len(p.backpack) == 0 {
		fmt.Printf("%s's backpack is empty.\n", p.name)
		return
	}
	fmt.Printf("%s's backpack contains:\n", p.name)
	for _, it := range p.backpack {
		fmt.Printf("- %s\n", it.name)
	}
}