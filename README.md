# solution-a

#create test table

CREATE TABLE test
(
id int NOT NULL auto_increment,
LastName varchar(255),
FirstName varchar(255),
gender ENUM('Man', 'Woman')
)

# env vars
export GOPATH=/Users/striker/web/solution-a
export MYSQL_DSN="root:my-secret-pw@tcp(localhost:3306)/test"
