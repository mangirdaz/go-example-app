MAKEFLAGS += --silent

#########################   docker build
go-build-api:
	echo "Build GO Single Binary" ; \
	echo "##########################" ;\
	echo "CGO_ENABLED=0 GOOS=linux go build -o api ." ; \
	echo "##########################" ;\
	cd api/ && CGO_ENABLED=0 GOOS=linux go build -o api . 

go-build-fe:
	echo "Build GO Single Binary" ; \
	echo "##########################" ;\
	echo "CGO_ENABLED=0 GOOS=linux go build -o fe ." ; \
	echo "##########################" ;\
	cd fe/ && CGO_ENABLED=0 GOOS=linux go build -o fe .

go-build: go-build-api go-build-fe
	
docker-build: go-build
	docker-compose build

docker-run: docker-build
	docker-compose up 	

########################### OCP builds

ocp-new-build-fe: go-build-fe
	echo "Create NEW EMPTY Build" ; \
	echo "##########################" ;\
	echo "oc new-build --name fe --binary"
	echo "##########################" ;\
	oc new-build --name fe --binary
	sleep 5 ; echo "" ;\
	make ocp-start-build-fe

ocp-start-build-fe: go-build-fe
	echo "Build FE Container" ; \
	echo "##########################" ;\
	echo "oc start-build fe --from-file=fe/ --follow"
	echo "##########################" ;\
	oc start-build fe --from-file=fe/ --follow
	echo "can now tag is desired oc tag fe:latest fe:x.y.z"

ocp-new-build-api:  go-build-api
	echo "Create NEW EMPTY Build" ; \
	echo "##########################" ;\
	echo "oc new-build --name api --binary"
	echo "##########################" ;\
	oc new-build --name api --binary
	sleep 5 ; echo "" ;\
	echo "Build API Container" ; \
	echo "##########################" ;\
	echo "oc start-build api --from-file=api/ --follow"
	echo "##########################" ;\
	oc start-build api --from-file=api/ --follow
	echo "can now tag as desired oc tag api:latest api:x.y.z"

ocp-start-build-api: go-build-api
	echo "Build API Container" ; \
	echo "##########################" ;\
	echo "oc start-build api --from-file=api/ --follow"
	echo "##########################" ;\
	oc start-build api --from-file=api/ --follow
	echo "can now tag as desired oc tag api:latest api:x.y.z"


ocp-build: ocp-build-api ocp-build-fe

ocp-run-api:
	echo "Run API image" ;\
	echo "##########################" ;\
	echo "oc new-app --name api --image-stream=api -e=API_IP=0.0.0.0 -e=API_PORT=8080" ;\
	echo "##########################" ;\
	oc new-app --name api --image-stream=api -e=API_IP=0.0.0.0 -e=API_PORT=8080
	#oc new-app --name logspout bekt/logspout-logstash:latest -e=ROUTE_URIS=logstash://logstash:5000 -e=DOCKER_LABELS=true
	#oc new-app --name logstash basi/logstash:v0.8.0 -e=ROUTE_URIS=logstash://logstash:5000 -e=DOCKER_LABELS=true

ocp-run-fe:
	echo "Run API image" ; \
	echo "##########################" ;\
	echo "oc new-app --name fe --image-stream=fe -e=FE_IP=0.0.0.0 -e=FE_PORT=8080 -e=API_SVC=http://api-test-project.beta-7.cor00005.cna.ukcloud.com" ;\
	echo "##########################" ;\
	oc new-app --name fe --image-stream=fe -e=FE_IP=0.0.0.0 -e=FE_PORT=8080 -e=API_SVC=http://api-myproject.127.0.0.1.nip.io

ocp-expose:
	echo "Expose services" ; \
	echo "##########################" ;\
	echo "oc expose service fe ; oc expose service api" ; \
	echo "##########################" ;\
	oc expose service fe ; \
	oc expose service api ; \
	oc get route

ocp-run: ocp-run-api ocp-run-fe ocp-expose

clean:
	oc delete all --all ;\
	docker rmi mangirdas/go-example-fe ;\
	docker rmi mangirdas/go-example-api