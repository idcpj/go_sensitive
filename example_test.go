package trie

import "testing"

func TestUse(t *testing.T) {
	sensitive := NewSensitive()
	sensitive.AddWords([]string{"中国"})
	sensitive.AddWords([]string{"美国"})
	sensitive.Reload([]string{"中国"})

	input := "这是中国,这是美国"
	res := sensitive.FindAllSensitive(input, 30)
	assert(t, len(res) != 1, "长度不等于2")
	assert(t, res[0].Target != "中国", "没有匹配中国")
}
