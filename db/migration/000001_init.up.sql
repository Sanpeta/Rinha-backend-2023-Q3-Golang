CREATE TABLE IF NOT EXISTS pessoas (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    apelido VARCHAR(32) NOT NULL,
    nome VARCHAR(100) NOT NULL,
    nascimento VARCHAR(10) NOT NULL,
    stack VARCHAR(32)[]
);