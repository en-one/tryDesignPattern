package chain

// 🌰 假设我们现在有个校园论坛，由于社区规章制度、广告、法律法规的原因需要对用户的发言进行敏感词过滤
//    如果被判定为敏感词，那么这篇帖子将会被封禁

// SensitiveWordFilter 敏感词过滤器，判定是否是敏感词
type SensitiveWordFilter interface {
	Filter(content string) bool
}

// Chain 职责链
type Chain struct {
	filters []SensitiveWordFilter
}

// AddFilter 给职责链添加过滤器
func (c *Chain) AddFilter(filter SensitiveWordFilter) {
	c.filters = append(c.filters, filter)
}

// Filter AddFilter 职责链执行过来
func (c *Chain) Filter(content string) bool {
	for _, filter := range c.filters {
		if filter.Filter(content) {
			return true
		}
	}
	return false
}

// AdSensitiveWordFilter 广告
type AdSensitiveWordFilter struct{}

// Filter 实现过滤算法
func (f *AdSensitiveWordFilter) Filter(content string) bool {
	// TODO: 实现算法
	return false
}

// PoliticalWordFilter 政治敏感
type PoliticalWordFilter struct{}

// Filter 实现过滤算法
func (f *PoliticalWordFilter) Filter(content string) bool {
	// TODO: 实现算法
	return true
}
