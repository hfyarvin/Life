### create table user_coll(name varchar(20), email varchar(40));
### 删 drop table public.user_coll;
### 设置从属数据库 alter table public.user_coll owner to postgres;

### alter table public.user_coll drop column name;
### alter table public.user_coll add column age int default(11);
### alter table public.user_coll add column age integer;
### alter table public.user_coll alter column age set default 11;

### alter table public.user_coll rename email to contract;  
### 改表名 alter table user_coll rename to user_info;
### drop table if exist user_coll;

## --------------------------------------------------------------------------
- alter aggregate name (type) rename to new_name
- alter aggregate name (type) owner to new_owner

### conversion
- alter conversion name rename to new_name
- alter conversion name owner to new_owner