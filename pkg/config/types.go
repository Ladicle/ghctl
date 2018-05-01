package config

const (
	defaultDirName = ".ghctl"
	configFileName = "config"
)

// Config contains configuration path and data.
type Config struct {
	ConfigDir string
	Ghctl     Ghctl
}

// Ghctl is configuration for ghctl.
type Ghctl struct {
	CurrentContext string    `json:"current_context"`
	Contexts       []Context `json:"contexts"`
}

// Context manages information to access to GitHub.
type Context struct {
	Name        string `json:"name"`
	AccessToken string `json:"access_token"`
	Endpoint    string `json:"endpoint"`
}
