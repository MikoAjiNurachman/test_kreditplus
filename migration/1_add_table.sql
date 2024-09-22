-- +migrate Up
create table if not exists customers (
    id serial primary key,
    nik varchar(16),
    full_name varchar(20),
    legal_name varchar(20),
    birth_place varchar(10),
    birth_date date,
    sallary bigint,
    image_ktp varchar(100),
    image_selfie varchar(100),
    created_at timestamp without time zone default current_timestamp,
    updated_at timestamp without time zone default null,
    deleted boolean default false,
    unique(nik)
);

create table if not exists transactions (
    id serial primary key,
    customer_id int references customers(id),
    contract_number varchar(8) unique not null,
    otr_price decimal,
    admin_fee decimal,
    installment_amount decimal,
    -- cicilan
    rate_amount decimal,
    -- bunga
    asset_name varchar(30),
    total_payment decimal,
    -- kalo ada otr & admin fee 
    created_at timestamp without time zone default current_timestamp,
    updated_at timestamp without time zone default null,
    deleted boolean default false
);

create table if not exists limit_customers (
    id serial primary key,
    customer_id int references customers(id),
    "year" int, 
    tenor_1 int,
    tenor_2 int,
    tenor_3 int,
    tenor_4 int,
    tenor_5 int,
    tenor_6 int,
    tenor_7 int,
    tenor_8 int,
    tenor_9 int,
    tenor_10 int,
    tenor_11 int,
    tenor_12 int,
    created_at timestamp without time zone default current_timestamp,
    updated_at timestamp without time zone default null,
    deleted boolean default false,
    unique(customer_id, "year")
)

-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
drop table if exists customers;
drop table if exists transactions;
drop table if exists limit_customers;