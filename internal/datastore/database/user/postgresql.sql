-- name: CreateUser :execresult
INSERT INTO users(username, password, salt, product_id)
VALUES ($1, $2, $3, $4);

-- name: GetUser :one
SELECT * 
FROM users
WHERE username = $1
LIMIT 1;

-- name: GetListUserByUsername :many
SELECT id, username
FROM users
WHERE username LIKE $1
LIMIT 10;

-- name: GetUserExistance :one
SELECT EXISTS(
    SELECT *
    FROM users
    WHERE username = $1
    LIMIT 1
);

-- name: GetProductByUserID :one
SELECT * FROM products
LEFT JOIN users ON users.product_id = products.id
WHERE users.id = $1
LIMIT 1;

-- name: GetListProduct :many
SELECT * FROM products;

-- name: GetListFeatureByProductID :many
SELECT * FROM features
LEFT JOIN product_feature ON features.id = product_feature.feature_id
WHERE product_feature.product_id = $1;