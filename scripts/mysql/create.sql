DROP TABLE IF EXISTS user;

CREATE TABLE IF NOT EXISTS
    user (
        id UUID PRIMARY KEY NOT NULL,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL,
    );

CREATE INDEX IF NOT EXISTS
    user_email_index ON user (email);

CREATE INDEX IF NOT EXISTS
    user_name_index ON user (name);