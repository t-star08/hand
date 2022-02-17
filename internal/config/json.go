package config

type ConfigJson struct {
	Path	string `json:"path"`
}

const (
	CONF_DIR_NAME = ".hand"
	CONF_FILE_NAME = "hand.json"
)
