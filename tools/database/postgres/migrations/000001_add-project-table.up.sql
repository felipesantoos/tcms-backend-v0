create extension if not exists "uuid-ossp";

create table if not exists project (
    id uuid not null
       constraint pk_project_id primary key
       constraint df_project_id default uuid_generate_v4(),
    name varchar(200) not null
        constraint uk_project_name unique,
    description text null,
    is_active bool not null
        constraint df_project_is_active default true,
    is_deleted bool not null
        constraint df_project_is_deleted default false,
    created_at timestamp not null
        constraint df_project_created_at default now(),
    updated_at timestamp null
);
