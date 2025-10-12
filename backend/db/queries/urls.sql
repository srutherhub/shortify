-- name: GetUrlFromID :one
SELECT urls.url, routes.route_name, urls.created_at, urls.expires_at
FROM urls INNER JOIN routes on urls.base_route = routes.route_id
WHERE urls.id = $1
;

-- name: InsertNewUrl :exec
INSERT INTO urls
  (id,url,created_at)
VALUES($1, $2, $3);

-- name: IncrementUrlClickCount :exec
UPDATE urls SET click_count = click_count + 1 WHERE id = $1;

-- name: DeleteExpiredUrls :exec
DELETE FROM urls WHERE id = $1;