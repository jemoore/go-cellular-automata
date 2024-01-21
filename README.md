# Description

This code demonstrates two types of Celluar Automata:
* Conway's Game of Life
* Simple Automata 

## Usage
### Conway's Game of Life
```
go run main.go
```
### Simple Automata
The *-simple* flag is required for the Simple Automata 'game' to run.
Optionally, a rule (0-255) can be specified which will determine the output.
```
go run main.go -simple -rule 147
```

# Notes
I created this as a way to learn Golang.  After creating the basic functionality, I began looking for ways to factor out some of the common code and functionality.  In many languages this would be done with classes and polymorphism.  However, Go does not have the concept of classes.  I used interfaces and composition to refactor the code.

To learn about packages, I divided the code up into several packages.

So if the code looks a bit complicated and over engineered, please remember I was using this as more of a learning exercise.