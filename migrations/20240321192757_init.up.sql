begin;

create schema if not exists news;

create table if not exists news.news
(
    id          bigserial
        constraint news_pk
            primary key,
    title       varchar(200)                      not null,
    description text                              not null,
    img         text,
    create_at   time with time zone default now() not null
);

end;