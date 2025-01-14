package unit

import(
	"por/test/entity"
	"testing"
	"fmt"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestName(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`testNO`, func(t *testing.T) {
		test := entity.Tests{
			Name: "",
		}

		ok, err := govalidator.ValidateStruct(test)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("testNO"))
		

	})

	// t.Run(`Name is invalid`, func(t *testing.T) {
	// 	test := entity.Tests {
	// 		Name: "B123456",
	// 	}

	// 	ok,err := govalidator.ValidateStruct(test)

	// 	g.Expect(ok).NotTo(BeTrue())
	// 	g.Expect(err).NotTo(BeNil())

	// 	g.Expect(err.Error()).To(Equal(fmt.Sprintf("Name: %s does not validate as matches(^[B]\\d{7}$)",test.Name)))
	// })

	t.Run(`Email is invalid`, func(t *testing.T) {
		test := entity.Tests {
			Name: "a",
		}

		ok, err := govalidator.ValidateStruct(test)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("Name: %s does not validate as email",test.Name)))
		
	})
}

func TestValid(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`Test valid`, func(t *testing.T) {
		test := entity.Tests {
			Name: "A@gmail.com",
		}

		ok, err := govalidator.ValidateStruct(test)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}