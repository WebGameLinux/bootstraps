package beego

import (
		"reflect"
		"strings"
		"time"
)

func ClassName(v interface{}) string {
		if v == nil {
				return "<nil>"
		}
		var t = reflect.TypeOf(v)
		t = RealType(t)
		className := t.PkgPath() + "::" + t.Name()
		if className == "::" {
				return t.String()
		}
		return className
}

func IsPointer(t reflect.Type) bool {
		return strings.Contains(t.String(), "*")
}

func RealType(t reflect.Type) reflect.Type {
		for IsPointer(t) {
				t = t.Elem()
		}
		return t
}

func Start(boot BootStrap) {
		boot.Container("bootstrap.startAt", time.Now())
		boot.Container("bootstrap.class", ClassName(boot))
		boot.Boot()
}
