#### Building

```bash
gvt fetch --branch master github.com/jroimartin/gocui
gvt fetch --revision v0.11.2 github.com/sirupsen/logrus
gvt fetch --revision v2.3.3 github.com/TykTechnologies/tyk

```

#### Kubernetes
```bash
# minikube only
eval $(minikube docker-env)

# influxdb
kubectl create configmap influxdb-conf --from-file=_kubernetes/influxdb/influxdb.conf
kubectl create -f _kubernetes/influxdb/influxdb-deployment.yaml
kubectl create -f _kubernetes/influxdb/influxdb-service.yaml
# minikube only
open http://`minikube ip`:30101

# telegraf
kubectl create configmap telegraf-conf --from-file=_kubernetes/telegraf/telegraf.conf
kubectl create -f _kubernetes/telegraf/telegraf-deployment.yaml
kubectl create -f _kubernetes/telegraf/telegraf-service.yaml

# grafana
kubectl create -f _kubernetes/grafana/grafana-deployment.yaml
kubectl create -f _kubernetes/grafana/grafana-service.yaml
# minikube only
open http://`minikube ip`:30102

# redis
kubectl create -f _kubernetes/redis/redis-deployment.yaml
kubectl create -f _kubernetes/redis/redis-service.yaml

# gateway
kubectl create configmap tyk-gateway-conf --from-file=_kubernetes/gateway/tyk.conf
kubectl create configmap tyk-gateway-apps --from-file=_kubernetes/gateway/apps/1.json
kubectl create -f _kubernetes/gateway/tyk-gateway-deployment.yaml
kubectl create -f _kubernetes/gateway/tyk-gateway-service.yaml

curl --header "x-tyk-authorization: ThisInNotTheSecretYouAreLookingFor" http://`minikube ip`:30002/tyk-api-test/get

# pump
kubectl create configmap tyk-pump-conf --from-file=_kubernetes/pump/pump.conf
kubectl create -f _kubernetes/pump/tyk-pump-deployment.yaml

```

#### docker-compose
