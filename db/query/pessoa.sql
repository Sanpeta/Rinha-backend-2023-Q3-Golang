-- name: CreatePessoa :one
INSERT INTO pessoas (apelido, nome, nascimento, stack)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetPessoa :one
SELECT * FROM pessoas WHERE id = $1;

-- name: GetPessoas :many
SELECT * FROM pessoas 
WHERE apelido LIKE '%' || $1 || '%'
OR nome LIKE '%' || $1 || '%'
OR stack @> ARRAY[$1]
LIMIT 50;

-- name: CountPessoa :one
SELECT COUNT(DISTINCT id) FROM pessoas;