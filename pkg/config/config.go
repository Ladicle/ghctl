package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Ladicle/ghctl/pkg/util"
	homedir "github.com/mitchellh/go-homedir"
)

var c = Config{}

// SetConfigFile sets configuration path to the variable.
func SetConfigFile(path string) error {
	return c.SetConfigFile(path)
}

// SetConfigFile sets configuration path to the variable.
func (c *Config) SetConfigFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("%v is not exists", path)
	}
	c.ConfigFile = path
	return nil
}

// SetDefaultConfigFile sets default configuration path to ConfigFile.
func SetDefaultConfigFile() error {
	return c.SetDefaultConfigFile()
}

// SetDefaultConfigFile sets default configuration path to ConfigFile.
func (c *Config) SetDefaultConfigFile() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}
	c.ConfigFile = filepath.Join(home, ".ghctl", "config")
	return nil
}

// LoadConfig loads configuration data from the ConfigFile.
func LoadConfig() error {
	return c.LoadConfig()
}

// LoadConfig loads configuration data from the ConfigFile.
func (c *Config) LoadConfig() error {
	return util.LoadYAML(c.ConfigFile, &c.Ghctl)
}

// SaveConfig saves configuration data to the ConfigFile.
func SaveConfig() (err error) {
	return c.SaveConfig()
}

// SaveConfig saves configuration data to the ConfigFile.
func (c *Config) SaveConfig() error {
	if err := util.MkDirAllIfNotExist(c.ConfigFile); err != nil {
		return err
	}
	return util.WriteYAML(c.ConfigFile, c.Ghctl)
}

// RegisterContext registers a new context to contexts list.
func RegisterContext(ctx Context) error {
	return c.RegisterContext(ctx)
}

// RegisterContext registers a new context to contexts list.
func (c *Config) RegisterContext(ctx Context) error {
	if ctx.Name == c.Ghctl.CurrentContext {
		return nil
	}
	for _, v := range c.Ghctl.Contexts {
		if v.Name == ctx.Name {
			return fmt.Errorf(
				"%s has already used by other contexts.",
				"The context name must be unique.",
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
	return GetContext(c.Ghctl.CurrentContext)
}

// SetCurrentContext sets context for using in ghctl.
func SetCurrentContext(name string) error {
	return c.SetCurrentContext(name)
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
		return fmt.Errorf("%s is not contained in contexts list.", name)
	}
	c.Ghctl.CurrentContext = name
	return nil
}

// GetContexts returns all contexts.
func GetContexts() []Context {
	return c.Ghctl.Contexts
}

// GetContext returns the context matched specified name.
func GetContext(name string) *Context {
	return c.GetContext(name)
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
