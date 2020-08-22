class CreateTableProjects < ActiveRecord::Migration[6.0]
  def up
    execute <<~SQL
      create table projects (
        id serial4 primary key,
        name varchar(200) not null,
        workspace_id integer not null references workspaces(id),
        created_at timestamp with time zone not null default now(),
        updated_at timestamp with time zone,
        deleted_at timestamp with time zone
      );
    SQL
  end

  def down
    execute 'drop table projects;'
  end
end
