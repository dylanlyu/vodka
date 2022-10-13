create table members_phone
(
    id           uuid      default uuid_generate_v4() not null
        primary key,
    phone_code   text      default ''::text           not null,
    phone_number text      default ''::text           not null,
    passwd       text      default ''::text           not null,
    is_enable    bool      default true               not null,
    created_at   timestamp default now()              not null,
    updated_at   timestamp default now()              not null,
    deleted_at   timestamp
);

create index idx_members_phone_id
    on members_phone using hash (id);

create index idx_members_phone_phone_number
    on members_phone using gin (phone_number gin_trgm_ops);

create index idx_members_phone_phone_code
    on members_phone using gin (phone_code gin_trgm_ops);

create index idx_members_phone_created_at
    on members_phone (created_at desc);

create index idx_members_phone_updated_at
    on members_phone (updated_at desc);

create index idx_members_phone_deleted_at
    on members_phone (deleted_at desc);

