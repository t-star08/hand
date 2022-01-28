package keepcmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/t-star08/hand/internal/logUtil"
	"github.com/t-star08/hand/internal/preparator"
	"github.com/t-star08/hand/pkg/io/ioJson"
)

var CMD = &cobra.Command {
	Use: "keep <name>",
	Run: run,
	Short: "log what you worked. if first log in the day, [name] is needless",
}

var (
	logger = log.New(os.Stderr, "keep: ", log.LstdFlags)
	pre = preparator.NewLogEnvPreparator()
)

func setPreWorkers() {
	pre.SetInputChecker(
		func(args []string) error {
			if len(args) == 1 {
				return nil
			} else if len(args) > 1 {
				return fmt.Errorf("too many args\n\n\"keep\" must be \"hand keep\" or \"hand keep <name>\"")
			}
			if _, err := os.Stat(fmt.Sprintf("%s/%s.json", pre.Conf.Path, pre.Now.ForLogFileName)); !os.IsNotExist(err) {
				return fmt.Errorf("too few args\n\n\"keep\" must be \"hand keep <name>\" after first log in the day")
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
			if first := logUtil.CheckLogFile(pre.LogFilePath) != nil; !first {
				return nil
			}
			return fmt.Errorf("first")		
		},
	)
	pre.SetLogFileNameGetter(
		func(_ []string) string {
			return fmt.Sprintf("%s/%s.json", pre.Conf.Path, pre.Now.ForLogFileName)
		},
	)
}

func firstKeep() error {
	if err := logUtil.CreateLogFile(pre.LogFilePath); err != nil {
		return err
	}
	if err := ioJson.Puts(pre.LogFilePath, keep("000", "Start")); err != nil {
		return err
	}
	return nil
}

func keep(id, name string) []*logUtil.LogUnit {
	newLog := &logUtil.LogUnit {
		ID: id,
		Name: name,
		Time: pre.Now.ForInLog,
	}
	return append(pre.Logs, newLog)
}

func run(c *cobra.Command, args []string) {
	setPreWorkers()
	if err := pre.Execute(args); err != nil {
		if err.Error() == "first" {
			if err := firstKeep(); err != nil {
				logger.Fatalln(err)
			}
			os.Exit(0)
		}
		logger.Fatalln(err)
	}

	pre.Logs = keep(logUtil.PublishNextID(pre.Logs[len(pre.Logs) - 1]), args[0])
	if err := ioJson.Puts(pre.LogFilePath, pre.Logs); err != nil {
		logger.Fatalln(err)
	}
}
