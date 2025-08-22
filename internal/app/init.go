package app

import (
	"encoding/json"
	"os"
)

func CreatePackage(pkg *Package) {
	pkgJson, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		os.Exit(1)
	}
	err = os.WriteFile("modpack.json", pkgJson, 0744)
	if err != nil {
		os.Exit(1)
	}
}
