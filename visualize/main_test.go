package main

import (
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("TestMain")
	setup()
	defer close()

	m.Run()
}
