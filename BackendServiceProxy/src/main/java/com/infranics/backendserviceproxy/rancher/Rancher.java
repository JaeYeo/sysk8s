package com.infranics.backendserviceproxy.rancher;

import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
import java.util.Base64;

import lombok.*;
import org.apache.http.conn.ssl.TrustStrategy;
import org.apache.http.impl.client.HttpClientBuilder;
import org.apache.http.ssl.SSLContexts;

import org.json.JSONObject;


import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.http.client.HttpComponentsClientHttpRequestFactory;
import org.springframework.web.client.RestTemplate;


/**
 * @author : se.ban
 * 2022. 3. 4.
 * 
 */
@NoArgsConstructor
@AllArgsConstructor
@Builder
@Getter
@Setter
@ToString
public class Rancher {

	private final Logger logger = LoggerFactory.getLogger(this.getClass());

	// endpoint to access rancher
	private String endpoint;

	// bearer-token == access-key:secret-key
	private String auth;

	// 임시 clusterId, 추후는 직접 만들어서 획득
	private String tempClusterId;


	///// rancher member function /////

	public String GetKubeConfig()
	{
		try {

			HttpHeaders headers = createHttpHeaders(this.auth);

			RestTemplate restTemplate = CreateRestTemplate(this.auth);

			String kubeconfig_url = this.getEndpoint() + "/clusters/" + this.getTempClusterId() + "?action=generateKubeconfig";

			// logger.info("get kubeconfig:" + kubeconfig_url);

			String kube_config = "invalid kubeconfig";
			{
				HttpEntity<String> entity = new HttpEntity<String>(headers);

				ResponseEntity<String> response = restTemplate.exchange(kubeconfig_url, HttpMethod.POST, entity, String.class);
				// logger.info("create_cluster Result - status (" + response.getStatusCode() + ") has body: " + response.hasBody());
				// logger.info("create_cluster Response Body:" + response.getBody());


				JSONObject cluster_json = new JSONObject(response.getBody());

				kube_config = cluster_json.getString("config");

				// logger.info("kubeconfig :" + kube_config);

			}
			return kube_config;
		}
		catch (Exception e)
		{
			return null;
		}


	}


	static public String makeClusterBrokerJson(String addonName, String addonUrl)
	{
		/* ----------------
		apiVersion: addons.kyma-project.io/v1alpha1
		kind: ClusterAddonsConfiguration
		metadata:
			name: redis-cfg
		spec:
			repositories:
				- url: "https://raw.githubusercontent.com/infranics/sysk8s/main/redis-0.1.0/index-redis.yaml"
		--------------------*/

		String rawJsonCustomResourceObj =
				"{\"apiVersion\":\"addons.kyma-project.io/v1alpha1\"," +
				 "\"kind\":\"ClusterAddonsConfiguration\"," +
				 "\"metadata\": " +
					"{\"name\": \""+addonName+"\"}," +
				 "\"spec\": " +
					"{\"repositories\": [{\"url\": \""+addonUrl+"\"}]" +
					"}" +
				"}";


		return rawJsonCustomResourceObj;
	}

	static public String makeNamespaceBrokerJson(String addonName, String namespace, String addonUrl)
	{
		/* ----------------
		apiVersion: addons.kyma-project.io/v1alpha1
		kind: AddonsConfiguration
		metadata:
			name: redis-cfg
			namespace: redis
		spec:
			repositories:
				- url: "https://raw.githubusercontent.com/infranics/sysk8s/main/redis-0.1.0/index-redis.yaml"
		--------------------*/

		String rawJsonCustomResourceObj =
				"{\"apiVersion\":\"addons.kyma-project.io/v1alpha1\"," +
						"\"kind\":\"AddonsConfiguration\"," +
						"\"metadata\": " +
							"{\"name\": \""+addonName+"\"," +
							 "\"namespace\": \""+namespace+"\"" +
							"}," +
						"\"spec\": " +
							"{\"repositories\": [{\"url\": \""+addonUrl+"\"}]" +
							"}" +
				"}";


		return rawJsonCustomResourceObj;
	}


