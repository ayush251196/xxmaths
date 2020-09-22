package greetings

import "testing"

func TestGetHelloMessage(t *testing.T) {
	got := GetHelloMessage()
	want := "Hello, world."
	if got != want {
		t.Errorf("got =%q and want=%q",got,want)
	}
}
