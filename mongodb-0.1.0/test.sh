
helm install mongodb-01 chart/mongodb -f ./plans/micro/values.yaml --namespace mongodb --set auth.rootPassword='master77!!' --set auth.database=sysk8s --set service.port=27000
