package lru

import (
	"errors"
	"time"
)

type lru struct {
	index map[string]*node
	ttl   int
	first *node
	last  *node
}

type node struct {
	id    string
	ts    int64
	value int
	prev  *node
	next  *node
}

func Init(ttl int) *lru {
	return &lru{
		index: make(map[string]*node, 0),
		ttl:   ttl,
	}
}

func (l *lru) Add(id string, value int) {
	n := &node{
		id:    id,
		value: value,
		prev:  l.first,
		ts:    time.Now().Unix(),
	}

	if l.last == nil {
		l.last = n
	}

	if l.first != nil {
		l.first.next = n
	}

	l.index[id] = n
	l.first = n
}

func (l *lru) Get(id string) (*node, error) {
	if v, ok := l.index[id]; ok {
		return v, nil
	}

	return nil, errors.New("not found")
}

func (l *lru) Delete(id string) error {
	n, err := l.Get(id)
	if err != nil {
		return err
	}
	if l.first.id == id {
		l.first = n.prev
	}
	if l.last.id == id {
		l.last = n.next
	}

	before := n.prev
	after := n.next

	if before != nil {
		before.next = after
	}
	if after != nil {
		after.prev = before
	}
	delete(l.index, id)

	return nil
}

func (l *lru) expire() {
	done := false
	node := l.last
	for !done {
		if node.ts+int64(l.ttl) <= time.Now().Unix() {
			if node.next == nil {
				done = true
			}
			nextNode := node.next
			l.Delete(node.id)
			node = nextNode
		} else {
			done = true
		}
	}
}
