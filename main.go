package main

import "fmt"

const Prefix = "http://"

func main() {
	link := "Hello, its my page //string//: http://localhost123.com See you http://localhost123.com "
	fmt.Println(doMaskForLink(link))

}

func doMaskForLink(s string) string {
	search := []byte(s)
	prefix := []byte(Prefix)
	prefixLength := len(prefix)
	i := 0
	for i < len(search) {
		if containsPrefix(i, search) {
			i += prefixLength
			for i < len(search) && !searchForSpace(search[i]) {
				search[i] = byte('*')
				i++
			}
		}
		i++
	}
	return string(search)
}

func containsPrefix(idx int, arrayByte []byte) bool {
	prefix := []byte(Prefix)
	prefixLength := len(prefix)
	if idx+prefixLength > len(arrayByte) {
		return false
	}
	for i := 0; i < prefixLength; i++ {
		if arrayByte[idx+i] != prefix[i] {
			return false
		}
	}
	return true
}

func searchForSpace(searchByte byte) bool {
	return searchByte == ' ' || searchByte == '\n' || searchByte == '\t' || searchByte == ',' || searchByte == '!'
}
