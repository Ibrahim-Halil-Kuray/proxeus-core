package validate

import (
	"bytes"
	"testing"
)

func TestStream(t *testing.T) {
	rf := bytes.NewBuffer([]byte(`
	Content...................................
.................................................
........................................................
...............................................................
...............................................................
...............................................................
...............................................................
...............................................................
...............................................................
...............................................................
...............................................................
`))
	rules := Rules{
		"file": Rules{
			"exact": true,
			"kind":  "pdf",
		},
		//"min": "220kb",
		//"max": "300kb",
	}
	_, err := File(rf, rules)
	if err == nil {
		t.Error("expected wrong file type")
	}
}
