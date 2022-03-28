package preparer

import "os"

type Options struct {
	logger      Logger
	sourceDir   string
	platformDir string
}

type Option func(opts *Options)

// WithLogger sets a custom logger
func WithLogger(logger Logger) func(*Options) {
	return func(o *Options) {
		o.logger = logger
	}
}

// ReadEnvOptions reads options from env vars
func ReadEnvOptions(o *Options) {
	o.sourceDir = getEnvOrDefault("CNB_APP_DIR", "/workspace")
	o.platformDir = getEnvOrDefault("CNB_PLATFORM_DIR", "/platform")
}

// WithEnvOptions loads options from env vars
//
// See `ReadEnvOptions`
func WithEnvOptions() func(*Options) {
	return func(o *Options) {
		ReadEnvOptions(o)
	}
}

func getEnvOrDefault(key string, defaultVal string) string {
	stringVal, found := os.LookupEnv(key)
	if !found {
		return defaultVal
	}

	return stringVal
}
