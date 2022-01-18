CREATE EXTENSION IF NOT EXISTS "pg_trgm";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
 
create table accounts
(
    account_id uuid not null default uuid_generate_v4(),
    organization_id uuid not null ,
    account varchar not null,
    name varchar not null,
    pwd varchar not null,
    role_id uuid not null,
    is_deleted bool default false not null,
    created_at timestamp default now() not null,
    created_by uuid null ,
    updated_at timestamp null ,
    updated_by uuid null
);

create unique index accounts_id_uindex
    on accounts (account_id);

alter table accounts
    add constraint accounts_pk
        primary key (account_id);

create index accounts_id_index
    on accounts USING hash(account_id);

create index accounts_account_index
    on accounts (account);

create index accounts_name_index
    on accounts using gin
        (name collate pg_catalog."default" gin_trgm_ops)
    tablespace pg_default;

create index accounts_role_id_index
    on accounts USING hash(role_id);

create index accounts_organization_id_index
    on accounts USING hash(organization_id);

create index accounts_created_at_index
    on accounts (created_at desc );