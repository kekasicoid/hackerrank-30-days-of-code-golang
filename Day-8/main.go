package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)
	inputs := []string{}

	// Add input to slice of inputs
	for scanner.Scan() {
		input := scanner.Text()
		if len(input) == 0 {
			break
		}

		inputs = append(inputs, input)
	}

	// Pull out the first item from inputs slice which is the number of inputs
	nString := inputs[0]
	n, err := strconv.ParseInt(nString, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	// Use index slicing to extract the inputs of key/values bases on variable n
	nameAndNums := inputs[1 : n+1]
	namesToCheck := inputs[n+1:]
	m := make(map[string]int64, n)

	// Split and separate key/value strings and insert into map
	for _, v := range nameAndNums {
		strArr := strings.Split(v, " ")
		name := strArr[0]
		nums := strArr[1]
		number, err := strconv.ParseInt(nums, 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		m[name] = number
	}

	// Check if key exists and print value if it does according to required format
	for _, name := range namesToCheck {
		num, ok := m[name]
		if !ok {
			fmt.Println("Not found")
		} else {
			fmt.Printf("%v=%v\n", name, num)
		}
	}
}
