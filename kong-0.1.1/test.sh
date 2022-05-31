
helm install kong-02 chart/kong --namespace kong2 -f ./plans/micro/values.yaml --set postgresql.auth.password='master77!!' --set postgresql.auth.postgresPassword='master77!!' --set adminPassword='master'
