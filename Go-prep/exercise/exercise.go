package exercise

import "fmt"

type Player struct {
	name      string
	inventory []Item
}

type Item struct {
	name     string
	itemtype string
}

//pickupitem -- adds item to the item inventory
//dropitem -- removes items from inventory
//useitem -- uses items from inventory if present

func (p *Player) PickUpitem(item Item) {
	//item added
	p.inventory = append(p.inventory, item)
}

func (p *Player) Dropitem(newItem Item) {
	for i,item := range p.inventory{
		if item == newItem{
			//found item 
			fmt.Printf("Dropped a %s named %s from %s's inventory",newItem.itemtype,newItem.name,p.name)
			p.inventory = append(p.inventory[:i], p.inventory[i+1:]...)
		}else{
			fmt.Println("This item is not present in the inventory")
		}
	}
}

func (p *Player) Useitem(itemName Item) {
	//checking if item is present or not
	for i, item := range p.inventory {
		if item == itemName {
			if itemName.itemtype == "potion" {
				//found potion
				fmt.Printf("%s used %s and feels good", p.name, itemName)
			} else {
				//found some other thing
				fmt.Printf("%s used a %s named %s", p.name, itemName.itemtype, itemName.name)
			}
			p.inventory = append(p.inventory[:i], p.inventory[i+1:]...)
		} else {
			fmt.Printf("Did not find %s in %s's inventory", itemName.name, p.name)
		}
	}

}
