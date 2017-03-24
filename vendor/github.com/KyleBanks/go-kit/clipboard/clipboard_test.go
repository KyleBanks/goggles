package clipboard

import (
	"strings"
	"testing"
)

func Test_Read(t *testing.T) {
	_, err := Read()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_ReadString(t *testing.T) {
	_, err := ReadString()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Write(t *testing.T) {
	expect := "Test_Write"

	if err := Write(strings.NewReader(expect)); err != nil {
		t.Fatal(err)
	}

	out, err := ReadString()
	if err != nil {
		t.Fatal(err)
	}

	if out != expect {
		t.Fatalf("Unexpected output, expected=%v, got=%v", expect, out)
	}
}

func Test_WriteString(t *testing.T) {
	expect := "Test_WriteString"

	if err := WriteString(expect); err != nil {
		t.Fatal(err)
	}

	out, err := ReadString()
	if err != nil {
		t.Fatal(err)
	}

	if out != expect {
		t.Fatalf("Unexpected output, expected=%v, got=%v", expect, out)
	}
}
