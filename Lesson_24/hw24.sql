-- git pull lesson24
ALTER TABLE book
ADD CONSTRAINT fk_author_id FOREIGN KEY (author_id) REFERENCES author(id);

-- inserting into author
insert into author(id,name) values(1,'Sam');
insert into author(id,name) values(2,'Carl');
insert into author(id,name) values(3,'Gabriel');
insert into author(id,name) values(4,'Elsa');
insert into author(id,name) values(5,'Jenny');
insert into author(id,name) values(6,'JK Rowling');

-- inserting into book
insert into book(id,name,page,author_name,author_id) Values(1,'Great',34,'Elsa',4);
insert into book(id,name,page,author_name,author_id) Values(2,'See you',77,'Gabriel',3);
insert into book(id,name,page,author_name,author_id) Values(3,'Love to hate you',245,'Sam',1);
insert into book(id,name,page,author_name,author_id) Values(4,'Indeed',50,'Jenny',5);
insert into book(id,name,page,author_name,author_id) Values(5,'Harry Potter and Philosphers stone',249,'JK Rowling',6);
insert into book(id,name,page,author_name,author_id) Values(6,'You',92,'Carl',2);
insert into book(id,name,page,author_name,author_id) Values(7,'Falling',27,'Sam',1);
insert into book(id,name,page,author_name,author_id) Values(8,'World',84,'Jenny',5);
insert into book(id,name,page,author_name,author_id) Values(9,'Again',123,'Elsa',4);
insert into book(id,name,page,author_name,author_id) Values(10,'Harry Potter and The order of Pheonix',340,'JK Rowling',6);

-- results
select * from author;
select * from book;
Select b.name, a.name
From book As b
Join author As a
On a.id = b.author_id;