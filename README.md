[![CI](https://github.com/i-sevostyanov/NanoDB/actions/workflows/go.yml/badge.svg)](https://github.com/i-sevostyanov/NanoDB/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/i-sevostyanov/NanoDB)](https://goreportcard.com/report/github.com/i-sevostyanov/NanoDB)
[![codecov](https://codecov.io/gh/i-sevostyanov/NanoDB/branch/main/graph/badge.svg?token=y0lxdfxXdT)](https://codecov.io/gh/i-sevostyanov/NanoDB)
[![GitHub license](https://img.shields.io/github/license/i-sevostyanov/NanoDB)](https://github.com/i-sevostyanov/NanoDB/blob/main/LICENSE)

# NanoDB

SQL database, written as a learning project to better understand the internals of a database.

## Features

* A good starting point to dive into database internals
* Modular and extendable architecture
* Implemented widely usable SQL statements (SELECT, INSERT, UPDATE, DELETE, [and more](docs/sql.md))
* Interactive terminal for easy commands execution and experiments

## Documentation

* [Architecture](docs/architecture.md)
* [SQL reference](docs/sql.md)
* [References](docs/references.md)

## Shell

The NanoDB command line provides an SQL shell that can be used to select, insert, delete, and modify data.

### Run locally

```shell
git clone https://github.com/i-sevostyanov/NanoDB.git
cd NanoDB
go run ./cmd/shell/main.go
#> \import <path to project>/testdata/demo.sql
```

### Examples:

Imagine that we have a table with the following definition:

```sql
CREATE TABLE aircrafts
(
    id    INTEGER PRIMARY KEY,
    code  TEXT    NOT NULL,
    model TEXT    NOT NULL,
    range INTEGER NOT NULL
);
```

#### SELECT

Select all columns:

```shell
demo #> SELECT * FROM aircrafts
+----+------+---------------------+-------+
| id | code |        model        | range |
+----+------+---------------------+-------+
| 1  | 773  | Boeing 777-300      | 11100 |
| 2  | 763  | Boeing 767-300      | 7900  |
| 3  | SU9  | Sukhoi Superjet-100 | 3000  |
| 4  | 320  | Airbus A320-200     | 5700  |
| 5  | 321  | Airbus A321-200     | 5600  |
| 6  | 319  | Airbus A319-100     | 6700  |
| 7  | 733  | Boeing 737-300      | 4200  |
| 8  | CN1  | Cessna 208 Caravan  | 1200  |
| 9  | CR2  | Bombardier CRJ-200  | 2700  |
+----+------+---------------------+-------+
(9 rows)

```

Select specified columns with order, limit and offset:

```shell
demo #> SELECT id, model, range FROM aircrafts WHERE range < 5000 ORDER BY range ASC LIMIT 5 OFFSET 2
+----+---------------------+-------+
| id |        model        | range |
+----+---------------------+-------+
| 3  | Sukhoi Superjet-100 | 3000  |
| 7  | Boeing 737-300      | 4200  |
+----+---------------------+-------+
(2 rows)
```

Select expression:

```shell
demo #> SELECT 6+(2^3)*5-3+4/(10-2)%3
+----------------------------------------------------+
| (((6 + ((2 ^ 3) * 5)) - 3) + ((4 / (10 - 2)) % 3)) |
+----------------------------------------------------+
| 43                                                 |
+----------------------------------------------------+
(1 rows)
```

#### INSERT

```shell
demo #> INSERT INTO aircrafts (code, model, range) VALUES ('773', 'Boeing 777-300', 11100);
Query OK, 1 row affected
```

#### UPDATE

```shell
demo #> UPDATE airports SET code = 'MRN' WHERE id = 2
Query OK, 1 row affected
```

#### DELETE

```shell
demo #> DELETE FROM airports WHERE id > 5
Query OK, 4 rows affected
```

## Roadmap

* Indices
* Joins
* Subqueries
* Aggregations (count, sum, avg, max, min)
* Explain and query optimization
* Transactions ([ACID](https://en.wikipedia.org/wiki/ACID),
  [MVCC](https://en.wikipedia.org/wiki/Multiversion_concurrency_control), 
  [WAL](https://en.wikipedia.org/wiki/Write-ahead_logging))
* Constraints (unique key, check, foreign key)
* Built-in functions (sin, cos, abs, floor, etc)

## Contributing

Contributions are welcome!

See [CONTRIBUTING.md](CONTRIBUTING.md) for more details.
