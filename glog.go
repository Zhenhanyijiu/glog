package glog

import ()

type Log struct {
	*Logger
}

//func New(level string, proname string, w io.Writer) *Log {
//	//log:=new(Log)
//	lg := logrus.New()
//	lv, err := logrus.ParseLevel(level)
//	if err != nil {
//		panic(err)
//	}
//	lg.SetLevel(lv)
//	lg.SetOutput(w)
//	lg.SetReportCaller(true) //显示行号等信息
//	lg.SetFormatter(&logrus.TextFormatter{
//		TimestampFormat: "2006-01-02 15:04:05", //时间格式化
//	})
//	//lg.WithFields(logrus.Fields{
//	//	"animal": "walrus",
//	//})
//	lg.AddHook(&DefaultFieldHook{prcName: proname})
//	log := &Log{lg}
//	return log
//}
//
//type DefaultFieldHook struct {
//	prcName string
//}
//
//func (hook *DefaultFieldHook) Fire(entry *logrus.Entry) error {
//	entry.Data["[ ProcName ]"] = "[" + hook.prcName + "]"
//	return nil
//}
//
//func (hook *DefaultFieldHook) Levels() []logrus.Level {
//	return logrus.AllLevels
//}
