package reflector_test

import (
	"testing"

	"github.com/jtarchie/reflector"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBenchmark(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reflector Suite")
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var _ = Describe("Reflector", func() {
	Describe("TypeOf", func() {
		It("should wrap the given type", func() {
			t := reflector.TypeOf(Person{})
			Expect(t).ToNot(BeNil())
			Expect(t.Name()).To(Equal("Person"))
		})
	})

	Describe("FieldByName", func() {
		var t *reflector.Type

		BeforeEach(func() {
			t = reflector.TypeOf(Person{})
		})

		It("should return the correct field", func() {
			field, ok := t.FieldByName("Name")
			Expect(ok).To(BeTrue())
			Expect(field.Name).To(Equal("Name"))

			// Second call should use cache
			field, _ = t.FieldByName("Name")
			Expect(field.Name).To(Equal("Name"))
		})
	})

	Describe("GetTag", func() {
		var t *reflector.Type
		var field *reflector.StructField

		BeforeEach(func() {
			t = reflector.TypeOf(Person{})
			field, _ = t.FieldByName("Name")
		})

		It("should return the correct tag", func() {
			tag := field.GetTag("json")
			Expect(tag).To(Equal("name"))

			// Second call should use cache
			tag = field.GetTag("json")
			Expect(tag).To(Equal("name"))
		})
	})
})