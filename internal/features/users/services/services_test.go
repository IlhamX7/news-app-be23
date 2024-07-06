package services_test

import (
	"news-app-be23/internal/features/users"
	"news-app-be23/internal/features/users/services"
	"news-app-be23/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestSignUp(t *testing.T) {
	qry := mocks.NewQuery(t)
	pu := mocks.NewPasswordUtilityInterface(t)
	jt := mocks.NewJwtUtilityInterface(t)
	srv := services.NewUserService(qry, pu, jt)
	input := users.User{Username: "ilham", Password: "54321", Email: "ilham77@gmail.com"}

	t.Run("Success Sign Up", func(t *testing.T) {
		inputQry := users.User{Username: "ilham", Password: "testpassword", Email: "ilham77@gmail.com"}

		pu.On("GeneratePassword", input.Password).Return([]byte("testpassword"), nil).Once()
		qry.On("SignUp", inputQry).Return(nil).Once()

		err := srv.SignUp(input)

		pu.AssertExpectations(t)
		qry.AssertExpectations(t)

		assert.Nil(t, err)
	})

	t.Run("Error Hash Password", func(t *testing.T) {
		pu.On("GeneratePassword", input.Password).Return(nil, bcrypt.ErrPasswordTooLong).Once()

		err := srv.SignUp(input)

		pu.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "input data tidak valid, data tidak bisa diproses")
	})

	t.Run("Error From Query", func(t *testing.T) {
		inputQry := users.User{Username: "ilham", Password: "examplepassword", Email: "ilham77@gmail.com"}

		pu.On("GeneratePassword", input.Password).Return([]byte("examplepassword"), nil).Once()
		qry.On("SignUp", inputQry).Return(gorm.ErrInvalidData).Once() // Change Register to SignUp

		err := srv.SignUp(input)

		pu.AssertExpectations(t)
		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "terjadi kesalahan pada server saat mengolah data")
	})
}
