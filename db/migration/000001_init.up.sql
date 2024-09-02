CREATE TABLE IF NOT EXISTS pessoas (
    id UUID DEFAULT gen_random_uuid(),
    apelido VARCHAR(32) NOT NULL,
    nome VARCHAR(100) NOT NULL,
    nascimento VARCHAR(10) NOT NULL,
    stack VARCHAR(32)[],
    PRIMARY KEY (id)
);