//add_test.go
package simplemath

import "testing"

func TestAdd1(t *testing.T){
	r := Add(1, 2)
	if r != 3 {
		t.Errorf("Add(1, 2) faild. Got %d, except 3.", r)
	}
	else {
		t.Errorf("successed %d", r)
	}
}