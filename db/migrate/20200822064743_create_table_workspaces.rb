class CreateTableWorkspaces < ActiveRecord::Migration[6.0]
  def up
    execute <<~SQL
      create table workspaces (
        id serial4 primary key,
        name varchar(310) not null, -- max length of user's name + 10 (for default user workspace)
        created_at timestamp with time zone not null default now(),
        updated_at timestamp with time zone,
        deleted_at timestamp with time zone
      );
    SQL
  end

  def down
    execute 'drop table workspaces;'
  end
end
