# message-poc

## Build

```sh
docker build . -t hugo08/message:your_version
```

Push to Docker Hub

```sh
docker push hugo08/message:your_version
```

## Deploy

Install helm chart:

```sh
helm upgrade --install message-poc ./deploy -n message
```

Uninstall helm chart:

```sh
helm uninstall message-poc -n message
```

## Mongo setup
Use the correct database:
```sh 
use message-db 
```

Create collection:
```sh 
db.createCollection("messages")
```

Create index:
```sh 
db.messages.createIndex({ content: "text" })
```

## Helpers link

https://www.mongodb.com/docs/manual/core/link-text-indexes/#std-label-text-search-on-premises
https://www.mongodb.com/resources/basics/full-text-search