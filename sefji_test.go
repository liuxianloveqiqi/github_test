package github_test

import (
	"fmt"
	"testing"
)

func FuzzName(f *testing.F) {
	f.Fuzz(func(t *testing.T) {
		a := 0
		fmt.Println(a)
	})
}
