### command
- mysql -hlocalhost -uroot -p
- show databases
- CREATE DATABASE IF NOT EXISTS yiibaidb DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
- use database; //选择要使用的数据库
- souce D:/kkkkk ;//导入数据
- CREATE DATABASE [IF NOT EXISTS] database_name;

- CREATE TABLE [IF NOT EXISTS] table_name(
        column_list
) engine=table_type;

### xorm
- go get github.com/go-xorm/cmd/xorm
- go get github.com/go-xorm/xorm
- go get github.com/go-sql-driver/mysql
