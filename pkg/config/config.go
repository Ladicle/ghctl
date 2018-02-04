package config

import (
	"fmt"
	"os"
	"path/filepath"

	"io/ioutil"

	homedir "github.com/mitchellh/go-homedir"
	yaml "gopkg.in/yaml.v2"
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
	if home, err := homedir.Dir(); err != nil {
		return err
	} else {
		c.ConfigFile = fmt.Sprintf("%s/.ghctl/config", home)
	}
	return nil
}

// LoadConfig loads configuration data from the ConfigFile.
func LoadConfig() error {
	return c.LoadConfig()
}

// LoadConfig loads configuration data from the ConfigFile.
func (c *Config) LoadConfig() error {
	if _, err := os.Stat(c.ConfigFile); os.IsNotExist(err) {
		return nil
	}
	d, err := ioutil.ReadFile(c.ConfigFile)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(d, c.Ghctl); err != nil {
		return err
	}
	return nil
}

// SaveConfig saves configuration data to the ConfigFile.
func SaveConfig() (err error) {
	return c.SaveConfig()
}

// SaveConfig saves configuration data to the ConfigFile.
func (c *Config) SaveConfig() (err error) {
	data, err := yaml.Marshal(c.Ghctl)
	if err != nil {
		return err
	}
	dir := filepath.Dir(c.ConfigFile)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0777); err != nil {
			return err
		}
	}

	f, err := os.OpenFile(c.ConfigFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() { err = f.Close() }()
	f.Write(data)

	return nil
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
