-- name: GetTags :one
SELECT * FROM tags 
WHERE id = $1 LIMIT 1;

-- name: GetAllTags :many
SELECT * FROM tags
ORDER BY id;

-- name: ListTags :many
SELECT * FROM tags
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateTags :one
INSERT INTO tags (name)
VALUES ($1)
RETURNING *;

-- name: UpdateTags :one
UPDATE tags
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTags :exec
DELETE FROM tags
WHERE id = $1;
