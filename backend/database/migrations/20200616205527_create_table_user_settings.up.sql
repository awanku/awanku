create table user_settings (
    id serial4 primary key,
    user_id integer not null references users(id),
    settings jsonb not null default '{}'
);
