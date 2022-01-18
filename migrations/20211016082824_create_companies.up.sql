create table companies
(
    company_id uuid default uuid_generate_v4() not null,
    name varchar not null,
    uniform_number int
);

create unique index companies_company_id_uindex
    on companies (company_id);

alter table companies
    add constraint companies_pk
        primary key (company_id);

create index companies_company_id_index
    on companies USING hash(company_id);

create index companies_name_index
    on companies using gin
        (name collate pg_catalog."default" gin_trgm_ops)
    tablespace pg_default;

create index companies_uniform_number_index
    on companies (uniform_number);
