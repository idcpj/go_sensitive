# go 实现的敏感词检测

特性
1. 敏感词可返回上下文,便于客户端显示高亮
2. 使用原始go实现,高性能


## 快速开始
``` go
sensitive := NewSensitive()
sensitive.AddWords([]string{"中国","美国"})
input := "这是中国,这是美国"

// 30  表示获取敏感词前后30字的内容
sensitive.FindAllSensitive(input, 30) 
//output:
// [
//     {
//     content:'这是中国,这是美国',
//     Indexes:2,
//     Len:2,
//     Target:"中国"
//     }
// ]
```