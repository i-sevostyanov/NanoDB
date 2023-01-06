# SQL

Detailed reference documentation for NanoDB's SQL dialect.

* [Data Types](#data-types)
* [SQL Syntax](#sql-syntax)
    * [Identifiers and Keywords](#identifiers-and-keywords)
    * [Numeric Constants](#numeric-constants)
    * [String Constants](#string-constants)
    * [Boolean Constants](#boolean-constants)
    * [Operators](#operators)
    * [Operator Precedence](#operator-precedence)
* [SQL Statements](#sql-statements)
    * Data Definition Language
      * [CREATE DATABASE](#create-database)
      * [DROP DATABASE](#drop-database)
      * [CREATE TABLE](#create-table)
      * [DROP TABLE](#drop-table)
    * Data Manipulation Language  
      * [SELECT](#select)
      * [INSERT](#insert)
      * [UPDATE](#update)
      * [DELETE](#delete)

## Data Types

* `BOOLEAN`: logical truth values, i.e. true and false.
* `FLOAT`: 64-bit signed floating-point numbers.
* `INTEGER`: 64-bit signed integer numbers.
* `TEXT`: UTF-8 encoded strings.
* `NULL`: unknown value ([three-valued logic](https://en.wikipedia.org/wiki/Three-valued_logic)).

## SQL Syntax

### Identifiers and Keywords

Identifiers and keywords must begin with a letter (a-z) or an underscore (_). Subsequent characters in an identifier or
keyword can be letters, digits (0-9) or underscores. For example:

```
abc
local_x
_xyz9
```

### Numeric Constants

Numeric constants are accepted in these general forms:

```
digits
digits.[digits]
```

where `digits` is one or more decimal digits (0-9).

These are some examples of valid numeric constants:

```
42
3.5
```

### String Constants

A string constant in SQL is an arbitrary sequence of characters bounded by single quotes ('), for
example `'This is a string'`.

### Boolean Constants

* `TRUE`: the boolean true value
* `FALSE`: the boolean false value

### Operators

Unary operators:

* `+` (prefix): identity, e.g. +1 yields 1
* `-` (prefix): negation, e.g. -2 yields -2

Binary operators:

* `+`: addition, e.g. 1 + 2 = 3
* `-`: subtraction, e.g. 3 - 2 = 1
* `*`: multiplication, e.g. 3 * 2 = 6
* `/`: division, e.g. 6 / 2 = 3
* `^`: exponentiation, e.g. 2 ^ 4 = 16
* `%`: modulo, e.g. 8 % 3 = 2

Comparison operators:

* `=`: equal
* `!=`: not equal
* `>`: greater than
* `<`: less than
* `>=`: greater or equal
* `<=`: less or equal

Logical operators:

* `AND`: conjunction
* `OR`: disjunction

### Operator Precedence

| Precedence | Operator                        | Associativity |
|------------|---------------------------------|---------------|
| 7          | `+`, `-` (unary plus/minus)     | Right         |
| 6          | `^`                             | Right         |
| 5          | `*`, `/`, `%`                   | Left          |
| 4          | `+`, `-`                        | Left          |
| 3          | `=`, `!=`, `>`, `>=`, `<`, `<=` | Left          |
| 2          | `AND`                           | Left          |
| 1          | `OR`                            | Left          |

## SQL Statements

### CREATE DATABASE

#### Syntax

```
CREATE DATABASE name
```

#### Description

CREATE DATABASE will create a new database.

#### Example

```
CREATE DATABASE films;
```

### DROP DATABASE

#### Syntax

```
DROP DATABASE name
```

#### Description

DROP DATABASE drops a database.

#### Example

```
DROP DATABASE films;
```

### CREATE TABLE

#### Syntax

```
CREATE TABLE table_name (
  [ column_name data_type [ column_constraint [ ... ] ]
  [, ... ]
)
```

where `column_constraint` is:

* NOT NULL
* NULL
* DEFAULT expr
* PRIMARY KEY

#### Description

CREATE TABLE will create a new empty table.

#### Example

```
CREATE TABLE films (
    id        INTEGER PRIMARY KEY,
    code      TEXT NOT NULL,
    title     TEXT NOT NULL,
    is_active BOOLEAN DEFAULT FALSE,
);
```

### DROP TABLE

#### Syntax

```
DROP TABLE table_name
```

#### Description

DROP TABLE removes tables from the database.

#### Example

```
DROP TABLE films;
```

### SELECT

#### Syntax

```
SELECT [ * | expression [ [ AS ] output_name [, ...] ] ]
    [ FROM table_name [, ...] ]
    [ WHERE predicate ]
    [ ORDER BY order_expr [ ASC | DESC ] [, ...] ]
    [ LIMIT count ]
    [ OFFSET start ]
```

#### Description

SELECT retrieves rows from zero or more tables.

#### Example

```
SELECT id, title FROM films LIMIT 2;

 id | title
----+----------------
  1 | Meet Joe Black 
  2 | Seven 
```

### INSERT

#### Syntax

```
INSERT INTO table_name [ ( column_name [, ... ] ) ] VALUES ( expression [, ... ] ) [, ... ]
```

#### Description

INSERT inserts new rows into a table.

#### Example

```
INSERT INTO films (id, code, title, is_active) VALUES (1, 'UA502', 'Bananas', true);
```

### UPDATE

#### Syntax

```
UPDATE table_name SET column_name = expression [, ... ] [ WHERE predicate ]
```

#### Description

UPDATE changes the values of the specified columns in all rows that satisfy the condition. Only the columns to be
modified need be mentioned in the SET clause; columns not explicitly modified retain their previous values.

#### Example

```
UPDATE films SET code = 'UW500' WHERE id = 42;
```

### DELETE

#### Syntax

```
DELETE FROM table_name [ WHERE predicate ]
```

#### Description

DELETE deletes rows that satisfy the WHERE clause from the specified table. If the WHERE clause is absent, the effect is
to delete all rows in the table.

#### Example

```
DELETE FROM films WHERE id = 10;
```
