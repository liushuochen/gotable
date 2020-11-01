# gotable

## Introduction
打印CLI表格框架。

## 效果
### 普通表格
![](https://tuocheng.oss-cn-beijing.aliyuncs.com/gotable_test_plain.png)
### 彩色表格
![](https://tuocheng.oss-cn-beijing.aliyuncs.com/gotable_test_color.png)
### 结构体数组转表格
![](https://tuocheng.oss-cn-beijing.aliyuncs.com/gotable_test_strcut.png)


## Bootstrap
直接运行`table_test.go`
或者参考[使用指南](https://blog.csdn.net/TCatTime/article/details/103068260#%E8%8E%B7%E5%8F%96gotable)

### Demo
```go
func TestTable(t *testing.T) {
	tbl, err := CreateTable(testHeader)
	assert.Nil(t, err)

	for _, value := range plainData {
		err := tbl.AddValue(value)
		assert.Nil(t, err)
	}

	tbl.PrintTable()
}

func TestColorTable(t *testing.T) {
	tbl, err := CreateTable(testHeader)
	assert.Nil(t, err)

	for _, value := range colorData {
		err := tbl.AddValue(value)
		assert.Nil(t, err)
	}

	tbl.PrintTable()
}

func TestStructTable(t *testing.T) {
	controller := func(field string, val reflect.Value) color.Color {
		switch field {
		case "Name":
			if strings.Contains(val.String(), "c") {
				return color.CYAN
			}
		case "Experience":
			if val.Int() > 5 {
				return color.MAGENTA
			}
		case "Salary":
			if val.Float() < 1000 {
				return color.YELLOW
			}
		}
		return ""
	}

	tbl, err := CreateTableFromStruct(MyStruct{}, WithColorController(controller))
	assert.Nil(t, err)

	err = tbl.AddValuesFromSlice(structSliceData)
	assert.Nil(t, err)

	tbl.PrintTable()
}
```
## 版本更新
1.3 解决打印彩色字符串时，计算长度时错误，表格对不齐
- 根本问题：把不可显示的字符也计算进长度了
- 解决方式：设计接口替换string类型。将表格支持的string替换成Sequence接口，提供Val()和Len()。用Len来进行长度计算，避免不可打印字符占用宽度计算。

1.4 支持struct slice转table
- 支持struct slice转table（基于反射）
- 支持单元格自定义颜色配置