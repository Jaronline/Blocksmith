package lib

import (
	"encoding/json"
	"errors"
	"os"
)

const ModpackFile = "modpack.json"
const ModpackFilePerms = 0744

type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type PartialPackage struct {
	Name    *string
	Version *string
}

func (p *Package) Write() error {
	pkgJson, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	if err = os.WriteFile(ModpackFile, pkgJson, ModpackFilePerms); err != nil {
		return err
	}
	return nil
}

func GetDefaultPackage() (*Package, error) {
	var currPkg *PartialPackage
	var err error

	if currPkg, err = GetCurrentPackage(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	var defName string

	if defName, err = getDefaultName(); err != nil {
		return nil, err
	}

	pkg := &Package{
		Name:    defName,
		Version: defaultVersion,
	}

	if currPkg != nil {
		if currPkg.Name != nil {
			pkg.Name = *currPkg.Name
		}
		if currPkg.Version != nil {
			pkg.Version = *currPkg.Version
		}
	}

	return pkg, nil
}

func GetCurrentPackage() (*PartialPackage, error) {
	pkgJson, err := os.ReadFile(ModpackFile)
	if err != nil {
		return nil, err
	}
	var pkg *PartialPackage
	if err = json.Unmarshal(pkgJson, &pkg); err != nil {
		return nil, err
	}
	return pkg, nil
}
