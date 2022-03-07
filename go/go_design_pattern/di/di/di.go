package di

import (
	"fmt"
	"reflect"
)

type Container struct {
	providers map[reflect.Type]provider
	results   map[reflect.Type]reflect.Value
}

type provider struct {
	value  reflect.Value
	params []reflect.Type
}

func New() *Container {
	return &Container{
		providers: map[reflect.Type]provider{},
		results:   map[reflect.Type]reflect.Value{},
	}
}

func isError(t reflect.Type) bool {
	if t.Kind() != reflect.Interface {
		return false
	}

	return t.Implements(reflect.TypeOf(reflect.TypeOf((*error)(nil)).Elem()))
}

func (c *Container) Provide(construct interface{}) error {
	v := reflect.ValueOf(construct)

	if v.Kind() != reflect.Func {
		return fmt.Errorf("constructor must be a func")
	}

	vt := v.Type()

	params := make([]reflect.Type, vt.NumIn())
	for i := 0; i < vt.NumIn(); i++ {
		params[i] = vt.In(i)
	}

	results := make([]reflect.Type, vt.NumOut())
	for i := 0; i < vt.NumOut(); i++ {
		results[i] = vt.Out(i)
	}

	provider := provider{
		value:  v,
		params: params,
	}

	for _, result := range results {
		if isError(result) {
			continue
		}

		if _, ok := c.providers[result]; ok {
			return fmt.Errorf("%s had a provider", result)
		}

		c.providers[result] = provider
	}
	return nil
}

func (c *Container) Invoke(function interface{}) error {
	v := reflect.ValueOf(function)

	if v.Kind() != reflect.Func {
		return fmt.Errorf("construct must be func")
	}

	vt := v.Type()

	var err error
	params := make([]reflect.Value, vt.NumIn())
	for i := 0; i < vt.NumIn(); i++ {
		params[i], err = c.buildParam(vt.In(i))
		if err != nil {
			return err
		}
	}

	v.Call(params)

	return nil
}

func (c *Container) buildParam(param reflect.Type) (val reflect.Value, err error) {
	if result, ok := c.results[param]; ok {
		return result, nil
	}

	provider, ok := c.providers[param]
	if !ok {
		return reflect.Value{}, fmt.Errorf("can not find provider: %s", param)
	}

	params := make([]reflect.Value, len(provider.params))
	for i, p := range provider.params {
		params[i], err = c.buildParam(p)
	}

	results := provider.value.Call(params)
	for _, result := range results {
		if isError(result.Type()) && !result.IsNil() {
			return reflect.Value{}, fmt.Errorf("%v call err: %+v", provider, result)
		}

		if !isError(result.Type()) && !result.IsNil() {
			c.results[result.Type()] = result
		}
	}
	return c.results[param], nil
}
