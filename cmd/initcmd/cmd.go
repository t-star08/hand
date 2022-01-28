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
	if _, err := os.Stat("./.hand"); !os.IsNotExist(err) {
		return fmt.Errorf("\".hand\" already exist")
	}

	os.Mkdir("./.hand", 0777)
	return nil
}

func run(c *cobra.Command, args []string) {
	if err := createSettingsDirectory(); err != nil {
		logger.Fatalln(err)
	}

	if err := ioJson.Puts("./.hand/hand.json", template); err != nil {
		logger.Fatalln(err)
	}
}
