create table users
(
    id          bigint auto_increment,
    user_id     bigint       not null,
    username    varchar(256) not null,
    encrypt_pwd varchar(256) not null,
    constraint user_id_uindex
        unique (id),
    constraint user_user_id_uindex
        unique (user_id),
    constraint user_username_uindex
        unique (username)
);

create table encrypt_info
(
    user_id bigint       not null,
    salt    varchar(256) not null,
    constraint encrypt_info_user_id_uindex
        unique (user_id),
    constraint encrypt_info_users_user_id_fk
        foreign key (user_id) references users (user_id)
);

alter table encrypt_info
    add primary key (user_id);
