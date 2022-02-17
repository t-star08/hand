package config

import (
	"github.com/t-star08/hand/pkg/evaluator/fileName"
	"github.com/t-star08/hand/pkg/fileUtil/fileSearcher"
)

var (
	evaluator = fileName.NewExactEvaluator()
)

func SearchConfigFilePath() (string, error) {
	if pathToConfDir, err := fileSearcher.BackwardSearch(evaluator, CONF_DIR_NAME, "."); err != nil {
		return pathToConfDir, err
	} else {
		return pathToConfDir + "/" + CONF_FILE_NAME, nil
	}
}
