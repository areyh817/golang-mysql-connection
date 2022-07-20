mysql> create table test(
    -> id int auto_increment primary key,
    -> name varchar(30),
    -> age int
    -> );
Query OK, 0 rows affected (0.02 sec)

mysql> show tables;
+---------------------+
| Tables_in_goproject |
+---------------------+
| test                |
+---------------------+