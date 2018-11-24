# Mock JSON

This is a temporary place to store json. Ultimately they need to go into the specification, but for now, this is easier to comprehand.

## General JSON API requirement

- `Content-Type: application/vnd.api+json` + no param for every requests and responses
  - 415 if it's there's any param
  - 406 for sth else (not sure what that means)



```
{
    "data": {}/[],      <- depends on the context of the endpoint (but I think for me I will just keep it [] at all time)     
    "error": [],
    "meta": {},
    "jsonapi": {},  <- optional
    "links": {},    <- optional
    "included": {}  <- optional, exists only if data exists
}
```

- Avoid responding with `200 + data: null/[]/{}`. It probably makes more sense for it to be `404`, but ultimately it depends on the context (e.g. search with no results is on the line)
- pagination lives in the `links` block [with 4 keywords](https://jsonapi.org/format/#fetching-pagination)
- POST
  - It's kinda interesting: it don't put all ids in the url, but in the `relationship` object in the body
  - [read this](https://jsonapi.org/format/#crud-creating)
- PATCH
  - It's gonna be interesting with Retrofit: that means I need to have a body with everything nullable + make sure moshi will remove those null from the body?

## Error code enum

To be use in `code` in [Error Objects](https://jsonapi.org/format/#errors). The general format is 

`{ERROR CODE}_{SHORT DESCRIPTION_SNAKE_CASE}`

Ultimately I think this should be for every single response, but for now this is mostly for errors with same HTTP status code, so that you can tell `404 notes not found` apart from `404 loft not found`.

These codes **SHOULD NEVER BE CHANGE** unless it's absolutely necessary (i.e. they get ossified the moment it got out), otherwise it can have huge impact to how client does error handling. Adding new code is allowed, but a long notification period must exists to ensure clients have the time to ensure they can understand and hahave accordingly

### 404 Not found

- `404_LOFT_NOT_FOUND`
- `404_NOTE_NOT_FOUND`
- `404_TASK_NOT_FOUND`
- `404_EVENT_NOT_FOUND`
- `404_MEMBER_NOT_FOUND`

### 403 Forbidden

- `403_FORBIDDEN` -> (need to rename) you don't have the permission to access/modify
- `403_CLIENT_GENERATED_LOFT_ID` -> (need to rename) you try to create loft with ID, which is just not how it works. It act as a reminder to the client that "you might be doing something completely wrong"
- `403_CLIENT_GENERATED_NOTE_ID` -> (need to rename)
- `403_CLIENT_GENERATED_TASK_ID` -> (need to rename)
- `403_CLIENT_GENERATED_EVENT_ID` -> (need to rename)
- `403_CLIENT_GENERATED_MEMBER_ID` -> (need to rename)
- (Other types of 403)

### 410 Gone

- `410_LOFT_GONE` -> 2 valid token requests try to delete loft together (is it even possible?)
- `410_NOTE_GONE` -> 2 valid token requests try to delete notes together
- `410_TASK_GONE`  -> ditto but for tasks
- `410_EVENT_GONE`  -> ditto but for events
- `410_MEMBER_GONE`  -> ditto but for members

### TBC (Stuff that I am not sure what the status code should be)

- `400_SORT_NOT_APPLICABLE` -> [as specified](https://jsonapi.org/format/#fetching-sorting)
- `401_UNAUTHORIZED`
- `405_METHOD_NOT_ALLOW`
- `406_NOT_ACCEPTABLE` -> server: I can't produce something that you understand
- `415_UNSUPPORTED_MEDIA_TYPE`  -> server: There is no way I can understand what you send
- `503_NO_DATABASE_CONNECTION`
- `504_MAINTANCE`   -> not sure about this, kinda too generic