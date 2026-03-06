package logger

import "strings"

// reset
var colorReset = "\033[0m"

// COLORS
var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorBlue = "\033[34m"

// Reserved Color Variables
var colorMagenta = "\033[35m"
var colorCyan = "\033[36m"
var colorGray = "\033[37m"
var colorWhite = "\033[97m"

// levelType
type LevelType int8

const (
	VERBOSE LevelType = iota - 2 // -2
	DEBUG                        // -1
	INFO                         // 0
	WARN                         // 1
	ERROR                        // 2
	FATAL                        // 3
)

func (l LevelType) String() string {
	switch l {
	case VERBOSE:
		return "VERBOSE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "ERROR"
	default:
		return "INFO"
	}
}

func (l LevelType) color() string {
	switch l {
	case VERBOSE:
		return colorCyan
	case DEBUG:
		return colorBlue
	case INFO:
		return colorGreen
	case WARN:
		return colorYellow
	case ERROR:
		return colorRed
	case FATAL:
		return colorRed
	default:
		return colorGreen
	}
}

// ParseLevel converts a string representation of a log level to a levelType.
// It defaults to INFO if the string is unrecognized or empty.
func ParseLevel(levelStr string) LevelType {
	switch strings.ToUpper(strings.TrimSpace(levelStr)) {
	case "VERBOSE":
		return VERBOSE // Assuming you added VERBOSE as -2
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	default:
		return INFO // Safe fallback
	}
}

// UnmarshalText teaches Viper/mapstructure how to decode a string into a levelType
func (l *LevelType) UnmarshalText(text []byte) error {
	*l = ParseLevel(string(text)) // Using the ParseLevel func from my previous message
	return nil
}
