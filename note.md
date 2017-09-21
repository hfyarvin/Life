```sql
UPDATE  s SET origin_per_price = (SELECT w.per_price from  w WHERE w.id = s.project_invest_way_id) where s.project_invest_way_id > 0;
```
```sql
UPDATE  s SET origin_per_price = (SELECT p.person_min_crowd_funding from  p WHERE p.id = s.project_id) where (s.project_invest_way_id = 0 or s.project_invest_way_id is null);
```
```sql
Add
alter table t1 add c1 int(11);
alter table t2 add p bigint(20);
```