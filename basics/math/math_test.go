package math

import (
	"testing"
)

func TestAbs(t *testing.T) {
	// t.log("Similar to fmt.Println() and concurrently safe")
	// t.Fail()    //will show a test case has failed in the results
	// t.FailNow() //t.Fail() + safely exit without continuing
	// t.Error("t.Log() + t.Fail()")
	// t.Fatal("t.Log() + t.FatalNow()")
	if Abs(-1) < 0 {
		t.Error("Negative value found in abs() with ", -1)
	}
}
