
helm install minio-02 chart/minio --namespace minio -f ./values.yaml --set fullnameOverride=minio-02
