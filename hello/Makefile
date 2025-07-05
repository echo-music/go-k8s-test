up:
	kubectl apply -f ../deploy/hello.yaml
down:
	kubectl delete -f ../deploy/hello.yaml
clear:down
	docker image prune && docker rmi echomusic/hello:v2.0
image:
	docker rmi echomusic/hello:v2.0 && docker  build . -t  echomusic/hello:v2.0