	static public String makeProvisioningWithClusterBrokerJson(String serviceName,
															   String namespace,
															   String clusterClassName,
															   String clusterPlanName,
															   String[] parameters)
	{
		/* ----------------
		apiVersion: servicecatalog.k8s.io/v1beta1
		kind: ServiceInstance
		metadata:
		  name: redis-01
		  namespace: redis
		spec:
		  clusterServiceClassExternalName: redis
		  clusterServicePlanExternalName: micro
		  parameters:
			fullnameOverride: redis-01
		--------------------*/

		String rawJsonCustomResourceObj =
				"{\"apiVersion\":\"servicecatalog.k8s.io/v1beta1\"," +
						"\"kind\":\"ServiceInstance\"," +
						"\"metadata\": " +
						"{\"name\": \""+serviceName+"\"," +
						" \"namespace\": \""+namespace+"\"}," +
						"\"spec\": " +
						"{\"clusterServiceClassExternalName\": \""+clusterClassName+"\"," +
						" \"clusterServicePlanExternalName\": \""+clusterPlanName+"\"," +
						" \"parameters\": { \"fullnameOverride\": \""+parameters[0]+"\" }" +
						"}" +
				"}";


		return rawJsonCustomResourceObj;
	}


	static public String makeProvisioningRedisWithClusterBrokerJson(String serviceName,
															   	String namespace,
															   	String clusterClassName,
															   	String clusterPlanName,
															   	String password,
																String redisPort)
	{
		/* ----------------
		apiVersion: servicecatalog.k8s.io/v1beta1
		kind: ServiceInstance
		metadata:
		  name: redis-03
		  namespace: redis01
		spec:
		  clusterServiceClassExternalName: redis
		  clusterServicePlanExternalName: micro
		  parameters:
			fullnameOverride: redis-03
			global:
			  redis:
				password: master77!!
			master:
			  service:
				ports:
				  redis: 6380


		--------------------*/

		String rawJsonCustomResourceObj =
				"{\"apiVersion\":\"servicecatalog.k8s.io/v1beta1\"," +
						"\"kind\":\"ServiceInstance\"," +
						"\"metadata\": " +
						"{\"name\": \""+serviceName+"\"," +
						" \"namespace\": \""+namespace+"\"}," +
						"\"spec\": " +
						"{\"clusterServiceClassExternalName\": \""+clusterClassName+"\"," +
						" \"clusterServicePlanExternalName\": \""+clusterPlanName+"\"," +
						" \"parameters\": { \"fullnameOverride\": \""+serviceName+"\", " +
											"\"global\": { \"redis\": { \"password\": \""+password+"\"}}," +
											"\"master\": { \"service\": { \"ports\": { \"redis\": \""+redisPort+"\"}}}" +
										"}" +
						"}" +
				"}";

		return rawJsonCustomResourceObj;
	}


	static public String makeProvisioningRedisWithNamespaceBrokerJson(String serviceName,
																	String namespace,
																	String namespaceClassName,
																	String namespacePlanName,
																	String password,
																	String redisPort)
	{
		/* ----------------
		apiVersion: servicecatalog.k8s.io/v1beta1
		kind: ServiceInstance
		metadata:
		  name: redis-03
		  namespace: redis01
		spec:
		  serviceClassExternalName: redis
		  servicePlanExternalName: micro
		  parameters:
			fullnameOverride: redis-03
			global:
			  redis:
				password: master77!!
			master:
			  service:
				ports:
				  redis: 6380
		--------------------*/

		String rawJsonCustomResourceObj =
				"{\"apiVersion\":\"servicecatalog.k8s.io/v1beta1\"," +
						"\"kind\":\"ServiceInstance\"," +
						"\"metadata\": " +
						"{\"name\": \""+serviceName+"\"," +
						" \"namespace\": \""+namespace+"\"}," +
						"\"spec\": " +
						"{\"serviceClassExternalName\": \""+namespaceClassName+"\"," +
						" \"servicePlanExternalName\": \""+namespacePlanName+"\"," +
						" \"parameters\": { \"fullnameOverride\": \""+serviceName+"\", " +
											"\"global\": { \"redis\": { \"password\": \""+password+"\"}}," +
										    "\"master\": { \"service\": { \"ports\": { \"redis\": \""+redisPort+"\"}}}" +
										"}" +
						"}" +
				"}";

		return rawJsonCustomResourceObj;
	}



