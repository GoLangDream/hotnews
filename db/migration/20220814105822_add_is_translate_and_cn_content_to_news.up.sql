alter table news
    add is_translate bool default false not null;

alter table news
    add cn_content mediumtext null;

update news set is_translate = true;