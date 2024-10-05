-- +goose Up
-- SQL в этой секции будет выполнен для обновления БД

create table contracts(
    id uuid PRIMARY KEY,
    name varchar(100) not null,
    sum numeric(10, 2) not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    signed_at TIMESTAMP,
    status_id uuid NOT NULL
);

create table employees(
    id uuid PRIMARY KEY,
    name text not null,
    position varchar(100) not null,
    employee_num varchar(20) not null
);

create table clients(
    id uuid PRIMARY KEY,
    employee_id uuid NOT NULL REFERENCES employees(id),
    name text not null,
    phone varchar(20) not null,
    has_documents boolean,
    passport varchar(10) not null,
    is_deleted boolean
);

CREATE TABLE receipts(
    id uuid PRIMARY KEY,
    contract_id uuid NOT NULL REFERENCES contracts(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    sum NUMERIC(10, 2) NOT NULL,
    is_deleted boolean
);

create table statuses(
    id uuid PRIMARY KEY,
    name varchar(100) not null
);

create table vehicles(
    id uuid PRIMARY KEY,
    client_id uuid NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
    vehicle_number varchar(10) not null,
    brand varchar(100) not null,
    model varchar(100) not null,
    is_deleted boolean
);

create table applications(
    id uuid PRIMARY KEY,
    employee_id uuid NOT NULL REFERENCES employees(id),
    client_id uuid not null REFERENCES clients(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    name text not null,
    status varchar(100) not null,
    contract_id uuid NOT NULL REFERENCES contracts(id),
    is_deleted boolean
);

create table services(
    id uuid PRIMARY KEY,
    name text not null
);

create table acts(
    id uuid PRIMARY KEY,
    name text not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    application_id uuid NOT NULL REFERENCES applications(id),
    service_id uuid  NOT NULL REFERENCES services(id),
    is_deleted boolean
);

create table storages(
    id uuid PRIMARY KEY,
    employee_id uuid NOT NULL REFERENCES employees(id),
    storage_num int not null
);

create table details(
    id uuid PRIMARY KEY,
    storage_id uuid NOT NULL REFERENCES storages(id),
    name text not null,
    price numeric(10, 2) not null
);

-- +goose Down
-- SQL в этой секции будет выполнен для отката изменений

drop table if EXISTS receipts;
drop table if EXISTS statuses;
drop table if EXISTS vehicles;
drop table if EXISTS contracts;
drop table if EXISTS clients;
drop table if EXISTS applications;
drop table if EXISTS employees;
drop table if EXISTS acts;
drop table if EXISTS storages;
drop table if EXISTS details;
drop table if EXISTS services;