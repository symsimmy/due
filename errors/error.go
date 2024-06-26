package errors

import (
	"fmt"
	"github.com/symsimmy/due/common/stack"
	"strings"
)

type Error struct {
	err   error
	code  int
	text  string
	stack *stack.Stack
}

// New 新建一个错误
// 可传入一下参数：
// text : 文本字符串
func New(text string) *Error {
	e := &Error{
		code:  -1,
		stack: stack.Callers(1, stack.Full),
	}

	e.text = text

	return e
}

// NewError 新建一个错误
// 可传入一下参数：
// text : 文本字符串
// code : 错误码
// error: 原生错误
func NewError(args ...interface{}) *Error {
	e := &Error{
		code:  -1,
		stack: stack.Callers(1, stack.Full),
	}

	for _, arg := range args {
		switch v := arg.(type) {
		case error:
			e.err = v
		case string:
			e.text = v
		case int:
			e.code = v
		}
	}

	return e
}

func (e *Error) Error() (text string) {
	if e == nil {
		return
	}

	text = e.text

	return
}

func Is(e error, err error) bool {
	if e != nil && err != nil {
		return strings.EqualFold(e.Error(), err.Error())
	}
	return e == nil && err == nil
}

func (e *Error) Is(err error) bool {
	if err != nil {
		return strings.EqualFold(e.Error(), err.Error())
	}
	return false
}

func (e *Error) ErrorWithArgs(args ...interface{}) (text string) {
	if e == nil {
		return
	}

	text = fmt.Sprintf(e.text, args)
	return
}

// Code 返回错误码
func (e *Error) Code() int32 {
	if e == nil {
		return -1
	}

	return int32(e.code)
}

// Next 返回下一个错误
func (e *Error) Next() error {
	if e == nil {
		return nil
	}

	return e.err
}
