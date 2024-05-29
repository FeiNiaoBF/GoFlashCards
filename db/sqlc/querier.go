// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateCards(ctx context.Context, arg CreateCardsParams) (Card, error)
	CreateTags(ctx context.Context, name string) (Tag, error)
	DeleteCards(ctx context.Context, id int64) error
	DeleteTags(ctx context.Context, id int64) error
	GetAllCard(ctx context.Context) ([]Card, error)
	GetAllTags(ctx context.Context) ([]Tag, error)
	GetCard(ctx context.Context, id int64) (Card, error)
	GetCardByTag(ctx context.Context, tagsID int32) ([]Card, error)
	GetTags(ctx context.Context, id int64) (Tag, error)
	ListCards(ctx context.Context, arg ListCardsParams) ([]Card, error)
	ListTags(ctx context.Context, arg ListTagsParams) ([]Tag, error)
	UpdateCards(ctx context.Context, arg UpdateCardsParams) (Card, error)
	UpdateTags(ctx context.Context, arg UpdateTagsParams) (Tag, error)
}

var _ Querier = (*Queries)(nil)
