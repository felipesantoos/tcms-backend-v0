-- name: InsertIntoProject :one
insert into project (id, name, description, is_active, is_deleted, created_at, updated_at)
values (default, $1, $2, default, default, default, default)
returning id;
