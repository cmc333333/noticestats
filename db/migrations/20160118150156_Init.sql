
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE notice (
  id VARCHAR(20) NOT NULL,
  published CHAR(10) NOT NULL,
  PRIMARY KEY (id)
);
CREATE TABLE notice_agency (
  notice_id VARCHAR(20) NOT NULL,
  agency VARCHAR(100) NOT NULL,
  FOREIGN KEY (notice_id) REFERENCES notice(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE notice;
DROP TABLE notice_agency;
