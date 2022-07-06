k8s-mongo: k8s-mongo-pvc k8s-mongo-pv k8s-mongo-deployment k8s-mongo-service

k8s-mongo-pvc:
	kubectl apply -f .k8s/mongodb/pvc.yaml

k8s-mongo-pv:
	kubectl apply -f .k8s/mongodb/pv.yaml

k8s-mongo-deployment:
	kubectl apply -f .k8s/mongodb/deployment.yaml

k8s-mongo-service:
	kubectl apply -f .k8s/mongodb/service.yaml

k8s-app: k8s-app-deployment k8s-app-service

k8s-app-deployment: 
	kubectl apply -f .k8s/app/deployment.yaml

k8s-app-service: 
	kubectl apply -f .k8s/app/service.yaml

k8s-app-port-forward:
	kubectl port-forward service/team-management 8080:80

k8s-mongodb-port-forward:
	kubectl port-forward service/mongodb 27017:27017

unit-test:
	go test ./members/...
