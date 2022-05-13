package core

import (
	"fmt"
	"github.com/depocket/lab/base"
	"reflect"
)

type Scheme struct {
	types map[string]reflect.Type
}

func (s *Scheme) RegisterService(name string, t interface{}) {
	a := reflect.TypeOf(t)
	s.types[name] = a
}

func (s *Scheme) New(name string) (base.AppService, error) {
	t, ok := s.types[name]
	if !ok {
		return nil, fmt.Errorf("unrecognized type name: %s", name)
	}
	return reflect.New(t).Interface().(base.AppService), nil
}

func NewScheme() *Scheme {
	return &Scheme{types: map[string]reflect.Type{}}
}
