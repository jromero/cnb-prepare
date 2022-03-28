package project

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"

	"github.com/jromero/cnb-prepare/pkg/project/types"
	v01 "github.com/jromero/cnb-prepare/pkg/project/v01"
	v02 "github.com/jromero/cnb-prepare/pkg/project/v02"
)

type Project struct {
	Version string `toml:"schema-version"`
}

type VersionDescriptor struct {
	Project Project `toml:"_"`
}

var parsers = map[string]func(string) (types.Descriptor, error){
	"0.1": v01.NewDescriptor,
	"0.2": v02.NewDescriptor,
}

func ReadProjectDescriptor(pathToFile string) (*types.Descriptor, error) {
	projectTomlContents, err := ioutil.ReadFile(filepath.Clean(pathToFile))
	if err != nil {
		return nil, err
	}

	var versionDescriptor struct {
		Project struct {
			Version string `toml:"schema-version"`
		} `toml:"_"`
	}

	_, err = toml.Decode(string(projectTomlContents), &versionDescriptor)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing schema version")
	}

	version := versionDescriptor.Project.Version
	if version == "" {
		version = "0.1"
	}

	if _, ok := parsers[version]; !ok {
		return nil, fmt.Errorf("unknown project descriptor schema version %s", version)
	}

	descriptor, err := parsers[version](string(projectTomlContents))
	if err != nil {
		return nil, err
	}

	return &descriptor, validate(descriptor)
}

func validate(p types.Descriptor) error {
	if p.Build.Exclude != nil && p.Build.Include != nil {
		return errors.New("descriptor cannot have both include and exclude defined")
	}

	if len(p.Project.Licenses) > 0 {
		for _, license := range p.Project.Licenses {
			if license.Type == "" && license.URI == "" {
				return errors.New("descriptor must have a type or uri defined for each license")
			}
		}
	}

	for _, bp := range p.Build.Buildpacks {
		if bp.ID == "" && bp.URI == "" {
			return errors.New("descriptor buildpacks must have an id or url defined")
		}
		if bp.URI != "" && bp.Version != "" {
			return errors.New("descriptor buildpacks cannot have both uri and version defined")
		}
	}

	return nil
}
