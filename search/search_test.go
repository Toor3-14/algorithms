package search

import "testing"
import "gihub.com/Toor3-14/algorithms/abc"

func TestBinaryCustom(t *testing.T) {
	
	// First test with default alphabet

	cases := []struct {
		in []string
	}{
		{[]string{"Bob", "Jaque", "Jack", "Chris", "Johne", "Nocola"}},
		{[]string{"Bob", "Jaque", "keller", "Phill", "scharlie", "Siegmut"}},
	}
	for _, c := range cases {
		got := BinaryCustom(c.in,"Bob")
		if got != 0 {
			t.Errorf("BinaryCustom(%v,\"Bob\") == %d, want %d", c.in, got, 0)
		}
	}

	// Second test with redact alphabet

	in := []string{"Zeus", "Ivan", "Aandrey", "Aavitya", "Akolya"}
	want := 1
	abc.ABC = []string{"Z","I","A"}

	got := BinaryCustom(in,"Ivan")

	if got != want {
		t.Errorf("BinaryCustom(%v,\"Ivan\") == %d, want %d", in, got, want)
	}
}

func TestBinary(t *testing.T) {

	// number test 

	cases := []struct {
		in, num ,want []int
	}{
		{[]int{1,3,3,4,6,6,7,8,9,123,354}, []int{9}, []int{8}},
		{[]int{-543,-1,0,42,654}, []int{654}, []int{4}},
		{[]int{0,1,2,3,4,5,6,7,8,9}, []int{0}, []int{0}},
	}
	for _, c := range cases {
		got := Binary(c.in,c.num[0])
		if got != c.want[0] {
			t.Errorf("Binary(%v,%v) == %d, want %d", c.in, c.num[0], got, c.want[0])
		}
	}

	// string test 

	in := []string{"Andrey", "Alex", "Charles", "Chris", "Johne", "Nocola"}
	want := 5
	got := Binary(in, "Nocola")

	if got != want {
		t.Errorf("Binary(%v,\"Nocola\") == %d, want %d", in, got, want)
	}
}