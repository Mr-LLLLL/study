package chain

type ISensitiveWordFilter interface {
	Filter(content string) bool
}

type SensitiveWordFilterChain struct {
	filters []ISensitiveWordFilter
}

func (c *SensitiveWordFilterChain) AddFilter(filter ISensitiveWordFilter) {
	c.filters = append(c.filters, filter)
}

func (c *SensitiveWordFilterChain) Filter(content string) bool {
	for _, filter := range c.filters {
		if filter.Filter(content) {
			return true
		}
	}
	return false
}

type AdsensitiveWordFilter struct{}

func (f *AdsensitiveWordFilter) Filter(content string) bool {
	return false
}

type PolitialWordFilter struct{}

func (f *PolitialWordFilter) Filter(context string) bool {
	return true
}
