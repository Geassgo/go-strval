/*
--------------------------------
@Create 2025/10/16 11:45
@Author lengpucheng<lpc@hll520.cn>
@Project go-strval
@Version 1.0.0 2025/10/16 11:45
@Description 数据库交互功能测试
--------------------------------
本文件包含对Bool、Int、Float类型数据库交互功能的测试，验证其driver.Valuer和sql.Scanner接口实现的正确性。
测试模拟了数据库读写操作，确保自定义类型可以与GORM等数据库框架正常配合使用。
*/

package strval

import (
	"database/sql/driver"
	"testing"
)

// TestBoolDatabase 测试Bool类型的数据库接口实现
func TestBoolDatabase(t *testing.T) {
	// 测试Value方法
	b1 := Bool(true)
	val, err := b1.Value()
	if err != nil {
		t.Errorf("Bool.Value() returned error: %v", err)
	}
	if boolVal, ok := val.(bool); !ok || boolVal != true {
		t.Errorf("Bool.Value() returned %v, want true", val)
	}

	// 测试Scan方法 - 从bool值
	var b2 Bool
	err = b2.Scan(true)
	if err != nil {
		t.Errorf("Bool.Scan(bool) returned error: %v", err)
	}
	if b2 != true {
		t.Errorf("Bool.Scan(true) resulted in %v, want true", b2)
	}

	// 测试Scan方法 - 从int值
	err = b2.Scan(int64(1))
	if err != nil {
		t.Errorf("Bool.Scan(int) returned error: %v", err)
	}
	if b2 != true {
		t.Errorf("Bool.Scan(1) resulted in %v, want true", b2)
	}

	// 测试Scan方法 - 从string值
	err = b2.Scan("true")
	if err != nil {
		t.Errorf("Bool.Scan(string) returned error: %v", err)
	}
	if b2 != true {
		t.Errorf("Bool.Scan(\"true\") resulted in %v, want true", b2)
	}

	// 测试Scan方法 - 空值
	err = b2.Scan(nil)
	if err != nil {
		t.Errorf("Bool.Scan(nil) returned error: %v", err)
	}
	if b2 != false {
		t.Errorf("Bool.Scan(nil) resulted in %v, want false", b2)
	}

	// 测试Scan方法 - 无效字符串
	err = b2.Scan("invalid")
	if err != nil {
		t.Errorf("Bool.Scan(invalid string) returned error: %v", err)
	}
	if b2 != false {
		t.Errorf("Bool.Scan(\"invalid\") resulted in %v, want false", b2)
	}
}

// TestIntDatabase 测试Int类型的数据库接口实现
func TestIntDatabase(t *testing.T) {
	// 测试Value方法
	i1 := Int(42)
	val, err := i1.Value()
	if err != nil {
		t.Errorf("Int.Value() returned error: %v", err)
	}
	if intVal, ok := val.(int); !ok || intVal != 42 {
		t.Errorf("Int.Value() returned %v, want 42", val)
	}

	// 测试Scan方法 - 从int64值
	var i2 Int
	err = i2.Scan(int64(42))
	if err != nil {
		t.Errorf("Int.Scan(int64) returned error: %v", err)
	}
	if i2 != 42 {
		t.Errorf("Int.Scan(42) resulted in %v, want 42", i2)
	}

	// 测试Scan方法 - 从float值
	err = i2.Scan(float64(42.5))
	if err != nil {
		t.Errorf("Int.Scan(float) returned error: %v", err)
	}
	if i2 != 42 {
		t.Errorf("Int.Scan(42.5) resulted in %v, want 42", i2)
	}

	// 测试Scan方法 - 从string值
	err = i2.Scan("42")
	if err != nil {
		t.Errorf("Int.Scan(string) returned error: %v", err)
	}
	if i2 != 42 {
		t.Errorf("Int.Scan(\"42\") resulted in %v, want 42", i2)
	}

	// 测试Scan方法 - 空值
	err = i2.Scan(nil)
	if err != nil {
		t.Errorf("Int.Scan(nil) returned error: %v", err)
	}
	if i2 != 0 {
		t.Errorf("Int.Scan(nil) resulted in %v, want 0", i2)
	}

	// 测试Scan方法 - 无效字符串
	err = i2.Scan("invalid")
	if err != nil {
		t.Errorf("Int.Scan(invalid string) returned error: %v", err)
	}
	if i2 != 0 {
		t.Errorf("Int.Scan(\"invalid\") resulted in %v, want 0", i2)
	}
}

