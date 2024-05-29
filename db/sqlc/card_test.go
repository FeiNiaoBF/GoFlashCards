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
		Front:  util.RandText(10),
		Back:   util.RandText(10),
		TagsID: 1,
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

// TODO: How do a good test here
func TestGetAllCard(t *testing.T) {
	cardsTest, err := testQueries.GetAllCard(context.Background())
	require.NoError(t, err)
	require.Len(t, cardsTest, len(cardsTest))
}

func TestUpdateCards(t *testing.T) {
	card := createRandCards(t)

	arg := UpdateCardsParams{
		ID:     card.ID,
		Front:  util.RandText(10),
		Back:   util.RandText(10),
		TagsID: 1,
	}

	card2, err := testQueries.UpdateCards(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, card2)
	// update
	require.Equal(t, card2.Front, arg.Front)
	require.Equal(t, card2.Back, arg.Back)
	require.Equal(t, card2.TagsID, arg.TagsID)

	require.Equal(t, card2.ID, card.ID)
	require.Equal(t, card2.Know, card.Know)
}

func TestDeleteCards(t *testing.T) {
	ctx := context.Background()
	card1 := createRandCards(t)

	err := testQueries.DeleteCards(ctx, card1.ID)
	require.NoError(t, err)

	card2, err := testQueries.GetCard(ctx, card1.ID)
	require.EqualError(t, err, "no rows in result set")

	require.Empty(t, card2)
}

func TestListCards(t *testing.T) {
	loog := 10
	for i := 0; i < loog; i++ {
		_ = createRandCards(t)
	}

	arg := ListCardsParams{
		Limit:  10,
		Offset: 10,
	}
	cards, err := testQueries.ListCards(context.Background(), arg)
	require.NoError(t, err)
	// log.Print(len(cards))
	require.Len(t, cards, 10)

	for _, card := range cards {
		require.NotEmpty(t, card)
	}
}

func createRandCards(t *testing.T) Card {
	// create a card
	arg := CreateCardsParams{
		Front:  util.RandText(10),
		Back:   util.RandText(20),
		TagsID: 1,
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
