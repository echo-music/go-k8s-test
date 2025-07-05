up:
	kubectl apply -f deploy/hello.yaml -f deploy/service.yaml
down:
	kubectl delete -f deploy/hello.yaml -f deploy/service.yaml


