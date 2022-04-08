package util

type Logger struct {
	component string
}

func NewLogger(component string) *Logger {
	return &Logger{
		component: component,
	}
}
