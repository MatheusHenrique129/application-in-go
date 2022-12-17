CREATE TABLE IF NOT EXISTS users (
    id BIGINT NOT NULL AUTO_INCREMENT,
    full_name VARCHAR(55) NOT NULL,
    cpf VARCHAR(15) NOT NULL,
    email TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    gender TEXT NOT NULL,
    passworld TEXT NOT NULL,
    birth_date TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,

    PRIMARY KEY id_pk (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;