#/bin/bash

## -- tag it with the concierge name
## docker build -t concierge ../../
##
docker build -t quay.io/wework/kong-smash-service  ../../
docker run -p 8282:8282 quay.io/wework/kong-smash-service
##
## do not commit the executed (run) container, run creates a new layer. The side effect of creating anything after 
## the build is added as a new layer - so if you produced 100MB of logs as a result of the run, that would get
## pushed up and not the intended pure build.
##
#docker commit $(docker ps -l |grep concierge |awk '{print $1}') quay.io/wework/hackathon-2018-07:concierge
docker push quay.io/wework/kong-smash-service
##
kubectl patch deployment smash -p \
  "{\"spec\":{\"template\":{\"metadata\":{\"annotations\":{\"date\":\"`date +'%s'`\"}}}}}"
