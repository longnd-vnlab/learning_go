package main

import "fmt"

// Animal interface defines the behavior for all animal types
type Animal interface {
	Speak()
	Type() string
}

// Dog implements Animal with value receiver methods
type Dog struct {
	Name  string
	Breed string
}

// Cat implements Animal with pointer receiver methods
type Cat struct {
	Name  string
	Color string
}

// Bird partially implements Animal (for demonstration purposes)
type Bird struct {
	Name    string
	Species string
}

// Speak implementation for Dog uses a value receiver
// This means both Dog values and pointers can be used as Animals
func (d Dog) Speak() {
	fmt.Println(d.Name + " says: Woof!")
}

// Type returns the animal type for Dog
func (d Dog) Type() string {
	return "Dog (" + d.Breed + ")"
}

// Speak implementation for Cat uses a pointer receiver
// This means ONLY Cat pointers can be used as Animals
func (c *Cat) Speak() {
	fmt.Println(c.Name + " says: Meow!")
}

// Type returns the animal type for Cat
func (c *Cat) Type() string {
	return "Cat (" + c.Color + ")"
}

// Speak implementation for Bird
func (b Bird) Speak() {
	fmt.Println(b.Name + " says: Chirp!")
}

// Note: Bird does not implement Type() method, so it cannot satisfy Animal interface

// MakeAnimalSpeak demonstrates polymorphic behavior through interfaces
func MakeAnimalSpeak(a Animal) {
	fmt.Printf("Animal type: %s\n", a.Type())
	a.Speak()
}

func main() {
	fmt.Println("=== Interface Method Receivers in Go ===")

	// Using Dog with value receiver methods
	dog := Dog{"Rex", "German Shepherd"}
	dogPtr := &Dog{"Buddy", "Golden Retriever"}

	// Both work because Dog has value receiver methods
	fmt.Println("--- Value receiver (Dog) ---")
	MakeAnimalSpeak(dog)    // Value works with value receiver
	MakeAnimalSpeak(dogPtr) // Pointer also works with value receiver

	// Using Cat with pointer receiver methods
	cat := &Cat{"Fluffy", "Tabby"} // Must use pointer for Cat

	fmt.Println("\n--- Pointer receiver (Cat) ---")
	MakeAnimalSpeak(cat) // Works with pointer

	// This would NOT compile:
	// directCat := Cat{"Whiskers", "Calico"}
	// MakeAnimalSpeak(directCat)  // Error: Cat does not implement Animal (Speak method has pointer receiver)

	fmt.Println("\n--- Interface Assignment Examples ---")

	// These work:
	var animal Animal
	animal = Dog{"Max", "Beagle"} // Value type with value receiver methods
	animal.Speak()

	animal = &Dog{"Cooper", "Poodle"} // Pointer type with value receiver methods
	animal.Speak()

	animal = &Cat{"Mittens", "Black"} // Pointer type with pointer receiver methods
	animal.Speak()

	// This would fail to compile:
	// animal = Cat{"Garfield", "Orange"}  // Value type with pointer receiver methods

	fmt.Println("\n--- Implementation Rule ---")
	fmt.Println("1. If a type has VALUE receiver methods, both values AND pointers implement the interface")
	fmt.Println("2. If a type has POINTER receiver methods, only pointers implement the interface")

	// Bird example - incomplete interface implementation
	// This would not compile:
	// animal = Bird{"Tweety", "Canary"}  // Missing Type() method
}
