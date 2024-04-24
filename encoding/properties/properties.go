package properties

import (
	"bytes"
	"github.com/magiconair/properties"
	"log"
	"reflect"
	"strings"
)

const Name = "properties"

var DefaultCodec = &codec{}

type codec struct{}

// Name 编解码器名称
func (codec) Name() string {
	return Name
}

// Marshal 编码
func (codec) Marshal(v interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	return buffer.Bytes(), nil
}

type Mys struct {
}

// Unmarshal 解码
func (codec) Unmarshal(data []byte, v interface{}) error {
	p, err := properties.Load(data, properties.UTF8)
	if mapPtr, ok := v.(*map[string]interface{}); ok {
		t := make(map[string]interface{})
		for key, value := range p.Map() {
			t[key] = value
		}
		*mapPtr = t
	} else {
		log.Println("[properties][encode] x is not a pointer to a map[string]interface{}")
	}
	c := make(map[string]interface{})
	mergeConfig(c, *v.(*map[string]interface{}))

	vPtr := reflect.ValueOf(v)
	if vPtr.Kind() == reflect.Ptr && vPtr.Elem().Kind() == reflect.Map {
		vPtr.Elem().Set(reflect.ValueOf(c))
	} else {
		log.Println("[properties][encode] v is not a pointer to a map[string]interface{}")
	}
	return err
}

func mergeConfig(config map[string]interface{}, keysAndValues map[string]interface{}) {
	for key, value := range keysAndValues {
		parts := strings.Split(key, ".")
		if len(parts) < 2 {
			continue
		}
		lastKey := parts[len(parts)-1]
		lastValue := value

		for i := len(parts) - 2; i >= 0; i-- {
			innerKey := parts[i]
			innerValue := make(map[string]interface{})
			innerValue[lastKey] = lastValue

			lastKey = innerKey
			lastValue = innerValue
		}

		// 判断最外层键是否存在，如果存在则进行合并，否则直接赋值
		if existingValue, ok := config[parts[0]]; ok {
			if existingMap, ok := existingValue.(map[string]interface{}); ok {
				// 合并现有的映射和新的嵌套映射
				mergeMaps(existingMap, lastValue.(map[string]interface{}))
			} else {
				// 如果最外层键存在但是值不是映射类型，则直接赋值
				config[parts[0]] = lastValue
			}
		} else {
			config[parts[0]] = lastValue
		}
	}
}

func mergeMaps(target map[string]interface{}, source map[string]interface{}) {
	for key, value := range source {
		if targetValue, ok := target[key]; ok {
			if targetMap, ok := targetValue.(map[string]interface{}); ok {
				// 合并现有的映射和新的嵌套映射
				if sourceMap, ok := value.(map[string]interface{}); ok {
					mergeMaps(targetMap, sourceMap)
				} else {
					// 如果现有的值是映射类型，但新的值不是映射类型，则直接覆盖
					targetMap[key] = value
				}
			} else {
				// 如果现有的值不是映射类型，则直接覆盖
				target[key] = value
			}
		} else {
			target[key] = value
		}
	}
}

// Marshal 编码
func Marshal(v interface{}) ([]byte, error) {
	return DefaultCodec.Marshal(v)
}

// Unmarshal 解码
func Unmarshal(data []byte, v interface{}) error {
	return DefaultCodec.Unmarshal(data, v)
}
