# Create docker ingress network, (Outcomment to create)
# docker network create showcase-indexer


# Stop and Remove all docker containers.
docker rm $(docker stop $(docker ps -aq))

# Rebuild the applicaiton
docker build -t metis-showcase-indexer .

# Run Elastic search
docker run -dit --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" --network showcase-indexer docker.elastic.co/elasticsearch/elasticsearch:7.10.1

## Give elasticsearch time to setup before workers try to connect with persistent tcp connection..
sleep 5

# Run 1 coordinator, 2 managers and 5 workers.
docker run -dit --name server -e host=server  -p 8080:8080 --network showcase-indexer -e mode=coordinator metis-showcase-indexer
docker run -dit --name manager1 -e host=server --network showcase-indexer -e mode=manager metis-showcase-indexer
docker run -dit --name manager2 -e host=server --network showcase-indexer -e mode=manager metis-showcase-indexer
docker run -dit --name worker1 -e host=server --network showcase-indexer -e ELASTICSEARCH_URL="http://elasticsearch:9200" -e mode=worker metis-showcase-indexer
docker run -dit --name worker2 -e host=server --network showcase-indexer -e ELASTICSEARCH_URL="http://elasticsearch:9200" -e mode=worker metis-showcase-indexer
docker run -dit --name worker3 -e host=server --network showcase-indexer -e ELASTICSEARCH_URL="http://elasticsearch:9200" -e mode=worker metis-showcase-indexer
docker run -dit --name worker4 -e host=server --network showcase-indexer -e ELASTICSEARCH_URL="http://elasticsearch:9200" -e mode=worker metis-showcase-indexer
docker run -dit --name worker5 -e host=server --network showcase-indexer -e ELASTICSEARCH_URL="http://elasticsearch:9200" -e mode=worker metis-showcase-indexer

# Adding a few repositories
curl -d "https://github.com/DanielHauge/02148-cda-exercises" http://localhost:8080/add
curl -d "https://github.com/DanielHauge/metis-showcase-indexer" http://localhost:8080/add
curl -d "https://github.com/DanielHauge/BlockLand" http://localhost:8080/add
curl -d "https://github.com/DanielHauge/metis-storage" http://localhost:8080/add
curl -d "https://github.com/DanielHauge/metis-showcase-api" http://localhost:8080/add
curl -d "https://github.com/DanielHauge/plex-folder-soldier" http://localhost:8080/add
curl -d "https://github.com/DanielHauge/Wowhub-Backend" http://localhost:8080/add
curl -d "https://github.com/DanielHauge/Wowhub-Frontend" http://localhost:8080/add
curl -d "https://github.com/DanielHauge/LanguageProject" http://localhost:8080/add

# Echo the control commands for easy copy paste.
echo "curl http://localhost:8080/log"
echo "curl http://localhost:8080/repo"
echo "curl http://localhost:8080/tasks"
echo "curl http://localhost:8080/status"
echo "docker rm \$(docker stop \$(docker ps -aq))"