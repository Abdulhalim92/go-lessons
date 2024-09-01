-- Создание базы данных
CREATE DATABASE my_practical_db;

-- Подключение к базе данных
\c my_practical_db;

-- Создание таблиц

-- Таблица компании
CREATE TABLE company (
    company_id SERIAL PRIMARY KEY,
    company_name VARCHAR(255) NOT NULL,
    industry VARCHAR(255)
);

-- Таблица сотрудников (отношение один ко многим с таблицей компании)
CREATE TABLE employee (
    employee_id SERIAL PRIMARY KEY,
    employee_name VARCHAR(255) NOT NULL,
    position VARCHAR(255),
    salary DECIMAL(10, 2),
    company_id INT REFERENCES company(company_id)
);

-- Таблица проектов
CREATE TABLE project (
    project_id SERIAL PRIMARY KEY,
    project_name VARCHAR(255) NOT NULL,
    budget DECIMAL(15, 2),
    deadline DATE,
    company_id INT REFERENCES company(company_id)
);

-- Связывающая таблица для сотрудников и проектов (много ко многим)
CREATE TABLE employee_project (
    employee_id INT REFERENCES employee(employee_id),
    project_id INT REFERENCES project(project_id),
    PRIMARY KEY (employee_id, project_id)
);

-- Таблица контактной информации (отношение один к одному с таблицей сотрудников)
CREATE TABLE contact_info (
    contact_id SERIAL PRIMARY KEY,
    phone_number VARCHAR(20),
    email VARCHAR(255),
    address VARCHAR(255),
    employee_id INT UNIQUE REFERENCES employee(employee_id)
);

-- Вставка данных в таблицы

-- Вставка компаний
INSERT INTO company (company_name, industry) VALUES
    ('Tech Innovators', 'Technology'),
    ('Health Solutions', 'Healthcare'),
    ('Eco Energy', 'Energy');

-- Вставка сотрудников
INSERT INTO employee (employee_name, position, salary, company_id) VALUES
    ('John Doe', 'Software Engineer', 80000, 1),
    ('Jane Smith', 'Project Manager', 95000, 1),
    ('Alice Johnson', 'Data Analyst', 65000, 2),
    ('Bob Brown', 'Sales Manager', 70000, 3),
    ('Charlie Green', 'Technician', 55000, 3);

-- Вставка проектов
INSERT INTO project (project_name, budget, deadline, company_id) VALUES
    ('AI Development', 500000, '2025-12-31', 1),
    ('Health Monitoring System', 200000, '2024-06-30', 2),
    ('Solar Panel Installation', 300000, '2024-09-15', 3);

-- Вставка связей между сотрудниками и проектами
INSERT INTO employee_project (employee_id, project_id) VALUES
    (1, 1),
    (2, 1),
    (3, 2),
    (4, 3),
    (5, 3);

-- Вставка контактной информации
INSERT INTO contact_info (phone_number, email, address, employee_id) VALUES
    ('+1234567890', 'johndoe@example.com', '123 Tech Lane', 1),
    ('+0987654321', 'janesmith@example.com', '456 Project Blvd', 2),
    ('+1122334455', 'alicejohnson@example.com', '789 Data Ave', 3),
    ('+1222333444', 'bobbrown@example.com', '321 Sales St', 4),
    ('+1555666777', 'charliegreen@example.com', '654 Tech Park', 5);
