package reflector_test

import (
	"reflect"
	"testing"

	"github.com/jtarchie/reflector"
)

func BenchmarkControl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := reflect.TypeOf(Person{})
		field, _ := t.FieldByName("Name")
		_ = field.Type
		_ = field.Tag.Get("json")
	}
}

func BenchmarkCacheField(b *testing.B) {
	t := reflect.TypeOf(Person{})

	for i := 0; i < b.N; i++ {
		field, _ := t.FieldByName("Name")
		_ = field.Type
		_ = field.Tag.Get("json")
	}
}

func BenchmarkCacheTypeTag(b *testing.B) {
	t := reflect.TypeOf(Person{})
	field, _ := t.FieldByName("Name")

	for i := 0; i < b.N; i++ {
		_ = field.Type
		_ = field.Tag.Get("json")
	}
}

func BenchmarkReflector(b *testing.B) {
	t := reflector.TypeOf(Person{})

	for i := 0; i < b.N; i++ {
		field, _ := t.FieldByName("Name")
		_ = field.GetTag("json")
	}
}
