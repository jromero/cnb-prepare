package main_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jromero/cnb-prepare/pkg/preparer"
	"github.com/jromero/cnb-prepare/pkg/testhelpers"
	cp "github.com/otiai10/copy"
	. "github.com/otiai10/mint"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func Test_CreateEnvVarFile(t *testing.T) {
	logger := testhelpers.NewLogger(t)

	logger.Debug("Creating temp dir...")
	tmpDir, err := os.MkdirTemp("", "test")
	Expect(t, err).ToBe(nil)

	logger.Info("Temp Dir: %s", tmpDir)

	logger.Debug("Scaffolding build environment...")
	err = cp.Copy("testdata/build_env_skeleton", tmpDir)
	Expect(t, err).ToBe(nil)

	logger.Debug("Providing app source: app_descriptor_0.2")
	appDir := filepath.Join(tmpDir, "workspace")
	err = cp.Copy("testdata/app_descriptor_0.2", appDir)
	Expect(t, err).ToBe(nil)

	os.Setenv("CNB_APP_DIR", appDir)
	defer os.Unsetenv("CNB_APP_DIR")

	os.Setenv("CNB_PLATFORM_DIR", filepath.Join(tmpDir, "platform"))
	defer os.Unsetenv("CNB_PLATFORM_DIR")

	orderPath := filepath.Join(tmpDir, "layers", "order.toml")
	os.Setenv("CNB_ORDER_PATH", orderPath)
	defer os.Unsetenv("CNB_ORDER_PATH")

	err = preparer.Preparer(
		preparer.WithLogger(logger),
		preparer.WithEnvOptions(),
	)
	Expect(t, err).ToBe(nil)

	javaEnvFilePath := filepath.Join(tmpDir, "platform", "env", "JAVA_OPTS")
	contents, err := ioutil.ReadFile(javaEnvFilePath)
	Expect(t, err).ToBe(nil)
	Expect(t, string(contents)).ToBe("-Xmx1g")
}

func Test_UpdateOrderToml(t *testing.T) {
	logger := testhelpers.NewLogger(t)

	logger.Debug("Creating temp dir...")
	tmpDir, err := os.MkdirTemp("", "test")
	Expect(t, err).ToBe(nil)

	logger.Info("Temp Dir: %s", tmpDir)

	logger.Debug("Scaffolding build environment...")
	err = cp.Copy("testdata/build_env_skeleton", tmpDir)
	Expect(t, err).ToBe(nil)

	logger.Debug("Providing app source: app_descriptor_0.2")
	appDir := filepath.Join(tmpDir, "workspace")
	err = cp.Copy("testdata/app_descriptor_0.2", appDir)
	Expect(t, err).ToBe(nil)

	os.Setenv("CNB_APP_DIR", appDir)
	defer os.Unsetenv("CNB_APP_DIR")

	os.Setenv("CNB_PLATFORM_DIR", filepath.Join(tmpDir, "platform"))
	defer os.Unsetenv("CNB_PLATFORM_DIR")

	orderPath := filepath.Join(tmpDir, "layers", "order.toml")
	os.Setenv("CNB_ORDER_PATH", orderPath)
	defer os.Unsetenv("CNB_ORDER_PATH")

	err = preparer.Preparer(
		preparer.WithLogger(logger),
		preparer.WithEnvOptions(),
	)
	Expect(t, err).ToBe(nil)

	contents, err := ioutil.ReadFile(orderPath)
	Expect(t, err).ToBe(nil)

	Expect(t, strings.TrimSpace(string(contents))).
		ToBe(strings.TrimSpace(`
[[order]]

  [[order.group]]
    id = "hello-world"
    version = "0.0.1"

  [[order.group]]
    id = "hello-comet"
`))
}
