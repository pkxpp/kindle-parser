CREATE TABLE books(
       id INT PRIMARY KEY,
       author text,
       title text,
       UNIQUE(author, title)
);

CREATE TABLE highlights(
       id INT PRIMARY KEY,
       'text' TEXT,
       location TEXT,
       book_id INT REFERENCES books(id),
       page INT,
       time DATETIMETZ,
       UNIQUE('text', location, book_id)
);
