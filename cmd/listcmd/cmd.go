package listcmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/t-star08/hand/internal/logUtil"
	"github.com/t-star08/hand/internal/preparator"
	"github.com/t-star08/hand/pkg/io/ioStd"
)

var CMD = &cobra.Command {
	Use: "list",
	Run: run,
	Short: "show all logged dates",
}

var (
	logger = log.New(os.Stderr, "list: ", log.LstdFlags)
	pre = preparator.NewLogEnvPreparator()
)

func setPreWorkers() {
	pre.SetInputChecker(
		func(args []string) error {
			if len(args) > 0 {
				return fmt.Errorf("too many args\n\n\"list\" must be \"hand list\"")
			}
			return nil
		},
	)
	pre.SetLogDirectoryChecker(
		func() error {
			return logUtil.CheckLogDirectory(pre.Conf.Path)
		},
	)
	pre.SetLogFileChecker(
		func() error {
			return nil
		},
	)
	pre.SetLogFileNameGetter(
		func(_ []string) string {
			pre.NoNeedLogFile = true
			return ""
		},
	)
}

func showLogList() error {
	logDates := make([]string, 0)
	lsLogDir, _ := ioutil.ReadDir(pre.Conf.Path)
	for _, info := range lsLogDir {
		logDates = append(logDates, strings.Replace(info.Name(), ".json", "", -1))
	}

	return ioStd.MonospacedPuts("", logDates)
}

func run(c *cobra.Command, args []string) {
	setPreWorkers()
	if err := pre.Execute(args); err != nil {
		logger.Fatalln(err)
	}

	fmt.Printf("log directory: %s\n", pre.Conf.Path)
	showLogList()
}
