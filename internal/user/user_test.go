package user_test

import (
	"flag"
	"testing"
)

var testflag = flag.Bool("testa", true, "fortesting")

func TestService_FindAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		if !*testflag {
			t.Fatalf("error gan")
		}
	})
}
