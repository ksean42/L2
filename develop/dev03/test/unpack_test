package unpack

import (
	"errors"
	// "fmt"
	"log"
	"testing"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

type testData struct {
	input string
	want  string
}

type errorTestData struct {
	input string
	want  string
	err   error
}

var errInvalidString = errors.New("некорректная строка")

var trueTest = []testData{
	{"a4bc2d5e", "aaaabccddddde"},
	{"a15bc2d5e", "aaaaaaaaaaaaaaabccddddde"},
	{"abcd", "abcd"},
	{"ab\\55cd", "ab55555cd"},
	{"qwe\\4\\5", "qwe45"},
	{"qwe\\45", "qwe44444"},
	{"qwe\\\\5", "qwe\\\\\\\\\\"},
}

var errorTest = []errorTestData{
	{"45", "", errInvalidString},
	{"", "", nil},
	{"45asassda", "", errInvalidString},
	{"a4bc2d5e", "aaaabccddddde", nil},
}

func TestUnpack(t *testing.T) {
	var expected string
	for _, v := range trueTest {
		expected = v.want
		if actual, err := unpack(v.input); err != nil {
			t.Fatalf("%s Error: %s\n", failed, err)
		} else if actual != expected {
			t.Fatalf("%s Wrong answer\n Expected: %15s | Actual: %15s\n", failed, expected, actual)
		} else {
			log.Printf("%s Input: %15s | Expected: %15s | Actual: %15s\n", success, v.input, expected, actual)
		}
	}
}

func TestErrorUnpack(t *testing.T) {

	for _, v := range errorTest {
		expected := v.want
		errExpected := v.err
		actual, err := unpack(v.input)
		if err != nil && err.Error() == errExpected.Error() && expected == actual {
			log.Printf("%s Input: %15s | Expected: %15s | Actual: %15s | ErrorExpected: %s | ErrorActual %s\n", success, v.input, expected, actual, errExpected, err)
		} else if err != nil && (expected != actual || err.Error() != errExpected.Error()) {
			t.Fatalf("%s Wrong answer Input: %15s | Expected: %15s | Actual: %15s\nErrorExpected: %s | ErrorActual %s\n", failed, v.input, expected, actual, errExpected, err)
		}
	}
}
