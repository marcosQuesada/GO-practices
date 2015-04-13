package lru

import (
	"testing"
	"time"
)

func TestToAddElementsOnLRU(t *testing.T) {
	lru := Init(1)

	lru.Add("a", 10)
	lru.Add("b", 1000)
	lru.Add("c", 100)

	if len(lru.index) != 3 {
		t.Fail()
	}

	b := lru.index["b"]
	if b.value != 1000 {
		t.Fail()
	}

	if lru.first.id != "c" {
		t.Error("First is not C", lru.first.id)
	}

	if lru.last.id != "a" {
		t.Error("First is not A", lru.last.id)
	}
}

func TestToAccessValuesOnLRU(t *testing.T) {
	lru := Init(1)

	lru.Add("a", 10)
	lru.Add("b", 1000)
	lru.Add("c", 100)

	b, err := lru.Get("b")

	if b.value != 1000 && err != nil {
		t.Fail()
	}

	if lru.first.id != "c" {
		t.Error("Bad First Node", lru.first.id)
	}

	if lru.last.id != "a" {
		t.Error("Bad Last Node", lru.last.id)
	}

	if b.next.id != "c" {
		t.Error("Node has not expected next ", b.next)
	}

	if b.prev.id != "a" {
		t.Fail()
	}
}

func TestDeleteElements(t *testing.T) {
	lru := Init(1)
	lru.Add("a", 10)
	lru.Add("b", 1000)
	lru.Add("c", 100)
	lru.Add("d", 10000)

	err := lru.Delete("c")
	if err != nil {
		t.Error("Error deleting D", err)
	}

	if len(lru.index) != 3 {
		t.Error("Index is not empty", lru.index)
	}

	if lru.first.id != "d" {
		t.Error("Bad First Node", lru.first.id)
	}

	if lru.last.id != "a" {
		t.Error("Bad Last Node", lru.last.id)
	}

	v, err := lru.Get("c")
	if err == nil {
		t.Error("Element not removed from index", v)
	}

	b, _ := lru.Get("b")
	if b.next.id != "d" {
		t.Error("Bad next", b.next.id)
	}

	d, _ := lru.Get("d")
	if d.prev.id != "b" {
		t.Error("Bad Prev", d.prev.id)
	}
}

func TestDeleteTopElement(t *testing.T) {
	lru := Init(1)
	lru.Add("a", 10)
	lru.Add("b", 1000)
	lru.Add("c", 100)
	lru.Add("d", 10000)

	err := lru.Delete("d")
	if err != nil {
		t.Error("Error deleting D", err)
	}

	if len(lru.index) != 3 {
		t.Error("Index is not empty", lru.index)
	}

	if lru.first.id != "c" {
		t.Error("Bad First Node", lru.first.id)
	}

	if lru.last.id != "a" {
		t.Error("Bad Last Node", lru.last.id)
	}

	v, err := lru.Get("d")
	if err == nil {
		t.Error("Element not removed from index", v)
	}

	if lru.first.id != "c" {
		t.Error("Bad first lru element", lru.first.id)
	}
}

func TestDeleteLastElement(t *testing.T) {
	lru := Init(1)
	lru.Add("a", 10)
	lru.Add("b", 1000)
	lru.Add("c", 100)
	lru.Add("d", 10000)

	err := lru.Delete("a")
	if err != nil {
		t.Error("Error deleting D", err)
	}

	if lru.first.id != "d" {
		t.Error("Bad First Node", lru.first.id)
	}

	if lru.last.id != "b" {
		t.Error("Bad Last Node", lru.last.id)
	}
}

func TestExpireNodes(t *testing.T) {
	lru := Init(1)

	lru.Add("a", 10)
	lru.Add("b", 1000)
	lru.Add("c", 100)
	time.Sleep(time.Second * time.Duration(2))
	lru.expire()

	if len(lru.index) != 0 {
		t.Error("Index is not empty", lru.index)
	}
}
