create table user_settings (
    id serial4 primary key,
    user_id integer not null references users(id),
    settings jsonb not null default '{}'
);

create unique index unique_user_id_on_user_settings on user_settings(user_id);
