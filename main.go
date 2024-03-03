package main

import (
	"fmt"
	"strconv"
)

func main() {
	p := "a"
	i := "a"

	page, _ := strconv.Atoi(p)
	items, _ := strconv.Atoi(i)

	fmt.Println(page)
	fmt.Println(items)

}
