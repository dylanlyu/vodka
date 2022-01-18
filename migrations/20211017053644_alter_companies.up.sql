alter table companies
    add is_deleted bool default false not null;

alter table companies
    add created_at timestamp default now() not null;

alter table companies
    add created_by uuid null;

alter table companies
    add updated_at timestamp null;

alter table companies
    add updated_by uuid null;