	static public String makeProvisioningMariadbWithClusterBrokerJson(String serviceName,
																	String namespace,
																	String namespaceClassName,
																	String namespacePlanName,
																	String password,
																	String database,
																	String port)
	{
		/* ----------------
		apiVersion: servicecatalog.k8s.io/v1beta1
		kind: ServiceInstance
		metadata:
		  name: mariadb-01
		  namespace: mariadb
		spec:
		  clusterServiceClassExternalName: mariadb
		  clusterServicePlanExternalName: micro
		  parameters:
			fullnameOverride: mariadb-01
			auth:
			  rootPassword: master77!!
			  database: sysk8s
			primary:
			  service:
				ports:
				  mysql: 3308
		--------------------*/

		String rawJsonCustomResourceObj =
				"{\"apiVersion\":\"servicecatalog.k8s.io/v1beta1\"," +
						"\"kind\":\"ServiceInstance\"," +
						"\"metadata\": " +
						"{\"name\": \""+serviceName+"\"," +
						" \"namespace\": \""+namespace+"\"}," +
						"\"spec\": " +
						"{\"clusterServiceClassExternalName\": \""+namespaceClassName+"\"," +
						" \"clusterServicePlanExternalName\": \""+namespacePlanName+"\"," +
						" \"parameters\": { \"fullnameOverride\": \""+serviceName+"\", " +
											"\"auth\": { \"rootPassword\": \""+password+"\",\"database\": \""+database+"\"}," +
										    "\"primary\": { \"service\": { \"ports\": { \"mysql\": \""+port+"\"}}}" +
										"}" +
						"}" +
				"}";

		return rawJsonCustomResourceObj;
	}



	static public String makeProvisioningMariadbWithNamespaceBrokerJson(String serviceName,
																	String namespace,
																	String namespaceClassName,
																	String namespacePlanName,
																	String password,
																	String database,
																	String port)
	{
		/* ----------------
		apiVersion: servicecatalog.k8s.io/v1beta1
		kind: ServiceInstance
		metadata:
		  name: mariadb-01
		  namespace: mariadb
		spec:
		  serviceClassExternalName: mariadb
		  servicePlanExternalName: micro
		  parameters:
			fullnameOverride: mariadb-01
			auth:
			  rootPassword: master77!!
			  database: sysk8s
			primary:
			  service:
				ports:
				  mysql: 3308
		--------------------*/

		String rawJsonCustomResourceObj =
				"{\"apiVersion\":\"servicecatalog.k8s.io/v1beta1\"," +
						"\"kind\":\"ServiceInstance\"," +
						"\"metadata\": " +
						"{\"name\": \""+serviceName+"\"," +
						" \"namespace\": \""+namespace+"\"}," +
						"\"spec\": " +
						"{\"serviceClassExternalName\": \""+namespaceClassName+"\"," +
						" \"servicePlanExternalName\": \""+namespacePlanName+"\"," +
						" \"parameters\": { \"fullnameOverride\": \""+serviceName+"\", " +
											"\"auth\": { \"rootPassword\": \""+password+"\",\"database\": \""+database+"\"}," +
										    "\"primary\": { \"service\": { \"ports\": { \"mysql\": \""+port+"\"}}}" +
										"}" +
						"}" +
				"}";

		return rawJsonCustomResourceObj;
	}

	static public String makeProvisioningPostgresqlWithClusterBrokerJson(String serviceName,
																	String namespace,
																	String namespaceClassName,
																	String namespacePlanName,
																	String password,
																	String database,
																	String port)
	{
		/* ----------------
		apiVersion: servicecatalog.k8s.io/v1beta1
		kind: ServiceInstance
		metadata:
		  name: postgresql-01
		  namespace: postgresql
		spec:
		  clusterServiceClassExternalName: postgresql
		  clusterServicePlanExternalName: micro
		  parameters:
			fullnameOverride: postgresql-01
			auth:
			  postgresPassword: master77!!
			  database: sysk8s
			global:
			  postgresql:
				service:
				  ports:
					postgresql: 5433

		--------------------*/

		String rawJsonCustomResourceObj =
				"{\"apiVersion\":\"servicecatalog.k8s.io/v1beta1\"," +
						"\"kind\":\"ServiceInstance\"," +
						"\"metadata\": " +
						"{\"name\": \""+serviceName+"\"," +
						" \"namespace\": \""+namespace+"\"}," +
						"\"spec\": " +
						"{\"clusterServiceClassExternalName\": \""+namespaceClassName+"\"," +
						" \"clusterServicePlanExternalName\": \""+namespacePlanName+"\"," +
						" \"parameters\": { \"fullnameOverride\": \""+serviceName+"\", " +
											"\"auth\": { \"postgresPassword\": \""+password+"\",\"database\": \""+database+"\"}," +
										    "\"global\": { \"postgresql\": {\"service\": { \"ports\": { \"postgresql\": \""+port+"\"}}}}" +
										"}" +
						"}" +
				"}";

		return rawJsonCustomResourceObj;
	}




