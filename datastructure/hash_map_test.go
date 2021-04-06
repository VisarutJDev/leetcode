package datastructure

import (
	"fmt"
	"testing"
)

func TestHashMap(t *testing.T) {
	testHashTable := Init()
	fmt.Println(testHashTable)
	list := []string{
		"ERIC",
		"KENNY",
		"RANDY",
		"BUTTERS",
		"STAN",
		"TOKEN",
	}

	for _, v := range list {
		testHashTable.Insert(v)
	}

	testHashTable.Delete("STAN")

	fmt.Println(testHashTable.Search("TOKEN"))
}
