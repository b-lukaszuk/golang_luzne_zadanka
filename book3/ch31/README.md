# Book 3

Selected bits/exercises from "Get Programming with Go" by Youngman N. and Peppe R.

The snippets/programs written at a whim, sometimes modified, and used for educational purposes.

# Lesson 31. Concurrent State

## Task 1

positionworker.go

[...] change the code so that the delay time gets a half a second longer with
each move.

The code:

<pre>
func worker() {
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}
	next := time.After(time.Second)
	for {
		select {
		case <-next:
			pos = pos.Add(direction)
			fmt.Println("current position is ", pos)
			next = time.After(time.Second)
		}
	}
}
</pre>

## Task 2

rover.go

Using the RoverDriver type [...], define `Start` and `Stop` methods and
associated commands to make the rover obey them.

## Caution note

**The content of this folder may be incorrect, erroneous and/or harmful. Use it at Your own risk.**
