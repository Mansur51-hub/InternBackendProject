# InternBackendProject

## Команды для создания таблиц:

create table users ( 
id int primary key unique not null, 
balance float not null
);

create table reservations ( 
order_id int primary key unique not null,
user_id int not null,
service_id int not null,
amount float not null
);

create table services ( 
id int primary key unique not null, 
description varchar(30)
);

INSERT INTO services VALUES (1, 'Service1');
INSERT INTO services VALUES (2, 'Service2');
INSERT INTO services VALUES (3, 'Service3');
INSERT INTO services VALUES (4, 'Service4');

create table reports ( 
order_id int primary key unique not null,
user_id int not null,
service_id int not null,
amount float not null,
date date not null
);

## Методы

| Method | Url        | Body | Description |
| ------ | ---------- | ---- | ----------- |
| Get    | /balances  | id   | Получить баланс пользователя | 
| Post   | /balances  | id, balance | Пополнить баланс пользователя |
| Post   | /reservations | order_id, user_id, service_id, amount | Зарезервировать деньги у пользователя |
| Delete   | /reservations | order_id, user_id, service_id, amount | Списать деньги с резерва и добавить в отчет |

# Примеры с postman

Получить баланс пользователя:

![image](https://user-images.githubusercontent.com/72014699/202284270-f72f2e57-1d6d-4e58-8f1a-699451a3e5e7.png)

Пополнить баланс:

![image](https://user-images.githubusercontent.com/72014699/202284423-fff7e206-7f6e-4192-b3fe-68faaf8af36e.png)

Смотрим, что в таблице:

![image](https://user-images.githubusercontent.com/72014699/202284531-747334aa-c535-4306-a04a-aa790c148982.png)




