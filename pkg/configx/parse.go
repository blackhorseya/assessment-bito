package configx

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

// C is a global config instance.
var C = new(Config)

// LoadWithPathAndName loads config from path and name.
func LoadWithPathAndName(path, name string) (err error) {
	v := viper.GetViper()

	if path != "" {
		v.SetConfigFile(path)
	} else {
		home, _ := os.UserHomeDir()
		v.AddConfigPath(home)
		v.SetConfigType("yaml")
		v.SetConfigName("." + name)
	}

	err = v.ReadInConfig()
	if err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		return err
	}

	err = v.Unmarshal(&C)
	if err != nil {
		return err
	}

	return nil
}
