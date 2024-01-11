-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id) 
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, name, url, user_id;

-- name: GetFeeds :many
SELECT id, created_at, updated_at, name, url, user_id
FROM feeds
ORDER BY created_at DESC;