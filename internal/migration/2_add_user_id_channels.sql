-- +goose Up
ALTER TABLE channels ADD COLUMN user_id VARCHAR(255) NOT NULL;
