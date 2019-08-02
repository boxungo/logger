package logger

import (
	"fmt"
	"io"

	"github.com/boxungo/now"
)

// defaultFormatHandler 默认的日志格式化函数
func defaultFormatHandler(out io.Writer, level int, arg string) {
	str := fmt.Sprintf("%v [%s] %s\n", now.String(), levelName[level], arg)

	_, err := out.Write([]byte(str))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
