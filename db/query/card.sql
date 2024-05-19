-- name: GetCard :one
SELECT * FROM cards
WHERE id = $1 LIMIT 1;

-- name: ListCards :many
SELECT * FROM cards
ORDER BY id;

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