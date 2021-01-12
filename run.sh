#docker network create showcase-indexer

docker stop server
docker stop worker1
docker stop worker2

docker rm server
docker rm worker1
docker rm worker2

docker build -t metis-showcase-indexer .

docker run -dit --name server -e host=server  -p 8080:8080 --network showcase-indexer -e mode=coordinator metis-showcase-indexer
docker run -dit --name worker1 -e host=server --network showcase-indexer -e mode=worker metis-showcase-indexer
docker run -dit --name worker2 -e host=server --network showcase-indexer -e mode=worker metis-showcase-indexer

curl -d "https://github.com/DanielHauge/02148-cda-exercises" http://localhost:8080/add
curl -d "https://github.com/DanielHauge/metis-showcase-indexer" http://localhost:8080/add
curl -d "https://github.com/DanielHauge/metis-showcase" http://localhost:8080/add

echo "curl http://localhost:8080/log"
echo "curl http://localhost:8080/repo"
echo "curl http://localhost:8080/tasks"
echo "curl http://localhost:8080/status"