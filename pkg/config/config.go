package config

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

var gConf = Config{}

// SetConfigDir sets configuration directory to the variable.
func SetConfigDir(path string) error {
	return gConf.SetConfigDir(path)
}

// SetConfigDir sets configuration directory to the variable.
func (c *Config) SetConfigDir(path string) error {
	if stat, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("%v is not exists", path)
	} else if !stat.IsDir() {
		return fmt.Errorf("%v is not directory", path)
	}
	c.ConfigDir = path
	return nil
}

// SetDefaultConfigDir sets default configuration directory to ConfigDir.
func SetDefaultConfigDir() error {
	return gConf.SetDefaultConfigDir()
}

// SetDefaultConfigDir sets default configuration directory to ConfigDir.
func (c *Config) SetDefaultConfigDir() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	c.ConfigDir = filepath.Join(home, defaultDirName)
	return nil
}

func getConfigFilePath(configDir string) string {
	return filepath.Join(configDir, configFileName)
}

// LoadConfig loads configuration data from the ConfigFile.
func LoadConfig() error {
	return gConf.LoadConfig()
}

// LoadConfig loads configuration data from the ConfigFile.
func (c *Config) LoadConfig() error {
	return LoadYAML(getConfigFilePath(c.ConfigDir), &c.Ghctl)
}

// SaveConfig saves configuration data to the ConfigFile.
func SaveConfig() (err error) {
	return gConf.SaveConfig()
}

// SaveConfig saves configuration data to the ConfigFile.
func (c *Config) SaveConfig() error {
	path := getConfigFilePath(c.ConfigDir)
	if err := MkDirAllIfNotExist(path); err != nil {
		return err
	}
	return WriteYAML(path, c.Ghctl)
}

// RegisterContext registers a new context to contexts list.
func RegisterContext(ctx Context) error {
	return gConf.RegisterContext(ctx)
}

// RegisterContext registers a new context to contexts list.
func (c *Config) RegisterContext(ctx Context) error {
	if ctx.Name == c.Ghctl.CurrentContext {
		return nil
	}
	for _, v := range c.Ghctl.Contexts {
		if v.Name == ctx.Name {
			return fmt.Errorf(
				"%s has already used by other contexts",
				ctx.Name)
		}
	}
	c.Ghctl.Contexts = append(c.Ghctl.Contexts, ctx)
	if c.Ghctl.CurrentContext == "" {
		if err := c.SetCurrentContext(ctx.Name); err != nil {
			return err
		}
	}
	return nil
}

// GetCurrentContext returns current context.
func GetCurrentContext() *Context {
	return gConf.GetCurrentContext()
}

// GetCurrentContext returns current context.
func (c *Config) GetCurrentContext() *Context {
	return c.GetContext(c.Ghctl.CurrentContext)
}

// SetCurrentContext sets context for using in ghctl.
func SetCurrentContext(name string) error {
	return gConf.SetCurrentContext(name)
}

// SetCurrentContext sets context for using in ghctl.
func (c *Config) SetCurrentContext(name string) error {
	exists := false
	for _, v := range c.Ghctl.Contexts {
		if v.Name == name {
			exists = true
			break
		}
	}
	if !exists {
		return fmt.Errorf("%s is not contained in contexts list", name)
	}
	c.Ghctl.CurrentContext = name
	return nil
}

// GetContexts returns all contexts.
func GetContexts() []Context {
	return gConf.Ghctl.Contexts
}

// GetContext returns the context matched specified name.
func GetContext(name string) *Context {
	return gConf.GetContext(name)
}

// GetContext returns the context matched specified name.
func (c *Config) GetContext(name string) *Context {
	for _, ctx := range c.Ghctl.Contexts {
		if name == ctx.Name {
			return &ctx
		}
	}
	return nil
}
