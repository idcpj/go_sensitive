package trie

type SensitiveMap struct {
	sensitiveNode map[string]*SensitiveMap
	isEnd         bool
}

func NewSensitive() *SensitiveMap {

	return &SensitiveMap{
		sensitiveNode: make(map[string]*SensitiveMap),
		isEnd:         false,
	}
}

func (s *SensitiveMap) AddWords(dictionary []string) {
	for _, words := range dictionary {
		sMapTmp := s
		w := []rune(words)
		wordsLength := len(w)
		for i := 0; i < wordsLength; i++ {
			t := string(w[i])
			isEnd := false
			//如果是敏感词的最后一个字，则确定状态
			if i == (wordsLength - 1) {
				isEnd = true
			}
			func(tx string) {
				if _, ok := sMapTmp.sensitiveNode[tx]; !ok { //如果该字在该层级索引中找不到，则创建新的层级
					sMapTemp := &SensitiveMap{}
					sMapTemp.sensitiveNode = make(map[string]*SensitiveMap)
					sMapTemp.isEnd = isEnd
					sMapTmp.sensitiveNode[tx] = sMapTemp
				}
				sMapTmp = sMapTmp.sensitiveNode[tx] //进入下一层级
				sMapTmp.isEnd = isEnd
			}(t)
		}
	}
}
func (s *SensitiveMap) Reload(dictionary []string) {
	s.sensitiveNode = make(map[string]*SensitiveMap)
	s.isEnd = false
	s.AddWords(dictionary)
}

/*
作用：检查是否含有敏感词，仅返回检查到的第一个敏感词
返回值：敏感词，是否含有敏感词
*/
func (s *SensitiveMap) CheckSensitive(text string) (string, bool) {
	content := []rune(text)
	contentLength := len(content)
	result := false
	ta := ""
	for index := range content {
		sMapTmp := s
		target := ""
		in := index
		for {
			wo := string(content[in])
			target += wo
			if _, ok := sMapTmp.sensitiveNode[wo]; ok {
				if sMapTmp.sensitiveNode[wo].isEnd {
					result = true
					break
				}
				if in == contentLength-1 {
					break
				}
				sMapTmp = sMapTmp.sensitiveNode[wo] //进入下一层级
				in++
			} else {
				break
			}
		}
		if result {
			ta = target
			break
		}
	}
	return ta, result
}

/*
作用：返回文本中的所有敏感词
返回值：数组，格式为“["敏感词"][敏感词在检测文本中的索引位置，敏感词长度]”
*/
type Target struct {
	Target  string
	Indexes int
	Len     int
	Content string
}

// Target 原文
// contextLength 上下文的长度,默认0,不返回上下文
func (s *SensitiveMap) FindAllSensitive(text string, correlation int) []*Target {
	content := []rune(text)
	contentLength := len(content)
	result := false

	ta := make([]*Target, 0)

	for index := range content {
		sMapTmp := s
		target := ""
		in := index
		result = false
		for {
			wo := string(content[in])
			target += wo
			if _, ok := sMapTmp.sensitiveNode[wo]; ok {
				if sMapTmp.sensitiveNode[wo].isEnd {
					result = true
					break
				}
				if in == contentLength-1 {
					break
				}
				sMapTmp = sMapTmp.sensitiveNode[wo] //进入下一层级
				in++
			} else {
				break
			}
		}
		if result {

			elems := &Target{
				Target:  target,
				Indexes: index,
				Len:     len([]rune(target)),
				Content: "",
			}

			if correlation > 0 {
				before := index - correlation
				if before < 0 {
					before = 0
				}
				after := index + elems.Len + +correlation
				if after > contentLength {
					after = contentLength
				}
				elems.Content = string(content[before:after])
			}

			ta = append(ta, elems)

		}
	}
	return ta
}
