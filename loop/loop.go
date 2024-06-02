package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)


	arr := []string{"A", "B", "C"}
	for k, v := range arr {
		fmt.Printf("Key: %v\n", k)
		fmt.Printf("Value: %v\n", v)
	}

	for _, v := range arr {
		fmt.Println(v)
	}

	myMap := map[string]int {
		"a": 10,
		"b": 55,
	}

	for k, v := range myMap {
		fmt.Printf("Key: %v\n", k)
		fmt.Printf("Value: %v\n", v)
	}
}
