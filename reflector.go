package reflector

import (
	"reflect"
	"sync"
)

type Type struct {
	reflect.Type

	fields typedMap[string, *StructField]
}

func TypeOf(i any) *Type {
	return &Type{
		Type:   reflect.TypeOf(i),
		fields: newTypedMap[string, *StructField](),
	}
}

type StructField struct {
	reflect.StructField

	tags typedMap[string, string]
}

func (t *Type) FieldByName(name string) (*StructField, bool) {
	if v, ok := t.fields.Get(name); ok {
		return v, ok
	}

	s, ok := t.Type.FieldByName(name)
	wrapped := &StructField{
		StructField: s,
		tags:        newTypedMap[string, string](),
	}
	t.fields.Set(name, wrapped)
	return wrapped, ok
}

func (s *StructField) GetTag(name string) string {
	if v, ok := s.tags.Get(name); ok {
		return v
	}

	tag := s.Tag.Get(name)
	s.tags.Set(name, tag)
	return tag
}

func newTypedMap[K comparable, V any]() typedMap[K, V] {
	return typedMap[K, V]{
		mu: sync.RWMutex{},
		m:  make(map[K]V),
	}
}

type typedMap[K comparable, V any] struct {
	mu sync.RWMutex
	m  map[K]V
}

func (m *typedMap[K, V]) Get(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.m[key]
	return v, ok
}

func (m *typedMap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	m.m[key] = value
	m.mu.Unlock()
}
