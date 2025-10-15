-- name: GetUrlFromID :one
SELECT urls.url, routes.route_name, urls.created_at, urls.expires_at
FROM urls INNER JOIN routes on urls.base_route = routes.route_id
WHERE urls.id = $1
;

-- name: InsertNewUrl :exec
INSERT INTO urls
  (id,url,created_at,expires_at,utm_source,utm_medium,utm_campaign,utm_term,utm_content)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: IncrementUrlClickCount :exec
UPDATE urls SET click_count = click_count + 1 WHERE id = $1;

-- name: DeleteExpiredUrl :exec
DELETE FROM urls WHERE id = $1 AND base_route IN (SELECT route_id
  FROM routes
  WHERE route_name = $2) ;