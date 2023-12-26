- [Crypto Alert System](#crypto-alert-system)
  - [Setup](#setup)
  - [System design](#system-design)
  - [Database Design](#database-design)
  - [Deep dive](#deep-dive)
- [todos](#todos)

# Crypto Alert System

## Setup

## System design

## Database Design

## Deep dive

# todos
- [X] Create a rest API endpoint for the user’s to create an alert `alerts/create/`
- [X] Create a rest API endpoint for the user’s to delete an alert `alerts/delete/`
- [X] Create a rest API endpoint to `fetch all the alerts` that the user has created.
- [X] The response should also include the `status of the alerts` (created, triggered, completed, deleted)
- [X] `Paginate` the response.
- [X] Include `filter` options based on the status of the alerts
- [X] binance’s `websocket` connection to get real time price updates
- [X]When the price of the coin reaches the price specified by the users, send an email to the user that the target price has been hit `smtp`
- [X] Add a caching layer `redis`
- [X] Add user authentication to the endpoints. Use `JWT` tokens.
- [X] Go
- [X] Use `Postgres` to store data and  `redis` for sorted sets (or any DB you feel that gets the job done)
- [X] `Kafka` as a message broker for the task to send emails /print the output
- [X] Bundle everything inside a docker-compose file so it’s easier for us to test.
- [X] Document your solution in README.md file. Consider adding the following details
- [] Steps to run the project (eg: docker-compose up)
- [] Document the endpoints
- [] Document the solution for sending alerts
