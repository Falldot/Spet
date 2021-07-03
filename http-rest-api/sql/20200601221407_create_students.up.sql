CREATE TABLE students (
    id bigserial not null primary key,
    surname varchar,
    middleName varchar,
    name varchar,
    date_b varchar,
    city varchar,
    street varchar,
    house varchar,
    flat varchar,
    phone varchar,
    info varchar,
    numGroup varchar,
    activs varchar,
    gender varchar,
    status varchar,
    orphan bool,
    invalid bool,
    low_income bool,
    low_parents bool,
    idn bool,
    kdn bool,
    many_children bool,
    login varchar,
    password varchar,
    budget varchar
);

CREATE TABLE users (
    id bigserial not null primary key,
    login       varchar,
    password    varchar,
    role        int
);

