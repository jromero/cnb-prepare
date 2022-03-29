package order

type Order struct {
	Groups []Group `toml:"order"`
}

type Group struct {
	Buildpacks []BuildpackEntry `toml:"group"`
}

type BuildpackEntry struct {
	ID       string `toml:"id,omitempty"`
	Version  string `toml:"version,omitempty"`
	Optional bool   `toml:"optional,omitempty"`
}
