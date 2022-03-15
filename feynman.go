package feynman

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

type solution interface{}

var basicTypes = [...]string{
	"int",
	"string",
	"float",
	"interface",
	"boolean",
}

func randomIntegerWithSeed(max int) int {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	num := r.Intn(max)
	return num
}

func getType(t string) interface{} {

	switch t {
	case "int":
		return 42
	case "string":
		return "the answer"
	case "float":
		return 42.0
	case "interface":
		return 42
	case "boolean":
		return true
	default:
		return "still 42!"
	}
}

// Step 1: Write the problem down!
func write(p string) { fmt.Printf("Problem statement: %s\n", p) }

// Step 2: Think really hard about the problem!
func think() solution {
	fmt.Println("Thinking really hard...")
	maxTime := randomIntegerWithSeed(big.MaxExp)

	thinkingTime := time.Duration(maxTime)
	time.Sleep(thinkingTime) // Think for maxTime in nanoseconds!

	n := randomIntegerWithSeed(len(basicTypes))

	s := getType(basicTypes[n])

	return s
}

// Step 3: Solve the problem!
func Solve(p string) {
	write(p)
	ans := think()
	fmt.Printf("The answer to your problem is %v!\n", ans)
}
