CREATE TABLE people (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL UNIQUE,
    date_of_birth TIMESTAMP NOT NULL
);

INSERT INTO people(name, date_of_birth) 
    VALUES  ('test1','2000-01-01'), 
            ('test2','2001-01-01'), 
            ('test3','2002-01-01'), 
            ('test4','2003-01-01'),
            ('test5','2004-01-01');
