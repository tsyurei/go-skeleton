CREATE TABLE user (
  id serial unique,
  name    varchar(40),
  email   varchar(40),
  hashed_password bytea
);