	static public String makeProvisioningPostgresqlWithNamespaceBrokerJson(String serviceName,
																	String namespace,
																	String namespaceClassName,
																	String namespacePlanName,
																	String password,
																	String database,
																	String port)
	{
		/* ----------------
		apiVersion: servicecatalog.k8s.io/v1beta1
		kind: ServiceInstance
		metadata:
		  name: postgresql-01
		  namespace: postgresql
		spec:
		  serviceClassExternalName: postgresql
		  servicePlanExternalName: micro
		  parameters:
			fullnameOverride: postgresql-01
			auth:
			  postgresPassword: master77!!
			  database: sysk8s
			global:
			  postgresql:
				service:
				  ports:
					postgresql: 5433

		--------------------*/

		String rawJsonCustomResourceObj =
				"{\"apiVersion\":\"servicecatalog.k8s.io/v1beta1\"," +
						"\"kind\":\"ServiceInstance\"," +
						"\"metadata\": " +
						"{\"name\": \""+serviceName+"\"," +
						" \"namespace\": \""+namespace+"\"}," +
						"\"spec\": " +
						"{\"serviceClassExternalName\": \""+namespaceClassName+"\"," +
						" \"servicePlanExternalName\": \""+namespacePlanName+"\"," +
						" \"parameters\": { \"fullnameOverride\": \""+serviceName+"\", " +
											"\"auth\": { \"postgresPassword\": \""+password+"\",\"database\": \""+database+"\"}," +
										    "\"global\": { \"postgresql\": {\"service\": { \"ports\": { \"postgresql\": \""+port+"\"}}}}" +
										"}" +
						"}" +
				"}";

		return rawJsonCustomResourceObj;
	}



	static public String makeProvisioningRabbitmqWithClusterBrokerJson(String serviceName,
																	String namespace,
																	String namespaceClassName,
																	String namespacePlanName,
																	String password,
																	String port,
																    String ingressHostname)
	{
		/* ----------------
		apiVersion: servicecatalog.k8s.io/v1beta1
		kind: ServiceInstance
		metadata:
		  name: rabbitmq-01
		  namespace: rabbitmq
		spec:
		  clusterServiceClassExternalName: rabbitmq
		  clusterServicePlanExternalName: micro
		  parameters:
			fullnameOverride: rabbitmq-01
			auth:
			  password: master77!!
			service:
			  type: NodePort
			  port: 5673
			ingress:
			  enabled: true
			  hostname: rabbitmq.spaasta.com

		--------------------*/

		String rawJsonCustomResourceObj =
				"{\"apiVersion\":\"servicecatalog.k8s.io/v1beta1\"," +
						"\"kind\":\"ServiceInstance\"," +
						"\"metadata\": " +
						"{\"name\": \""+serviceName+"\"," +
						" \"namespace\": \""+namespace+"\"}," +
						"\"spec\": " +
						"{\"clusterServiceClassExternalName\": \""+namespaceClassName+"\"," +
						" \"clusterServicePlanExternalName\": \""+namespacePlanName+"\"," +
						" \"parameters\": { \"fullnameOverride\": \""+serviceName+"\", " +
											"\"auth\": { \"password\": \""+password+"\"}," +
										    "\"service\": { \"type\": \"NodePort\", \"port\": \""+port+"\"}," +
											"\"ingress\": { \"enabled\": \"true\", \"hostname\": \""+ingressHostname+"\"}"+
										"}" +
						"}" +
				"}";

		return rawJsonCustomResourceObj;
	}

