
insert into permissions (name) values('read'),('write'),('execute');
insert into roles (name) values('user'),('supervisor'),('admin');

insert into role_permissions (role_id, permission_id) values (1, 1),(2, 1),(2, 2),(3, 1),(3, 2), (3, 3);

insert into accounts (email, password) values('test@email.com', '$2a$10$yiXzoytWf0kjOhQDxTZ8hOToikT.etROBNSwu0xZRB/OwO7bztym6');

