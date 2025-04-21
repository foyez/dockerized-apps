-- name: CreateTodo :one
INSERT INTO todos (
  text
) VALUES (
  $1
) RETURNING *;

-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTodo :one
UPDATE todos
SET is_done = $1
WHERE id = $2
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1;