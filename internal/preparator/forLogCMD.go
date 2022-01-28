package preparator

import (
	"github.com/t-star08/hand/internal/config"
	"github.com/t-star08/hand/internal/logUtil"
	"github.com/t-star08/hand/internal/thisTime"
	"github.com/t-star08/hand/pkg/io/ioJson"
)

type LogEnvPreparator struct {
	Conf				*config.ConfigJson
	Now					*thisTime.ThisTime
	LogFilePath			string
	NoNeedLogFile		bool
	Logs				[]*logUtil.LogUnit
	CheckInput			func(args []string) error
	CheckLogDirectory	func() error
	CheckLogFile		func() error
	GetLogFileName		func(args []string) string
}

func NewLogEnvPreparator() *LogEnvPreparator {
	return &LogEnvPreparator{}
}

func (pre *LogEnvPreparator) SetInputChecker(inputChecker func(args []string) error) {
	pre.CheckInput = inputChecker
}

func (pre *LogEnvPreparator) SetLogDirectoryChecker(logDirectoryChecker func() error) {
	pre.CheckLogDirectory = logDirectoryChecker
}

func (pre *LogEnvPreparator) SetLogFileChecker(logFileChecker func() error) {
	pre.CheckLogFile = logFileChecker
}

func (pre *LogEnvPreparator) SetLogFileNameGetter(logFileNameGetter func(args []string) string) {
	pre.GetLogFileName = logFileNameGetter
}

func (pre *LogEnvPreparator) Execute(args []string) error {
	pre.Now = thisTime.Lock()

	if conf, err := ReadConf(); err != nil {
		return err
	} else {
		pre.Conf = conf
	}
	
	if err := pre.CheckInput(args); err != nil {
		return err
	}

	pre.LogFilePath = pre.GetLogFileName(args)
	
	if err := pre.CheckLogDirectory(); err != nil {
		return err
	}

	if err := pre.CheckLogFile(); err != nil {
		return err
	}

	if !pre.NoNeedLogFile {
		if err := ioJson.Gets(pre.LogFilePath, &pre.Logs); err != nil {
			return err
		}	
	}

	return nil
}

