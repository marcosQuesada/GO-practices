package main

import (
	"fmt"
	"testing"
)

func TestConcept(t *testing.T) {
	b := Init("foo")
	b.Handle("bar")
	fmt.Println("Hooo")

	Floor(b)

	f := &FooBroker{
		name:  "fooo",
		field: 10,
	}

	f.Handle("hiiii")
	Floor(f)

	c := &Composition{
		Broker: f,
		bar:    "hi there",
	}

	c.Run("run forest")
}
