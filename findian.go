package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MainFindian() {
	var ianlover string
	fmt.Println("String to look for ian")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		ianlover = scanner.Text()
	}
	ianlover = strings.ToLower(ianlover)
	if strings.HasPrefix(ianlover, "i") && strings.Contains(ianlover, "a") && strings.HasSuffix(ianlover, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
