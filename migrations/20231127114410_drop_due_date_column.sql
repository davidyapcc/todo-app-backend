-- +goose Up
-- +goose StatementBegin
alter table todos drop column due_date;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table todos add column due_date datetime;
-- +goose StatementEnd
