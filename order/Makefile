up:
	kubectl apply -f ../deploy/order.yaml
down:
	kubectl delete -f ../deploy/order.yaml
clear:down
	docker image prune && docker rmi echomusic/order:v2.0
image:
	docker rmi echomusic/order:v2.0 && docker  build . -t  echomusic/order:v2.0



