# [MongoDB Enterprise with Docker](https://docs.mongodb.com/manual/tutorial/install-mongodb-enterprise-with-docker/)

Set Mongo version and your Docker username
```bash
export MONGODB_VERSION=5.0
export DOCKER_USERNAME=pedrokunz
```

To build the enterprise image
```bash
curl -O --remote-name-all https://raw.githubusercontent.com/docker-library/mongo/master/$MONGODB_VERSION/{Dockerfile,docker-entrypoint.sh}

chmod 755 ./docker-entrypoint.sh

docker build --build-arg MONGO_PACKAGE=mongodb-enterprise --build-arg MONGO_REPO=repo.mongodb.com -t $DOCKER_USERNAME/mongo-enterprise:$MONGODB_VERSION .

docker run --name mymongo -itd $DOCKER_USERNAME/mongo-enterprise:$MONGODB_VERSION

docker exec -it mymongo /usr/bin/mongo --eval "db.version()"
```