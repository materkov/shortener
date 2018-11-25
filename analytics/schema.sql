CREATE SCHEMA analytics;

CREATE TABLE analytics.click
(
  id     serial,
  url_id int,
  date   timestamp
);
