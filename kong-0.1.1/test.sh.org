
helm install kong-01 chart/kong --namespace kong -f ./plans/micro/values.yaml --set service.exposeAdmin=true --set service.type=NodePort --set postgresql.auth.password='master77!!' --set postgresql.auth.postgresPassword='master77!!'
