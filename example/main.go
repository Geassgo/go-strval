/*
--------------------------------
@Create 2025/10/16 11:45
@Author lengpucheng<lpc@hll520.cn>
@Project go-strval
@Version 1.0.0 2025/10/16 11:45
@Description strval库使用示例
--------------------------------
本文件展示了如何使用strval库中的Bool、Int和Float类型，
包括JSON和YAML的序列化与反序列化示例。
*/
package main

import (
	"encoding/json"
	"fmt"

	"github.com/lengpucheng/go-strval"
	"gopkg.in/yaml.v3"
)

// JSONConfig JSON配置示例结构体
type JSONConfig struct {
	// Enabled 是否启用功能
	Enabled strval.Bool `json:"enabled"`
	// MaxConnections 最大连接数
	MaxConnections strval.Int `json:"maxConnections"`
	// TimeoutSeconds 超时时间（秒）
	TimeoutSeconds strval.Float `json:"timeoutSeconds"`
}

// YAMLConfig YAML配置示例结构体
type YAMLConfig struct {
	// Enabled 是否启用功能
	Enabled strval.Bool `yaml:"enabled"`
	// Port 服务端口
	Port strval.Int `yaml:"port"`
	// RetryDelay 重试延迟（秒）
	RetryDelay strval.Float `yaml:"retryDelay"`
}

func main() {
	fmt.Println("=== JSON 示例 ===")
	jsonExample()

	fmt.Println("\n=== YAML 示例 ===")
	yamlExample()
}

// jsonExample 演示JSON序列化和反序列化
func jsonExample() {
	// JSON字符串形式的配置（注意字段值都是字符串）
	jsonStr := `{
		"enabled": "true",
		"maxConnections": "100",
		"timeoutSeconds": "3.5"
	}`

	// 反序列化
	var config JSONConfig
	err := json.Unmarshal([]byte(jsonStr), &config)
	if err != nil {
		fmt.Printf("JSON反序列化失败: %v\n", err)
		return
	}

	// 访问解析后的值（需要转换为基本类型）
	fmt.Printf("Enabled: %v (类型: %T)\n", bool(config.Enabled), bool(config.Enabled))
	fmt.Printf("MaxConnections: %v (类型: %T)\n", int(config.MaxConnections), int(config.MaxConnections))
	fmt.Printf("TimeoutSeconds: %v (类型: %T)\n", float64(config.TimeoutSeconds), float64(config.TimeoutSeconds))

	// 序列化（输出原始类型）
	output, err := json.Marshal(config)
	if err != nil {
		fmt.Printf("JSON序列化失败: %v\n", err)
		return
	}
	fmt.Printf("序列化结果: %s\n", string(output))

	// 测试不同形式的布尔值
	differentBoolExample()
}

// yamlExample 演示YAML序列化和反序列化
func yamlExample() {
	// YAML字符串形式的配置
	yamlStr := `
enabled: "false"
port: "8080"
retryDelay: "2.5"
`

	// 反序列化
	var config YAMLConfig
	err := yaml.Unmarshal([]byte(yamlStr), &config)
	if err != nil {
		fmt.Printf("YAML反序列化失败: %v\n", err)
		return
	}

	// 访问解析后的值
	fmt.Printf("Enabled: %v\n", bool(config.Enabled))
	fmt.Printf("Port: %v\n", int(config.Port))
	fmt.Printf("RetryDelay: %v\n", float64(config.RetryDelay))

	// 序列化
	output, err := yaml.Marshal(config)
	if err != nil {
		fmt.Printf("YAML序列化失败: %v\n", err)
		return
	}
	fmt.Printf("序列化结果:\n%s\n", string(output))
}

// differentBoolExample 演示不同形式的布尔值解析
func differentBoolExample() {
	fmt.Println("\n=== 不同形式的布尔值解析 ===")

	type BoolTest struct {
		Value strval.Bool `json:"value"`
	}

	// 测试各种布尔值形式
	testValues := []string{
		`{"value": true}`,        // 直接布尔值
		`{"value": "true"}`,    // 字符串true
		`{"value": "yes"}`,     // 字符串yes
		`{"value": "1"}`,       // 字符串1
		`{"value": "false"}`,   // 字符串false
		`{"value": "no"}`,      // 字符串no
		`{"value": "0"}`,       // 字符串0
	}

	for _, test := range testValues {
		var bt BoolTest
		err := json.Unmarshal([]byte(test), &bt)
		if err != nil {
			fmt.Printf("解析 %s 失败: %v\n", test, err)
		} else {
			fmt.Printf("解析 %s 结果: %v\n", test, bool(bt.Value))
		}
	}
}