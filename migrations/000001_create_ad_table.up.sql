CREATE TABLE ads (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  start_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  end_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  age_start INT NOT NULL DEFAULT 0,
  age_end INT NOT NULL DEFAULT 100,
  gender ENUM('M', 'F')
);