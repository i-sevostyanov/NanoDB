# SQL
Detailed reference documentation for NanoDB's SQL dialect.

## Data Types
* BOOLEAN: logical truth values, i.e. true and false.
* FLOAT: 64-bit signed floating-point numbers.
* INTEGER: 64-bit signed integer numbers.
* STRING: UTF-8 encoded strings.
* NULL: unknown value ([three-valued logic](https://en.wikipedia.org/wiki/Three-valued_logic)).

## SQL Syntax

### Identifiers and Keywords
Identifiers and keywords must begin with a letter (a-z) or an underscore (_). 
Subsequent characters in an identifier or keyword can be letters, underscores or digits (0-9).

### Numeric Constants
Numeric constants are accepted in these general forms:
```
digits
digits.[digits]
```
where digits is one or more decimal digits (0-9).

These are some examples of valid numeric constants:
```
42
3.5
```

### String Constants
A string constant in SQL is an arbitrary sequence of characters bounded by single quotes ('), for example 'This is a string'.

### Boolean Constants

* `TRUE`: the boolean true value.
* `FALSE`: the boolean false value.

### Operators

Unary operators:

* `+` (prefix): identity, e.g. +1 yields 1.
* `-` (prefix): negation, e.g. -2 yields -2.

Binary operators:

* `+`: addition, e.g. 1 + 2 yields 3.
* `-`: subtraction, e.g. 3 - 2 yields 1.
* `*`: multiplication, e.g. 3 * 2 yields 6.
* `/`: division, e.g. 6 / 2 yields 3.
* `^`: exponentiation, e.g. 2 ^ 4 yields 16.
* `%`: modulo or remainder, e.g. 8 % 3 yields 2. The result has the sign of the divisor.

### Operator Precedence (highest to lowest)

| Precedence | Operator                        | Associativity |
| ---------- | ------------------------------- | ------------- |
| 7          | `+`, `-` (unary plus/minus)     | Right         |
| 6          | `^`                             | Right         |
| 5          | `*`, `/`, `%`                   | Left          |
| 4          | `+`, `-`                        | Left          |
| 3          | `=`, `!=`, `>`, `>=`, `<`, `<=` | Left          |
| 2          | `AND`                           | Left          |
| 1          | `OR`                            | Left          |

## SQL Statements


### CREATE TABLE

#### Syntax
```
CREATE TABLE table_name (
    [ column_name data_type ] [, ... ]
)
```

#### Description
CREATE TABLE will create a new empty table.

#### Example
```
CREATE TABLE films (
    id        INTEGER,
    code      STRING,
    title     STRING,
    is_active BOOLEAN,
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
UPDATE changes the values of the specified columns in all rows that satisfy the condition. 
Only the columns to be modified need be mentioned in the SET clause; columns not explicitly modified retain their previous values.

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
DELETE deletes rows that satisfy the WHERE clause from the specified table. 
If the WHERE clause is absent, the effect is to delete all rows in the table. The result is a valid, but empty table.

#### Example
```
DELETE FROM films WHERE id = 10;
```
