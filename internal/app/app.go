package app

const Prefix = "http://"

func DoMaskForLink(input string) string {
	search := []byte(input)
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