// TestFloatDatabase 测试Float类型的数据库接口实现
func TestFloatDatabase(t *testing.T) {
	// 测试Value方法
	f1 := Float(3.14)
	val, err := f1.Value()
	if err != nil {
		t.Errorf("Float.Value() returned error: %v", err)
	}
	if floatVal, ok := val.(float64); !ok || floatVal != 3.14 {
		t.Errorf("Float.Value() returned %v, want 3.14", val)
	}

	// 测试Scan方法 - 从float64值
	var f2 Float
	err = f2.Scan(float64(3.14))
	if err != nil {
		t.Errorf("Float.Scan(float64) returned error: %v", err)
	}
	if f2 != 3.14 {
		t.Errorf("Float.Scan(3.14) resulted in %v, want 3.14", f2)
	}

	// 测试Scan方法 - 从int值
	err = f2.Scan(int64(42))
	if err != nil {
		t.Errorf("Float.Scan(int) returned error: %v", err)
	}
	if f2 != 42 {
		t.Errorf("Float.Scan(42) resulted in %v, want 42", f2)
	}

	// 测试Scan方法 - 从string值
	err = f2.Scan("3.14")
	if err != nil {
		t.Errorf("Float.Scan(string) returned error: %v", err)
	}
	if f2 != 3.14 {
		t.Errorf("Float.Scan(\"3.14\") resulted in %v, want 3.14", f2)
	}

	// 测试Scan方法 - 空值
	err = f2.Scan(nil)
	if err != nil {
		t.Errorf("Float.Scan(nil) returned error: %v", err)
	}
	if f2 != 0 {
		t.Errorf("Float.Scan(nil) resulted in %v, want 0", f2)
	}

	// 测试Scan方法 - 无效字符串
	err = f2.Scan("invalid")
	if err != nil {
		t.Errorf("Float.Scan(invalid string) returned error: %v", err)
	}
	if f2 != 0 {
		t.Errorf("Float.Scan(\"invalid\") resulted in %v, want 0", f2)
	}
}

// TestDriverValuerInterface 验证类型是否正确实现了driver.Valuer接口
func TestDriverValuerInterface(t *testing.T) {
	// 验证Bool实现了driver.Valuer接口
	var valuer driver.Valuer = Bool(true)
	if valuer == nil {
		t.Errorf("Bool does not implement driver.Valuer interface")
	}

	// 验证Int实现了driver.Valuer接口
	valuer = Int(42)
	if valuer == nil {
		t.Errorf("Int does not implement driver.Valuer interface")
	}

	// 验证Float实现了driver.Valuer接口
	valuer = Float(3.14)
	if valuer == nil {
		t.Errorf("Float does not implement driver.Valuer interface")
	}
}

// TestSQLScannerInterface 验证类型是否正确实现了sql.Scanner接口
func TestSQLScannerInterface(t *testing.T) {
	// 验证Bool实现了sql.Scanner接口
	var boolVal Bool
	var scanner interface{ Scan(interface{}) error } = &boolVal
	if scanner == nil {
		t.Errorf("*Bool does not implement sql.Scanner interface")
	}

	// 验证Int实现了sql.Scanner接口
	var intVal Int
	scanner = &intVal
	if scanner == nil {
		t.Errorf("*Int does not implement sql.Scanner interface")
	}

	// 验证Float实现了sql.Scanner接口
	var floatVal Float
	scanner = &floatVal
	if scanner == nil {
		t.Errorf("*Float does not implement sql.Scanner interface")
	}
}