
-- +goose Up
CREATE TABLE IF NOT EXISTS task (
  id int(11) NOT NULL AUTO_INCREMENT,
  title varchar(256) DEFAULT NULL,
  description varchar(256) DEFAULT NULL,
  priority int(11) DEFAULT NULL,
  created_at datetime DEFAULT NULL,
  updated_at datetime DEFAULT NULL,
  completed_at datetime DEFAULT NULL,
  is_deleted tinyint(1) NOT NULL,
  is_completed tinyint(1) NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- SQL in section 'Up' is executed when this migration is applied


-- +goose Down
  DROP TABLE IF EXISTS task;
-- SQL section 'Down' is executed when this migration is rolled back

