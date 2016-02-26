

frontend-up:
		kubectl create -f frontend-controller.yaml

frontend-down:
		kubectl delete rc frontend

refront:	frontend-down frontend-up

frontend-service-up:
		kubectl create -f frontend-service.yaml

frontend-service-down:
		kubectl delete service frontend

redis-up:
		kubectl create -f redis-master-controller.yaml
		kubectl create -f redis-master-service.yaml
		kubectl create -f redis-slave-controller.yaml
		kubectl create -f redis-slave-service.yaml

redis-down:
		-kubectl delete rc redis-master
		-kubectl delete service redis-master
		-kubectl delete rc redis-slave
		-kubectl delete service redis-slave

client-up:
		kubectl create -f client-pod.yaml
client-down:
		kubectl delete pod satnam-client

get:
		kubectl get pods,rc,service
