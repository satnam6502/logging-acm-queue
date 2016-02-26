TAG = v1

client:	main.go
		go get && go build -o $@

build:
		docker build -t satnam6502/client:$(TAG) .

push:
		docker push satnam6502/client:$(TAG)
