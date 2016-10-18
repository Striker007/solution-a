# solution-a

**solution-a** - *modern data faker.*

**current status**: *under active development*.

[![N|Solid](http://ic.pics.livejournal.com/droids_life/68035718/427/427_600.jpg)](http://github.com/Striker007/solution-a)

#create test table

```sql
CREATE TABLE test
(
id int NOT NULL auto_increment,
LastName varchar(255),
FirstName varchar(255),
gender ENUM('Man', 'Woman')
)
```

# env vars
```sh
export GOPATH=/Users/striker/web/solution-a
export MYSQL_DSN="root:my-secret-pw@tcp(localhost:3306)/test"
```
