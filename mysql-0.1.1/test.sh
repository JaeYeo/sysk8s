
helm install mysql-01 chart/mysql -f ./plans/micro/values.yaml --namespace mysql --set auth.rootPassword='master77!!' --set auth.database=sysk8s --set primary.service.ports.mysql=3307
