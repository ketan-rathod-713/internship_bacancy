begin;

create table if not exists gs_info (
    id uuid not null,
    user_id uuid not null,
    server_name text not null,
    game_name text not null,
    image text not null,
    command text not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    primary key (id)
    -- constraint gs_info_user_id_fk foreign key (user_id) references users (id) on delete cascade
);

create unique index if not exists gs_info_user_id_server_name_uindex on gs_info (user_id, server_name);

commit;
