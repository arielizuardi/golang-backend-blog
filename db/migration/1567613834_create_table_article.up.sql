CREATE TABLE article (
    id SERIAL NOT NULL,
    title TEXT,
    content TEXT, 
    author VARCHAR(100),
    PRIMARY KEY (id)
);