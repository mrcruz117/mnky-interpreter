package object

import (
	"testing"
)

func TestStringHashkey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is Michael"}
	diff2 := &String{Value: "My name is Michael"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys: %v != %v", hello1.HashKey(), hello2.HashKey())
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys: %v != %v", diff1.HashKey(), diff2.HashKey())
	}
	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have the same hash keys: %v == %v", hello1.HashKey(), diff1.HashKey())
	}

}
