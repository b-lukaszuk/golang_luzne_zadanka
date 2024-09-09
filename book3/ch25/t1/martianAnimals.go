// simple solution,
// not idiot/error-proof, not optimized for speed
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Animal interface {
	move()
	eat()
}

type Bee struct {
	name string
}

func (b *Bee) String() string {
	return fmt.Sprintf("Bee %s", b.name)
}

func (b *Bee) move() {
	fmt.Printf("%s flies around, bzz, bzz.\n", b)
}

func (b *Bee) eat() {
	fmt.Printf("%s eats nectar.\n", b)
}

type Sheep struct {
	name string
}

func (b *Sheep) String() string {
	return fmt.Sprintf("Sheep %s", b.name)
}

func (s *Sheep) move() {
	fmt.Printf("%s runs around.\n", s)
}

func (s *Sheep) eat() {
	fmt.Printf("%s eats grass.\n", s)
}

type Cow struct {
	name string
}

func (b *Cow) String() string {
	return fmt.Sprintf("Cow %s", b.name)
}

func (c *Cow) move() {
	fmt.Printf("%s walks around.\n", c)
}

func (c *Cow) eat() {
	fmt.Printf("%s chews grass.\n", c)
}

func getRandNumber(minIncl float64, maxExcl float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(maxExcl-minIncl) + minIncl
}

func getRandInt(maxExcl uint) uint {
	return uint(getRandNumber(0, float64(maxExcl)))
}

func runCycle(dayNum uint) {
	animals := []Animal{&Bee{"Maya"}, &Sheep{"Dolly"}, &Cow{"Molly"}}
	var animalInd uint = 0
	var action uint = 0
	fmt.Printf("\n===Day %d===\n", dayNum)
	// shorter day, shorter output
	for i := 9; i < 18; i++ {
		animalInd = getRandInt(uint(len(animals)))
		action = getRandInt(2)
		fmt.Printf("%d o'clock.\n", i)
		if action == 0 {
			animals[animalInd].eat()
		} else {
			animals[animalInd].move()
		}
	}
	fmt.Printf("It's night. All animals are sleeping.\n")
}

func runNCycles(nCycles uint) {
	for i := 0; i < int(nCycles); i++ {
		runCycle(uint(i + 1))
	}
}

func main() {
	fmt.Printf("Toy program.\n")
	fmt.Printf("It describes the life of a few animals on a terraformed Mars.\n")

	runNCycles(3)

	fmt.Printf("\nThat's all. Goodbye!\n")
}
