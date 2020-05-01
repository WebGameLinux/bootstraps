package functions

import (
		"errors"
		"fmt"
		"github.com/WebGameLinux/bootstraps/beego"
		"reflect"
		"regexp"
)

//var argsFuncRegexp = regexp.MustCompile(`func\((.+)\)`)
var returnAnyRegexp = regexp.MustCompile(`func\(.*\)(.*)`)

type FunctionBinder interface {
		GetError() error
		Bind(methodTypePoint interface{}, methodInstancePoint interface{})
		Wrap(ctx interface{}, function interface{}, returnStack ...[]interface{}) func()
}

type FunctionBindDto struct {
		Err error
}

type FunctionBindWrapper struct {
		FunctionBindDto
}

var function FunctionBinder

func (this *FunctionBindWrapper) GetError() error {
		defer func() {
				this.Err = nil
		}()
		return this.Err
}

// 函数绑定器
// @param methodTypePoint interface : 函数地址容器 (函数变量的地址)
// @param methodInstancePoint interface : 函数指针 (函数的地址)
func (this *FunctionBindWrapper) Bind(methodTypePoint interface{}, methodInstancePoint interface{}) {
		if !IsFunc(methodInstancePoint) || !IsFunc(methodInstancePoint) {
				this.Err = errors.New("TypeError:bind param type has not function point type")
				return
		}
		if methodTypePoint == nil || methodInstancePoint == nil {
				this.Err = errors.New("TypeError:bind param type has not function point type,but give nil point params")
				return
		}
		if !BindFunc(methodTypePoint, methodInstancePoint) {
				logTxt := fmt.Sprintf("T1:%T , T2:%T", methodTypePoint, methodInstancePoint)
				this.Err = errors.New("TypeError:bind function type not match, " + logTxt)
		}
}

// 函数包装器
func (this *FunctionBindWrapper) Wrap(ctx interface{}, function interface{}, returnStack ...[]interface{}) func() {
		if !IsFunc(function) {
				this.Err = errors.New("TypeError:bind param type has not function point type")
				return nil
		}
		return func() {
				var args []reflect.Value
				fn := reflect.ValueOf(function)
				if arg, ok := ctx.([]reflect.Value); ok {
						fn.Call(arg)
						return
				}
				if arg, ok := ctx.(reflect.Value); ok {
						args = append(args, arg)
				} else {
						args = append(args, reflect.ValueOf(ctx))
				}
				args = fn.Call(args)
				if len(args) > 0 && len(returnStack) > 0 {
						stack := returnStack[0]
						// 空间足时
						if cap(stack) < len(args) {
								stack[0] = len(args)
								stack[1] = args
								return
						}
						// 正常返回
						for i, v := range args {
								stack[i] = v.Interface()
						}
				}
		}
}

// 是否函数类型
func IsFunc(v interface{}) bool {
		var t = reflect.TypeOf(v)
		t = beego.RealType(t)
		if t.Kind() != reflect.Func {
				return false
		}
		return true
}

// 函数绑定
func BindFunc(methodTypePoint, methodInstancePoint interface{}) bool {
		var (
				mP      = reflect.TypeOf(methodTypePoint)
				mI      = reflect.TypeOf(methodInstancePoint)
				TyKind  = mP.Kind()
				InsKind = mI.Kind()
		)
		if TyKind != InsKind {
				v := reflect.ValueOf(methodTypePoint).Elem()
				ins := reflect.ValueOf(methodInstancePoint)
				if v.Kind() == InsKind && v.String() == ins.String() {
						v.Set(ins)
						return true
				}
		}
		return false
}

func HasReturn(t reflect.Type) bool {
		if t.Kind() == reflect.Func {
				return returnAnyRegexp.MatchString(t.Kind().String())
		}
		return false
}

func Bind(methodTypePoint interface{}, methodInstancePoint interface{}) {
		function.Bind(methodTypePoint, methodInstancePoint)
}

func GetError() error {
		return function.GetError()
}

func BindOk() bool {
		return function.GetError() == nil
}

func Wrap(ctx interface{}, fun interface{}, returnStack ...[]interface{}) func() {
		return function.Wrap(ctx, fun, returnStack...)
}

func GetFunctionBinder() FunctionBinder {
		return function
}
