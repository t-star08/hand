package insertcmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/t-star08/hand/internal/logUtil"
	"github.com/t-star08/hand/internal/preparator"
	"github.com/t-star08/hand/internal/thisTime"
	"github.com/t-star08/hand/pkg/io/ioJson"
)

var CMD = &cobra.Command {
	Use: "insert <name> <time>",
	Run: run,
	Short: "insert log",
}

var (
	logger = log.New(os.Stderr, "insert: ", log.LstdFlags)
	pre = preparator.NewLogEnvPreparator()
)

func Insert(pre *preparator.LogEnvPreparator, name, time string) ([]*logUtil.LogUnit, error) {
	var wanna *thisTime.ThisClock
	if c, err := thisTime.ParseClock(time); err != nil {
		return pre.Logs, err
	} else {
		wanna = c
	}

	var insertPoint int
	if ip, err := searchInsertPoint(pre, wanna); err != nil {
		return pre.Logs, err
	} else {
		insertPoint = ip
	}

	newLog := &logUtil.LogUnit {
		ID: logUtil.PublishID(insertPoint),
		Name: name,
		Time: wanna.Format,
	}
	if insertPoint == len(pre.Logs) {
		return append(pre.Logs[:insertPoint], newLog), nil
	} else {
		return append(pre.Logs[:insertPoint], append([]*logUtil.LogUnit{newLog}, shiftIDRight(pre.Logs[insertPoint:])...)...), nil
	}
}

func setPreWorkers() {
	pre.SetInputChecker(
		func(args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("too few args\n\n\"insert\" must be \"hand insert <name> <time>\"")
			}
			if len(args) > 2 {
				return fmt.Errorf("too many args\n\n\"insert\" must be \"hand insert <name> <time>\"")
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

func searchInsertPoint(pre *preparator.LogEnvPreparator, wanna *thisTime.ThisClock) (int, error) {
	var one *thisTime.ThisClock
	for i, log := range pre.Logs {
		if c, err := thisTime.ParseClock(log.Time); err != nil {
			return -1, err
		} else {
			one = c
		}
		if thisTime.LessThisClock(wanna, one) {
			if i == 0 {
				return i, fmt.Errorf("cannot insert before \"START\" work")
			}
			return i, nil
		}
	}
	if thisTime.LessThisClock(pre.Now.Clock, wanna) {
		return -1, fmt.Errorf("cannot log the future")
	}
	return len(pre.Logs), nil
}

func shiftIDRight(logs []*logUtil.LogUnit) []*logUtil.LogUnit {
	for _, log := range logs {
		log.ID = logUtil.PublishNextID(log)
	}
	return logs
}

func run(c *cobra.Command, args []string) {
	setPreWorkers()
	if err := pre.Execute(args); err != nil {
		logger.Fatalln(err)
	}

	var err error
	if pre.Logs, err = Insert(pre, args[0], args[1]); err != nil {
		logger.Fatalln(err)
	}

	if err := ioJson.Puts(pre.LogFilePath, pre.Logs); err != nil {
		logger.Fatalln(err)
	}
}
