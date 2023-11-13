# **SysMasterK8sVelero(Velero Web) Helm**

## Introduction
백업관리 서비스 [Velero](https://velero.io/)를 웹 서비스 형태로 제공하기 위한 Velero 웹 서비스입니다.<br> velero-api, velero-app으로 2개의 pod로 구성되어 있으며 helm을 통해 동시배포를 제공합니다.<br>

## Prerequisites
-   Helm 3.0+
-   Kubernetes v1.20.15
-   Kustomize v4.5.4
-   Client: v1.24.2

## Installing the Chart
차트를 설치하기 위해 릴리즈 이름 `veleros`, 네임스페이스 이름 `velero-service`라고 하면 명령어는 다음과 같다:

```bash
$ helm install veleros -n velero-service --create-namespace . -f values.yaml
```

## Uninstalling the Chart
`velero-service` 네임스페이스, `veleros`를 삭제:

```bash
$ helm uninstall veleros -n velero-service
```

## Configuration

SysMasterK8sVelero 차트의 적용되는 values.yaml 설정목록

| Parameter                                  | Description                               | Default                            |
| ------------------------------------------ | ----------------------------------------- | ---------------------------------- |
| `veleroApi.image.repository`                         | Docker image repo                         | `harbor.spaasta.com/velero/velero-api`|
| `veleroApi.image.tag`                                | Docker image                              | `1.5`                    |
| `veleroApi.image.pullPolicy`                         | Docker image pull policy                  | `IfNotPresent`                     |
| `veleroApi.imagePullSecrets`                        | Docker image pull secrets                 | `[]`                               |
| `veleroApi.service.type`                             | Service type                              | `ClusterIP`                        |
| `veleroApi.service.port`                             | Service port of veleroApi| `80`                               |
| `veleroApi.resources`                                | Resource limits for the pod               | `{}`                               |
| `veleroApi.ingress.enabled`                          | Ingress enabled                           | `true`                            |
| `veleroApi.ingress.hosts`                            | Ingress hosts                             | `velero.saas.sysmasterk8s.com`                               |
| `veleroApi.ingress.tls`                              | Ingress TLS                               | `[]`                               |
| `veleroApi.nodeSelector`                             | NodeSelector                              | `{}`                               |
| `veleroApi.tolerations`                              | Tolerations                               | `[]`                               |
| `veleroApi.affinity`                                 | Affinity                                  | `{}`                               |
| `veleroApi.secret.spring_datasource_url`                                 | 플랫폼 DB접속 URL                                  | `jdbc:mysql://IP:3306/DBNAME?characterEncoding=UTF-8&autoReconnect=true&failOverReadOnly=false&maxReconnects=10`                               |
| `veleroApi.secret.spring_datasource_username`                                 | 플랫폼 DB접속 아이디                                  | `None`                               |
| `veleroApi.secret.spring_datasource_password`                                 | 플랫폼 DB접속 비밀번호                                  | `None`                               |
| `veleroApi.secret.server_port`                                 | API서버포트                                  | `None`                               |
| `veleroApi.secret.server_servlet_context_path`                                 | API configPath                                  | `80`                               |
| `veleroApi.secret.spring_cloud_vault_k8s_configpath`                                 | kubeconfig 저장경로                                   | `/api`                               |
| `veleroApi.secret.spring_cloud_vault_host`                                 | 암호키저장소 접속URL                                  | `vault.saas.sysmasterk8s.com`                               |
| `veleroApi.secret.spring_cloud_vault_port`                                 | 암호키저장소 접속포트                                  | `80`                               |
| `veleroApi.secret.spring_cloud_vault_token`                                 | 암호키저장소 토큰                                  | `None`                               |
| `veleroApi.secret.spring_cloud_vault_k8s_projectid`                                 | project_id 정보                                  | `None`                               |
| `veleroApp.configmap.authentication_url`                                 | 통합인증 접속URL                                  | `https://keycloak2.spaasta.com/auth`                               |
| `veleroApp.configmap.authentication_client_id`                                 | 통합인증 Client Id                                  | `velero`                               |
| `veleroApp.configmap.authentication_realm`                                 | 통합인증 realm정보                                  | `sysK8s`                               |

각각의 특정 파라미터를 `--set key=value[,key=value]`를 통해 사용할 수 있다. <br>
또는 차트를 설치하는 동안 매개변수 값을 지정하는 YAML 파일을 제공할 수 있다.

```bash
$ helm install veleros -n velero-service . -f values.yaml \
  --create-namespace \
  --set veleroApi.ingress.enabled=false
```


