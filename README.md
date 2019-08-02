# 日志


## 示例

```golang
import "github.com/boxungo/logger"

log := logger.New(os.Stdout, logger.DebugLevel)

log.Debugf("Timestamp: %v", "time")
log.Infof("Timestamp: %v", "time")
log.Warnf("Timestamp: %v", "time")
log.Errorf("Timestamp: %v", "time")
log.Fatalf("Timestamp: %v", "time")

log.Debug(" out Debug")
log.Info(" out Info")
log.Warn(" out Warn")
log.Error(" out Error")
log.Fatal(" out Fatal")
```

```text
2019-08-02 16:50:26.453078 [DEBUG] Timestamp: time
2019-08-02 16:50:26.453195 [INFO] Timestamp: time
2019-08-02 16:50:26.453208 [WARN] Timestamp: time
2019-08-02 16:50:26.453218 [ERROR] Timestamp: time
2019-08-02 16:50:26.453221 [FATAL] Timestamp: time
2019-08-02 16:50:26.453224 [DEBUG]  out Debug
2019-08-02 16:50:26.453250 [INFO]  out Info
2019-08-02 16:50:26.453258 [WARN]  out Warn
2019-08-02 16:50:26.453268 [ERROR]  out Error
2019-08-02 16:50:26.453272 [FATAL]  out Fatal
```

## 自定义输出格式

```golang
func customFormatHandler(out io.Writer, level int, arg string) {
    str := fmt.Sprintf("%v - %s", level, arg)
    out.Write([]byte(str))
}

log.SetFormatHandler(customFormatHandler)
```