package sqlc

import (
	"context"
	"github.com/FeiNiaoBF/GoFlashCards/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateCards(t *testing.T) {
	// create a card
	arg := CreateCardsParams{
		Front: util.RandText(10),
		Back:  util.RandText(10),
	}

	card, err := testQueries.CreateCards(context.Background(), arg)
	// [testify package](https://github.com/stretchr/testify)
	require.NoError(t, err)
	require.NotEmpty(t, card)
	// content
	require.Equal(t, arg.Front, card.Front)
	require.Equal(t, arg.Back, card.Back)
}

func TestGetCard(t *testing.T) {
	tCard := createRandCards(t)
	card, err := testQueries.GetCard(context.Background(), tCard.ID)
	require.NoError(t, err)
	require.NotEmpty(t, card)
	// content
	require.Equal(t, tCard.Front, card.Front)
	require.Equal(t, tCard.Back, card.Back)
	require.Equal(t, tCard.ID, card.ID)
	require.Equal(t, tCard.Know, card.Know)
	require.WithinDuration(t, tCard.AddTime.Time, card.AddTime.Time, time.Second)

}
func TestGetAllCard(t *testing.T) {
	var loop = 5
	cards := make([]Card, 0)
	for i := 0; i < loop; i++ {
		cards[i] = createRandCards(t)
	}

}

func createRandCards(t *testing.T) Card {
	// create a card
	arg := CreateCardsParams{
		Front: util.RandText(10),
		Back:  util.RandText(20),
	}

	card, err := testQueries.CreateCards(context.Background(), arg)
	// [testify package](https://github.com/stretchr/testify)
	require.NoError(t, err)
	require.NotEmpty(t, card)
	// content
	require.Equal(t, arg.Front, card.Front)
	require.Equal(t, arg.Back, card.Back)

	require.NotZero(t, card.ID)
	require.NotZero(t, card.AddTime)
	require.False(t, card.Know)
	return card
}
