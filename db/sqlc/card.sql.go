// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: card.sql

package sqlc

import (
	"context"
)

const createCards = `-- name: CreateCards :one
INSERT INTO cards (front, back, tags_id)
VALUES ($1, $2, $3)
RETURNING id, front, back, know, tags_id, "add_Time"
`

type CreateCardsParams struct {
	Front  string `json:"front"`
	Back   string `json:"back"`
	TagsID int32  `json:"tags_id"`
}

func (q *Queries) CreateCards(ctx context.Context, arg CreateCardsParams) (Card, error) {
	row := q.db.QueryRow(ctx, createCards, arg.Front, arg.Back, arg.TagsID)
	var i Card
	err := row.Scan(
		&i.ID,
		&i.Front,
		&i.Back,
		&i.Know,
		&i.TagsID,
		&i.AddTime,
	)
	return i, err
}

const deleteCards = `-- name: DeleteCards :exec
DELETE FROM cards
WHERE id = $1
`

func (q *Queries) DeleteCards(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteCards, id)
	return err
}

const getAllCard = `-- name: GetAllCard :many
SELECT id, front, back, know, tags_id, "add_Time"
FROM cards
ORDER BY id
`

func (q *Queries) GetAllCard(ctx context.Context) ([]Card, error) {
	rows, err := q.db.Query(ctx, getAllCard)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Card
	for rows.Next() {
		var i Card
		if err := rows.Scan(
			&i.ID,
			&i.Front,
			&i.Back,
			&i.Know,
			&i.TagsID,
			&i.AddTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCard = `-- name: GetCard :one
SELECT id, front, back, know, tags_id, "add_Time"
FROM cards
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetCard(ctx context.Context, id int64) (Card, error) {
	row := q.db.QueryRow(ctx, getCard, id)
	var i Card
	err := row.Scan(
		&i.ID,
		&i.Front,
		&i.Back,
		&i.Know,
		&i.TagsID,
		&i.AddTime,
	)
	return i, err
}

const getCardByTag = `-- name: GetCardByTag :many
SELECT id, front, back, know, tags_id, "add_Time"
FROM cards
WHERE tags_id = $1
`

func (q *Queries) GetCardByTag(ctx context.Context, tagsID int32) ([]Card, error) {
	rows, err := q.db.Query(ctx, getCardByTag, tagsID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Card
	for rows.Next() {
		var i Card
		if err := rows.Scan(
			&i.ID,
			&i.Front,
			&i.Back,
			&i.Know,
			&i.TagsID,
			&i.AddTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCards = `-- name: ListCards :many
SELECT id, front, back, know, tags_id, "add_Time"
FROM cards
ORDER BY id
LIMIT $1 OFFSET $2
`

type ListCardsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCards(ctx context.Context, arg ListCardsParams) ([]Card, error) {
	rows, err := q.db.Query(ctx, listCards, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Card
	for rows.Next() {
		var i Card
		if err := rows.Scan(
			&i.ID,
			&i.Front,
			&i.Back,
			&i.Know,
			&i.TagsID,
			&i.AddTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCards = `-- name: UpdateCards :one
UPDATE cards
SET front = $2,
    back = $3,
    tags_id = $4
WHERE id = $1
RETURNING id, front, back, know, tags_id, "add_Time"
`

type UpdateCardsParams struct {
	ID     int64  `json:"id"`
	Front  string `json:"front"`
	Back   string `json:"back"`
	TagsID int32  `json:"tags_id"`
}

func (q *Queries) UpdateCards(ctx context.Context, arg UpdateCardsParams) (Card, error) {
	row := q.db.QueryRow(ctx, updateCards,
		arg.ID,
		arg.Front,
		arg.Back,
		arg.TagsID,
	)
	var i Card
	err := row.Scan(
		&i.ID,
		&i.Front,
		&i.Back,
		&i.Know,
		&i.TagsID,
		&i.AddTime,
	)
	return i, err
}
