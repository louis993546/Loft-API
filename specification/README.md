# Specification

This folder contains all the necessary details of the Loft REST API.

## Base rules

- Follow rules from [JSON API](http://jsonapi.org/) v1.0 Stable

## Sign in to a loft

Method: `POST`
Endpoint: `/loft/?`
Body: ?

## Sign out of a loft

Method: `POST`
Endpoint: `/loft/?`
Body: ?

## Get to-do list of a loft

Method: `POST`
Endpoint: `/loft/{loft_id}/todos`
```
{
    "data": [
        {
            "type": "todo",
            "id": "some uuid",
            "attributes": {
                "name": "blah blah blah blah"
            }
        }
    ]
}
```

## Add an item to to-do list of a loft

Method: `POST`
Endpoint: `/loft/{loft_id}/todos`
Body:
```
{
    "task": {
        "name": ""
    }
}
```

Response: TODO