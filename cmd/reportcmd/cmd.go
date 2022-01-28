package reportcmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/t-star08/hand/internal/logUtil"
	"github.com/t-star08/hand/internal/preparator"
	"github.com/t-star08/hand/internal/thisTime"
)

var CMD = &cobra.Command {
	Use: "report <date>",
	Run: run,
	Short: "report what you worked in the <date>. if wanna today's, no args required",
}

var (
	logger = log.New(os.Stderr, "report: ", log.LstdFlags)
	pre = preparator.NewLogEnvPreparator()
)

func setPreWorkers() {
	pre.SetInputChecker(
		func(args []string) error {
			if len(args) == 0 {
				return nil
			} else if len(args) > 1 {
				return fmt.Errorf("too many args\n\n\"report\" must be \"hand report\" or \"hand report <date>\"")
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

func detectTargetDate(logFilePath string) string {
	splitedPath := strings.Split(logFilePath, "/")
	return strings.Replace(strings.Replace(splitedPath[len(splitedPath)-1], "-", "/", -1), ".json", "", -1)
}

func parse(targetDate string) (string, error) {
	var (
		report = fmt.Sprintf("--- Report[%s] ---", targetDate)
		prefix = "  "
		prevLog, thisLog *logUtil.LogUnit
		prev, this *thisTime.ThisClock
		workDurations = make(map[string]*thisTime.ThisDuration)
	)

	prevLog = pre.Logs[0]
	for i := 1; i < len(pre.Logs); i++ {
		thisLog = pre.Logs[i]
		if p, err := thisTime.ParseClock(prevLog.Time); err != nil {
			return "", err
		} else {
			prev = p
		}
		if p, err := thisTime.ParseClock(thisLog.Time); err != nil {
			return "", err
		} else {
			this = p
		}

		if d, exist := workDurations[thisLog.Name]; exist {
			workDurations[thisLog.Name] = d.Add(this.Sub(prev))
		} else {
			workDurations[thisLog.Name] = this.Sub(prev)
		}
		prevLog = thisLog
	}

	sum := thisTime.NewThisDuration(0, 0, 0)
	for name, duration := range workDurations {
		report += fmt.Sprintf("\n%s・%s: %s", prefix, name, duration.Format)
		sum = sum.Add(duration)
	}

	report += fmt.Sprintf("\n%s・Sum: %s", prefix, sum.Format)
	report += "\n--------------------------"
	return report, nil
}

func run(c *cobra.Command, args []string) {
	setPreWorkers()
	if err := pre.Execute(args); err != nil {
		logger.Fatalln(err)
	}

	if repo, err := parse(detectTargetDate(pre.LogFilePath)); err != nil {
		logger.Fatalln(err)
	} else {
		fmt.Println(repo)
	}
}
