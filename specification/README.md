# Specification (v1)

This folder contains all the necessary details of the Loft REST API.

## Base rules

- Follow rules from [JSON API](http://jsonapi.org/) v1.0 Stable

## Loft

### Add a new loft (along with a new user)

Method: `POST`

Endpoint: `/loft/`

Params: None

Body: [JSON sample](request/post-loft.json)

The body consists of

- `loft`: the loft that the user is trying to create
- `user`: the user that is creating the loft.

TODO: user needs some kind of authentication data (password-like)

Response: [JSON sample](response/post-loft.json)

## To-Do list

### Get tasks

Method: `GET`

Endpoint: `/loft/{loft_id}/todos`

Params: TODO

Body: None

Response: [JSON sample](response/get-todos.json)

TODO: authentication token missing, pagination

### Add a task

Method: `POST`

Endpoint: `/loft/{loft_id}/todos`

Params: None

Body: [JSON sample](request/post-todos.json)

Response: [JSON sample](response/post-todos.json)

TODO: authentication token missing

## Update a task

Method: `PUT`

Endpoint: `/loft/{loft_id}/todos/{task_id}`

Params: None

Body: [JSON sample](request/put-todos.json)

Response: [JSON sample](response/put-todos.json)

TODO: authentication token missing

## Delete a task

Method: `DELETE`

Endpoint: `loft/{loft_id}/todos/{task_id}`

Params: None

Body: None

Response: [JSON sample](response/delete-todos.json)