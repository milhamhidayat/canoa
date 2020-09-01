package repository_test

import (
	"context"
	"testing"
	"time"

	sq "github.com/huandu/go-sqlbuilder"
	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"soccer-api/db"
	"soccer-api/internal/entity"
	r "soccer-api/internal/team/repository"
)

type repositorySuite struct {
	db.Suite
}

func TestTeamRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skip team repository test")
	}
	suite.Run(t, new(repositorySuite))
}

func (s *repositorySuite) SetupTest() {
	log.Info("starting test. migrating the database")
	ok, err := s.Migration.Up()
	require.True(s.T(), ok)
	require.NoError(s.T(), err)
	log.Info("database migrated successfully")
}

func (s *repositorySuite) TearDownTest() {
	log.Info("finishing test. migrate down the database")
	ok, err := s.Migration.Down()
	require.True(s.T(), ok)
	require.NoError(s.T(), err)
}

func (s *repositorySuite) seedTeams(teams ...entity.Team) {
	ib := sq.NewInsertBuilder()
	ib.InsertInto("teams").Cols("id", "name", "founded", "stadium", "created_at")
	for _, v := range teams {
		ib.Values(v.ID, v.Name, v.Founded, v.Stadium, v.CreatedAt)
	}
	sql, args := ib.Build()

	_, err := s.DBConn.Exec(sql, args...)
	require.NoError(s.T(), err)
}

func (s *repositorySuite) TestTeamRepo_Get() {
	mockTeams := []entity.Team{
		{
			ID:        1,
			Name:      "London FC",
			Founded:   1970,
			Stadium:   "London Stadium",
			CreatedAt: time.Date(2020, time.August, 10, 9, 20, 0, 0, time.UTC),
		},
	}
	s.seedTeams(mockTeams[0])

	mockTeam := mockTeams[0]
	mockTeam.CreatedAt = time.Time{}

	tests := map[string]struct {
		id             int64
		expectedResult entity.Team
		expectedError  error
	}{
		"success": {
			id:             mockTeam.ID,
			expectedResult: mockTeam,
			expectedError:  nil,
		},
		"not found": {
			id:             5,
			expectedResult: entity.Team{},
			expectedError:  entity.ErrNotFound,
		},
	}

	for testName, tc := range tests {
		s.T().Run(testName, func(t *testing.T) {
			repo := r.New(s.DBConn)
			res, err := repo.Get(context.Background(), tc.id)
			if err != nil {
				require.EqualError(t, err, tc.expectedError.Error())
				return
			}

			require.Equal(s.T(), tc.expectedResult, res)
			require.NoError(s.T(), err)
		})
	}
}

func (s *repositorySuite) TestTeamRepo_Insert() {
	mockTeam := entity.Team{
		Name:      "London FC",
		Founded:   1970,
		Stadium:   "London Stadium",
		CreatedAt: time.Time{},
	}

	tests := map[string]struct {
		input          entity.Team
		expectedResult entity.Team
		expectedError  error
	}{
		"success": {
			input:          mockTeam,
			expectedResult: mockTeam,
			expectedError:  nil,
		},
	}

	for testName, tc := range tests {
		s.T().Run(testName, func(t *testing.T) {
			repo := r.New(s.DBConn)
			res, err := repo.Insert(context.Background(), tc.input)
			if err != nil {
				require.EqualError(t, err, tc.expectedError.Error())
				return
			}

			res.ID = 0
			require.Equal(s.T(), tc.expectedResult, res)
			require.NoError(s.T(), err)
		})
	}
}

func (s *repositorySuite) TestTeamRepo_Fetch() {
	createdAt1, err := time.Parse(time.RFC3339, "2019-03-15T09:30:02+07:00")
	require.NoError(s.T(), err)

	loc, _ := time.LoadLocation("Asia/Jakarta")

	mockTeams := []entity.Team{
		{
			ID:        1,
			Name:      "London FC",
			Founded:   1970,
			Stadium:   "London Stadium",
			CreatedAt: createdAt1.In(loc),
		},
		{
			ID:        2,
			Name:      "The Rhinos",
			Founded:   1990,
			Stadium:   "Aurora Park",
			CreatedAt: createdAt1.In(loc).AddDate(0, 0, 1),
		},
	}
	s.seedTeams(mockTeams...)

	result := []entity.Team{mockTeams[1], mockTeams[0]}

	tests := map[string]struct {
		filter         entity.QueryFilter
		expectedResult []entity.Team
		expectedCursor string
		expectedError  error
	}{
		"success with num": {
			filter:         entity.QueryFilter{Num: 20},
			expectedResult: result,
			expectedCursor: "MjAxOS0wMy0xNVQwOTozMDowMiswNzowMA==",
			expectedError:  nil,
		},
		"success with num and cursor": {
			filter:         entity.QueryFilter{Num: 20, Cursor: "MjAxOS0wMy0xNVQwOTozMDowMiswNzowMA=="},
			expectedResult: nil,
			expectedCursor: "MjAxOS0wMy0xNVQwOTozMDowMiswNzowMA==",
			expectedError:  nil,
		},
	}

	for testName, tc := range tests {
		s.T().Run(testName, func(t *testing.T) {
			repo := r.New(s.DBConn)
			res, csr, err := repo.Fetch(context.Background(), tc.filter)
			if err != nil {
				require.EqualError(t, errors.Cause(err), tc.expectedError.Error())
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.expectedResult, res)
			require.Equal(t, tc.expectedCursor, csr)
		})
	}
}
