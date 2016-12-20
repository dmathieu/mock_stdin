package safebuffer

import (
	"io/ioutil"
	"testing"
)

func TestReadWrite(t *testing.T) {
	mock := NewMock()

	_, err := mock.Write([]byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}

	c, err := ioutil.ReadAll(mock)
	if err != nil {
		t.Fatal(err)
	}
	if string(c) != "hello world" {
		t.Fatalf("Unexpected read content. Got %q, expected `hello world`", c)
	}
}

func TestConcurrentReadWrite(t *testing.T) {
	mock := NewMock()

	go func() {
		c, err := ioutil.ReadAll(mock)
		if err != nil {
			t.Fatal(err)
		}
		if string(c) != "hello world" {
			t.Fatalf("Unexpected read content. Got %q, expected `hello world`", c)
		}
	}()

	_, err := mock.Write([]byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}
}
