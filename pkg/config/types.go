package config

// Config contains configuration path and data.
type Config struct {
	ConfigFile string
	Ghctl      Ghctl
}

// Ghctl is configuration for ghctl.
type Ghctl struct {
	CurrentContext string    `yaml:"current_context"`
	Contexts       []Context `yaml:"contexts"`
}

// Context manages information to access to GitHub.
type Context struct {
	Name        string `yaml:"name"`
	AccessToken string `yaml:"access_token"`
	Endpoint    string `yaml:"endpoint"`
}
