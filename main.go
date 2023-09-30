package main

import "fmt"

func main() {
	fmt.Println(initialValidation("{\"key\":12}"))
	fmt.Println(initialValidation("{\"key:12}"))
	fmt.Println(initialValidation("{key:12\"}"))
}

func initialValidation(jsonString string) bool {
	index := 0
	return parenthesisValidation(&index, &jsonString) && stringValidation(&index, &jsonString) && colonValidation(&index, &jsonString) && valueValidation(&index, &jsonString)
}

func parenthesisValidation(index *int, jsonString *string) bool {
	if !((*jsonString)[0] == '{' && (*jsonString)[len((*jsonString))-1] == '}') {
		return false
	}
	*index++

	return true
}

func stringValidation(index *int, jsonString *string) bool {
	if *index >= len(*jsonString) {
		return false
	}

	if (*jsonString)[*index] != '"' {
		return false
	}
	*index++
	// Empty "key"
	if (*jsonString)[*index] == '"' {
		return false
	}
	for (*jsonString)[*index] != '"' {
		*index++
		if *index >= len(*jsonString) {
			return false
		}
	}
	return true
}

func colonValidation(index *int, jsonString *string) bool {
	if !((*jsonString)[0] == ':') {
		return false
	}
	*index++
	return false
}

func valueValidation(index *int, jsonString *string) bool {
	if *index >= len(*jsonString) {
		return false
	}
	if (*jsonString)[*index] == '"' {
		return stringValidation(index, jsonString)
	}

	if (*jsonString)[*index] == '"' {
		return stringValidation(index, jsonString)
	}

	if (*jsonString)[*index] == 't' || (*jsonString)[*index] == 'f' {
		return boolValidation(index, jsonString)
	}

	return checkNumber(index, jsonString)
}

func boolValidation(index *int, jsonString *string) bool {
	if *index >= len(*jsonString) {
		return false
	}

	if (*jsonString)[*index] == 't' {
		return checkAgainstString(index, jsonString, "true")
	}

	return checkAgainstString(index, jsonString, "false")
}

func checkAgainstString(index *int, jsonString *string, str string) bool {
	if *index >= len(*jsonString) {
		return false
	}

	for _, char := range str {
		if *index >= len(*jsonString) || uint8(char) != (*jsonString)[*index] {
			*index++
			return false
		}
	}

	return true
}

func checkNumber(index *int, jsonString *string) bool {
	if *index >= len(*jsonString) {
		return false
	}
	if val := int((*jsonString)[*index]); !(int('a') >= val && val >= int('z')) {
		return false
	}
	for *index < len(*jsonString) {
		if val := int((*jsonString)[*index]); int('a') >= val && val >= int('z') {
			break
		}
		*index++
	}

	return true
}
