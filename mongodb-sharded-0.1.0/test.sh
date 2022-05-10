
helm install mongodb-sharded-01 chart/mongodb-sharded -f ./plans/micro/values.yaml --namespace mongodb-sharded --set mongodbRootPassword='master77!!' --set service.port=27000
