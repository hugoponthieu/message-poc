# Permify setup

Post the schema at ./schema.perm on /v1/tenants/{tenant_id}/schemas/write :

```json
{
    "schema": "entity user {}\n\nentity server {\n    relation member @user\n}\n\nentity channel {\n    relation dmUser @user\n    relation server @server\n\n    action view = dmUser or server.member\n}"
}
```


Post the data at /v1/tenants/{tenant_id}/data/write

We will have a user that is called:
```json
{
    "metadata": {
        "schema_version": ""
    },
    "tuples": [
        {
            "entity": {
                "type": "channel",
                "id": "discussion"
            },
            "relation": "dmUser",
            "subject": {
                "type":"user",
                "id":"hugo"
            }

        }
    ]
}
```

Post data to perform checks on /v1/tenants/{tenant_id}/permissions/check:

```json
{
  "metadata": {
    "snap_token": "",
    "schema_version": "",
    "depth": "5"
  },
  "entity": {
    "type": "channel",
    "id": "discussion"
  },
  "permission": "view",
  "subject": {
    "type": "user",
    "id": "hug",
    "relation": ""
  }
}
```

