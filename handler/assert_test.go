package handler

import "testing"

func assertHttpStatus(t testing.TB, statusExpected, statusActual int) {
	t.Helper()
	if statusExpected != statusActual {
		t.Errorf("Expected %v actual %v ", statusExpected, statusActual)
	}
}
