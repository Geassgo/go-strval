/*
--------------------------------
@Create 2025/10/16 11:45
@Author lengpucheng<lpc@hll520.cn>
@Project go-strval
@Version 1.0.0 2025/10/16 11:45
@Description 增强的基本类型支持，实现从字符串形式的JSON/YAML反序列化功能
--------------------------------
本文件实现了三个主要类型：Bool、Int和Float，这些类型分别包装了Go的基本类型bool、int和float64。
主要功能包括：
1. 支持从字符串形式的JSON/YAML值反序列化为对应的基本类型
2. 提供友好的错误处理机制，当解析失败时返回零值并记录错误日志
3. 序列化为JSON/YAML时保持原始类型格式
*/

package strval

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"log/slog"

	"gopkg.in/yaml.v3"
)

// StringValuer 定义所有字符串值包装类型的共同泛型接口
type StringValuer[T any] interface {
	// GetValue 获取包装的原始值
	GetValue() T
}

// Bool 增强的布尔类型，支持从字符串形式的JSON/YAML反序列化
type Bool bool

// MarshalJSON 实现json.Marshaler接口，将Bool序列化为JSON布尔值
// 返回值:
//   - []byte: 序列化后的JSON字节
//   - error: 序列化过程中的错误
func (b Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(b))
}

// UnmarshalJSON 实现json.Unmarshaler接口，支持从JSON布尔值或字符串反序列化为Bool
// 参数:
//   - data: JSON数据字节
// 返回值:
//   - error: 反序列化过程中的错误
// 说明:
//   - 支持直接解析JSON布尔值
//   - 支持解析字符串形式的布尔值（如"true"、"false"、"yes"、"no"、"1"、"0"）
//   - 解析失败时返回false并记录错误日志
func (b *Bool) UnmarshalJSON(data []byte) error {
	// 尝试直接解析为bool
	var boolVal bool
	if err := json.Unmarshal(data, &boolVal); err == nil {
		*b = Bool(boolVal)
		return nil
	}

	// 尝试解析为字符串
	var strVal string
	if err := json.Unmarshal(data, &strVal); err != nil {
		*b = false
		slog.Error("invalid Bool value: not a bool or string", "error", err)
		return nil
	}

	// 解析字符串形式的bool值
	boolVal, err2 := parseBool(strVal)
	if err2 != nil {
		*b = false
		slog.Error("invalid Bool string value", "value", strVal, "error", err2)
		return nil
	}

	*b = Bool(boolVal)
	return nil
}

// MarshalYAML 实现yaml.Marshaler接口，将Bool序列化为YAML布尔值
// 返回值:
//   - interface{}: 序列化后的值
//   - error: 序列化过程中的错误
func (b Bool) MarshalYAML() (interface{}, error) {
	return bool(b), nil
}

// GetValue 实现StringValuer[bool]接口，获取包装的原始布尔值
// 返回值:
//   - bool: 原始的bool值
func (b Bool) GetValue() bool {
	return bool(b)
}

// Value 实现driver.Valuer接口，用于数据库写入操作
// 返回值:
//   - driver.Value: 数据库可接受的值
//   - error: 转换过程中的错误
func (b Bool) Value() (driver.Value, error) {
	return b.GetValue(), nil
}

// Scan 实现sql.Scanner接口，用于数据库读取操作
// 参数:
//   - value: 从数据库读取的值
// 返回值:
//   - error: 扫描过程中的错误
func (b *Bool) Scan(value interface{}) error {
	if value == nil {
		*b = false
		return nil
	}

	// 尝试直接转换为bool
	if boolVal, ok := value.(bool); ok {
		*b = Bool(boolVal)
		return nil
	}

	// 尝试从int转换
	if intVal, ok := value.(int64); ok {
		*b = Bool(intVal != 0)
		return nil
	}

	// 尝试从string转换
	if strVal, ok := value.(string); ok {
		boolVal, err := parseBool(strVal)
		if err != nil {
			*b = false
			slog.Error("invalid Bool value from database", "value", strVal, "error", err)
			return nil
		}
		*b = Bool(boolVal)
		return nil
	}

	*b = false
	slog.Error("unsupported Bool value type from database", "type", fmt.Sprintf("%T", value))
	return nil
}

// UnmarshalYAML 实现yaml.Unmarshaler接口，支持从YAML布尔值或字符串反序列化为Bool
// 参数:
//   - node: YAML节点
// 返回值:
//   - error: 反序列化过程中的错误
// 说明:
//   - 支持直接解析YAML布尔值
//   - 支持解析字符串形式的布尔值
//   - 解析失败时返回false并记录错误日志
func (b *Bool) UnmarshalYAML(node *yaml.Node) error {
	// 尝试直接解析为bool
	var boolVal bool
	if err := node.Decode(&boolVal); err == nil {
		*b = Bool(boolVal)
		return nil
	}

	// 尝试解析为字符串
	var strVal string
	if err := node.Decode(&strVal); err != nil {
		*b = false
		slog.Error("invalid Bool value: not a bool or string", "error", err)
		return nil
	}

	// 解析字符串形式的bool值
	boolVal, err2 := parseBool(strVal)
	if err2 != nil {
		*b = false
		slog.Error("invalid Bool string value", "value", strVal, "error", err2)
		return nil
	}

	*b = Bool(boolVal)
	return nil
}

