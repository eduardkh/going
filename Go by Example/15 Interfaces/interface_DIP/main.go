package main

import "fmt"

// SoundMaker interface for animals that can make a sound
type SoundMaker interface {
	MakeSound() string
}

// Renamable interface for entities that can be renamed
type Renamable interface {
	Rename(newName string)
	GetName() string // Added to safely retrieve names without assuming underlying type
}

// NamedEntity provides a base implementation for entities that have a name
type NamedEntity struct {
	Name string
}

// Rename changes the name of the entity
func (n *NamedEntity) Rename(newName string) {
	n.Name = newName
}

func (n *NamedEntity) GetName() string {
	return n.Name
}

// Dog struct with a NamedEntity embedded for common functionality
type Dog struct {
	NamedEntity
}

func (d *Dog) MakeSound() string {
	return "Woof!"
}

// Cat struct with a NamedEntity embedded for common functionality
type Cat struct {
	NamedEntity
}

func (c *Cat) MakeSound() string {
	return "Meow!"
}

// Bird struct with a NamedEntity embedded for common functionality
type Bird struct {
	NamedEntity
}

func (b *Bird) MakeSound() string {
	return "Tweet!"
}

func main() {
	animals := []SoundMaker{
		&Dog{NamedEntity{Name: "Buddy"}},
		&Cat{NamedEntity{Name: "Whiskers"}},
		&Bird{NamedEntity{Name: "Sunny"}},
	}

	renamables := []Renamable{
		&Dog{NamedEntity{Name: "Rex"}},
		&Cat{NamedEntity{Name: "Felix"}},
		&Bird{NamedEntity{Name: "Polly"}},
	}

	// Interact with animals through the SoundMaker interface
	for _, animal := range animals {
		fmt.Println(animal.MakeSound())
	}

	// // Interact with entities through the Renamable interface
	// for _, entity := range renamables {
	// 	// entity.Rename("New Name")
	// 	fmt.Println("Name:", entity.GetName()) // Safely retrieve and print the new name
	// 	entity.Rename(fmt.Sprintf("%s 2.0", entity.GetName()))
	// 	fmt.Println("Renamed Entity to:", entity.GetName()) // Safely retrieve and print the new name
	// }
	// Interact with entities through the Renamable interface
	for _, entity := range renamables {
		fmt.Println("Name:", entity.GetName()) // Safely retrieve and print the current name

		// Type assertion to check if the entity is a Dog
		if dog, isDog := entity.(*Dog); isDog {
			// Rename only if the entity is a Dog
			dog.Rename(fmt.Sprintf("%s 2.0", dog.GetName()))
			fmt.Println("Renamed Dog to:", dog.GetName())
		} else {
			// For non-dog entities, just print the name without renaming
			fmt.Println("Non-dog entity, name remains:", entity.GetName())
		}
	}

}
