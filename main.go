package main

import "fmt"

func main() {
	fmt.Println(initialValidation("{\"key\":12}"))
	fmt.Println(initialValidation("{\"key:12}"))
	fmt.Println(initialValidation("{key:12\"}"))
}

func validateJsonObject(index *int, jsonString *string) bool {
	if !parenthesisValidation(index, jsonString) {
		return false
	}
	duplicateKey := map[string]bool{}
	prevIndex := 0
	for true {
		skipSpaces(index, jsonString)
		prevIndex = *index
		if !stringValidation(index, jsonString) {
			return false
		}
		// Used to validate duplicate key.
		if _, ok := duplicateKey[(*jsonString)[prevIndex:*index]]; ok {
			return false
		}
		duplicateKey[(*jsonString)[prevIndex:*index]] = true

		skipSpaces(index, jsonString)
		if !charValidation(index, jsonString, ':') {
			return false
		}
		skipSpaces(index, jsonString)
		if !valueValidation(index, jsonString) {
			return false
		}
		skipSpaces(index, jsonString)
		if !charValidation(index, jsonString, ',') {
			break
		}
	}

	skipSpaces(index, jsonString)
	if !charValidation(index, jsonString, '}') {
		return false
	}

	return true
}

func initialValidation(jsonString string) bool {
	index := 0
	return validateJsonObject(&index, &jsonString)
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
	// We need to move string after " so that it can match next value
	*index++
	return true
}

func charValidation(index *int, jsonString *string, char uint8) bool {
	if *index >= len(*jsonString) {
		return false
	}

	if (*jsonString)[*index] != char {
		return false
	}
	*index++
	return true
}

func valueValidation(index *int, jsonString *string) bool {
	if *index >= len(*jsonString) {
		return false
	}
	if (*jsonString)[*index] == '"' {
		return stringValidation(index, jsonString)
	}

	if (*jsonString)[*index] == 't' || (*jsonString)[*index] == 'f' {
		return boolValidation(index, jsonString)
	}

	if (*jsonString)[*index] == 'n' {
		return checkAgainstString(index, jsonString, "null")
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
			return false
		}
		*index++
	}

	return true
}

func checkNumber(index *int, jsonString *string) bool {
	if *index >= len(*jsonString) {
		return false
	}
	if val := int((*jsonString)[*index]); !(int('0') <= val && val <= int('9')) {
		return false
	}
	for *index < len(*jsonString) {
		if val := int((*jsonString)[*index]); int('0') <= val && val <= int('9') {
			break
		}
		*index++
	}

	return true
}

func skipSpaces(index *int, jsonString *string) {
	if *index >= len(*jsonString) {
		return
	}
	for (*jsonString)[*index] == uint8(' ') {
		*index += 1
		if *index >= len(*jsonString) || (*jsonString)[*index] != uint8(' ') {
			return
		}
	}
}
