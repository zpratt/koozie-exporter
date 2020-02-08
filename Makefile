SHELL := /bin/bash -e
version := 1.0.0
host := 0.0.0.0
port := 8080
current_cluster := "$$(kind get clusters)"

clean:
	kind delete cluster
	docker rm $$(docker ps -a -q)
	docker rmi $$(docker images -f "reference=topokube:*" -q)
	rm -rf node_modules

docker:
	docker build . -t topokube:$(version)
	docker run -e "PORT=$(port)" -e "HOST=$(host)" -p $(port):$(port) topokube:$(version) &

inkind:
	docker build . -t topokube:$(version)
	@echo cluster is $(current_cluster)
	if [[ -z $$(kind get clusters) ]]; then kind create cluster --config cluster-config.yaml; fi
	kind load docker-image topokube:$(version)
	kubectl config set current-context kind-kind
	helmfile apply

testkind:
	curl -H "HOST:topokube.local" http://0.0.0.0:30080/

nodocker:
	npm i
	PORT=$(port) HOST=$(host) npm start