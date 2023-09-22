-- name: CreateRemoteConfig :execresult
INSERT INTO remote_config(key, value)
VALUES($1, $2);

-- name: GetRemoteConfigByKey :one
SELECT *
FROM remote_config
WHERE remote_config.key = $1
LIMIT 1;

-- name: UpdateRemoteConfig :execresult
UPDATE remote_config
SET value = $2
WHERE key = $1; 
