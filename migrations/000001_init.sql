-- +goose Up
CREATE TABLE users (
                       id BIGSERIAL PRIMARY KEY,
                       username VARCHAR(64) NOT NULL UNIQUE,
                       hashed_password TEXT NOT NULL,
                       is_admin BOOLEAN NOT NULL DEFAULT FALSE,
                       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE resources (
                           id BIGSERIAL PRIMARY KEY,
                           title VARCHAR(255) NOT NULL,
                           title_en VARCHAR(255) NOT NULL DEFAULT '',
                           description TEXT NOT NULL,
                           resource_type VARCHAR(32) NOT NULL,
                           status VARCHAR(32) NOT NULL DEFAULT 'PENDING',
                           poster_image TEXT,
                           tmdb_id BIGINT,
                           media_type VARCHAR(32),
                           likes_count INTEGER NOT NULL DEFAULT 0,
                           created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                           updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE approval_records (
                                  id BIGSERIAL PRIMARY KEY,
                                  resource_id BIGINT NOT NULL REFERENCES resources(id) ON DELETE CASCADE,
                                  status VARCHAR(32) NOT NULL,
                                  notes TEXT NOT NULL DEFAULT '',
                                  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE site_settings (
                               id BIGSERIAL PRIMARY KEY,
                               setting_key VARCHAR(100) NOT NULL UNIQUE,
                               setting_value JSONB NOT NULL DEFAULT '{}'::jsonb,
                               created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                               updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS site_settings;
DROP TABLE IF EXISTS approval_records;
DROP TABLE IF EXISTS resources;
DROP TABLE IF EXISTS users;