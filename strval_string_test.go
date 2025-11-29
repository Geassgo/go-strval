package strval

import (
	"encoding/json"
	"testing"
)

// TestStringJSONMarshal 测试String类型的JSON序列化功能
func TestStringJSONMarshal(t *testing.T) {
	t.Run("String serialization", func(t *testing.T) {
		// 测试字符串值的序列化
		s := String("hello")
		data, err := json.Marshal(s)
		if err != nil {
			t.Fatalf("failed to marshal String: %v", err)
		}
		if string(data) != `"hello"` {
			t.Errorf("expected \"hello\", got %s", data)
		}
	})

	t.Run("Number serialization", func(t *testing.T) {
		// 测试从数值创建的字符串序列化
		s := String("123")
		data, err := json.Marshal(s)
		if err != nil {
			t.Fatalf("failed to marshal String: %v", err)
		}
		if string(data) != `"123"` {
			t.Errorf("expected \"123\", got %s", data)
		}
	})

	t.Run("Bool serialization", func(t *testing.T) {
		// 测试从布尔值创建的字符串序列化
		s := String("true")
		data, err := json.Marshal(s)
		if err != nil {
			t.Fatalf("failed to marshal String: %v", err)
		}
		if string(data) != `"true"` {
			t.Errorf("expected \"true\", got %s", data)
		}
	})
}

// TestStringJSONUnmarshal 测试String类型的JSON反序列化功能
func TestStringJSONUnmarshal(t *testing.T) {
	t.Run("From string", func(t *testing.T) {
		// 测试从JSON字符串反序列化
		var s String
		data := []byte(`"hello"`)
		err := json.Unmarshal(data, &s)
		if err != nil {
			t.Fatalf("failed to unmarshal String from string: %v", err)
		}
		if string(s) != "hello" {
			t.Errorf("expected hello, got %s", s)
		}
	})

	t.Run("From number", func(t *testing.T) {
		// 测试从JSON数值反序列化 - 关键测试点，确保数值能转为字符串
		var s String
		data := []byte(`123`)
		err := json.Unmarshal(data, &s)
		if err != nil {
			t.Fatalf("failed to unmarshal String from number: %v", err)
		}
		if string(s) != "123" {
			t.Errorf("expected \"123\", got %s", s)
		}
	})

	t.Run("From bool", func(t *testing.T) {
		// 测试从JSON布尔值反序列化
		var s String
		data := []byte(`true`)
		err := json.Unmarshal(data, &s)
		if err != nil {
			t.Fatalf("failed to unmarshal String from bool: %v", err)
		}
		if string(s) != "true" {
			t.Errorf("expected \"true\", got %s", s)
		}
	})

	t.Run("From float", func(t *testing.T) {
		// 测试从JSON浮点数反序列化
		var s String
		data := []byte(`123.45`)
		err := json.Unmarshal(data, &s)
		if err != nil {
			t.Fatalf("failed to unmarshal String from float: %v", err)
		}
		if string(s) != "123.45" {
			t.Errorf("expected \"123.45\", got %s", s)
		}
	})
}

// TestStringComplexStruct 测试String类型在复杂结构体中的使用
func TestStringComplexStruct(t *testing.T) {
	// 定义一个包含String类型的结构体
	type TestStruct struct {
		Name  String `json:"name"`
		Age   String `json:"age"`
		Score String `json:"score"`
	}

	// 测试结构体的序列化
	t.Run("Struct serialization", func(t *testing.T) {
		// 验证用户需求的关键点：确保数值类型字段序列化为字符串
		input := TestStruct{
			Name:  String("test"),
			Age:   String("25"),
			Score: String("98.5"),
		}

		data, err := json.Marshal(input)
		if err != nil {
			t.Fatalf("failed to marshal TestStruct: %v", err)
		}

		// 验证所有字段都序列化为字符串格式
		expected := `{"name":"test","age":"25","score":"98.5"}`
		if string(data) != expected {
			t.Errorf("expected %s, got %s", expected, data)
		}
	})

	// 测试结构体的反序列化
	t.Run("Struct unmarshalling with mixed types", func(t *testing.T) {
		// 使用混合类型的JSON进行反序列化测试
		data := []byte(`{"name":"test","age":25,"score":98.5}`)

		var result TestStruct
		err := json.Unmarshal(data, &result)
		if err != nil {
			t.Fatalf("failed to unmarshal TestStruct with mixed types: %v", err)
		}

		// 验证数值类型被正确转换为字符串
		if string(result.Name) != "test" {
			t.Errorf("expected name to be 'test', got '%s'", result.Name)
		}
		if string(result.Age) != "25" {
			t.Errorf("expected age to be '25', got '%s'", result.Age)
		}
		if string(result.Score) != "98.5" {
			t.Errorf("expected score to be '98.5', got '%s'", result.Score)
		}
	})

	// 直接测试用户需求中的例子："test": 1 应该序列化为字符串
	t.Run("User example test", func(t *testing.T) {
		type Example struct {
			Test String `json:"test"`
		}

		// 1. 测试从数值反序列化
		data := []byte(`{"test":1}`)
		var result Example
		err := json.Unmarshal(data, &result)
		if err != nil {
			t.Fatalf("failed to unmarshal example: %v", err)
		}

		// 2. 验证反序列化后的值
		if string(result.Test) != "1" {
			t.Errorf("expected test to be '1', got '%s'", result.Test)
		}

		// 3. 验证序列化后保持字符串格式
		serialized, err := json.Marshal(result)
		if err != nil {
			t.Fatalf("failed to marshal example: %v", err)
		}

		// 4. 最终验证：确保"test"字段被序列化为字符串格式"1"而不是数值1
		expected := `{"test":"1"}`
		if string(serialized) != expected {
			t.Errorf("expected %s, got %s", expected, serialized)
		}
	})
}

// TestStringGetValue 测试String类型的GetValue方法
func TestStringGetValue(t *testing.T) {
	s := String("test value")
	val := s.GetValue()

	if val != "test value" {
		t.Errorf("expected 'test value', got '%s'", val)
	}

	if _, ok := interface{}(s).(StringValuer[string]); !ok {
		t.Error("String should implement StringValuer[string]")
	}
}