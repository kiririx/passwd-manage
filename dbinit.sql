create table t_passwords
(
    id          int auto_increment
        primary key,
    topic       varchar(1000) not null,
    user_id     int           not null,
    username    varchar(200)  null,
    password    varchar(1000) null,
    description varchar(1000) null,
    createTime  mediumtext    null,
    updateTime  mediumtext    null
);

create table t_users
(
    id         int auto_increment
        primary key,
    username   varchar(200)  not null,
    password   varchar(1000) not null,
    createTime mediumtext    null,
    updateTime mediumtext    null
);

