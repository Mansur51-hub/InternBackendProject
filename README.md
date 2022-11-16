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

![image](https://user-images.githubusercontent.com/72014699/202283550-48a9035d-590d-4a84-85f6-ce449de8016c.png)

Пополнить баланс:

![image](https://user-images.githubusercontent.com/72014699/202283670-e2ebeac1-cc1c-411d-aff7-a3e142efda94.png)

Смотрим, что в таблице:

![image](https://user-images.githubusercontent.com/72014699/202283829-98b398a5-a74b-42bc-ba3e-64122a761371.png)



