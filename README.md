# go-strval

Go 库提供了增强的基本类型（Bool、Int、Float），支持从字符串形式的 JSON/YAML 反序列化到相应的基本类型。

## 功能特点

- **Bool 类型**：支持从字符串形式（如 "true", "false", "yes", "no", "1", "0"）反序列化为 bool 值
- **Int 类型**：支持从字符串形式反序列化为 int 值
- **Float 类型**：支持从字符串形式反序列化为 float64 值
- **优雅处理错误**：当格式异常时，会将值设置为零值，并使用 slog 记录详细错误信息
- **标准序列化**：序列化为 JSON/YAML 时输出原始类型值，而不是字符串

## 安装

```bash
go get github.com/geassgo/go-strval
```

## 使用示例

### 基本使用

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/geassgo/go-strval"
)

type Config struct {
	Enabled  strval.Bool   `json:"enabled"`
	MaxCount strval.Int    `json:"maxCount"`
	Timeout  strval.Float  `json:"timeout"`
}

func main() {
	// JSON 反序列化（字符串形式）
	jsonData := []byte(`{
		"enabled": "true",
		"maxCount": "100",
		"timeout": "3.5"
	}`)

	var config Config
	err := json.Unmarshal(jsonData, &config)
	if err != nil {
		panic(err)
	}

	// 直接使用转换后的基本类型值
	fmt.Printf("Enabled: %v\n", bool(config.Enabled))   // true
	fmt.Printf("MaxCount: %v\n", int(config.MaxCount))   // 100
	fmt.Printf("Timeout: %v\n", float64(config.Timeout)) // 3.5

	// JSON 序列化（输出原始类型）
	output, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output)) // {"enabled":true,"maxCount":100,"timeout":3.5}
}
```

### YAML 支持

```go
package main

import (
	"fmt"

	"github.com/geassgo/go-strval"
	"gopkg.in/yaml.v3"
)

type YAMLConfig struct {
	Enabled  strval.Bool   `yaml:"enabled"`
	MaxCount strval.Int    `yaml:"maxCount"`
	Timeout  strval.Float  `yaml:"timeout"`
}

func main() {
	// YAML 反序列化
	yamlData := []byte(`
enabled: "false"
maxCount: "200"
timeout: "5.25"
`)

	var config YAMLConfig
	err := yaml.Unmarshal(yamlData, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Enabled: %v\n", bool(config.Enabled))   // false
	fmt.Printf("MaxCount: %v\n", int(config.MaxCount))   // 200
	fmt.Printf("Timeout: %v\n", float64(config.Timeout)) // 5.25
}
```

## 错误处理

当解析失败时，库会：
1. 将值设置为对应类型的零值（false, 0, 0.0）
2. 使用 slog 记录详细的错误信息，包括原始值和具体错误

例如，当解析无效的布尔值字符串时：
```
ERROR invalid Bool string value value=invalid error="cannot parse 'invalid' as bool"
```

## 测试

运行测试以验证功能：

```bash
go test -v ./...
```

## 类型转换

所有类型都可以直接转换为对应的基本类型：

```go
var b strval.Bool = true
basicBool := bool(b)

var i strval.Int = 42
basicInt := int(i)

var f strval.Float = 3.14
basicFloat := float64(f)
```