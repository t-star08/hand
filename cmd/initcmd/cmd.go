package initcmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/t-star08/hand/internal/config"
	"github.com/t-star08/hand/pkg/io/ioJson"
)

var CMD = &cobra.Command {
	Use: "init",
	Run: run,
	Short: "make \"hand\" config",
}

var (
	logger = log.New(os.Stderr, "init: ", log.LstdFlags)
	template = &config.ConfigJson {
		Path: "./log",
	}
)

func createSettingsDirectory() error {
	if _, err := os.Stat(fmt.Sprintf("./%s", config.CONF_DIR_NAME)); !os.IsNotExist(err) {
		return fmt.Errorf("\"%s\" already exist", config.CONF_DIR_NAME)
	}

	os.Mkdir(fmt.Sprintf("./%s", config.CONF_DIR_NAME), 0777)
	return nil
}

func run(c *cobra.Command, args []string) {
	if err := createSettingsDirectory(); err != nil {
		logger.Fatalln(err)
	}

	if err := ioJson.Puts(fmt.Sprintf("./%s/%s", config.CONF_DIR_NAME, config.CONF_FILE_NAME), template); err != nil {
		logger.Fatalln(err)
	}
}
