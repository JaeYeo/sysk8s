
helm install mariadb-galera-03 chart/mariadb-galera -f ./plans/micro/values.yaml --namespace mariadb-galera --set rootUser.password='master77!!' --set db.name=sysk8s
