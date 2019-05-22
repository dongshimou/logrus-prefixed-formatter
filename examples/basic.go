package main

import (
	"github.com/dongshimou/logrus"
	prefixed "github.com/dongshimou/logrus-prefixed-formatter"
)

var log = logrus.New()

func init() {
	formatter := new(prefixed.TextFormatter)
	//formatter.CallerPrettyfier= func(frame *runtime.Frame) (function string, file string) {
	//	return frame.Function,fmt.Sprintf("%s:%d",frame.File,frame.Line)
	//}
	formatter.FullTimestamp=true
	formatter.TimestampFormat="15:04:05"
	log.Formatter = formatter
	log.Level = logrus.DebugLevel
	log.SetReportCaller(true)

}
func Info(args ...interface{}){
	l:=log
	l.CallerSkip=1
	l.Info(args...)
}
func main() {
	defer func() {
		err := recover()
		if err != nil {
			// Fatal message
			log.WithFields(logrus.Fields{
				"omg":    true,
				"number": 100,
			}).Fatal("[main] The ice breaks!")
		}
	}()

	Info("test skip")

	// You could either provide a map key called `prefix` to add prefix
	log.WithFields(logrus.Fields{
		"prefix": "main",
		"animal": "walrus",
		"number": 8,
	}).Debug("Started observing beach")

	// Or you can simply add prefix in square brackets within message itself
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Debug("[main] A group of walrus emerges from the ocean")

	// Warning message
	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("[main] The group's number increased tremendously!")

	// Information message
	log.WithFields(logrus.Fields{
		"prefix":      "sensor",
		"temperature": -4,
	}).Info("Temperature changes")

	// Panic message
	log.WithFields(logrus.Fields{
		"prefix": "sensor",
		"animal": "orca",
		"size":   9009,
	}).Panic("It's over 9000!")
}
