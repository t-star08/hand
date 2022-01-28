package statuscmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/t-star08/hand/internal/logUtil"
	"github.com/t-star08/hand/internal/preparator"
)

var CMD = &cobra.Command {
	Use: "status <date>",
	Run: run,
	Short: "show <date>'s log status. if wanna today's, no args required",
}

var (
	logger = log.New(os.Stderr, "status: ", log.LstdFlags)
	pre = preparator.NewLogEnvPreparator()
)

func setPreWorkers() {
	pre.SetInputChecker(
		func(args []string) error {
			if len(args) == 0 {
				return nil
			} else if len(args) > 1 {
				return fmt.Errorf("too many args\n\n\"status\" must be \"hand status\" or \"hand status [date]\"")
			}
			if _, err := os.Stat(fmt.Sprintf("%s/%s.json", pre.Conf.Path, args[0])); !os.IsNotExist(err) {
				return nil
			} else if _, err := os.Stat(fmt.Sprintf("%s/%s.json", pre.Conf.Path, strings.Replace(args[0], "/", "-", -1))); !os.IsNotExist(err) {
				return nil
			} else {
				return err
			}
		},
	)
	pre.SetLogDirectoryChecker(
		func() error {
			return logUtil.CheckLogDirectory(pre.Conf.Path)
		},
	)
	pre.SetLogFileChecker(
		func() error {
			return logUtil.CheckLogFile(pre.LogFilePath)
		},
	)
	pre.SetLogFileNameGetter(
		func(args []string) string {
			if len(args) == 0 {
				return fmt.Sprintf("%s/%s.json", pre.Conf.Path, pre.Now.ForLogFileName)
			}
			if _, err := os.Stat(fmt.Sprintf("%s/%s.json", pre.Conf.Path, args[0])); !os.IsNotExist(err) {
				return fmt.Sprintf("%s/%s.json", pre.Conf.Path, args[0])
			}
			return fmt.Sprintf("%s/%s.json", pre.Conf.Path, strings.Replace(args[0], "/", "-", -1))
		},
	)
}

func showLog() {
	for _, log := range pre.Logs {
		fmt.Printf("ID: %s, Time: %s, Work: %s\n", log.ID, log.Time, log.Name)
	}
}

func run(c * cobra.Command, args []string) {
	setPreWorkers()
	if err := pre.Execute(args); err != nil {
		logger.Fatalln(err)
	}

	showLog()
}
