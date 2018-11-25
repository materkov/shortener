CREATE SCHEMA redirecter;

CREATE TABLE redirecter.shortened_url
(
  id       serial,
  url      varchar(255)
);