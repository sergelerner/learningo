// Go slices
package main

import (
	"fmt"
)

func slices() {
	// Same type
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	// fmt.Println(len(loons)) // 3
	// fmt.Println(loons[1]) // daffy
	// fmt.Println(loons[1:]) // [daffy taz]

	// for
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	// Single value range
	for i := range loons {
		fmt.Println(i)
	}

	// Double value range
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	// Double value range, ignore index by using _
	for _, name := range loons {
		fmt.Println(name)
	}

	// append
	loons = append(loons, "elmer")
	fmt.Println(loons) // [bugs daffy taz elmer]
}
