package preparer

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jromero/cnb-prepare/pkg/project"
	"github.com/jromero/cnb-prepare/pkg/project/types"
	"github.com/pkg/errors"
)

func Preparer(opts ...Option) error {
	var options = &Options{
		logger: &defaultLogger{},
	}

	for _, opt := range opts {
		opt(options)
	}

	options.logger.Debug("Options: %+v", options)

	descriptor, err := project.ReadProjectDescriptor(filepath.Join(options.sourceDir, "project.toml"))
	if err != nil {
		if err != os.ErrNotExist {
			return errors.Wrap(err, "project.toml")
		}
	}

	if descriptor != nil {
		err = processDescriptor(*options, *descriptor)
		if err != nil {
			return err
		}
	}

	return nil
}

func processDescriptor(options Options, descriptor types.Descriptor) error {
	options.logger.Debug("Found descriptor:\n %s", prettyPrint(descriptor))

	options.logger.Debug("Creating environment variable files...")
	platformEnvDir := filepath.Join(options.platformDir, "env")
	err := os.MkdirAll(platformEnvDir, os.ModePerm)
	if err != nil {
		return err
	}

	for _, env := range descriptor.Build.Env {
		envFilePath := filepath.Join(platformEnvDir, env.Name)
		options.logger.Debug("Creating environment variable file: %s", envFilePath)
		err := ioutil.WriteFile(envFilePath, []byte(env.Value), os.ModePerm)
		if err != nil {
			return errors.Wrapf(err, "creating env file for variable '%s'", env.Name)
		}
	}

	return nil
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}
