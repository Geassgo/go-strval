/*
--------------------------------
@Create 2025/10/16 12:30
@Author lengpucheng<lpc@hll520.cn>
@Project go-strval
@Version 1.0.0 2025/10/16 12:30
@Description 泛型接口StringValuer测试
--------------------------------
本文件包含对StringValuer泛型接口实现的测试，验证Bool、Int、Float类型是否正确实现了该接口，
以及GetValue方法能否正确返回相应类型的原始值。
*/

package strval

import (
	"testing"
)

// TestBoolGetValueGeneric 测试Bool类型的GetValue方法与泛型接口
func TestBoolGetValueGeneric(t *testing.T) {
	// 测试true值
	b1 := Bool(true)
	val := b1.GetValue()
	if val != true {
		t.Errorf("Bool(true).GetValue() returned %v, want true", val)
	}

	// 测试false值
	b2 := Bool(false)
	val = b2.GetValue()
	if val != false {
		t.Errorf("Bool(false).GetValue() returned %v, want false", val)
	}

	// 验证返回类型
	if _, ok := interface{}(val).(bool); !ok {
		t.Errorf("Bool.GetValue() did not return bool type")
	}
}

// TestIntGetValueGeneric 测试Int类型的GetValue方法与泛型接口
func TestIntGetValueGeneric(t *testing.T) {
	// 测试正整数值
	i1 := Int(42)
	val := i1.GetValue()
	if val != 42 {
		t.Errorf("Int(42).GetValue() returned %v, want 42", val)
	}

	// 测试负整数值
	i2 := Int(-10)
	val = i2.GetValue()
	if val != -10 {
		t.Errorf("Int(-10).GetValue() returned %v, want -10", val)
	}

	// 测试零值
	i3 := Int(0)
	val = i3.GetValue()
	if val != 0 {
		t.Errorf("Int(0).GetValue() returned %v, want 0", val)
	}

	// 验证返回类型
	if _, ok := interface{}(val).(int); !ok {
		t.Errorf("Int.GetValue() did not return int type")
	}
}

// TestFloatGetValueGeneric 测试Float类型的GetValue方法与泛型接口
func TestFloatGetValueGeneric(t *testing.T) {
	// 测试正浮点数值
	f1 := Float(3.14)
	val := f1.GetValue()
	if val != 3.14 {
		t.Errorf("Float(3.14).GetValue() returned %v, want 3.14", val)
	}

	// 测试负浮点数值
	f2 := Float(-2.718)
	val = f2.GetValue()
	if val != -2.718 {
		t.Errorf("Float(-2.718).GetValue() returned %v, want -2.718", val)
	}

	// 测试零值
	f3 := Float(0)
	val = f3.GetValue()
	if val != 0 {
		t.Errorf("Float(0).GetValue() returned %v, want 0", val)
	}

	// 验证返回类型
	if _, ok := interface{}(val).(float64); !ok {
		t.Errorf("Float.GetValue() did not return float64 type")
	}
}

// TestStringValuerInterface 测试各类型是否正确实现了泛型接口
func TestStringValuerInterface(t *testing.T) {
	// 创建类型变量用于编译期类型检查
	var _ StringValuer[bool] = (*Bool)(nil)
	var _ StringValuer[int] = (*Int)(nil)
	var _ StringValuer[float64] = (*Float)(nil)

	// 通过接口使用Bool类型
	var boolVal StringValuer[bool] = Bool(true)
	if boolVal.GetValue() != true {
		t.Errorf("StringValuer[bool].GetValue() returned %v, want true", boolVal.GetValue())
	}

	// 通过接口使用Int类型
	var intVal StringValuer[int] = Int(42)
	if intVal.GetValue() != 42 {
		t.Errorf("StringValuer[int].GetValue() returned %v, want 42", intVal.GetValue())
	}

	// 通过接口使用Float类型
	var floatVal StringValuer[float64] = Float(3.14)
	if floatVal.GetValue() != 3.14 {
		t.Errorf("StringValuer[float64].GetValue() returned %v, want 3.14", floatVal.GetValue())
	}
}

// printValue 是一个泛型函数，可以处理任何StringValuer实现
func printValue[T any](v StringValuer[T]) T {
	return v.GetValue()
}

// TestStringValuerGenericUsage 测试泛型接口的实际使用场景
func TestStringValuerGenericUsage(t *testing.T) {

	// 测试Bool类型
	b := Bool(true)
	boolVal := printValue(b)
	if boolVal != true {
		t.Errorf("printValue(b) returned %v, want true", boolVal)
	}

	// 测试Int类型
	i := Int(42)
	intVal := printValue(i)
	if intVal != 42 {
		t.Errorf("printValue(i) returned %v, want 42", intVal)
	}

	// 测试Float类型
	f := Float(3.14)
	floatVal := printValue(f)
	if floatVal != 3.14 {
		t.Errorf("printValue(f) returned %v, want 3.14", floatVal)
	}

	// 测试在表达式中直接使用GetValue返回值
	result := i.GetValue() + 10
	if result != 52 {
		t.Errorf("Expected 52, got %v", result)
	}

	resultFloat := f.GetValue() * 2
	if resultFloat != 6.28 {
		t.Errorf("Expected 6.28, got %v", resultFloat)
	}

	// 测试在条件语句中直接使用GetValue返回值
	if b.GetValue() {
		// 应该执行这里
	} else {
		t.Errorf("Condition should have been true")
	}
}