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
	rm -f topokube
	go clean -cache
	go clean
	kind delete cluster --name koozie
	docker rmi -f $$(docker images -f "reference=topokube:*" -q) || echo "no matching images"
	docker rmi -f $$(docker images -f "reference=topokube-ui:*" -q) || echo "no matching images"
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
	kind load docker-image --name koozie topokube:$(version)
	kind load docker-image  --name koozie topokube-ui:$(version)
	kubectl config set current-context kind-koozie
	helmfile apply --skip-diff-on-install --include-needs

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

cause-deploy:
	kubectl run --restart=Never --rm -i -n default api-test --image=alpine -- /bin/ash -c "echo hello"
	curl -k https://topokube.local:30443/metrics | grep koozie

verify:
	golangci-lint run ./...
	go test -v -count=1 -cover ./...