// Int 增强的整型，支持从字符串形式的JSON/YAML反序列化
type Int int

// MarshalJSON 实现json.Marshaler接口，将Int序列化为JSON数值
// 返回值:
//   - []byte: 序列化后的JSON字节
//   - error: 序列化过程中的错误
func (i Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(i))
}

// UnmarshalJSON 实现json.Unmarshaler接口，支持从JSON数值或字符串反序列化为Int
// 参数:
//   - data: JSON数据字节
// 返回值:
//   - error: 反序列化过程中的错误
// 说明:
//   - 支持直接解析JSON数值
//   - 支持解析字符串形式的整数值
//   - 解析失败时返回0并记录错误日志
func (i *Int) UnmarshalJSON(data []byte) error {
	// 尝试直接解析为int
	var intVal int
	if err := json.Unmarshal(data, &intVal); err == nil {
		*i = Int(intVal)
		return nil
	}

	// 尝试解析为字符串
	var strVal string
	if err := json.Unmarshal(data, &strVal); err != nil {
		*i = 0
		slog.Error("invalid Int value: not an int or string", "error", err)
		return nil
	}

	// 解析字符串形式的int值
	intVal, err2 := strconv.Atoi(strVal)
	if err2 != nil {
		*i = 0
		slog.Error("invalid Int string value", "value", strVal, "error", err2)
		return nil
	}

	*i = Int(intVal)
	return nil
}

// MarshalYAML 实现yaml.Marshaler接口，将Int序列化为YAML数值
// 返回值:
//   - interface{}: 序列化后的值
//   - error: 序列化过程中的错误
func (i Int) MarshalYAML() (interface{}, error) {
	return int(i), nil
}

// GetValue 实现StringValuer[int]接口，获取包装的原始整数值
// 返回值:
//   - int: 原始的int值
func (i Int) GetValue() int {
	return int(i)
}

// Value 实现driver.Valuer接口，用于数据库写入操作
// 返回值:
//   - driver.Value: 数据库可接受的值
//   - error: 转换过程中的错误
func (i Int) Value() (driver.Value, error) {
	return i.GetValue(), nil
}

// Scan 实现sql.Scanner接口，用于数据库读取操作
// 参数:
//   - value: 从数据库读取的值
// 返回值:
//   - error: 扫描过程中的错误
func (i *Int) Scan(value interface{}) error {
	if value == nil {
		*i = 0
		return nil
	}

	// 尝试直接转换为int64
	if int64Val, ok := value.(int64); ok {
		*i = Int(int64Val)
		return nil
	}

	// 尝试从float转换
	if floatVal, ok := value.(float64); ok {
		*i = Int(floatVal)
		return nil
	}

	// 尝试从string转换
	if strVal, ok := value.(string); ok {
		intVal, err := strconv.Atoi(strVal)
		if err != nil {
			*i = 0
			slog.Error("invalid Int value from database", "value", strVal, "error", err)
			return nil
		}
		*i = Int(intVal)
		return nil
	}

	*i = 0
	slog.Error("unsupported Int value type from database", "type", fmt.Sprintf("%T", value))
	return nil
}

// UnmarshalYAML 实现yaml.Unmarshaler接口，支持从YAML数值或字符串反序列化为Int
// 参数:
//   - node: YAML节点
// 返回值:
//   - error: 反序列化过程中的错误
// 说明:
//   - 支持直接解析YAML数值
//   - 支持解析字符串形式的整数值
//   - 解析失败时返回0并记录错误日志
func (i *Int) UnmarshalYAML(node *yaml.Node) error {
	// 尝试直接解析为int
	var intVal int
	if err := node.Decode(&intVal); err == nil {
		*i = Int(intVal)
		return nil
	}

	// 尝试解析为字符串
	var strVal string
	if err := node.Decode(&strVal); err != nil {
		*i = 0
		slog.Error("invalid Int value: not an int or string", "error", err)
		return nil
	}

	// 解析字符串形式的int值
	intVal, err2 := strconv.Atoi(strVal)
	if err2 != nil {
		*i = 0
		slog.Error("invalid Int string value", "value", strVal, "error", err2)
		return nil
	}

	*i = Int(intVal)
	return nil
}

// Float 增强的浮点型，支持从字符串形式的JSON/YAML反序列化
type Float float64

