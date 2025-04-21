package db

import (
	"context"
	"testing"

	"github.com/foyez/go/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTodo(t *testing.T) {
	text := util.RandomText()
	todo, err := testQueries.CreateTodo(context.Background(), text)
	require.NoError(t, err)
	require.NotEmpty(t, todo)

	require.Equal(t, text, todo.Text)

	require.NotZero(t, todo.ID)
	require.NotZero(t, todo.CreatedAt)
}
