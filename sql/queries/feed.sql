-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, feed_name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :one
SELECT * FROM feeds
WHERE user_id = $1;

-- name: GetAllFeeds :many
SELECT * FROM feeds
ORDER BY user_id;
