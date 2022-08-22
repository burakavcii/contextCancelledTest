package httpClients

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_HTTPContextClient_CallCancelBefore(t *testing.T) {
	t.Log("TestCancelBefore")
	client := NewHTTPContextClient()

	err := client.CallCancelBefore()
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}

func Test_HTTPContextClient_CallCancelBefore2(t *testing.T) {
	t.Log("TestCancelBefore")
	client := NewHTTPContextClient()

	err := client.CallCancelBefore2()
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}

func Test_HTTPContextClient_CallCancelBefore3(t *testing.T) {
	t.Log("TestCancelBefore")
	client := NewHTTPContextClient3(URL)

	err := client.CallCancelBefore3()
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}

func Test_HTTPContextClient_CallCancelBefore4(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 15)
	}))
	defer ts.Close()
	t.Log("TestCancelBefore")
	client := NewHTTPContextClient3(ts.URL)

	err := client.CallCancelBefore3()
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}

func Test_HTTPContextClient_CallCancelAfter(t *testing.T) {
	t.Log("TestCancelAfter")
	client := NewHTTPContextClient()

	err := client.CallCancelAfter()
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}
