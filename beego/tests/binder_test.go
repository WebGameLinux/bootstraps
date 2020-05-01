package tests

import (
		"fmt"
		"github.com/WebGameLinux/bootstraps/beego/container/functions"
		"testing"
)

func Name() {
		fmt.Println("name123")
}

func TestIsFunc(t *testing.T) {
		var (
				fn func(t *testing.T)
		)
		fn = TestIsFunc
		if functions.IsFunc(fn) != true {
				t.Errorf("match function type error")
		}
}

func TestFunctionBindWrapper_Bind(t *testing.T) {
		var (
				fn  func(t *testing.T)
				fn2 func(t *testing.T)
				fn3 func()
		)
		fn3 = Name
		fn = TestIsFunc
		binder := functions.GetFunctionBinder()
		binder.Bind(&fn2, fn)
		err := binder.GetError()
		if err != nil {
				t.Errorf("%s", err)
		}
		binder.Bind(&fn2, fn3)
		err = binder.GetError()
		if err == nil {
				t.Errorf("绑定异常")
		}
		fmt.Println(err)

}

func TestFunctionBindWrapper_Wrap(t *testing.T) {
		var (
				fn  func(t *testing.T)
				fn2 func(t *testing.T)
		)
		fn = TestIsFunc
		functions.Bind(&fn2, fn)
		err := functions.GetError()
		if err != nil {
				t.Errorf("%s", err)
		}
		functions.Wrap(t, fn)()
}
