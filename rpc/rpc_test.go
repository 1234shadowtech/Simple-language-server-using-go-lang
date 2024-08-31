package rpc_test

import (
	"educationalsp/rpc"
	"testing"
)

type Message struct {
	Message string `json:"message"`
}

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":{\"Testing\":\"hello\"}}"
	actual := rpc.EncodeMessage(EncodingExample(EncodingExample{Testing: true}))
	if expected != actual {
		t.Fatalf("expected: %s, actual: %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
	contentLength := len(content)
	if err != nil {
		t.Fatal(err)
	}
	if contentLength != 15 {
		t.Fatalf("expected :15, got : %d", contentLength)
	}
	if method != "hi" {
		t.Fatalf("expected : 'hi' , got: %s ", method)
	}
}
