package configure

import (
	"fmt"
	"reflect"
)

// VersionParseInfo 用来定义一个版本，如果进行解析
// 不同的版本，使用的解析器是不同的
type VersionParseInfo struct {
	Version Version

	AsParse reflect.Type

	ConverConfigFunc func(interface{}) (interface{}, error)
}

func MajorMinorVersion(major, minor int) Version {
	return Version(fmt.Sprintf("%d.%d", major, minor))
}
