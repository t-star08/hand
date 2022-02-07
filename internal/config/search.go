package config

import (
	"github.com/t-star08/hand/pkg/evaluator/fileName"
	"github.com/t-star08/hand/pkg/fileUtil/fileSearcher"
)

var (
	configDirectory = ".hand"
	configFile = "hand.json"
	evaluator = fileName.NewExactEvaluator()
)

func SearchConfigFilePath() (string, error) {
	if configPath, err := fileSearcher.BackwardSearch(evaluator, configDirectory, "."); err != nil {
		return configPath, err
	} else {
		return configPath + "/" + configFile, nil
	}
}
