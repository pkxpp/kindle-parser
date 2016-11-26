CREATE TABLE books(id int primary key, author text, title text);


CREATE TABLE highlights(
id INT PRIMARY KEY,
'text' TEXT,
location TEXT,
book_id INT,
page INT,
time DATETIMETZ
);
