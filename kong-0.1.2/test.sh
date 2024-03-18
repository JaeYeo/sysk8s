
helm install kong-02 chart/kong --namespace kong2 -f ./values.yaml --set fullnameOverride=kong-02
