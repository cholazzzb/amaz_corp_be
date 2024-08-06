-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- +migrate Up
CREATE TABLE IF NOT EXISTS products (
  id serial UNIQUE NOT NULL PRIMARY KEY,
  name text NOT NULL
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS roles (
  id serial UNIQUE NOT NULL PRIMARY KEY,
  name text NOT NULL
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
  id uuid UNIQUE NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
  username text UNIQUE NOT NULL,
  password text NOT NULL,
  salt text NOT NULL,
  product_id integer NOT NULL,
  role_id integer NOT NULL,
  CONSTRAINT fk_product_id FOREIGN KEY(product_id) REFERENCES products(id),
  CONSTRAINT fk_role_id FOREIGN KEY(role_id) REFERENCES roles(id) ON DELETE CASCADE
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS features (
  id uuid UNIQUE NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
  name text NOT NULL,
  max_limit integer 
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS product_feature (
  product_id serial NOT NULL,
  feature_id uuid NOT NULL,
  CONSTRAINT fk_product_id FOREIGN KEY(product_id) REFERENCES products(id),
  CONSTRAINT fk_feature_id FOREIGN KEY(feature_id) REFERENCES features(id)
);