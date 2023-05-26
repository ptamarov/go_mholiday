package hello

import (
	"testing"
)

func TestSayHello(t *testing.T) {

	substests := []struct {
		items  []string
		result string
	}{
		{result: "Hello, world!"},
		{items: []string{"Matt"},
			result: "Hello, Matt!"},
	}

	for _, st := range substests {
		if s := Say(st.items); s != st.result {
			t.Errorf("Wanted %s (%v) but got %s", st.result, st.items, s)
			// %v variable format
			// %s string format
		}
	}

}
