CREATE TABLE IF NOT EXISTS Opening (
  Role     BIGSERIAL PRIMARY KEY, 
	Company  varchar(100) NOT NULL UNIQUE, 
	Location varchar(255),
	Remote   boolean DEFAULT TRUE,
	Link     varchar(255),
	Salary   INT
)