build:
		go build -o ./bin/fibonacci ./main.go

start:
		sudo ./bin/fibonacci

docker-build:
		docker build -t fibonacci:latest -f Dockerfile .

docker-run:
		docker run --link memcached:cache -d -p 8080:8080 -p 3000:3000 --rm --name fibonacci fibonacci
docker-stop:
		docker stop fibonacci 
docker-cache:
		docker pull memcached
		docker run -d --rm --name memcached memcached

docker-cache-stop:
		docker stop memcached