package main

import (
	"testing"
)

func contains(s []interface{}, e interface{}) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func TestMultiplex(t *testing.T) {
	input1 := make(chan interface{})
	input2 := make(chan interface{})
	input3 := make(chan interface{})

	output := Multiplex(input1, input2, input3)

	// Отправляем значения в случайном порядке во все входные каналы
	go func() {
		defer close(input1)
		defer close(input2)
		defer close(input3)
		input1 <- "Hello"
		input2 <- 42
		input3 <- true
		input2 <- "World"
		input3 <- 3.14
		input1 <- false
		input3 <- 10
		input1 <- "School21"
		input2 <- 1.618
	}()

	// Ожидаем получения всех значений на выходном канале
	var received []interface{}
	for i := 0; i < 9; i++ {
		msg := <-output
		received = append(received, msg)
	}

	expected := []interface{}{"Hello", 42, true, "World", 3.14, false, 10, "School21", 1.618}
	for _, elem := range expected {
		if !contains(received, elem) {
			t.Errorf("Received does not contain %v", elem)
		}
	}
}
