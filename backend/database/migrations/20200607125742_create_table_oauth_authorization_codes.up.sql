create table oauth_authorization_codes (
    id varchar(100) not null,
    user_id integer not null references users(id),
    expires_at timestamp with time zone not null
);

create unique index unique_active_oauth_authorization_codes on oauth_authorization_codes(id, user_id);
create index order_active_oauth_authorization_codes on oauth_authorization_codes(id, user_id, expires_at desc);
