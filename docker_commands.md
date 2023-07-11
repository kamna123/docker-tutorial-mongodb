# commands

## create docker network
docker network create mongo-network

## start mongodb and add to docker network
docker run -p 27017:27017 -d \
-e MONGO_INITDB_ROOT_USERNAME=mongoadmin \
-e MONGO_INITDB_ROOT_PASSWORD=secret \
--name mongodb --net mongo-network mongo

## start mongo-express and add to docker network
docker run -d \
-p 8081:8081 \
-e ME_CONFIG_MONGODB_ADMINUSERNAME=mongoadmin \
-e ME_CONFIG_MONGODB_ADMINPASSWORD=secret \
--net mongo-network \
--name mongo-express \
-e ME_CONFIG_MONGODB_SERVER=mongodb \
mongo-express
