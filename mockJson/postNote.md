# POST /notes

## Request

`POST /photos HTTP/1.1`  
`Content-Type: application/vnd.api+json`  
`Accept: application/vnd.api+json`  

```json
{
    "data": {
        "type": "notes",
        "attributes": {
            "title": "",
            "description": ""
        }
    },
    "relationships": {
        "loft": {
            "data": {
                "type": "loft",
                "id": "loft id"
            }
        }
    }
}
```

## 200

`Content-Type: application/vnd.api+json`

```json
{
    "data": {
        "type": "notes",
        "id": "newly generated UUID",
        "attributes": {
            "title": "",
            "description": ""
        }
    },
    "relationships": {
        "loft": {
            "data": {
                "type": "loft",
                "id": "loft id"
            }
        }
    }
}
```

### Other possible generic responses

- `404_LOFT_NOT_FOUND`
- Any 5XX
- 415
- 406