package go_builder

import "fmt"

type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

type ResourcePoolConfigOption struct {
	MaxTotal int
	maxIdle  int
	minIdle  int
}

type ResourcePoolConfigOptFunc func(option *ResourcePoolConfigOption)

func NewResourcePoolConfig(name string, opts ...ResourcePoolConfigOptFunc) (*ResourcePoolConfig, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	option := &ResourcePoolConfigOption{
		MaxTotal: 10,
		maxIdle:  9,
		minIdle:  0,
	}

	for _, opt := range opts {
		opt(option)
	}

	if option.MaxTotal < 0 || option.maxIdle < 0 || option.minIdle < 0 {
		return nil, fmt.Errorf("args err, option:%v", option)
	}

	if option.MaxTotal < option.maxIdle || option.maxIdle < option.minIdle {
		return nil, fmt.Errorf("args err, option: %v", option)
	}

	return &ResourcePoolConfig{
		name:     name,
		maxTotal: option.MaxTotal,
		maxIdle:  option.maxIdle,
		minIdle:  option.minIdle,
	}, nil
}

