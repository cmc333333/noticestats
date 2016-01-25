
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE stats (
  id VARCHAR(20) NOT NULL,
  xml_len INT NOT NULL,
  regtext_len INT NOT NULL,
  page_len INT NOT NULL,
  is_correction BOOL NOT NULL,
  PRIMARY KEY(id)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE stats;
