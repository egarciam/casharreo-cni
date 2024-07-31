# Install multus
https://github.com/k8snetworkplumbingwg/multus-cni/blob/master/docs%2Fhow-to-use.md

## download cni plugins
Obtener cni plugins curl -OL https://github.com/containernetworking/plugins/releases/download/v1.5.1/cni-plugins-linux-amd64-v1.5.1.tgz


## Construir pod egress
docker build -f ./Dockerfile.alpine -t localhost:5000/egress-gateway:latest .

docker build -f ./Dockerfile.alpine-curl -t localhost:5000/alpine-curl:latest . 
docker push localhost:5000/alpine-curl:latest


## echo server original tkng labs makefile (ingress.mk row 48 target egress prep )
```docker run -d --rm --network kind --name echo mpolden/echoip -l ":80" ``` <br/>
**Usando --network=kind por el k8s**

## Run echo server 2 in docker (copiado de )
```docker run --name echo-2 -d --net=kind ghcr.io/spidernet-io/egressgateway-nettools:latest /usr/bin/nettools-server -protocol web -webPort 8080``` <br/>
**Usando --network=kind por el k8s**

## echo server 3 desde (egress gateway project)
``` docker run -d --name echo-3 --rm --net=kind registry.k8s.io/echoserver:1.4 ``` <br/>
**Usando --network=kind por el k8s**

Invocar desde el sample pod curl echo-3:8080

