up:
	kubectl apply -f deploy/hello.yaml -f deploy/service.yaml
down:
	kubectl delete -f deploy/hello.yaml -f deploy/service.yaml
clear:down
	docker image prune && docker rmi echomusic/hello:v2.0




