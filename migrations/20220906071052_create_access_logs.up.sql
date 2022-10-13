create table access_logs
(
    id         uuid      default uuid_generate_v4() not null
        primary key,
    ip_address text      default ''::text           not null,
    member_id  uuid                                 not null,
    method     text      default ''::text           not null,
    api        text      default ''::text           not null,
    msg        text      default ''::text           not null,
    status     text      default ''::text           not null,
    created_at timestamp default now()              not null,
    updated_at timestamp default now()              not null,
    deleted_at timestamp
);

create index idx_access_logs_id
    on access_logs using hash (id);

create index idx_access_logs_ip_address
    on access_logs using gin (ip_address gin_trgm_ops);

create index idx_access_logs_member_id
    on access_logs using hash (member_id);

create index idx_access_logs_method
    on access_logs using gin (method gin_trgm_ops);

create index idx_access_logs_api
    on access_logs using gin (api gin_trgm_ops);

create index idx_access_logs_msg
    on access_logs using gin (msg gin_trgm_ops);

create index idx_access_logs_status
    on access_logs using gin (status gin_trgm_ops);

create index idx_access_logs_created_at
    on access_logs (created_at desc);

create index idx_access_logs_updated_at
    on access_logs (updated_at desc);

create index idx_access_logs_deleted_at
    on access_logs (deleted_at desc);