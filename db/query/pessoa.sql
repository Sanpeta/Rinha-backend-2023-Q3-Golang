-- name: CreatePessoa :one
INSERT INTO pessoas (id, apelido, nome, nascimento, stack, search_index)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;

-- name: GetPessoa :one
SELECT * FROM pessoas WHERE id = $1;

-- name: GetPessoas :many
SELECT * FROM pessoas
WHERE search_index ILIKE '%' || $1 || '%'
LIMIT 50;

-- name: CountPessoas :one
SELECT COUNT(DISTINCT id) FROM pessoas;