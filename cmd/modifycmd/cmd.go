package modifycmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/t-star08/hand/cmd/insertcmd"
	"github.com/t-star08/hand/cmd/removecmd"
	"github.com/t-star08/hand/internal/logUtil"
	"github.com/t-star08/hand/internal/preparator"
	"github.com/t-star08/hand/pkg/io/ioJson"
)

var CMD = &cobra.Command {
	Use: "modify <log ID> <name> <time>",
	Run: run,
	Short: "modify log which specified by <log ID>",
}

var (
	logger = log.New(os.Stderr, "modify: ", log.LstdFlags)
	pre = preparator.NewLogEnvPreparator()
)

func setPreWorkers() {
	pre.SetInputChecker(
		func(args []string) error {
			if len(args) < 3 {
				return fmt.Errorf("too few args\n\n\"modify\" must Be \"hand modify <log ID> <name> <time>\"\nNotes: \"\" being specified, the property will not be changed")
			}
			if len(args) > 3 {
				return fmt.Errorf("too many args\n\n\"modify\" must Be \"hand modify <log ID> <name> <time>\"\nNotes: \"\" being specified, the property will not be changed")
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
			return logUtil.CheckLogFile(pre.LogFilePath)
		},
	)
	pre.SetLogFileNameGetter(
		func(_ []string) string {
			return fmt.Sprintf("%s/%s.json", pre.Conf.Path, pre.Now.ForLogFileName)
		},
	)
}

func modify(logID, name, time string) ([]*logUtil.LogUnit, error) {
	var (
		err error
		removedLog *logUtil.LogUnit
	)
	if removedLog, pre.Logs, err = removecmd.Remove(pre, logID); err != nil {
		return pre.Logs, err
	}

	if name != "" {
		removedLog.Name = name
	}
	if time != "" {
		removedLog.Time = time
	}

	return insertcmd.Insert(pre, removedLog.Name, removedLog.Time)
}

func run(c *cobra.Command, args []string) {
	setPreWorkers()
	if err := pre.Execute(args); err != nil {
		logger.Fatalln(err)
	}

	var err error
	if pre.Logs, err = modify(args[0], args[1], args[2]); err != nil {
		logger.Fatalln(err)
	}

	if err := ioJson.Puts(pre.LogFilePath, pre.Logs); err != nil {
		logger.Fatalln(err)
	}
}
