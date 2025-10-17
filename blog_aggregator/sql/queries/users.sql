-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE name = $1 LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT * FROM users;

-- name: CreateFeed :one
INSERT INTO feeds (id, 
                   created_at,
                   updated_at,
                   name,
                   url,
                   user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.name, url, users.name FROM feeds
INNER JOIN users ON feeds.user_id = users.id;

-- name: CreateFeedFollows :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follow (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1, $2, $3, $4, $5)
    RETURNING *   
) SELECT inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users ON inserted_feed_follow.user_id = users.id
INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedByUrl :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetFeedFollowsForUser :many
SELECT feed_follow.*, users.name AS user_name, feeds.name AS feed_name
FROM feed_follow
INNER JOIN users ON feed_follow.user_id = users.id
INNER JOIN feeds on feed_follow.feed_id = feeds.id
WHERE feed_follow.user_id = $1;


-- name: DeleteFeedFollowByUser :exec
DELETE FROM feed_follow
WHERE user_id = $1 AND feed_id = $2;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $1, updated_at = $2
WHERE id = $3;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST 
LIMIT 1;

-- name: CreatePosts :one
INSERT INTO posts (
    id,
    created_at,
    update_at,
    title,
    url,
    description,
    published_at,
    feed_id
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
) RETURNING *;

-- name: GetPostsForUsers :many
SELECT posts.*, feeds.name AS feed_name FROM posts
JOIN feed_follow ON feed_follow.feed_id = posts.feed_id
JOIN feeds ON posts.feed_id = feeds.id
WHERE feed_follow.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;