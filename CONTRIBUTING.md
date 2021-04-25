# Contributing

You can participate in NanoDB in several ways, there is so much to do to make this project awesome!

- [Reporting a bug](#reporting-a-bug)
- [Proposing an SQL feature or a general design change](#proposing-an-sql-feature-or-a-general-design-change)
- [Other contributions](#other-contributions)

The goal of this document is to provide guidelines to contributors to make working with NanoDB a smooth experience.
NanoDB is a complex project that has many moving parts, and we want to make sure that everyone's working towards the
same goal.

We want to avoid situations where you would put energy on a contribution that gets rejected because it doesn't match the
direction this project is taking.

## Reporting a bug

We use GitHub issues to track bugs. Report a bug
by [opening an issue](https://github.com/i-sevostyanov/NanoDB/issues/new), it's that easy!

Make sure you give enough context for us to reproduce the bug:

- The version of NanoDB you are using
- The observed behavior
- The expected behavior
- Ideally, some code like a main file or a Go [playground](https://play.golang.org) link

## Proposing an SQL feature, or a general design change

NanoDB is not a common database because it mixes documents with SQL. Some SQL features that make sense in mainstream
relational databases may not be a good fit for this project and vice versa. Every feature can have a big impact on the
project, including performance, ordering, encoding, etc. To propose a consistent solution to users, every SQL feature
needs to be thoroughly designed and discussed before even starting implementing it.

- Open an issue describing the SQL feature you want NanoDB to include
- Discuss it with maintainers and contributors
- If necessary, someone will be in charge of writing an RFC
- Once the feature is accepted, someone will be in charge of creating a Pull Request

Depending on the complexity and importance of the feature, the RFC and the PR may or may not be assigned to you.
However, we will always do our best to let you contribute code if you are eager to do so and the feature doesn't require
too much context.

## Other contributions

Anything that can help improve NanoDB is a welcome contribution. Here are a few examples of things that can be improved:

- CI/Tooling
- Documentation
- Performance

In most cases, contributing to NanoDB must follow these guidelines:

- Open an issue
- Discuss with maintainers and contributors
- Open a PR to propose your solution to the problem

If the change is tiny, like correcting a typo, or a very obvious bug fix, you may open a PR directly.
