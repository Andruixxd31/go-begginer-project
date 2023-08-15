CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

Create Table IF NOT EXISTS account(
    id uuid DEFAULT uuid_generate_v4(),
    name VARCHAR(15) NOT NULL,
    create_at Date NOT NULL DEFAULT CURRENT_DATE,
    update_at Date NOT NULL,
    deleted_at DATE NOT NULL,
    CONSTRAINT account_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS book(
    id uuid DEFAULT uuid_generate_v4(),
    account_id uuid NOT NULL,
    title TEXT NOT NULL ,
    author TEXT NOT NULL,
    year INT,
    likes INT DEFAULT 0,
    create_at Date NOT NULL DEFAULT CURRENT_DATE,
    update_at Date NOT NULL,
    deleted_at DATE NOT NULL,
    CONSTRAINT book_pkey PRIMARY KEY (id),
    CONSTRAINT account_fk FOREIGN KEY (account_id) REFERENCES account(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS upvote(
    account_id uuid NOT NULL,
    book_id uuid NOT NULL,
    
    CONSTRAINT upvote_key PRIMARY KEY (account_id, book_id),
    CONSTRAINT account_fk FOREIGN KEY (account_id) REFERENCES account(id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT book_fk FOREIGN KEY (book_id) REFERENCES book(id) ON DELETE CASCADE ON UPDATE CASCADE
);