	static public String makeProvisioningRabbitmqWithNamespaceBrokerJson(String serviceName,
																		 String namespace,
																		 String namespaceClassName,
																		 String namespacePlanName,
																		 String password,
																		 String port,
																		 String ingressHostname)
	{
		/* ----------------
		apiVersion: servicecatalog.k8s.io/v1beta1
		kind: ServiceInstance
		metadata:
		  name: rabbitmq-01
		  namespace: rabbitmq
		spec:
		  serviceClassExternalName: rabbitmq
		  servicePlanExternalName: micro
		  parameters:
			fullnameOverride: rabbitmq-01
			auth:
			  password: master77!!
			service:
			  type: NodePort
			  port: 5673
			ingress:
			  enabled: true
			  hostname: rabbitmq.spaasta.com
		--------------------*/

		String rawJsonCustomResourceObj =
				"{\"apiVersion\":\"servicecatalog.k8s.io/v1beta1\"," +
						"\"kind\":\"ServiceInstance\"," +
						"\"metadata\": " +
						"{\"name\": \""+serviceName+"\"," +
						" \"namespace\": \""+namespace+"\"}," +
						"\"spec\": " +
						"{\"serviceClassExternalName\": \""+namespaceClassName+"\"," +
						" \"servicePlanExternalName\": \""+namespacePlanName+"\"," +
						" \"parameters\": { \"fullnameOverride\": \""+serviceName+"\", " +
											"\"auth\": { \"password\": \""+password+"\"}," +
											"\"service\": { \"type\": \"NodePort\", \"port\": \""+port+"\"}," +
											"\"ingress\": { \"enabled\": \"true\", \"hostname\": \""+ingressHostname+"\"}"+
										"}" +
						"}" +
				"}";

		return rawJsonCustomResourceObj;
	}

	static public String makeProvisioningJenkinsWithClusterBrokerJson(String serviceName,
																		String namespace,
																		String namespaceClassName,
																		String namespacePlanName,
																		String password,
																		String ingressHostname)
	{
		/* ----------------
		apiVersion: servicecatalog.k8s.io/v1beta1
		kind: ServiceInstance
		metadata:
		  name: jenkins-02
		  namespace: jenkins
		spec:
		  clusterServiceClassExternalName: jenkins
		  clusterServicePlanExternalName: micro
		  parameters:
			fullnameOverride: jenkins-02
			jenkinsPassword: master77!!
			ingress:
			  enabled: true
			  hostname: jenkins2.spaasta.com

		--------------------*/

		String rawJsonCustomResourceObj =
				"{\"apiVersion\":\"servicecatalog.k8s.io/v1beta1\"," +
						"\"kind\":\"ServiceInstance\"," +
						"\"metadata\": " +
						"{\"name\": \""+serviceName+"\"," +
						" \"namespace\": \""+namespace+"\"}," +
						"\"spec\": " +
						"{\"clusterServiceClassExternalName\": \""+namespaceClassName+"\"," +
						" \"clusterServicePlanExternalName\": \""+namespacePlanName+"\"," +
						" \"parameters\": { \"fullnameOverride\": \""+serviceName+"\", " +
											"\"jenkinsPassword\": \""+password+"\"," +
											"\"ingress\": { \"enabled\": \"true\", \"hostname\": \""+ingressHostname+"\"}"+
										"}" +
						"}" +
				"}";

		return rawJsonCustomResourceObj;
	}




	static public String makeProvisioningJenkinsWithNamespaceBrokerJson(String serviceName,
																		 String namespace,
																		 String namespaceClassName,
																		 String namespacePlanName,
																		 String password,
																		 String ingressHostname)
	{
		/* ----------------
		apiVersion: servicecatalog.k8s.io/v1beta1
		kind: ServiceInstance
		metadata:
		  name: jenkins-02
		  namespace: jenkins
		spec:
		  serviceClassExternalName: jenkins
		  servicePlanExternalName: micro
		  parameters:
			fullnameOverride: jenkins-02
			jenkinsPassword: master77!!
			ingress:
			  enabled: true
			  hostname: jenkins2.spaasta.com

		--------------------*/

		String rawJsonCustomResourceObj =
				"{\"apiVersion\":\"servicecatalog.k8s.io/v1beta1\"," +
						"\"kind\":\"ServiceInstance\"," +
						"\"metadata\": " +
						"{\"name\": \""+serviceName+"\"," +
						" \"namespace\": \""+namespace+"\"}," +
						"\"spec\": " +
						"{\"serviceClassExternalName\": \""+namespaceClassName+"\"," +
						" \"servicePlanExternalName\": \""+namespacePlanName+"\"," +
						" \"parameters\": { \"fullnameOverride\": \""+serviceName+"\", " +
											"\"jenkinsPassword\": \""+password+"\"," +
											"\"ingress\": { \"enabled\": \"true\", \"hostname\": \""+ingressHostname+"\"}"+
										"}" +
						"}" +
				"}";

		return rawJsonCustomResourceObj;
	}


