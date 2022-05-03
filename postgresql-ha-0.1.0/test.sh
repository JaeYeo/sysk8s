
helm install postgresql-ha-01 chart/postgresql-ha -f ./plans/enterprise/values.yaml --namespace postgresql-ha --set postgresql.password='master77!!' --set postgresql.database=sysk8s --set service.port=5433 
