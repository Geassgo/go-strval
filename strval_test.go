/*
--------------------------------
@Create 2025/10/16 11:45
@Author lengpucheng<lpc@hll520.cn>
@Project go-strval
@Version 1.0.0 2025/10/16 11:45
@Description strval包的测试用例
--------------------------------
本文件包含对strval包中Bool、Int和Float类型的测试用例，
测试内容包括JSON/YAML的序列化和反序列化功能，
以及结构体中嵌套使用这些类型的功能验证。
*/

package strval

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"
)

// TestBoolSerialization 测试Bool类型的序列化和反序列化功能
// 测试内容包括：
// 1. JSON序列化
// 2. JSON反序列化（直接布尔值）
// 3. JSON反序列化（字符串形式）
// 4. JSON反序列化（其他真值形式）
// 5. JSON反序列化（错误值）
// 6. YAML序列化
// 7. YAML反序列化（直接布尔值）
// 8. YAML反序列化（字符串形式）
func TestBoolSerialization(t *testing.T) {
	// 测试JSON序列化
	b := Bool(true)
	jsonData, err := json.Marshal(b)
	if err != nil {
		t.Fatalf("JSON marshal failed: %v", err)
	}
	if string(jsonData) != "true" {
		t.Errorf("Expected 'true', got '%s'", string(jsonData))
	}

	// 测试JSON反序列化 - 直接布尔值
	var b1 Bool
	err = json.Unmarshal([]byte("true"), &b1)
	if err != nil || !bool(b1) {
		t.Errorf("Failed to unmarshal boolean 'true': %v, value: %v", err, b1)
	}

	// 测试JSON反序列化 - 字符串形式
	var b2 Bool
	err = json.Unmarshal([]byte(`"true"`), &b2)
	if err != nil || !bool(b2) {
		t.Errorf("Failed to unmarshal string 'true': %v, value: %v", err, b2)
	}

	// 测试JSON反序列化 - 其他真值
	var b3 Bool
	err = json.Unmarshal([]byte(`"yes"`), &b3)
	if err != nil || !bool(b3) {
		t.Errorf("Failed to unmarshal string 'yes': %v, value: %v", err, b3)
	}

	// 测试JSON反序列化 - 错误值
	var b4 Bool
	err = json.Unmarshal([]byte(`"invalid"`), &b4)
	if err != nil || bool(b4) {
		t.Errorf("Failed to handle invalid value: %v, value: %v", err, b4)
	}

	// 测试YAML序列化
	yamlData, err := yaml.Marshal(b)
	if err != nil {
		t.Fatalf("YAML marshal failed: %v", err)
	}
	if string(yamlData) != "true\n" {
		t.Errorf("Expected 'true\n', got '%s'", string(yamlData))
	}

	// 测试YAML反序列化 - 直接布尔值
	var b5 Bool
	err = yaml.Unmarshal([]byte("true"), &b5)
	if err != nil || !bool(b5) {
		t.Errorf("Failed to unmarshal YAML boolean 'true': %v, value: %v", err, b5)
	}

	// 测试YAML反序列化 - 字符串形式
	var b6 Bool
	err = yaml.Unmarshal([]byte(`"false"`), &b6)
	if err != nil || bool(b6) {
		t.Errorf("Failed to unmarshal YAML string 'false': %v, value: %v", err, b6)
	}
}

// TestIntSerialization 测试Int类型的序列化和反序列化功能
// 测试内容包括：
// 1. JSON序列化
// 2. JSON反序列化（直接数值）
// 3. JSON反序列化（字符串形式）
// 4. JSON反序列化（错误值）
// 5. YAML序列化
// 6. YAML反序列化（直接数值）
// 7. YAML反序列化（字符串形式）
func TestIntSerialization(t *testing.T) {
	// 测试JSON序列化
	i := Int(42)
	jsonData, err := json.Marshal(i)
	if err != nil {
		t.Fatalf("JSON marshal failed: %v", err)
	}
	if string(jsonData) != "42" {
		t.Errorf("Expected '42', got '%s'", string(jsonData))
	}

	// 测试JSON反序列化 - 直接数值
	var i1 Int
	err = json.Unmarshal([]byte("123"), &i1)
	if err != nil || int(i1) != 123 {
		t.Errorf("Failed to unmarshal int '123': %v, value: %v", err, i1)
	}

	// 测试JSON反序列化 - 字符串形式
	var i2 Int
	err = json.Unmarshal([]byte(`"456"`), &i2)
	if err != nil || int(i2) != 456 {
		t.Errorf("Failed to unmarshal string '456': %v, value: %v", err, i2)
	}

	// 测试JSON反序列化 - 错误值
	var i3 Int
	err = json.Unmarshal([]byte(`"invalid"`), &i3)
	if err != nil || int(i3) != 0 {
		t.Errorf("Failed to handle invalid value: %v, value: %v", err, i3)
	}

	// 测试YAML序列化
	yamlData, err := yaml.Marshal(i)
	if err != nil {
		t.Fatalf("YAML marshal failed: %v", err)
	}
	if string(yamlData) != "42\n" {
		t.Errorf("Expected '42\n', got '%s'", string(yamlData))
	}

	// 测试YAML反序列化 - 直接数值
	var i4 Int
	err = yaml.Unmarshal([]byte("789"), &i4)
	if err != nil || int(i4) != 789 {
		t.Errorf("Failed to unmarshal YAML int '789': %v, value: %v", err, i4)
	}

	// 测试YAML反序列化 - 字符串形式
	var i5 Int
	err = yaml.Unmarshal([]byte(`"999"`), &i5)
	if err != nil || int(i5) != 999 {
		t.Errorf("Failed to unmarshal YAML string '999': %v, value: %v", err, i5)
	}
}

