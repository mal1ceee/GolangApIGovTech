# GolangApIGovTech

Technical assessment for GovTech Software Engineer internship using Golang.
- Building a backend API for assessment
- Postman collection(Desktop to Test it locally): https://www.postman.com/kevvvvinn/workspace/govtech/collection/30008705-b10c384c-3def-49f7-b499-c7cf54cdb430?action=share&creator=30008705
- 

## Prerequitesities 

- GO 1.22
- PostgreSQL 15
- Gin Web Framework
- sqlx Package
- pq Package
- (optional) Postman

## Installation guide 

1. Run git clone https://github.com/mal1ceee/GolangApIGovTech.git

2. After clonning the repo, run go mody tidy in root directory to install dependancies

3. Run CREATE DATABASE postgres

4. Run -U postgres -d postgres -a -f (Path to database_schema.sql)

5. psql -U postgres -d postgres -a -f (Path to seed.sql) to populate table

6. After clonning the repo, run go mod tidy to install dependancies

7. To run the server, cd cmd/server/main.go to start the application 



## Change log

This is a record of all past commits for easy access and documentation.

| Date | Changes |
|--------|--------|
| 1402024 | 1. Initialize the file structure and added boilerplate <br> 2. Created and Tested the db connections using postgre 15 |
| 15022024 | 1. Testing db connection to extract the the data from the table and printing it |
| 16022024 | 1. Implemented the handler, service and repository code + bug fixes |
| 18022024 | Fixed notifications and register API + refactor and bug fixes |

## Notes

1. If there is any changes to the username, password or port, the config can be changed in the config.go file.
2. Please install the necessary version as there are features that might not be available in older/newer version
3. Still relatively new to Go so please dont mind the log messages and weird code (Had to google alot) :D
4. Thanks for looking through the code!!

