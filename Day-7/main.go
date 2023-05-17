package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	a := int32(len(arrTemp))
	if a > n {
		panic("out of array length")
	}

	for i, j := 0, len(arrTemp)-1; i < j; i, j = i+1, j-1 {
		arrTemp[i], arrTemp[j] = arrTemp[j], arrTemp[i]
	}

	fmt.Println(strings.Join(arrTemp, " "))
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