// TestFloatSerialization 测试Float类型的序列化和反序列化功能
// 测试内容包括：
// 1. JSON序列化
// 2. JSON反序列化（直接数值）
// 3. JSON反序列化（字符串形式）
// 4. JSON反序列化（错误值）
// 5. YAML序列化
// 6. YAML反序列化（直接数值）
// 7. YAML反序列化（字符串形式）
func TestFloatSerialization(t *testing.T) {
	// 测试JSON序列化
	f := Float(3.14)
	jsonData, err := json.Marshal(f)
	if err != nil {
		t.Fatalf("JSON marshal failed: %v", err)
	}
	if string(jsonData) != "3.14" {
		t.Errorf("Expected '3.14', got '%s'", string(jsonData))
	}

	// 测试JSON反序列化 - 直接数值
	var f1 Float
	err = json.Unmarshal([]byte("2.718"), &f1)
	if err != nil || float64(f1) != 2.718 {
		t.Errorf("Failed to unmarshal float '2.718': %v, value: %v", err, f1)
	}

	// 测试JSON反序列化 - 字符串形式
	var f2 Float
	err = json.Unmarshal([]byte(`"1.618"`), &f2)
	if err != nil || float64(f2) != 1.618 {
		t.Errorf("Failed to unmarshal string '1.618': %v, value: %v", err, f2)
	}

	// 测试JSON反序列化 - 错误值
	var f3 Float
	err = json.Unmarshal([]byte(`"invalid"`), &f3)
	if err != nil || float64(f3) != 0 {
		t.Errorf("Failed to handle invalid value: %v, value: %v", err, f3)
	}

	// 测试YAML序列化
	yamlData, err := yaml.Marshal(f)
	if err != nil {
		t.Fatalf("YAML marshal failed: %v", err)
	}
	if string(yamlData) != "3.14\n" {
		t.Errorf("Expected '3.14\n', got '%s'", string(yamlData))
	}

	// 测试YAML反序列化 - 直接数值
	var f4 Float
	err = yaml.Unmarshal([]byte("0.5"), &f4)
	if err != nil || float64(f4) != 0.5 {
		t.Errorf("Failed to unmarshal YAML float '0.5': %v, value: %v", err, f4)
	}

	// 测试YAML反序列化 - 字符串形式
	var f5 Float
	err = yaml.Unmarshal([]byte(`"0.25"`), &f5)
	if err != nil || float64(f5) != 0.25 {
		t.Errorf("Failed to unmarshal YAML string '0.25': %v, value: %v", err, f5)
	}
}

// TestStructSerialization 测试在结构体中嵌套使用Bool、Int和Float类型的功能
// 测试内容包括：
// 1. 结构体的JSON序列化
// 2. 结构体的JSON反序列化（字符串形式）
// 3. 结构体的YAML序列化
// 4. 结构体的YAML反序列化（字符串形式）
func TestStructSerialization(t *testing.T) {
	type TestStruct struct {
		B Bool   `json:"b" yaml:"b"`
		I Int    `json:"i" yaml:"i"`
		F Float  `json:"f" yaml:"f"`
	}

	// 测试结构体JSON序列化
	s := TestStruct{
		B: true,
		I: 100,
		F: 2.5,
	}
	jsonData, err := json.Marshal(s)
	if err != nil {
		t.Fatalf("Struct JSON marshal failed: %v", err)
	}
	if string(jsonData) != `{"b":true,"i":100,"f":2.5}` {
		t.Errorf("Unexpected JSON output: %s", string(jsonData))
	}

	// 测试结构体JSON反序列化（字符串形式）
	jsonStr := `{"b":"false","i":"200","f":"3.5"}`
	var s1 TestStruct
	err = json.Unmarshal([]byte(jsonStr), &s1)
	if err != nil {
		t.Fatalf("Struct JSON unmarshal failed: %v", err)
	}
	if bool(s1.B) != false || int(s1.I) != 200 || float64(s1.F) != 3.5 {
		t.Errorf("Unexpected values after unmarshal: %+v", s1)
	}

	// 测试结构体YAML序列化
	_, err = yaml.Marshal(s)
	if err != nil {
		t.Fatalf("Struct YAML marshal failed: %v", err)
	}

	// 测试结构体YAML反序列化（字符串形式）
	yamlStr := `
 b: "true"
 i: "300"
 f: "4.5"
`
	var s2 TestStruct
	err = yaml.Unmarshal([]byte(yamlStr), &s2)
	if err != nil {
		t.Fatalf("Struct YAML unmarshal failed: %v", err)
	}
	if bool(s2.B) != true || int(s2.I) != 300 || float64(s2.F) != 4.5 {
		t.Errorf("Unexpected values after YAML unmarshal: %+v", s2)
	}
}