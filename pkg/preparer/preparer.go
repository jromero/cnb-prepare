package preparer

import (
	"encoding/json"
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
		processDescriptor(options.logger, *descriptor)
	}

	return nil
}

func processDescriptor(logger Logger, descriptor types.Descriptor) {
	logger.Debug("Found descriptor:\n %s", prettyPrint(descriptor))
	// TODO: https://github.com/jromero/cnb-prepare/issues/4
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}
