# Loft API design -> please follow json api standard

https://editor.swagger.io/
https://github.com/go-swagger/go-swagger


## Ground rules
- GET is stateless
- GET is body-less

## Things I need to figure out
- PUT vs PATCH: when to use which?
- 

1. create new loft
POST /loft
{
    "loft_name": "",
    "user_name": ""
}

200
{
    "lifetime_token": ""
}

2. request to join loft
POST /loft/{loft_id}/request
{
    "name": "",
    "message": ""
}

200


3. get notes
GET /loft/{loft_id}/notes

200
{
    "notes": [
        {

        }
    ]
}

4. add note
POST /loft/{loft_id}/notes
{

}

5. update note (the whole object)
PUT /loft/{loft_id}/notes/{note_id}
{

}

6. add task

7. delete task

8. edit task (the whole object)

9. edit task (patch)
PATCH /loft/{loft_id}/tasks/{task_id}
(authentication in header?)
{

}