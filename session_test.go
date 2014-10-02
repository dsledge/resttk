package resttk

import (
	"testing"
)

func TestSessionAdd(t *testing.T) {
	key := "abcd1234"
	sessionStore := NewSessionMemoryStore()
	sessionStore.Add(key, "efgh5678")
	_, ok := sessionStore.Get(key)
	if !ok {
		t.Fail()
	}
}

func TestSessionDelete(t *testing.T) {
	t.Skip()
}
