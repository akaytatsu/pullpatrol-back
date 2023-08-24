package usecase_user_test

import (
	"app/entity"
	"app/mocks"
	usecase_user "app/usecase/user"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestUsecaseUser_LoginUser(t *testing.T) {

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	password, _ := entity.GeneratePassword("password33")

	mockUserRepo := mocks.NewMockIRepositoryUser(ctrl)
	mockUserRepo.EXPECT().GetByMail(gomock.Any()).Return(&entity.EntityUser{
		Email:    "mailer@mailer.com",
		Password: password,
	}, nil)

	_, err := usecase_user.NewService(mockUserRepo).LoginUser("mailer@mailer.com", "password33")

	assert.Nil(t, err)
}

func TestUsecaseUser_CreateUser(t *testing.T) {

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockIRepositoryUser(ctrl)
	mockUserRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)

	Convey("User can't be created", t, func() {

		err := usecase_user.NewService(mockUserRepo).Create(&entity.EntityUser{})

		So(err, ShouldNotBeNil)
	})

	Convey("User can be created", t, func() {

		err := usecase_user.NewService(mockUserRepo).Create(&entity.EntityUser{
			Email:    "mailer@mailer.com",
			Name:     "Mailer",
			Password: "password33",
		})

		So(err, ShouldBeNil)
	})
}

func TestUsecaseUser_GetUserByToken(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockIRepositoryUser(ctrl)
	usercaseUser := usecase_user.NewService(mockUserRepo)

	Convey("Unvalid token can't be found", t, func() {
		user, err := usercaseUser.GetUserByToken("token")

		So(user, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("Valid token but user can be found", t, func() {
		user := entity.EntityUser{
			ID: 33,
		}

		token, _, _ := usecase_user.JWTTokenGenerator(user)

		mockUserRepo.EXPECT().GetByID(33).Return(nil, errors.New("error"))

		userResp, err := usercaseUser.GetUserByToken(token)

		So(userResp, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("Valid token and user can be found", t, func() {
		user := entity.EntityUser{
			ID: 33,
		}

		token, _, _ := usecase_user.JWTTokenGenerator(user)

		mockUserRepo.EXPECT().GetByID(33).Return(&user, nil)

		userResp, err := usercaseUser.GetUserByToken(token)

		So(userResp, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

}