// MarshalJSON 实现json.Marshaler接口，将Float序列化为JSON数值
// 返回值:
//   - []byte: 序列化后的JSON字节
//   - error: 序列化过程中的错误
func (f Float) MarshalJSON() ([]byte, error) {
	return json.Marshal(float64(f))
}

// UnmarshalJSON 实现json.Unmarshaler接口，支持从JSON数值或字符串反序列化为Float
// 参数:
//   - data: JSON数据字节
// 返回值:
//   - error: 反序列化过程中的错误
// 说明:
//   - 支持直接解析JSON数值
//   - 支持解析字符串形式的浮点数值
//   - 解析失败时返回0并记录错误日志
func (f *Float) UnmarshalJSON(data []byte) error {
	// 尝试直接解析为float64
	var floatVal float64
	if err := json.Unmarshal(data, &floatVal); err == nil {
		*f = Float(floatVal)
		return nil
	}

	// 尝试解析为字符串
	var strVal string
	if err := json.Unmarshal(data, &strVal); err != nil {
		*f = 0
		slog.Error("invalid Float value: not a float or string", "error", err)
		return nil
	}

	// 解析字符串形式的float值
	floatVal, err2 := strconv.ParseFloat(strVal, 64)
	if err2 != nil {
		*f = 0
		slog.Error("invalid Float string value", "value", strVal, "error", err2)
		return nil
	}

	*f = Float(floatVal)
	return nil
}

// MarshalYAML 实现yaml.Marshaler接口，将Float序列化为YAML数值
// 返回值:
//   - interface{}: 序列化后的值
//   - error: 序列化过程中的错误
func (f Float) MarshalYAML() (interface{}, error) {
	return float64(f), nil
}

// GetValue 实现StringValuer[float64]接口，获取包装的原始浮点值
// 返回值:
//   - float64: 原始的float64值
func (f Float) GetValue() float64 {
	return float64(f)
}



// Value 实现driver.Valuer接口，用于数据库写入操作
// 返回值:
//   - driver.Value: 数据库可接受的值
//   - error: 转换过程中的错误
func (f Float) Value() (driver.Value, error) {
	return f.GetValue(), nil
}

// Scan 实现sql.Scanner接口，用于数据库读取操作
// 参数:
//   - value: 从数据库读取的值
// 返回值:
//   - error: 扫描过程中的错误
func (f *Float) Scan(value interface{}) error {
	if value == nil {
		*f = 0
		return nil
	}

	// 尝试直接转换为float64
	if floatVal, ok := value.(float64); ok {
		*f = Float(floatVal)
		return nil
	}

	// 尝试从int转换
	if intVal, ok := value.(int64); ok {
		*f = Float(intVal)
		return nil
	}

	// 尝试从string转换
	if strVal, ok := value.(string); ok {
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			*f = 0
			slog.Error("invalid Float value from database", "value", strVal, "error", err)
			return nil
		}
		*f = Float(floatVal)
		return nil
	}

	*f = 0
	slog.Error("unsupported Float value type from database", "type", fmt.Sprintf("%T", value))
	return nil
}

// UnmarshalYAML 实现yaml.Unmarshaler接口，支持从YAML数值或字符串反序列化为Float
// 参数:
//   - node: YAML节点
// 返回值:
//   - error: 反序列化过程中的错误
// 说明:
//   - 支持直接解析YAML数值
//   - 支持解析字符串形式的浮点数值
//   - 解析失败时返回0并记录错误日志
func (f *Float) UnmarshalYAML(node *yaml.Node) error {
	// 尝试直接解析为float64
	var floatVal float64
	if err := node.Decode(&floatVal); err == nil {
		*f = Float(floatVal)
		return nil
	}

	// 尝试解析为字符串
	var strVal string
	if err := node.Decode(&strVal); err != nil {
		*f = 0
		slog.Error("invalid Float value: not a float or string", "error", err)
		return nil
	}

	// 解析字符串形式的float值
	floatVal, err2 := strconv.ParseFloat(strVal, 64)
	if err2 != nil {
		*f = 0
		slog.Error("invalid Float string value", "value", strVal, "error", err2)
		return nil
	}

	*f = Float(floatVal)
	return nil
}

// parseBool 解析字符串形式的布尔值
// 参数:
//   - s: 输入字符串
// 返回值:
//   - bool: 解析后的布尔值
//   - error: 解析过程中的错误
// 支持的值:
//   - 真值: "true", "yes", "y", "1"
//   - 假值: "false", "no", "n", "0"
// 所有值不区分大小写，会自动去除前后空格
func parseBool(s string) (bool, error) {
	s = strings.ToLower(strings.TrimSpace(s))
	switch s {
	case "true", "yes", "y", "1":
		return true, nil
	case "false", "no", "n", "0":
		return false, nil
	default:
		return false, fmt.Errorf("cannot parse '%s' as bool", s)
	}
}