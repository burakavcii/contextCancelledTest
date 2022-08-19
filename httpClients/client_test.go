package httpClients

import "testing"

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

func Test_HTTPContextClient_CallCancelAfter(t *testing.T) {
	t.Log("TestCancelAfter")
	client := NewHTTPContextClient()

	err := client.CallCancelAfter()
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
}
