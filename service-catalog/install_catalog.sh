helm install catalog charts/catalog \
	--namespace catalog --set asyncBindingOperationsEnabled=true \
	--set 'podLabels.platform=platform_service' \
	--set 'podLabels.taskse=service_catalog' \
	--set 'podLabels.taskcl=backend_service' \
	--wait --create-namespace
