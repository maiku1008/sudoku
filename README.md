# sk1u/doku

## Intro

A sudoku solver written in Go, that can solve *any* puzzle;
It uses [Constraint Propagation](https://en.wikipedia.org/wiki/Constraint_satisfaction) and [Search](https://en.wikipedia.org/wiki/Search_algorithm) algorithms.


## Installation

## Use

We interface with this application with a simple cli app for running locally, or with a server app where we can wrap its main functions in REST Api endpoints, useful for creating a full stack web application. Such interfaces are described below.

### Cli

### Server

## TODO:
- [ ] Complete Readme
- [ ] Add Go CI
- [ ] Write simple CLI app
- [x] Validate if hash exists in server app
- [ ] Write a DisplayString() that will simply output the solved sudoku in string form
- [ ] Write a test and a handler for the above DisplayString()
- [ ] Consider splitting up the session struct into a request and response struct instead
- [ ] Dockerize server app
- [ ] Write tests for remaining helper functions
- [ ] Use different storage other than a map

---

This project has been inspired by [this](https://norvig.com/sudoku.html) blog post.