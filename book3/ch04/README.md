# Book 3

Selected bits/exercises from "Get Programming with Go" by Youngman N. and Peppe R.

The snippets/programs written at a whim, sometimes modified, and used for educational purposes.

# Lesson 04. Variable Scope.

## Task 1

Modify listing 4.7 to handle leap years.

```go
var era = "AD"

func main() {
    year := 2018
    month := rand.Intn(12)+1
    daysInMonth := 31

    switch month {
    case 2:
        daysInMonth = 28
    case 4, 6, 9, 11:
        daysInMonth = 30
    }
    day := rand.Intn(daysInMonth) + 1
    fmt.Println(era, year, month, day)
}
```

## Caution note

**The content of this folder may be incorrect, erroneous and/or harmful. Use it at Your own risk.**
