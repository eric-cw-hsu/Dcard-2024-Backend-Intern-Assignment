CREATE TABLE platforms (
  name VARCHAR(25) PRIMARY KEY,

  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ads_to_platforms (
  ad_id BIGINT UNSIGNED NOT NULL,
  platform VARCHAR(25) NOT NULL,

  PRIMARY KEY (ad_id, platform),
  FOREIGN KEY (ad_id) REFERENCES ads (id) ON DELETE CASCADE,
  FOREIGN KEY (platform) REFERENCES platforms (name) ON DELETE CASCADE
);

INSERT INTO platforms (name) 
VALUES ('web'), ('ios'), ('android');