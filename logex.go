package logs

import "fmt"

var log *Logger

func InitLog() {
	log = GetLogger()
}

func Debug(format string, v ...interface{}) {
	if log == nil {
		fmt.Println(fmt.Sprintf(format , v ...))
		return
	}
	if DEBUG < log.level {
		return
	}
	msg := fmt.Sprintf("[DEBUG] "+format, v...)
	log.writeMsg(DEBUG, msg)
}

func Info(format string, v ...interface{}) {
	if log == nil {
		fmt.Println(fmt.Sprintf(format , v ...))
		return
	}
	if INFO < log.level {
		return
	}
	msg := fmt.Sprintf("[INFO] "+format, v...)
	log.writeMsg(INFO, msg)
}

func Warn(format string, v ...interface{}) {
	if log == nil {
		fmt.Println(fmt.Sprintf(format , v ...))
		return
	}
	if WARN < log.level {
		return
	}
	msg := fmt.Sprintf("[WARN] "+format, v...)
	log.writeMsg(WARN, msg)
}

func Error(format string, v ...interface{}) {
	if log == nil {
		fmt.Println(fmt.Sprintf(format , v ...))
		return
	}
	if ERROR < log.level {
		return
	}
	msg := fmt.Sprintf("[ERROR] "+format, v...)
	log.writeMsg(ERROR, msg)
}

func Fatal(format string, v ...interface{}) {
	if log == nil {
		fmt.Println(fmt.Sprintf(format , v ...))
		return
	}
	if FATAL < log.level {
		return
	}
	msg := fmt.Sprintf("[FATAL] "+format, v...)
	log.writeMsg(FATAL, msg)
}
