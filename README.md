# InternBackendProject

## Запуск

Нужно установить данные package
`go get -u github.com/gorilla/mux`
`go get -u github.com/go-sql-driver/mysql`

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

![image](https://user-images.githubusercontent.com/72014699/202286006-0f984eff-5966-41b4-92c1-7124d6034aa9.png)

Пополнить баланс:

![image](https://user-images.githubusercontent.com/72014699/202286101-57675495-0e7b-4c1e-b0e6-0104af537bec.png)

Зарезервировать данные:

![image](https://user-images.githubusercontent.com/72014699/202287489-6b4eba01-39b6-4be7-95c8-e5df03729877.png)


В табличке reservations появились данные:

![image](https://user-images.githubusercontent.com/72014699/202287737-c7165b70-f557-4c6f-862b-e60a9edf3b31.png)


Списать из резерва деньги и добавить данные в отчет:

![image](https://user-images.githubusercontent.com/72014699/202288484-e6985cba-55f2-48d6-aec3-bfc3eef3edd9.png)


Проверяем, что таблица с резервацией стала пустой и данные добавились в отчет:

![image](https://user-images.githubusercontent.com/72014699/202288717-03099eb1-934d-40be-abd1-45be6a95965d.png)



