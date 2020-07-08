create database todo_app;
use todo_app;
create table todo(
    id integer primary key AUTO_INCREMENT,
    title varchar(200) not null,
    created_on datetime default CURRENT_TIMESTAMP,
    limit_on datetime not null,
    is_done boolean default false,
    is_important boolean default false
);