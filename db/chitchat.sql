CREATE TABLE users (
    id serial primary key,
    uuid varchar(64) not null unique,
    name varchar(64),
    email varchar(255),
    password varchar(255),
    created_at timestamp not null
);

CREATE TABLE sessions (
    id serial primary key,
    uuid varchar(64) not null unique,
    user_id integer references users(id),
    email varchar(255),
    created_at timestamp not null
);

CREATE TABLE threads (
    id serial primary key,
    uuid varchar(64) not null unique,
    topic text,
    user_id integer references users(id),
    created_at timestamp not null
);

CREATE TABLE posts (
    id serial primary key,
    uuid varchar(64) not null unique,
    user_id integer references users(id)
    thread_id integer references threads(id)
    body text,
    created_at timestamp not null
);