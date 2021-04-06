package datastructure

import (
	"fmt"
	"testing"
)

func TestTries(t *testing.T) {
	tries := InitTries()
	fmt.Println(tries)
	toAdd := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}
	for i := range toAdd {
		tries.Insert(toAdd[i])
	}

	fmt.Println(tries.Search("aragorn"))
}
