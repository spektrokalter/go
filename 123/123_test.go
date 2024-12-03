package main_test

import "testing"

func Test(t *testing.T) {
	f := func(t *testing.T, name string) {
		t.Helper()

		t.Run(name, func(t *testing.T) {
			t.Helper()
		})
	}
}
