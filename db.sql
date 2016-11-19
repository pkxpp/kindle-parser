CREATE TABLE books(
id PRIMARY KEY,
author TEXT,
title TEXT,
);

CREATE TABLE highlights(
id,
'text' TEXT,
location TEXT,
book_id INT,
page INT,
time DATETIMETZ
);
