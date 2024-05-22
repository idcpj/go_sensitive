package trie

import (
	"testing"
)

func assert(t *testing.T, b bool, err string) {
	if b {
		t.Error(err)
	}
}
func TestDemo1(t *testing.T) {
	sensitive := NewSensitive()
	sensitive.AddWords([]string{"中国"})
	sensitive.AddWords([]string{"美国"})
	input := "这是中国,这是美国"
	res := sensitive.FindAllSensitive(input, 0)
	assert(t, len(res) != 2, "长度不等于2")
	assert(t, res[0].Target != "中国", "没有匹配中国")
	assert(t, res[1].Target != "美国", "没有匹配美国")
}
func TestDemo2(t *testing.T) {
	sensitive := NewSensitive()
	sensitive.AddWords([]string{"中国"})
	sensitive.AddWords([]string{"美国"})
	input := "这是中国,这是美"
	res := sensitive.FindAllSensitive(input, 0)
	assert(t, len(res) != 1, "长度不等于1")
	assert(t, res[0].Target != "中国", "没有匹配中国")
}
func TestDemo3(t *testing.T) {
	sensitive := NewSensitive()
	sensitive.AddWords([]string{"中国"})
	sensitive.AddWords([]string{"美国"})
	input := "这是中国,这是中国"
	res := sensitive.FindAllSensitive(input, 0)
	assert(t, len(res) != 2, "长度不等于2")
	assert(t, res[0].Target != "中国", "没有匹配中国")
	assert(t, res[1].Target != "中国", "没有匹配中国")
}
func TestDemoWithLength(t *testing.T) {
	sensitive := NewSensitive()
	sensitive.AddWords([]string{"中国"})
	sensitive.AddWords([]string{"美国"})
	input := "这是中国,这是美国"
	res := sensitive.FindAllSensitive(input, 30)
	assert(t, len(res) != 2, "长度不等于2")
	assert(t, res[0].Target != "中国", "没有匹配中国")
	assert(t, res[0].Content != "这是中国,这是美国", "没有匹配上下文")
	assert(t, res[1].Target != "美国", "没有匹配美国")
	assert(t, res[1].Content != "这是中国,这是美国", "没有匹配上下文")
}

func TestDemo10(t *testing.T) {

	sensitive := NewSensitive()
	sensitive.AddWords([]string{"中国", "美国"})

	input := "这是中国\nim.conf\n奥萨蒂啊十大的啊的哈三菱电机啊啊时间的的代价阿斯顿阿达阿斯蒂芬奥萨蒂啊十大的啊的哈三菱电机啊啊时间的的代价阿斯顿阿达阿斯蒂芬\n奥萨蒂啊十大的啊的哈三菱电机啊啊时间的的代价阿斯顿阿达阿斯蒂芬\n奥萨蒂啊十大的啊的哈三菱电机啊这是美国啊时间的的代价阿斯顿阿达阿斯蒂芬\n奥萨蒂啊十大的啊的哈三大日本帝国菱电机啊啊时间的的代价阿斯顿阿达阿斯蒂芬\n奥萨蒂啊十大的啊的哈三菱电机啊啊时间的的代价阿斯顿阿达阿斯蒂芬\n_1777708677/im.conf\n"

	res := sensitive.FindAllSensitive(input, 30)

	assert(t, len(res) != 2, "长度不等于1")
	assert(t, res[0].Target != "中国", "没有匹配中国")
	assert(t, res[0].Content != "这是中国\nim.conf\n奥萨蒂啊十大的啊的哈三菱电机啊啊时间的的代", "没有匹配上下文")
	assert(t, res[1].Target != "美国", "没有匹配美国")
	assert(t, res[1].Content != "的代价阿斯顿阿达阿斯蒂芬\n奥萨蒂啊十大的啊的哈三菱电机啊这是美国啊时间的的代价阿斯顿阿达阿斯蒂芬\n奥萨蒂啊十大的啊的哈三大日", "没有匹配上下文")

}

func TestReload(t *testing.T) {
	sensitive := NewSensitive()
	sensitive.AddWords([]string{"中国"})
	sensitive.AddWords([]string{"美国"})
	sensitive.Reload([]string{"中国"})

	input := "这是中国,这是美国"
	res := sensitive.FindAllSensitive(input, 0)
	assert(t, len(res) != 1, "长度不等于2")
	assert(t, res[0].Target != "中国", "没有匹配中国")
}

func BenchmarkDemo1(b *testing.B) {
	sensitive := NewSensitive()
	sensitive.AddWords([]string{"中国", "美国"})

	input := "这是中国\nim.conf\n奥萨蒂啊十大的啊的哈三菱电机啊啊时间的的代价阿斯顿阿达阿斯蒂芬奥萨蒂啊十大的啊的哈三菱电机啊啊时间的的代价阿斯顿阿达阿斯蒂芬\n奥萨蒂啊十大的啊的哈三菱电机啊啊时间的的代价阿斯顿阿达阿斯蒂芬\n奥萨蒂啊十大的啊的哈三菱电机啊这是美国啊时间的的代价阿斯顿阿达阿斯蒂芬\n奥萨蒂啊十大的啊的哈三大日本帝国菱电机啊啊时间的的代价阿斯顿阿达阿斯蒂芬\n奥萨蒂啊十大的啊的哈三菱电机啊啊时间的的代价阿斯顿阿达阿斯蒂芬\n_1777708677/im.conf\n"

	for i := 0; i < b.N; i++ {
		_ = sensitive.FindAllSensitive(input, 30)

	}
}
