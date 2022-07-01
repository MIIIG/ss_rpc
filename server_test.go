package main

import "testing"

func TestHandler(t *testing.T) {
	expectedHandlerResult := "somthng"
	if actualHandlerResult := Handler(mock1, mock2); actualHandlerResult != expectedHandlerResult {
		t.Errorf("expected %s; got: %s", expectedHandlerResult, &actualHandlerResult)
	}
}

// // файл foo_test.go
// package foo

// import (
//     "testing"
// )

// func TestFooFunc(t *testing.T) {
//     expectedFooResult := "bar"
//     if actualFooResult := Foo(); actualFooResult != expectedFooResult {
//         t.Errorf("expected %s; got: %s", expectedFooResult, actualFooResult)
//     }
// }
