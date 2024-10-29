# Book 3

Selected bits/exercises from "Get Programming with Go" by Youngman N. and Peppe R.

The snippets/programs written at a whim, sometimes modified, and used for educational purposes.

# Lesson 30. Goroutines and Concurrency

## Task 1

remove-identical.go

Write a pipeline [...] (a goroutine) that remembers the previous value and only
sends the value to the next stage of the pipeline if it’s different from the one
that came before. [...] assume that the first value is never the empty string.

## Task 2

split-words.go

Write a pipeline element that takes strings, splits them up into words (you can
use the `Fields` function from the `strings` package), and sends all the words,
one by one, to the next pipeline stage.

## Caution note

**The content of this folder may be incorrect, erroneous and/or harmful. Use it at Your own risk.**
