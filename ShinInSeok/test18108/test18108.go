package test18108

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test18108() {

	br := bufio.NewReader(os.Stdin)
	str, _ := br.ReadString('\n')

	split := strings.Fields(str)

	first := split[0]

	num1, _ := strconv.Atoi(first)

	fmt.Print(num1 - 543)
}
