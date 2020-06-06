create table workspaces (
    id serial4 primary key,
    name varchar(200) not null,
    created_at timestamp with time zone not null default 'now()',
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);
