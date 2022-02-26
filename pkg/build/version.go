package build

import (
	"encoding/json"
	"os"
)

type PackageJSON struct {
	Version string `json:"version"`
}

// Opens the package.json file in the provided directory and returns a struct that represents its contents
func OpenPackageJSON(dir string) (PackageJSON, error) {
	reader, err := os.Open("package.json")
	if err != nil {
		return PackageJSON{}, err
	}

	defer logAndClose(reader)

	jsonObj := PackageJSON{}
	if err := json.NewDecoder(reader).Decode(&jsonObj); err != nil {
		return PackageJSON{}, err
	}

	return jsonObj, nil
}
