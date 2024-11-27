package unit

import (
	
	"Lab5/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestValidInput(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`all fields are valid`, func(t *testing.T) {
		user := entity.Member{
			Username: "ValidUser",
			Password: "ValidPassword123",
			Email:    "valid@example.com",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())      
		g.Expect(err).To(BeNil())      
	})
}

func TestUsername(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`username is required`, func(t *testing.T) {
		user := entity.Member{
			Username: "" , //ผิดตรงนี้
			Password:  "12345" ,
			Email:  "se@gmail.com",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Username is required"))
	})

}

func TestEmail(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`email is required`, func(t *testing.T) {
		user := entity.Member{
			Username: "UnitTest" ,
			Password:  "12345" ,
			Email:  "",//ผิดตรงนี้
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Email is required"))
	})

	t.Run(`email is invalid`, func(t *testing.T) {
		user := entity.Member{
			Username: "UnitTest" ,
			Password:  "12345" ,
			Email:  "se",//ผิดตรงนี้
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Email is invalid"))
	})

}

