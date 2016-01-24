
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE notice (
  id VARCHAR(20) NOT NULL,
  published DATE NOT NULL,
  PRIMARY KEY (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE notice;
