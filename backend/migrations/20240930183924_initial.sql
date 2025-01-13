-- +goose Up
-- SQL в этой секции будет выполнен для обновления БД

create table contracts(
    id uuid PRIMARY KEY,
    name varchar(100) not null,
    sum numeric(10, 2) not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    signed_at TIMESTAMP,
    status varchar(100) NOT NULL,
    is_deleted boolean
);

create table employees(
    id uuid PRIMARY KEY,
    name text not null,
    position varchar(100) not null,
    employee_num varchar(20) not null,
    is_deleted boolean
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

create table details(
    id uuid PRIMARY KEY,
    type text not null,
    name text not null,
    price numeric(10, 2) not null,
    is_deleted boolean
);

create table storages(
    id uuid PRIMARY KEY,
    storage_num varchar(20) not null,
    employee_id uuid NOT NULL REFERENCES employees(id),
    detail_id uuid NOT NULL REFERENCES details(id),
    detail_count int not null,
    is_deleted boolean
);


insert into services (id, name) values 
(gen_random_uuid(), 'Починка двигателя'),
(gen_random_uuid(), 'Замена подвески'),
(gen_random_uuid(), 'Замена шин');

INSERT INTO contracts (id, name, sum, signed_at, status, is_deleted) VALUES
(gen_random_uuid(), 'Договор на обслуживание', 10000.00, '2025-01-10 14:00:00', 'Подписан', false),
(gen_random_uuid(), 'Договор аренды', 20000.00, '2025-01-11 15:30:00', 'В обработке', false),
(gen_random_uuid(), 'Договор поставки', 15000.00, NULL, 'Черновик', false),
(gen_random_uuid(), 'Договор на ремонт', 5000.00, '2025-01-12 10:00:00', 'Подписан', false),
(gen_random_uuid(), 'Договор на услуги', 7000.00, '2025-01-13 11:15:00', 'Завершён', false),
(gen_random_uuid(), 'Договор покупки', 25000.00, '2025-01-09 09:00:00', 'Подписан', false),
(gen_random_uuid(), 'Договор подряда', 12000.00, NULL, 'В обработке', false),
(gen_random_uuid(), 'Договор продажи', 18000.00, '2025-01-08 13:45:00', 'Черновик', false),
(gen_random_uuid(), 'Договор на перевозку', 30000.00, '2025-01-10 14:00:00', 'Подписан', false),
(gen_random_uuid(), 'Договор аренды оборудования', 17000.00, '2025-01-12 12:00:00', 'Завершён', false);


INSERT INTO employees (id, name, position, employee_num, is_deleted) VALUES
(gen_random_uuid(), 'Иванов Иван', 'Главный менеджер по работе с клиентами', 'EMP001', false),
(gen_random_uuid(), 'Петров Петр', 'Главный менеджер по работе с клиентами', 'EMP002', false),
(gen_random_uuid(), 'Сидорова Анна', 'Менеджер по работе с клиентами', 'EMP003', false),
(gen_random_uuid(), 'Кузнецов Сергей', 'Менеджер по работе с клиентами', 'EMP004', false),
(gen_random_uuid(), 'Михайлова Ольга', 'Менеджер по работе с клиентами', 'EMP005', false),
(gen_random_uuid(), 'Смирнов Алексей', 'Младший механик', 'EMP006', false),
(gen_random_uuid(), 'Федорова Мария', 'Младший механик', 'EMP007', false),
(gen_random_uuid(), 'Тихонов Дмитрий', 'Старший механик', 'EMP008', false),
(gen_random_uuid(), 'Борисова Елена', 'Младший механик', 'EMP009', false),
(gen_random_uuid(), 'Гордеев Николай', 'Старший механик', 'EMP010', false);

INSERT INTO clients (id, employee_id, name, phone, has_documents, passport, is_deleted) VALUES
(gen_random_uuid(), (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), 'Клиент 1', '+79991234567', true, '1234567890', false),
(gen_random_uuid(), (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), 'Клиент 2', '+79997654321', false, '0987654321', false),
(gen_random_uuid(), (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), 'Клиент 3', '+79993456789', true, '1112223334', false),
(gen_random_uuid(), (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), 'Клиент 4', '+79991237890', false, '3334445556', false),
(gen_random_uuid(), (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), 'Клиент 5', '+79998765432', true, '5556667778', false),
(gen_random_uuid(), (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), 'Клиент 6', '+79990001122', false, '7778889990', false),
(gen_random_uuid(), (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), 'Клиент 7', '+79994445566', true, '0001112223', false),
(gen_random_uuid(), (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), 'Клиент 8', '+79993334455', false, '4445556667', false),
(gen_random_uuid(), (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), 'Клиент 9', '+79992223344', true, '6667778889', false),
(gen_random_uuid(), (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), 'Клиент 10', '+79991112233', false, '9990001112', false);

INSERT INTO receipts (id, contract_id, sum, is_deleted) VALUES
(gen_random_uuid(), (SELECT id FROM contracts ORDER BY RANDOM() LIMIT 1), 5000.00, false),
(gen_random_uuid(), (SELECT id FROM contracts ORDER BY RANDOM() LIMIT 1), 15000.00, false),
(gen_random_uuid(), (SELECT id FROM contracts ORDER BY RANDOM() LIMIT 1), 10000.00, false),
(gen_random_uuid(), (SELECT id FROM contracts ORDER BY RANDOM() LIMIT 1), 2500.00, false),
(gen_random_uuid(), (SELECT id FROM contracts ORDER BY RANDOM() LIMIT 1), 3000.00, false),
(gen_random_uuid(), (SELECT id FROM contracts ORDER BY RANDOM() LIMIT 1), 12000.00, false),
(gen_random_uuid(), (SELECT id FROM contracts ORDER BY RANDOM() LIMIT 1), 8000.00, false),
(gen_random_uuid(), (SELECT id FROM contracts ORDER BY RANDOM() LIMIT 1), 9000.00, false),
(gen_random_uuid(), (SELECT id FROM contracts ORDER BY RANDOM() LIMIT 1), 20000.00, false),
(gen_random_uuid(), (SELECT id FROM contracts ORDER BY RANDOM() LIMIT 1), 17000.00, false);

INSERT INTO vehicles (id, client_id, vehicle_number, brand, model, is_deleted) VALUES
(gen_random_uuid(), (SELECT id FROM clients ORDER BY RANDOM() LIMIT 1), 'А123ВС77', 'Toyota', 'Camry', false),
(gen_random_uuid(), (SELECT id FROM clients ORDER BY RANDOM() LIMIT 1), 'В456ОР78', 'BMW', 'X5', false),
(gen_random_uuid(), (SELECT id FROM clients ORDER BY RANDOM() LIMIT 1), 'С789ЕХ99', 'Audi', 'A4', false),
(gen_random_uuid(), (SELECT id FROM clients ORDER BY RANDOM() LIMIT 1), 'Е111КУ77', 'Mercedes', 'C-Class', false),
(gen_random_uuid(), (SELECT id FROM clients ORDER BY RANDOM() LIMIT 1), 'К222ОР78', 'Hyundai', 'Solaris', false),
(gen_random_uuid(), (SELECT id FROM clients ORDER BY RANDOM() LIMIT 1), 'М333СР99', 'Ford', 'Focus', false),
(gen_random_uuid(), (SELECT id FROM clients ORDER BY RANDOM() LIMIT 1), 'Н444ТР77', 'Lada', 'Vesta', false),
(gen_random_uuid(), (SELECT id FROM clients ORDER BY RANDOM() LIMIT 1), 'О555УР78', 'Nissan', 'Qashqai', false),
(gen_random_uuid(), (SELECT id FROM clients ORDER BY RANDOM() LIMIT 1), 'П666ФР99', 'Volkswagen', 'Tiguan', false),
(gen_random_uuid(), (SELECT id FROM clients ORDER BY RANDOM() LIMIT 1), 'Р777ЦР77', 'Mazda', 'CX-5', false);

INSERT INTO applications (id, employee_id, client_id, name, status, contract_id, is_deleted) VALUES
(gen_random_uuid(), (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1),
(SELECT id FROM clients ORDER BY RANDOM() LIMIT 1), 'Заявка 1', 'В процессе', (SELECT id FROM contracts ORDER BY RANDOM() LIMIT 1), false),
(gen_random_uuid(), (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), (SELECT id FROM clients ORDER BY RANDOM() LIMIT 1), 'Заявка 2', 'Завершена', (SELECT id FROM contracts ORDER BY RANDOM() LIMIT 1), false),
(gen_random_uuid(), (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), (SELECT id FROM clients ORDER BY RANDOM() LIMIT 1), 'Заявка 3', 'Отменена', (SELECT id FROM contracts ORDER BY RANDOM() LIMIT 1), false),
(gen_random_uuid(), (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), (SELECT id FROM clients ORDER BY RANDOM() LIMIT 1), 'Заявка 4', 'В процессе', (SELECT id FROM contracts ORDER BY RANDOM() LIMIT 1), false),
(gen_random_uuid(), (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), (SELECT id FROM clients ORDER BY RANDOM() LIMIT 1), 'Заявка 5', 'Завершена', (SELECT id FROM contracts ORDER BY RANDOM() LIMIT 1), false);

INSERT INTO details (id, type, name, price, is_deleted) VALUES
(gen_random_uuid(), 'Двигатель', 'Двигатель V8', 250000.00, false),
(gen_random_uuid(), 'Подвеска', 'Передняя подвеска', 70000.00, false),
(gen_random_uuid(), 'Шины', 'Зимние шины', 40000.00, false),
(gen_random_uuid(), 'Масло', 'Масло моторное', 5000.00, false),
(gen_random_uuid(), 'Фильтр', 'Фильтр воздушный', 3000.00, false),
(gen_random_uuid(), 'Тормоза', 'Тормозные диски', 15000.00, false),
(gen_random_uuid(), 'Сцепление', 'Комплект сцепления', 20000.00, false),
(gen_random_uuid(), 'Аккумулятор', 'Аккумулятор 12В', 10000.00, false),
(gen_random_uuid(), 'Радиатор', 'Радиатор охлаждения', 25000.00, false),
(gen_random_uuid(), 'Фары', 'Галогенные фары', 8000.00, false);

-- Таблица storages
INSERT INTO storages (id, storage_num, employee_id, detail_id, detail_count, is_deleted) VALUES
(gen_random_uuid(), 'СТ001', (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), (SELECT id FROM details ORDER BY RANDOM() LIMIT 1), 10, false),
(gen_random_uuid(), 'СТ002', (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), (SELECT id FROM details ORDER BY RANDOM() LIMIT 1), 15, false),
(gen_random_uuid(), 'СТ003', (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), (SELECT id FROM details ORDER BY RANDOM() LIMIT 1), 20, false),
(gen_random_uuid(), 'СТ004', (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), (SELECT id FROM details ORDER BY RANDOM() LIMIT 1), 5, false),
(gen_random_uuid(), 'СТ005', (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), (SELECT id FROM details ORDER BY RANDOM() LIMIT 1), 7, false),
(gen_random_uuid(), 'СТ006', (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), (SELECT id FROM details ORDER BY RANDOM() LIMIT 1), 12, false),
(gen_random_uuid(), 'СТ007', (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), (SELECT id FROM details ORDER BY RANDOM() LIMIT 1), 9, false),
(gen_random_uuid(), 'СТ008', (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), (SELECT id FROM details ORDER BY RANDOM() LIMIT 1), 6, false),
(gen_random_uuid(), 'СТ009', (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), (SELECT id FROM details ORDER BY RANDOM() LIMIT 1), 8, false),
(gen_random_uuid(), 'СТ010', (SELECT id FROM employees ORDER BY RANDOM() LIMIT 1), (SELECT id FROM details ORDER BY RANDOM() LIMIT 1), 11, false);

-- Таблица acts
INSERT INTO acts (id, name, created_at, application_id, service_id, is_deleted) VALUES
(gen_random_uuid(), 'Акт выполнения работ 1', '2025-01-10 14:00:00', (SELECT id FROM applications ORDER BY RANDOM() LIMIT 1), (SELECT id FROM services ORDER BY RANDOM() LIMIT 1), false),
(gen_random_uuid(), 'Акт выполнения работ 2', '2025-01-11 10:30:00', (SELECT id FROM applications ORDER BY RANDOM() LIMIT 1), (SELECT id FROM services ORDER BY RANDOM() LIMIT 1), false),
(gen_random_uuid(), 'Акт выполнения работ 3', '2025-01-12 09:15:00', (SELECT id FROM applications ORDER BY RANDOM() LIMIT 1), (SELECT id FROM services ORDER BY RANDOM() LIMIT 1), false),
(gen_random_uuid(), 'Акт выполнения работ 4', '2025-01-13 16:45:00', (SELECT id FROM applications ORDER BY RANDOM() LIMIT 1), (SELECT id FROM services ORDER BY RANDOM() LIMIT 1), false),
(gen_random_uuid(), 'Акт выполнения работ 5', '2025-01-14 11:20:00', (SELECT id FROM applications ORDER BY RANDOM() LIMIT 1), (SELECT id FROM services ORDER BY RANDOM() LIMIT 1), false),
(gen_random_uuid(), 'Акт выполнения работ 6', '2025-01-15 08:10:00', (SELECT id FROM applications ORDER BY RANDOM() LIMIT 1), (SELECT id FROM services ORDER BY RANDOM() LIMIT 1), false),
(gen_random_uuid(), 'Акт выполнения работ 7', '2025-01-16 14:50:00', (SELECT id FROM applications ORDER BY RANDOM() LIMIT 1), (SELECT id FROM services ORDER BY RANDOM() LIMIT 1), false),
(gen_random_uuid(), 'Акт выполнения работ 8', '2025-01-17 09:30:00', (SELECT id FROM applications ORDER BY RANDOM() LIMIT 1), (SELECT id FROM services ORDER BY RANDOM() LIMIT 1), false),
(gen_random_uuid(), 'Акт выполнения работ 9', '2025-01-18 12:40:00',
(SELECT id FROM applications ORDER BY RANDOM() LIMIT 1), (SELECT id FROM services ORDER BY RANDOM() LIMIT 1), false),
(gen_random_uuid(), 'Акт выполнения работ 10', '2025-01-19 10:25:00', (SELECT id FROM applications ORDER BY RANDOM() LIMIT 1), (SELECT id FROM services ORDER BY RANDOM() LIMIT 1), false);

-- +goose Down
-- SQL в этой секции будет выполнен для отката изменений

drop table if EXISTS receipts;
drop table if EXISTS vehicles;
drop table if EXISTS contracts;
drop table if EXISTS clients;
drop table if EXISTS applications;
drop table if EXISTS employees;
drop table if EXISTS acts;
drop table if EXISTS storages;
drop table if EXISTS details;
drop table if EXISTS services;