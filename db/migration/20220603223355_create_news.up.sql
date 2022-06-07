create table news
(
    id          int auto_increment,
    title       varchar(255) not null,
    cn_title    varchar(255) not null,
    content     mediumtext   not null,
    url         varchar(255) not null,
    source_id   varchar(255) not null,
    source_name varchar(255) not null,
    created_at  datetime     null,
    updated_at  datetime     null,
    deleted_at  datetime     null,
    constraint news_pk
        primary key (id)
);

create unique index news_id_uindex
    on news (id);

create unique index news_source_id_source_name_uindex
    on news (source_id, source_name);

