package main

import (
	"testing"
)

func TestCharValidation(t *testing.T) {
	number := ":"
	index := 0
	want := true
	if want != charValidation(&index, &number, ':') {
		t.Errorf("got wrong value wanted %t", want)
	}
	number = "W"
	want = false
	index = 0
	if want != charValidation(&index, &number, ':') {
		t.Errorf("got wrong value wanted %t", want)
	}
	number = ""
	want = false
	index = 0
	if want != charValidation(&index, &number, ':') {
		t.Errorf("got wrong value wanted %t", want)
	}
}

func TestCheckNumber(t *testing.T) {
	number := "123"
	index := 0
	want := true
	if want != checkNumber(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}
	number = "A1223"
	want = false
	index = 0
	if want != checkNumber(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}
	// In Line check this will be failed: number,
	number = "1223A"
	want = true
	index = 0
	if want != checkNumber(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}
}

func TestCheckAgainstString(t *testing.T) {
	number := "123"
	index := 0
	want := true
	if want != checkAgainstString(&index, &number, "123") {
		t.Errorf("got wrong value wanted %t", want)
	}

	number = "1234"
	index = 0
	want = true
	if want != checkAgainstString(&index, &number, "123") {
		t.Errorf("got wrong value wanted %t", want)
	}

	number = "1234"
	index = 0
	want = false
	if want != checkAgainstString(&index, &number, "12345") {
		t.Errorf("got wrong value wanted %t", want)
	}

	number = "1234"
	index = 0
	want = true
	if want != checkAgainstString(&index, &number, "") {
		t.Errorf("got wrong value wanted %t", want)
	}

	number = ""
	index = 0
	want = false
	if want != checkAgainstString(&index, &number, "123") {
		t.Errorf("got wrong value wanted %t", want)
	}
}

func TestBoolValidation(t *testing.T) {
	number := "true"
	index := 0
	want := true
	if want != boolValidation(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}

	number = "false"
	index = 0
	want = true
	if want != boolValidation(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}

	number = "false2"
	index = 0
	want = true
	if want != boolValidation(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}

	number = "true2"
	index = 0
	want = true
	if want != boolValidation(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}

	number = "testNoob"
	index = 0
	want = false
	if want != boolValidation(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}
}

func TestStringValidation(t *testing.T) {
	number := "true"
	index := 0
	want := false
	if want != stringValidation(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}

	number = "\"true"
	index = 0
	want = false
	if want != stringValidation(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}

	number = "\"true\""
	index = 0
	want = true
	if want != stringValidation(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}
}

func TestParenthesisValidation(t *testing.T) {
	number := "{}"
	index := 0
	want := true
	if want != parenthesisValidation(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}
	number = "{n"
	index = 0
	want = false
	if want != parenthesisValidation(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}
}

func TestValueValidation(t *testing.T) {
	number := "\"test\""
	index := 0
	want := true
	if want != valueValidation(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}
	number = "123"
	index = 0
	want = true
	if want != valueValidation(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}

	number = "123A"
	index = 0
	want = true
	if want != valueValidation(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}

	number = "true"
	index = 0
	want = true
	if want != valueValidation(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}

	number = "null"
	index = 0
	want = true
	if want != valueValidation(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}

	number = "Atrue"
	index = 0
	want = false
	if want != valueValidation(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}

	number = "\"test"
	index = 0
	want = false
	if want != valueValidation(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}

	number = "nul"
	index = 0
	want = false
	if want != valueValidation(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}
}

func TestValidateJsonObject(t *testing.T) {
	number := "{\"key1\":true,\"key2\":null,\"key3\":true}"
	index := 0
	want := true
	if want != validateJsonObject(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}

	number = "{ \"key1\" : true , \"key2\"  :null ,\"key3\" : true }"
	index = 0
	want = true
	if want != validateJsonObject(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}

	// Duplicate key
	number = "{ \"key\" : true1 , \"key\"  :null ,\"key\" : true }"
	index = 0
	want = false
	if want != validateJsonObject(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}

	number = "{ \"key1\" : true1 , \"key2\"  :null ,\"ke3y\" : true }"
	index = 0
	want = false
	if want != validateJsonObject(&index, &number) {
		t.Errorf("got wrong value wanted %t", want)
	}
}

func TestSkipSpaces(t *testing.T) {
	number := "   test"
	index := 0
	want := 3
	skipSpaces(&index, &number)
	if want != index {
		t.Errorf("got wrong value wanted %d got index %d", want, index)
	}
}
