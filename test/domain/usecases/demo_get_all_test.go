package usecases

import (
	"context"
	"fmt"
	"testing"

	"github.com/elielbatiston/application-name/src/domain/models"
	"github.com/elielbatiston/application-name/src/domain/usecases"
	"github.com/elielbatiston/application-name/src/infra/repositories/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewDemoGetAll(t *testing.T) {
	uc := usecases.NewDemoGetAll()
	assert.NotNil(t, uc)
}

func TestDemoGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock.NewMockDemoRepository(ctrl)
	uc := usecases.DemoGetAll{Repo: repo}

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		expected := []models.Demo{
			{ID: 1, Name: "1"},
			{ID: 2, Name: "2"},
		}
		repo.EXPECT().FindAll(ctx).Return(expected, nil)

		result, err := uc.Execute(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expected, result)
	})

	t.Run("fail", func(t *testing.T) {
		ctx := context.Background()
		repo.EXPECT().FindAll(ctx).Return(nil, fmt.Errorf("could not find"))

		result, err := uc.Execute(context.Background())
		assert.Error(t, err)
		assert.Empty(t, result)
	})
}
