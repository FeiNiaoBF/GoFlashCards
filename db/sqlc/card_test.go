package sqlc

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateCards(t *testing.T) {
	// create a card
	arg := CreateCardsParams{
		Front: "what is MySQL",
		Back:  "This is a relational database",
	}

	card, err := testQueries.CreateCards(context.Background(), arg)
	// [testify package](https://github.com/stretchr/testify)
	require.NoError(t, err)
	require.NotEmpty(t, card)
	// content
	require.Equal(t, arg.Front, card.Front)
	require.Equal(t, arg.Back, card.Back)
}
