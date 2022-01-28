package preparator

import (
	"github.com/t-star08/hand/internal/config"
	"github.com/t-star08/hand/pkg/io/ioJson"
)

func ReadConf() (*config.ConfigJson, error) {
	var (
		conf *config.ConfigJson
		configPath string
	)

	if path, err := config.SearchConfigFilePath(); err != nil {
		return conf, err
	} else {
		configPath = path
	}

	if err := ioJson.Gets(configPath, &conf); err != nil {
		return conf, err
	}

	return conf, nil
}
