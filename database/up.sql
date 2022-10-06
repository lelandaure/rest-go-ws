CREATE TABLE IF NOT EXISTS users(
                                    id VARCHAR(32) PRIMARY KEY,
                                    password VARCHAR(255) NOT NULL,
                                    email VARCHAR(255) NOT NULL,
                                    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);