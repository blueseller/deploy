package configure

// parse: 解析一个yaml文件,并且可以处理相关环境变量
// 1. 支持yaml 解析能力
// 2. 支持环境变量的配置使用
// 3. 支持多版本 解析
// 4. 环境变量的配置, 优先yaml配置

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type EnvVar struct {
	name  string
	value string
}

type envVals []EnvVar

type Parser struct {
	EnvPrefix string

	mapping map[Version]VersionParseInfo

	env envVals
}

func NewParser(prefix string, parseInfos []VersionParseInfo) *Parser {
	parser := &Parser{
		EnvPrefix: prefix,
		mapping:   make(map[Version]VersionParseInfo),
	}

	for _, parseInfo := range parseInfos {
		parser.mapping[parseInfo.Version] = parseInfo
	}

	for _, env := range os.Environ() {
		envVal := strings.SplitN(env, "=", 2)
		parser.env = append(parser.env, EnvVar{envVal[0], envVal[1]})
	}

	//sort(parser.env)
	return parser
}

func (p *Parser) overwriteFields(v reflect.Value, fullPath string, paths []string, value string) error {
	// 如果reflect.Value 是一个指针, 则一直循环, 直到Indirect 取出的值为struct value 后退出
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			panic("encountered nil pointer while handling enviroment variable," + fullPath)
		}
		v = reflect.Indirect(v)
	}

	switch v.Kind() {
	case reflect.Struct:
		p.overwriteStruct(v, fullPath, paths, value)
	case reflect.Map:
		p.overwriteMap(v, fullPath, paths, value)
	case reflect.Interface:
		if v.NumMethod() == 0 {
			if !v.IsNil() {
				return p.overwriteFields(v.Elem(), fullPath, paths, value)
			}

			// 如果interface 为空，则初始化一个
			var template map[string]interface{}
			wrappedV := reflect.MakeMap(reflect.TypeOf(template))
			v.Set(wrappedV)
			return p.overwriteMap(v, fullPath, paths, value)
		}
	}
	return nil
}

func (p *Parser) overwriteStruct(m reflect.Value, fullPath string, paths []string, value string) error {
	flag := false

	for i := 0; i < m.NumField(); i++ {
		if paths[0] == strings.ToUpper(m.Type().Field(i).Name) {
			field := m.Field(i)
			if len(paths) == 1 {
				strVal := reflect.New(m.Type().Field(i).Type)
				err := yaml.Unmarshal([]byte(value), strVal.Interface)
				if err != nil {
					return err
				}
				field.Set(reflect.Indirect(strVal))
			} else {
				if m.Type().Field(i).Type.Kind() == reflect.Map {
					if field.IsNil() {
						field.Set(reflect.MakeMap(m.Type().Field(i).Type))
					}
				}

				if m.Type().Field(i).Type.Kind() == reflect.Ptr {
					if field.IsNil() {
						field.Set(reflect.New(field.Elem().Type()))
					}
				}

				err := p.overwriteFields(field, fullPath, paths[1:], value)
				if err != nil {
					return err
				}
			}
			flag = true
		}
	}

	if !flag {
		logrus.Warnf("ignoring unrecognized enviroment variable %s", fullPath)
	}
	return nil

}

func (p *Parser) overwriteMap(m reflect.Value, fullPath string, paths []string, value string) error {
	// key 必须为string
	if m.Type().Key().Kind() != reflect.String {
		logrus.Warnf("ignoring enviroment variable %s invoving map with non-string keys", fullPath)
		return nil
	}

	for _, k := range m.MapKeys() {
		if strings.ToUpper(k.String()) == paths[0] {
			mapValue := m.MapIndex(k)
			var err error
			if len(paths) > 1 {
				// 如果value 为空的话, 我们需要重新创建出来一个
				if (mapValue.Kind() == reflect.Ptr ||
					mapValue.Kind() == reflect.Interface ||
					mapValue.Kind() == reflect.Map) &&
					mapValue.IsNil() {

					templateMap := reflect.MakeMap(m.Type().Elem())
					err = p.overwriteFields(templateMap, fullPath, paths[1:], value)
				} else {
					err = p.overwriteFields(mapValue, fullPath, paths[1:], value)
				}
			} else {
				mapValue = reflect.New(m.Type().Elem())
				err = yaml.Unmarshal([]byte(value), mapValue.Interface())
			}
			if err != nil {
				return err
			}

			m.SetMapIndex(reflect.ValueOf(strings.ToLower(paths[0])), reflect.Indirect(mapValue))
			break
		}
	}

	return nil
}

func (p *Parser) Parse(in []byte, v interface{}) error {
	var versionStruct struct {
		Version Version
	}

	if err := yaml.Unmarshal(in, &versionStruct); err != nil {
		return err
	}

	parseInfo, ok := p.mapping[versionStruct.Version]
	if !ok {
		return fmt.Errorf("unsupported version: %q", versionStruct.Version)
	}

	parseAs := reflect.New(parseInfo.AsParse)
	if err := yaml.Unmarshal(in, parseAs.Interface()); err != nil {
		return err
	}

	for _, envVal := range p.env {
		name := envVal.name

		if strings.HasPrefix(name, strings.ToUpper(p.EnvPrefix)+"_") {
			paths := strings.Split(name, "_")
			err := p.overwriteFields(parseAs, name, paths[1:], envVal.value)

			if err != nil {
				return err
			}
		}
	}

	c, err := parseInfo.ConverConfigFunc(parseAs.Interface())
	if err != nil {
		return err
	}
	reflect.ValueOf(v).Elem().Set(reflect.Indirect(reflect.ValueOf(c)))
	return nil
}

func Parse(fp io.Reader) (*Configuration, error) {
	in, err := ioutil.ReadAll(fp)
	if err != nil {
		return nil, err
	}

	p := NewParser("deploy", []VersionParseInfo{
		{
			Version: MajorMinorVersion(1, 0),
			AsParse: reflect.TypeOf(v0_1Configuration{}),
			ConverConfigFunc: func(c interface{}) (interface{}, error) {
				if v0_1, ok := c.(*v0_1Configuration); ok {
					if v0_1.Log.LogLevel == LogLevel("") {
						if v0_1.LogLevel != LogLevel("") {
							v0_1.Log.LogLevel = v0_1.LogLevel
						} else {
							v0_1.Log.LogLevel = LogLevel("info")
						}
					}

					if v0_1.LogLevel != LogLevel("") {
						v0_1.LogLevel = LogLevel("")
					}
					return (*Configuration)(v0_1), nil
				}
				return nil, fmt.Errorf("unexported *v0_1Configuration,received %#v", c)
			},
		},
	})

	config := new(Configuration)
	err = p.Parse(in, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
