TAG=v2

frontend:	main.go
		go get && go build -o $@

runfrontend:
		./frontend --port=:3000 --redis-master=:6379 --redis-slave=:6380

build:
		docker build -t satnam6502/frontend:$(TAG) .

push:
		docker push satnam6502/frontend:$(TAG)

redis-master:
		redis-server &

redis-slave:
		redis-server --port 6380 --slaveof localhost 6379 &

clean:
		docker rm -f redis-master
		docker rm -f redis-slave
