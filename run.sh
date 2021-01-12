docker network create showcase-indexer

docker stop server
docker stop worker1
docker stop worker2

docker rm server
docker rm worker1
docker rm worker2

docker build -t metis-showcase-indexer .

docker run -dit --name server -e host=server --network showcase-indexer -e mode=coordinator metis-showcase-indexer
docker run -dit --name worker1 -e host=server --network showcase-indexer -e mode=worker metis-showcase-indexer
docker run -dit --name worker2 -e host=server --network showcase-indexer -e mode=worker metis-showcase-indexer