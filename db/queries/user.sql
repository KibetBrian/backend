-- name: RegisterUser :one
INSERT INTO users (full_name, email, password)
VALUES($1, $2, $3) RETURNING *;

-- name: UpdateUser :one
UPDATE users SET email = $1, password = $2 WHERE email = $3 RETURNING *;

