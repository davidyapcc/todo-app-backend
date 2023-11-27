-- +goose Up
-- +goose StatementBegin
create table if not exists todos
(
    id          int auto_increment primary key,
    name        varchar(255)             null,
    description text                     null,
    due_date    datetime                 null,
    status      varchar(20)              null,
    created_at  datetime default (now()) null,
    updated_at  datetime                 null on update CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists todos;
-- +goose StatementEnd
