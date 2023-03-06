package sort

import "testing"
import "gihub.com/Toor3-14/algorithms/abc"

func TestBubleCustom(t *testing.T) {
	
	// First test with default alphabet

	cases := []struct {
		in, want []string
	}{
		{[]string{"Andrey", "Charles", "Johne", "Chris", "Alex", "Nocola"}, []string{"Andrey", "Alex", "Charles", "Chris", "Johne", "Nocola"}},
		{[]string{"keller", "Jaque", "Phill", "bob", "Siegmut", "scharlie"}, []string{"Jaque", "Phill", "Siegmut","bob", "keller", "scharlie"}},
	} 
	for _, c := range cases {
		BubleCustom(c.in)
		for i := 0; i < len(c.in); i++ {
			if c.in[i] != c.want[i] {
				t.Errorf("BubleCustom() == %v, want %v", c.in, c.want)
				return
			}
		}
	}

	// Second test with redact alphabet

	in := []string{"Aandrey", "Aavitya", "Ivan", "Zeus", "Akolya"}
	want := []string{"Zeus", "Aandrey", "Aavitya", "Ivan", "Akolya"}
	abc.ABC = []string{"Z","B","C","D"}

	BubleCustom(in)

	for i := 0; i < len(in); i++ {
		if in[i] != want[i] {
			t.Errorf("BubleCustom() == %v, want %v", in, want)
		}
	}
}

func TestBuble(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{3,6,9,8,123,6,354,4,1,7,3}, []int{1,3,3,4,6,6,7,8,9,123,354}},
		{[]int{-1,654,42,0,-543}, []int{-543,-1,0,42,654}},
		{[]int{0,9,8,7,6,5,4,3,2,1}, []int{0,1,2,3,4,5,6,7,8,9}},
	}
	for _, c := range cases {
		Buble(c.in)
		for i := 0; i < len(c.in); i++ {
			if c.in[i] != c.want[i] {
				t.Errorf("Buble() == %v, want %v", c.in, c.want)
			}
		}
	}
	in := []string{"Andrey", "Charles", "Johne", "Chris", "Alex", "Nocola"}
	want := []string{"Alex", "Andrey", "Charles", "Chris", "Johne", "Nocola"}

	Buble(in)

	for i := 0; i < len(in); i++ {
		if in[i] != want[i] {
			t.Errorf("Buble() == %v, want %v", in, want)
			return
		}
	}
}