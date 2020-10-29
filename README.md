# gotable
## Introduction
打印CLI表格框架。

## Bootstrap
直接运行`table_test.go`
或者参考[使用指南](https://blog.csdn.net/TCatTime/article/details/103068260#%E8%8E%B7%E5%8F%96gotable)
## Test
### 原始数据
```json
[
    {
        "Name":"Alice",
        "Experience":"Three year.",
        "Salary":"2300.00"
    },
    {
        "Name":"Bob",
        "Experience":"Ten year.",
        "Salary":"900.00"
    },
    {
        "Name":"Coco",
        "Experience":"One year.",
        "Salary":"9000.00"
    }
]
```

### Command

### 打印表格
#### 普通表格
![](https://tuocheng.oss-cn-beijing.aliyuncs.com/test_result.jpg)

#### 彩色表格
![](https://tuocheng.oss-cn-beijing.aliyuncs.com/test_color_result.jpg)



## 版本更新
1.3 解决打印彩色字符串时，计算长度时错误，表格对不齐
- 根本问题：把不可显示的字符也计算进长度了
- 解决方式：设计接口替换string类型。将表格支持的string替换成Sequence接口，提供Val()和Len()。用Len来进行长度计算，避免不可打印字符占用宽度计算。

1.4 支持struct slice转table
- 支持struct slice转table（基于反射）
- 规划中