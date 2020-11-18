package service_test

import (
	"context"
	"testing"
	"time"

	"soccer-api/internal/entity"
	mock "soccer-api/internal/team/mocks"
	"soccer-api/internal/team/service"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGet_Team(t *testing.T) {
	mockTeam := entity.Team{
		ID:        int64(1),
		Name:      "London FC",
		Founded:   1970,
		Stadium:   "London Stadium",
		CreatedAt: time.Date(2020, time.August, 10, 9, 20, 0, 0, time.UTC),
	}

	Convey("Given a team ID", t, func() {
		id := 1

		Convey("When user get by id", func() {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mock.NewMockTeam(ctrl)
			mockRepo.EXPECT().Get(gomock.Any(), int64(id)).Return(mockTeam, nil)

			teamService := service.New(mockRepo)
			_, err := teamService.Get(context.TODO(), int64(id))

			Convey("Then should return empty team", func() {
				// if !reflect.DeepEqual(resp, mockTeam) {
				// 	t.Error("resp is not equal with mockteam")
				// }
				if err != nil {
					t.Error("there is an error")
				}
			})
		})
	})
}