	static public String makeProvisioningKafkaWithClusterBrokerJson(String serviceName,
																	  String namespace,
																	  String namespaceClassName,
																	  String namespacePlanName)

	{
		/* ----------------
		apiVersion: servicecatalog.k8s.io/v1beta1
		kind: ServiceInstance
		metadata:
		  name: kafka-03
		  namespace: kafka
		spec:
		  clusterServiceClassExternalName: kafka
		  clusterServicePlanExternalName: micro
		  parameters:
			fullnameOverride: kafka-03



		--------------------*/

		String rawJsonCustomResourceObj =
				"{\"apiVersion\":\"servicecatalog.k8s.io/v1beta1\"," +
						"\"kind\":\"ServiceInstance\"," +
						"\"metadata\": " +
						"{\"name\": \""+serviceName+"\"," +
						" \"namespace\": \""+namespace+"\"}," +
						"\"spec\": " +
						"{\"clusterServiceClassExternalName\": \""+namespaceClassName+"\"," +
						" \"clusterServicePlanExternalName\": \""+namespacePlanName+"\"," +
						" \"parameters\": { \"fullnameOverride\": \""+serviceName+"\" " +
						"}" +
				"}";

		return rawJsonCustomResourceObj;
	}



	static public String makeProvisioningKafkaWithNamespaceBrokerJson(String serviceName,
																		String namespace,
																		String namespaceClassName,
																		String namespacePlanName)
	{
		/* ----------------
		apiVersion: servicecatalog.k8s.io/v1beta1
		kind: ServiceInstance
		metadata:
		  name: kafka-03
		  namespace: kafka
		spec:
		  serviceClassExternalName: kafka
		  servicePlanExternalName: micro
		  parameters:
			fullnameOverride: kafka-03



		--------------------*/

		String rawJsonCustomResourceObj =
				"{\"apiVersion\":\"servicecatalog.k8s.io/v1beta1\"," +
						"\"kind\":\"ServiceInstance\"," +
						"\"metadata\": " +
						"{\"name\": \""+serviceName+"\"," +
						" \"namespace\": \""+namespace+"\"}," +
						"\"spec\": " +
						"{\"serviceClassExternalName\": \""+namespaceClassName+"\"," +
						" \"servicePlanExternalName\": \""+namespacePlanName+"\"," +
						" \"parameters\": { \"fullnameOverride\": \""+serviceName+"\" " +
						"}" +
				"}";

		return rawJsonCustomResourceObj;
	}


	// bindings
	static public String makeBindingJson(String serviceName,
									  String namespace)
	{
		/* ----------------
		apiVersion: servicecatalog.k8s.io/v1beta1
		kind: ServiceBinding
		metadata:
		  name: redis-03-binding
		  namespace: redis
		spec:
		  instanceRef:
			name: redis-03
		--------------------*/

		String rawJsonCustomResourceObj =
				"{\"apiVersion\":\"servicecatalog.k8s.io/v1beta1\"," +
						"\"kind\":\"ServiceBinding\"," +
						"\"metadata\": " +
							"{\"name\": \""+serviceName+"-binding\"," +
							" \"namespace\": \""+namespace+"\" }," +
						"\"spec\": " +
							"{ \"instanceRef\": { \"name\": \""+serviceName+"\" }}" +
				"}";

		return rawJsonCustomResourceObj;
	}



	// util function
	static private RestTemplate CreateRestTemplate(String auth)
	{
		try
		{
			RestTemplate restTemplate = new RestTemplate();

			TrustStrategy acceptingTrustStrategy = new TrustStrategy() {
				@Override
				public boolean isTrusted(X509Certificate[] x509Certificates, String s) throws CertificateException {
					return true;
				}
			};

			restTemplate.setRequestFactory(new HttpComponentsClientHttpRequestFactory(
					HttpClientBuilder
							.create()
							.setSSLContext(SSLContexts.custom().loadTrustMaterial(null, acceptingTrustStrategy).build())
							.build()));


			return restTemplate;
		}
		catch(Exception e)
		{
			return null;
		}
	}

	static private HttpHeaders createHttpHeaders(String auth)
	{
		String notEncoded = auth;
		String encodedAuth = "Basic " + Base64.getEncoder().encodeToString(notEncoded.getBytes());
		HttpHeaders headers = new HttpHeaders();
		headers.setContentType(MediaType.APPLICATION_JSON);
		headers.add("Authorization", encodedAuth);
		return headers;
	}




}
