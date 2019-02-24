package main

import "fmt"

// Dog has a name and a list of people caring for him and watching
// what he does
type Dog struct {
	name    string
	sitters map[string][]chan string
}

// AddSitter adds an event listener to the Dog struct instance
func (d *Dog) AddSitter(e string, ch chan string) {
	if d.sitters == nil {
		d.sitters = make(map[string][]chan string)
	}

	if _, ok := d.sitters[e]; ok {
		d.sitters[e] = append(d.sitters[e], ch)
	} else {
		d.sitters[e] = []chan string{ch}
	}

}

// RemoveSitter removes an event listener from the Dog struct instance
func (d *Dog) RemoveSitter(e string, ch chan string) {
	if _, ok := d.sitters[e]; ok {
		for i, ch := range d.sitters[e] {
			if d.sitters[e][i] == ch {
				d.sitters[e] = append(d.sitters[e][:i], d.sitters[e][i+1:]...)
				break
			}
		}
	}
}

// Emit emits an event on the Dog struct instance
func (d *Dog) Emit(e string, response string) {
	if _, ok := d.sitters[e]; ok {
		for _, handler := range d.sitters[e] {
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}

func main() {
	balto := Dog{"Balto", nil}

	abhi := make(chan string)
	pan := make(chan string)
	balto.AddSitter("bark", abhi)
	balto.AddSitter("hungry", abhi)
	balto.AddSitter("poop", abhi)

	go func() {
		for {
			msg := <-abhi
			fmt.Println("Abhi: " + msg)
		}
	}()

	go func() {
		for {
			msg := <-pan
			fmt.Println("Pandey: " + msg)
		}
	}()

	fmt.Println("The dog barked!")
	balto.Emit("bark", "Do not bark punk !!! ")

	fmt.Println("The dog pooped")
	balto.Emit("poop", "Picked up the poop !!!")

	fmt.Println("The dog is hungry")
	balto.Emit("hungry", "The dog has been fed !!!")

	balto.RemoveSitter("poop", abhi)

	balto.AddSitter("poop", pan)

	fmt.Println("The dog pooped")
	balto.Emit("poop", "Pickep Up the poop !!!")

	fmt.Println("The dog pooped again")
	balto.Emit("poop", "Pickep Up the poop again !!!")

	fmt.Println("The dog is hungry again")
	balto.Emit("hungry", "The dog is fed again !!!")

	fmt.Scanln()

}
