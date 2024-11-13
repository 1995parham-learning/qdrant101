# Qdrant 101

## Introduction

I am going to deploy [Qdrant](https://qdrant.tech/) and then working on its scaling.
I want to use it in LLM as a service architecture.

## Deploy

Deploy a single-node setup using docker:

```bash
docker pull qdrant/qdrant
docker run -p 6333:6333 -p 6334:6334 \
    -v $(pwd)/qdrant_storage:/qdrant/storage:z \
    qdrant/qdrant
```

Under the default configuration all data will be stored in the `./qdrant_storage` directory.
This will also be the only directory that both the Container and the host machine can both see.

Qdrant is now accessible:

    REST API: localhost:6333
    Web UI: localhost:6333/dashboard
    GRPC API: localhost:6334
