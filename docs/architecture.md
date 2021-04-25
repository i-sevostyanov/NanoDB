# Architecture (Work-in-Progress)

## Components

### SQL Engine

The engine contains the following components:

* **Lexer**: Breaks down the SQL statement into tokens.
* **Parser**: Builds an AST (Abstract Syntax Tree) from tokens and performs semantic analysis of the SQL query.
* **Planner**: Builds an execution plan from the AST and optimizes it if possible.
* **Executor**: Executes the plan and collects results.

### Storage

NanoDB enables storing data in tables like you would in a traditional SQL database with a strict schema.
