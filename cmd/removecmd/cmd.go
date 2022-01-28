package removecmd

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
	Use: "remove <log ID>",
	Run: run,
	Short: "remove log by specifying <log ID>",
}

var (
	logger = log.New(os.Stderr, "remove: ", log.LstdFlags)
	pre = preparator.NewLogEnvPreparator()
)

func Remove(pre *preparator.LogEnvPreparator, logID string) (*logUtil.LogUnit, []*logUtil.LogUnit, error) {
	var targetLogIndex int
	if i, err := findTargetLogByLogID(pre.Logs, logID); err != nil {
		return pre.Logs[0], pre.Logs, err
	} else {
		targetLogIndex = i
	}

	targetLog := pre.Logs[targetLogIndex]
	return targetLog, append(pre.Logs[:targetLogIndex], shiftIDLeft(pre.Logs[targetLogIndex+1:])...), nil
}

func setPreWorkers() {
	pre.SetInputChecker(
		func(args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("too few args\n\n\"remove\" must be \"hand remove <log ID>\"")
			}
			if len(args) > 1 {
				return fmt.Errorf("too many args\n\n\"remove\" must be \"hand remove <log ID>\"")
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

func findTargetLogByLogID(logs []*logUtil.LogUnit, targetLogID string) (int, error) {
	for i, log := range logs {
		if log.ID == targetLogID {
			return i, nil
		}
	}
	return -1, fmt.Errorf("log ID \"%s\" not found", targetLogID)
}

func shiftIDLeft(logs []*logUtil.LogUnit) []*logUtil.LogUnit {
	for _, log := range logs {
		log.ID = logUtil.PublishPrevID(log)
	}
	return logs
}

func run(c *cobra.Command, args []string) {
	setPreWorkers()
	if err := pre.Execute(args); err != nil {
		logger.Fatalln(err)
	}

	var err error
	if _, pre.Logs, err = Remove(pre, args[0]); err != nil {
		logger.Fatalln(err)
	}

	if err := ioJson.Puts(pre.LogFilePath, pre.Logs); err != nil {
		logger.Fatalln(err)
	}
}
