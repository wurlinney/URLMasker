package main

import (
	"fmt"
	"urlmasker/internal/app"
)

const Prefix = "http://"

func main() {
	link := "Hello, its my page //string//: http://localhost123.com See you http://localhost123.com "
	fmt.Println(app.DoMaskForLink(link))

}
