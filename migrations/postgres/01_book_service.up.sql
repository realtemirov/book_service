CREATE TABLE IF NOT EXISTS "books" (
    "id" text PRIMARY KEY,
    "title" varchar(50) NOT NULL,
    "description" varchar(255) NOT NULL,
    "author" varchar(50) NOT NULL,
    "price" int NOT NULL    
);

-- Path: migrations/postgres/02_book_service.down.sql
-- INSERT INTO "books" ("id", "title", "description", "author", "price") VALUES ('1', 'The Lord of the Rings', 'The Lord of the Rings is an epic high fantasy novel written by English author and scholar J. R. R. Tolkien. Written in stages between 1937 and 1949, The Lord of the Rings is one of the best-selling novels ever written, with over 150 million copies sold.', 'J. R. R. Tolkien', 100);
-- INSERT INTO "books" ("id", "title", "description", "author", "price") VALUES ('2', 'The Hobbit', 'The Hobbit, or There and Back Again is a childrens fantasy novel by English author J. R. R. Tolkien. It was published on 21 September 1937 to wide critical acclaim, being nominated for the Carnegie Medal and awarded a prize from the New York', 'J. R. R. Tolkien', 50);
-- INSERT INTO "books" ("id", "title", "description", "author", "price") VALUES ('3', 'Albert', 'The Chronicles of Narnia is a series of seven fantasy novels by C. S. Lewis. It is considered a classic of childrens literature and is the author best known work, having sold over 100 million copies in 47 languages.', 'C. S. Lewis', 75);


CREATE TABLE IF NOT EXISTS "orders" (
    "id" text PRIMARY KEY,
    "user_id" text NOT NULL,
    "book_id" text NOT NULL,
    "quantity" text NOT NULL,
    "total" text NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp 
);

--INSERT INTO "orders"(id,user_id,book_id,quantity,total,created_at ) VALUES ('1','1','1','1','100','2020-01-01 00:00:00');