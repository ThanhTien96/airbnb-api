# PostgreSQL Example
## Model Inventory
![alt](https://www.w3resource.com/w3r_images/movie-database.png)

## Requirements
- PostgreSQL
- TablePlus or something alike
- Gorm

## Usage

1. First, do the migration of the model to database:
```
go run main.go -m=true
```

2. Then, generate the database
```
go run main.go -gd=true
```

3. If you want to clear the database
Do your query
```
go run main.go -c=true
```

To run the file, we have many options:
```
Options:
c: to clean the fake data in database                           (DEFAULT: false)
gd: generate fake data                  (DEFAULT: false)
m: migrate database table from model    (DEFAULT: false)
``` 

## References
+ https://gorm.io/docs/
+ https://gorm.io/gen/index.html
+ https://www.w3resource.com/sql-exercises/