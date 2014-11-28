package registry

import (
	"testing"
)

func TestSetGetEntryOnRegister(t *testing.T) {
	Start()
	Set("foo", "bar")
	value := Get("foo")
	if value != "bar" {
		t.Fail()
	}
}
