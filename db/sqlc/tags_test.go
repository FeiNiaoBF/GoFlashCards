package sqlc

import (
	"context"
	"testing"

	"github.com/FeiNiaoBF/GoFlashCards/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTag(t *testing.T) {
	// create a tag
	name := util.RandText(10)

	tag, err := testQueries.CreateTags(context.Background(), name)
	// [testify package](https://github.com/stretchr/testify)
	require.NoError(t, err)
	require.NotEmpty(t, tag)
	// content
	require.Equal(t, tag.Name, name)
}

func TestGetTags(t *testing.T) {
	tag1 := createRandTag(t)
	tag2, err := testQueries.GetTags(context.Background(), tag1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tag2)
	// content
	require.Equal(t, tag2.Name, tag1.Name)
	require.Equal(t, tag2.ID, tag1.ID)

}

func TestGetAllTags(t *testing.T) {
	// TODO: How do a good test here
}

func TestUpdateTags(t *testing.T) {
	tag := createRandTag(t)

	arg := UpdateTagsParams{
		ID:   tag.ID,
		Name: util.RandText(5),
	}

	tag2, err := testQueries.UpdateTags(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tag2)
	// update
	require.Equal(t, tag2.Name, arg.Name)

	require.Equal(t, tag2.ID, tag.ID)
}

func TestDeleteTags(t *testing.T) {
	ctx := context.Background()
	tag1 := createRandTag(t)

	err := testQueries.DeleteTags(ctx, tag1.ID)
	require.NoError(t, err)

	tag2, err := testQueries.GetTags(ctx, tag1.ID)
	require.EqualError(t, err, "no rows in result set")

	require.Empty(t, tag2)
}

func TestListTags(t *testing.T) {
	loog := 5
	for i := 0; i < loog; i++ {
		_ = createRandTag(t)
	}

	arg := ListTagsParams{
		Limit:  5,
		Offset: 5,
	}
	tags, err := testQueries.ListTags(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, tags, 5)

	for _, tag := range tags {
		require.NotEmpty(t, tag)
	}
}

func createRandTag(t *testing.T) Tag {
	// create a tag
	name := util.RandText(10)

	tag, err := testQueries.CreateTags(context.Background(), name)
	// [testify package](https://github.com/stretchr/testify)
	require.NoError(t, err)
	require.NotEmpty(t, tag)
	// content
	require.Equal(t, tag.Name, name)

	require.NotZero(t, tag.ID)

	return tag
}
