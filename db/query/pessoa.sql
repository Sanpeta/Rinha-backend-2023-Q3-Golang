-- name: CreatePessoa :one
INSERT INTO pessoas (apelido, nome, nascimento, stack)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetPessoa :one
SELECT * FROM pessoas WHERE id = $1;

-- name: GetPessoas :many
SELECT * FROM pessoas 
WHERE apelido LIKE '%' || @t || '%'
OR nome LIKE '%' || @t || '%'
OR stack @> ARRAY[@t]::VARCHAR[]
LIMIT 50;

-- name: CountPessoas :one
SELECT COUNT(DISTINCT id) FROM pessoas;