/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2022/9/1 8:50 下午
 * @Desc: TODO
 */

package encoder

import (
	"fmt"
	"strings"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type AsyncTextEncoder struct {
	zapcore.ObjectEncoder
	bufferPool     buffer.Pool
	timeFormat     string
	callerFullPath bool
	isTerminal     bool
}

var _ zapcore.Encoder = &AsyncTextEncoder{}

func NewAsyncTextEncoder(timeFormat string, callerFullPath, isTerminal bool) zapcore.Encoder {
	return &AsyncTextEncoder{
		bufferPool:     buffer.NewPool(),
		timeFormat:     timeFormat,
		callerFullPath: callerFullPath,
		isTerminal:     isTerminal,
	}
}

func (e *AsyncTextEncoder) Clone() zapcore.Encoder {
	return nil
}

func (e *AsyncTextEncoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	line := e.bufferPool.Get()

	levelText := ent.Level.CapitalString()[0:4]
	if e.isTerminal {
		var levelColor int
		switch ent.Level {
		case zapcore.DebugLevel:
			levelColor = gray
		case zapcore.WarnLevel:
			levelColor = yellow
		case zapcore.ErrorLevel, zapcore.FatalLevel, zapcore.PanicLevel:
			levelColor = red
		case zapcore.InfoLevel:
			levelColor = blue
		default:
			levelColor = blue
		}
		line.AppendString(fmt.Sprintf("\x1b[%dm%s", levelColor, levelText))
		line.AppendString(fmt.Sprintf("\x1b[0m[%s] ", ent.Time.Format(e.timeFormat)))
	} else {
		line.AppendString(levelText)
		line.AppendString(fmt.Sprintf("[%s] ", ent.Time.Format(e.timeFormat)))
	}

	line.AppendString(strings.TrimSuffix(ent.Message, "\n"))

	line.AppendString("\n")

	return line, nil
}
