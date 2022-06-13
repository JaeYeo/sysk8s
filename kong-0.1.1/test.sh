
helm install kong-02 chart/kong --namespace kong2 -f ./plans/micro/values.yaml --set fullnameOverride=kong-02
