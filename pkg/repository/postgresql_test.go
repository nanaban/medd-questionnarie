package repository

import (
	"context"
	"os"
	"testing"

	"questionnarie/pkg/model"

	"github.com/stretchr/testify/require"
)

var (
	testRepo *Repository
)

func setup(t *testing.T) {
	//conn := "postgresql://postgres:postgrespass@localhost:5433/test4?sslmode=disable"
	conn := os.Getenv("TEST_DB_CONN")

	r, err := New(conn)
	require.NoError(t, err)

	err = r.createSchema()
	require.NoError(t, err)

	testRepo = r
}

func teardown(t *testing.T) {
	err := testRepo.Close()
	require.NoError(t, err)
}

func TestRepository_Question(t *testing.T) {
	setup(t)
	defer teardown(t)

	t.Run("CreateSuccess", func(t *testing.T) {
		tests := []*model.Question{
			{
				Status:       1,
				QuestionType: 2,
			},
		}

		for _, tt := range tests {
			ctx := context.Background()
			err := testRepo.CreateQuestion(ctx, tt)
			require.NoError(t, err)
		}
	})
}
