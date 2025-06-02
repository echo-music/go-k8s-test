up:
	kubectl apply -f deploy/hello.yaml -f deploy/service.yaml -f deploy/ingress.yaml
down:
	kubectl delete -f deploy/hello.yaml -f deploy/service.yaml -f deploy/ingress.yaml


