package pkg

import "testing"

func TestReadPrintFile(t *testing.T) {
	if got := ReadPrintFile("./file.txt"); got != nil {
		t.Errorf("Hello() = %v, want %v", got, nil)
	}
}
