package gologger

type LogLevel uint8

var logLevelNum int = 4

const (
	Debug LogLevel = iota
	Info
	Error
	Fatal
)

func defaultLog() []LogLevel {
	return []LogLevel{Error, Fatal}
}

func (t LogLevel) String() string {
	return [...]string{"Debug", "Info", "Error", "Fatal"}[t]
}

func getLevel(i int) LogLevel {
	return [...]LogLevel{Debug, Info, Error, Fatal}[i]
}

func (t LogLevel) EnumIndex() int {
	return int(t)
}

func getLevelList(l LogLevel) []LogLevel {
	var logLevelList []LogLevel
	for i := l.EnumIndex(); i < logLevelNum; i++ {
		logLevelList = append(logLevelList, getLevel(i))
	}
	return logLevelList
}
