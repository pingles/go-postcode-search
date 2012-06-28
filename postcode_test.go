package postcode

import "testing"

func TestDoesSomething(t *testing.T) {
  if myFunc() != "Paul" {
    t.Error("Name didn't match")
  }
}