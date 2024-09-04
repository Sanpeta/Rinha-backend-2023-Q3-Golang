-- ALTER SYSTEM SET max_connections = 300;
-- ALTER SYSTEM SET log_error_verbosity = 'TERSE';

-- Use an extension to enable trigram similarity search and improve LIKE performance
-- https://www.postgresql.org/docs/current/runtime-config-connection.htmlhttps://mazeez.dev/posts/pg-trgm-similarity-search-and-fast-like
CREATE EXTENSION pg_trgm;

ALTER DATABASE rinha_backend_2023 SET synchronous_commit=OFF;
-- using 25% of memory as suggested in the docs:
--    https://www.postgresql.org/docs/9.1/runtime-config-resource.html
-- ALTER SYSTEM SET shared_buffers TO "425MB";

-- debug slow queries, run \d pg_stat_statements
-- docs: 
--    https://www.postgresql.org/docs/current/pgstatstatements.html
-- CREATE EXTENSION pg_stat_statements;
-- ALTER SYSTEM SET shared_preload_libraries = 'pg_stat_statements';

CREATE TABLE IF NOT EXISTS pessoas (
    id UUID NOT NULL PRIMARY KEY,
    apelido VARCHAR(32) NOT NULL,
    nome VARCHAR(100) NOT NULL,
    nascimento VARCHAR(10) NOT NULL,
    stack VARCHAR(32)[] NULL,
    search_index TEXT NOT NULL
);

CREATE INDEX pessoas_search_index_idx ON pessoas USING gin (search_index gin_trgm_ops);