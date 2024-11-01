package main

import (
	"fmt"
	"image"
	"log"
	"time"
)

type command int

const (
	right = command(0)
	left  = command(1)
	start = command(3)
	stop  = command(4)
)

type RoverDriver struct {
	commandc chan command
}

// drive is responsible for driving the rover. It
// is expected to be started in a goroutine
func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	prevDirection := image.Point{X: 1, Y: 0}
	updateInterval := time.Microsecond * 250
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.Y,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			case stop:
				prevDirection = direction
				direction = image.Point{
					X: 0,
					Y: 0,
				}
				log.Printf("rover stopped")
			case start:
				log.Printf("starting the rover")
				direction = prevDirection
			}
			log.Printf("new direction %v", direction)
		case <-nextMove:
			pos = pos.Add(direction)
			log.Printf("moved to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

// turns the rover left (90 deg counterclockwise)
func (r *RoverDriver) Left() {
	r.commandc <- left
}

// turns the rover right (90 deg clockwise)
func (r *RoverDriver) Right() {
	r.commandc <- right
}

func (r *RoverDriver) Start() {
	r.commandc <- start
}

func (r *RoverDriver) Stop() {
	r.commandc <- stop
}

func NewRowerDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

func main() {
	fmt.Printf("Toy program. It operates Mars rover.\n\n")
	r := NewRowerDriver()
	time.Sleep(time.Second * 1)
	r.Left()
	time.Sleep(time.Second * 1)
	r.Right()
	time.Sleep(time.Second * 1)
	r.Stop()
	time.Sleep(time.Second * 1)
	r.Start()
	time.Sleep(time.Second * 1)
	fmt.Println("\nEnd of program")
}
