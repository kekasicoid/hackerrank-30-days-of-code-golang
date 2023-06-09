package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	// Convert/format int as string representation of binary number
	binary := strconv.FormatInt(nTemp, 2)
	binaryArr := strings.Split(binary, "0")
	var max float64

	for _, val := range binaryArr {
		num, _ := strconv.ParseFloat(val, 64)
		max = math.Max(num, max)
	}

	// Output number of consecutive 1s
	consecutiveOnes := strconv.Itoa(int(max))
	fmt.Println(len(consecutiveOnes))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
