package gpx_test

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/progressmate/gpx"
)

//go:embed testdata/valid_1.1.gpx
var validFile []byte

func TestUnmarshal(t *testing.T) {
	t.Parallel()

	_, err := gpx.Unmarshal(validFile)
	if err != nil {
		t.Fatal(err)
		return
	}
}

func TestDecode(t *testing.T) {
	t.Parallel()

	reader := bytes.NewReader(validFile)
	_, err := gpx.Decode(reader)
	if err != nil {
		t.Fatal(err)
		return
	}
}

func TestReadFile(t *testing.T) {
	t.Parallel()

	t.Run("file exists", func(t *testing.T) {
		_, err := gpx.ReadFile("testdata/valid_1.1.gpx")
		if err != nil {
			t.Fatal(err)
			return
		}
	})

	t.Run("file does not exist", func(t *testing.T) {
		_, err := gpx.ReadFile("testdata/does-not-exist.gpx")
		if err == nil {
			t.Fatal("expected error")
			return
		}
	})
}
