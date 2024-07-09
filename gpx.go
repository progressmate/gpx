//go:generate xsdgen -o types.go -f -pkg gpx -r "gpx -> GPX" -r "Type -> " schema/1.1.xsd

// Package gpx provides parsing functionality for GPX 1.1 files into go types. See https://www.topografix.com/gpx.asp for
// more information on the GPX file format.
package gpx

import (
	"encoding/xml"
	"io"
	"os"
)

// Decode the contents of the io.Reader into a GPX type. The reader expects a valid GPX 1.1 file in XML format.
func Decode(r io.Reader) (GPX, error) {
	var gpx GPX
	err := xml.NewDecoder(r).Decode(&gpx)
	return gpx, err
}

// Unmarshal the provided bytes into a GPX type. The bytes provided should be a valid GPX 1.1 file in XML format.
func Unmarshal(data []byte) (GPX, error) {
	var gpx GPX
	err := xml.Unmarshal(data, &gpx)
	return gpx, err
}

// ReadFile attempts to decode the contents of a file on-disk into a GPX type. The provided path is expected to lead to
// a valid GPX 1.1 file in XML format.
func ReadFile(path string) (GPX, error) {
	f, err := os.Open(path)
	if err != nil {
		return GPX{}, err
	}

	defer f.Close()
	return Decode(f)
}
