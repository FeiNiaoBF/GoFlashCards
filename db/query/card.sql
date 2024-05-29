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
    front, back, tags_id
)VALUES (
             $1, $2, $3
         )
RETURNING *;

-- name: UpdateCards :one
UPDATE cards
SET front = $2,
    back = $3,
    tags_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteCards :exec
DELETE FROM cards
WHERE id = $1;

-- name: GetCardByTag :many
SELECT *
FROM cards
WHERE tags_id = $1;
