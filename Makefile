SHELL := /bin/bash -e
version := 1.0.0
host := 0.0.0.0
port := 8080
current_cluster := "$$(kind get clusters)"

clean:
	rm -rf **/node_modules
	rm -rf **/dist
	rm -rf **/out
	rm -rf **/.cache
	kind delete cluster
	docker rmi $$(docker images -f "reference=topokube:*" -q)
	docker rmi $$(docker images -f "reference=topokube-ui:*" -q)
	docker system prune -f
docker:
	cd ui && \
	docker build . -t topokube-ui:$(version)
	docker build . -t topokube:$(version)
	docker run -e "PORT=$(port)" -e "HOST=$(host)" -p $(port):$(port) topokube:$(version) &

inkind:
	cd ui && \
	docker build . -t topokube-ui:$(version)
	docker build . -t topokube:$(version)
	if [[ -z $(current_cluster) ]]; then kind create cluster --config cluster-config.yaml; fi
	kind load docker-image topokube:$(version)
	kind load docker-image topokube-ui:$(version)
	kubectl config set current-context kind-kind
	helmfile apply

testkind:
	curl -H "HOST:topokube.local" http://0.0.0.0:30080/

nodocker:
	npm i
	PORT=$(port) HOST=$(host) npm start

ui:
	cd ui && \
	docker build . -t topokube-ui:$(version)

docker-ui:
	cd ui && \
	docker build . -t topokube-ui:$(version) && \
	docker run -p $(port):80 topokube-ui:$(version) &