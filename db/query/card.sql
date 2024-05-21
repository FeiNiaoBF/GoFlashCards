-- name: GetCard :one
SELECT * FROM cards
WHERE id = $1 LIMIT 1;

-- name: GetAllCard :many
SELECT * FROM cards
ORDER BY id;

-- name: ListCards :many
SELECT * FROM cards
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateCards :one
INSERT INTO cards (
    front, back
)VALUES (
             $1, $2
         )
RETURNING *;

-- name: UpdateCards :one
UPDATE cards
SET front = $2,
    back = $3
WHERE id = $1
RETURNING *;

-- name: DeleteCards :exec
DELETE FROM cards
WHERE id = $1;

-- name: GetCardByTag :many
SELECT *
FROM cards c
JOIN card_tags ct
    ON c.id = ct.cards_id
JOIN tags t
    ON t.id = ct.tags_id
WHERE t.id = $1;

-- name: GetAllCardsWithTags :many
SELECT c.*, t.name AS tag_name
FROM cards c
         LEFT JOIN card_tags ct ON c.id = ct.cards_id
         LEFT JOIN tags t ON ct.tags_id = t.id
ORDER BY c.id;