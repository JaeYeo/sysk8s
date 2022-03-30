package com.infranics.backendserviceproxy;

import com.infranics.backendserviceproxy.rancher.Rancher;
import com.infranics.backendserviceproxy.rancher.RancherConfig;

import io.fabric8.kubernetes.api.model.*;
import io.fabric8.kubernetes.client.Config;
import io.fabric8.kubernetes.client.DefaultKubernetesClient;
import io.fabric8.kubernetes.client.KubernetesClient;
import io.fabric8.kubernetes.client.dsl.base.ResourceDefinitionContext;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;


import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import java.io.ByteArrayInputStream;
import java.util.Base64;
import java.util.Map;


@Controller
public class BackendserviceproxyController {

	private static final Logger logger = LoggerFactory.getLogger(BackendserviceproxyController.class);

	@Autowired
	RancherConfig rancherConfig;


	@ResponseBody
	@RequestMapping("/getkubeconfig")
	public String getkubeconfig(Model model) {
		System.out.println("request /getkubeconfig");
		System.out.println(rancherConfig.getRancher().getEndpoint());

		String kube_config = rancherConfig.getRancher().GetKubeConfig();

		System.out.println(kube_config);

		return kube_config;
	}

	@ResponseBody
	@RequestMapping("/createclusterbroker")
	public String createclusterbroker(Model model) {
		System.out.println("request /createclusterbroker");

		try {


			String kube_config = rancherConfig.getRancher().GetKubeConfig();

			// create cluster testing

			Config config = Config.fromKubeconfig(kube_config);

			KubernetesClient k8s = new DefaultKubernetesClient(config);

			String addon_name = "redis-cfg";
			String addon_url = "https://raw.githubusercontent.com/infranics/sysk8s/main/redis-0.1.0/index-redis.yaml";

			// yaml file에서 가저올때
			// ClassPathResource resource = new ClassPathResource("redis-cluster-addons.yaml");
			//k8s.load(new FileInputStream(resource.getFile())).createOrReplace();

			String rawJsonCustomResourceObj = Rancher.makeClusterBrokerJson(addon_name, addon_url);
			k8s.load(new ByteArrayInputStream(rawJsonCustomResourceObj.getBytes())).createOrReplace();




			return "success:create "+addon_name;
		}
		catch (Exception e)
		{
			return "fail:"+e.toString();
		}

	}


	@ResponseBody
	@RequestMapping("/deleteclusterbroker")
	public String deleteclusterbroker(Model model) {
		System.out.println("request /deleteclusterbroker");

		try {

			String kube_config = rancherConfig.getRancher().GetKubeConfig();

			// create cluster testing

			Config config = Config.fromKubeconfig(kube_config);

			KubernetesClient k8s = new DefaultKubernetesClient(config);

			String addon_name = "redis-cfg";
			// delete할때는 아래 내용이 필요없다. 필요하면 수정하시요.
			String addon_url = "https://raw.githubusercontent.com/infranics/sysk8s/main/redis-0.1.0/index-redis.yaml";

			String rawJsonCustomResourceObj = Rancher.makeClusterBrokerJson(addon_name,addon_url);

			k8s.load(new ByteArrayInputStream(rawJsonCustomResourceObj.getBytes())).delete();


			return "success:delete "+addon_name;
		}
		catch (Exception e)
		{
			return "fail:"+e.toString();
		}

	}

	@ResponseBody
	@RequestMapping("/provisioning")
	public String provisioning(Model model) {
		System.out.println("request /provisioning");

		try {

			String kube_config = rancherConfig.getRancher().GetKubeConfig();

			// create cluster testing

			Config config = Config.fromKubeconfig(kube_config);

			KubernetesClient k8s = new DefaultKubernetesClient(config);

			String serviceName = "redis-01";
			String namespace = "redis";
			String clusterClassName = "redis";
			String clusterPlanName = "micro";
			String[] parameters = {"redis-01"};

			String rawJsonCustomResourceObj = Rancher.makeProvisioningWithClusterBrokerJson(serviceName, namespace,clusterClassName,clusterPlanName,parameters);
			k8s.load(new ByteArrayInputStream(rawJsonCustomResourceObj.getBytes())).createOrReplace();

			return "success:provisioning "+serviceName + " in " +namespace ;
		}
		catch (Exception e)
		{
			return "fail:"+e.toString();
		}

	}
	@ResponseBody
	@RequestMapping("/deprovisioning")
	public String deprovisioning(Model model) {
		System.out.println("request /deprovisioning");

		try {

			String kube_config = rancherConfig.getRancher().GetKubeConfig();

			// create cluster testing

			Config config = Config.fromKubeconfig(kube_config);

			KubernetesClient k8s = new DefaultKubernetesClient(config);

			String serviceName = "redis-01";
			String namespace = "redis";

			// delete할때는 아래 내용이 필요없다. 필요하면 수정하시요.
			String clusterClassName = "redis";
			String clusterPlanName = "micro";
			String[] parameters = {"redis-01"};

			String rawJsonCustomResourceObj = Rancher.makeProvisioningWithClusterBrokerJson(serviceName, namespace,clusterClassName,clusterPlanName,parameters);
			k8s.load(new ByteArrayInputStream(rawJsonCustomResourceObj.getBytes())).delete();

			return "success:deprovisioning "+serviceName + " in " +namespace ;
		}
		catch (Exception e)
		{
			return "fail:"+e.toString();
		}

	}


	@ResponseBody
	@RequestMapping("/list")
	public String list(Model model) {
		System.out.println("request /list");

		try {

			String kube_config = rancherConfig.getRancher().GetKubeConfig();

			Config config = Config.fromKubeconfig(kube_config);

			KubernetesClient k8s = new DefaultKubernetesClient(config);



			ResourceDefinitionContext resourceDefinitionContext = new ResourceDefinitionContext.Builder()
					.withGroup("servicecatalog.k8s.io")
					.withVersion("v1beta1")
					.withKind("ServiceInstance")
					.withPlural("serviceinstances")
					.withNamespaced(true)
					.build();


			GenericKubernetesResourceList list = k8s.genericKubernetesResources(resourceDefinitionContext).inNamespace("redis").list();

			String info = "[";
			for(GenericKubernetesResource item : list.getItems())
			{
				String instance_name = item.getMetadata().getName();

				//System.out.println(item);

				Map<String,Object> specs = (Map<String, Object>)item.getAdditionalProperties().get("spec");

				String class_name = specs.get("clusterServiceClassExternalName").toString();
				String class_id = ((Map<String, Object>)specs.get("clusterServiceClassRef")).get("name").toString();

				String plan_name = specs.get("clusterServicePlanExternalName").toString();
				String plan_id = ((Map<String, Object>)specs.get("clusterServicePlanRef")).get("name").toString();


				Map<String,Object> status = (Map<String, Object>)item.getAdditionalProperties().get("status");
				String last_state = status.get("lastConditionState").toString();


				info = info + "instance="+instance_name+"/class="+class_name+"("+class_id+")/plan="+plan_name+"("+plan_id+")/state="+last_state+", ";

			}

			info += "]";

			return "success:"+info;

		}
		catch (Exception e)
		{
			return "fail:"+e.toString();
		}

	}


	@ResponseBody
	@RequestMapping("/instance")
	public String item(Model model) {
		System.out.println("request /instance");

		try {

			String kube_config = rancherConfig.getRancher().GetKubeConfig();

			Config config = Config.fromKubeconfig(kube_config);

			KubernetesClient k8s = new DefaultKubernetesClient(config);



			ResourceDefinitionContext resourceDefinitionContext = new ResourceDefinitionContext.Builder()
					.withGroup("servicecatalog.k8s.io")
					.withVersion("v1beta1")
					.withKind("ServiceInstance")
					.withPlural("serviceinstances")
					.withNamespaced(true)
					.build();

			String instance_name = "redis-03";

			GenericKubernetesResource  instance = k8s.genericKubernetesResources(resourceDefinitionContext).inNamespace("redis").withName(instance_name).get();

			System.out.println(instance);

			String info = "[";

			String instance_name_from = instance.getMetadata().getName();

			Map<String,Object> specs = (Map<String, Object>)instance.getAdditionalProperties().get("spec");

			String class_name = specs.get("clusterServiceClassExternalName").toString();
			String class_id = ((Map<String, Object>)specs.get("clusterServiceClassRef")).get("name").toString();

			String plan_name = specs.get("clusterServicePlanExternalName").toString();
			String plan_id = ((Map<String, Object>)specs.get("clusterServicePlanRef")).get("name").toString();


			Map<String,Object> status = (Map<String, Object>)instance.getAdditionalProperties().get("status");
			String last_state = status.get("lastConditionState").toString();


			info = info + "instance="+instance_name+"/class="+class_name+"("+class_id+")/plan="+plan_name+"("+plan_id+")/state="+last_state+", ";



			info += "]";


			return "success:"+info;

		}
		catch (Exception e)
		{
			return "fail:"+e.toString();
		}

	}


	@ResponseBody
	@RequestMapping("/nodeport")
	public String nodeport(Model model) {
		System.out.println("request /nodeport");

		try {

			String kube_config = rancherConfig.getRancher().GetKubeConfig();

			// create cluster testing

			Config config = Config.fromKubeconfig(kube_config);

			KubernetesClient k8s = new DefaultKubernetesClient(config);

			String serviceName = "redis-03";
			String namespace = "redis";

			String nodePort = "";


			ServiceList list = k8s.services().inNamespace("redis")
					.withLabelSelector("app.kubernetes.io/component=master,app.kubernetes.io/instance=redis-03,app.kubernetes.io/name=redis")
					.list();

			//System.out.println(list);

			for(Service item : list.getItems()) {


				//System.out.println(item);

				String instance_name = item.getMetadata().getName();


				nodePort = item.getSpec().getPorts().get(0).getNodePort().toString();


			}


			return "success:nodePort="+nodePort + " of "+serviceName+" in " +namespace  ;

		}
		catch (Exception e)
		{
			return "fail:"+e.toString();
		}

	}

	@ResponseBody
	@RequestMapping("/provisioningredis")
	public String provisioningredis(Model model) {
		System.out.println("request /provisioningredis");

		try {

			String kube_config = rancherConfig.getRancher().GetKubeConfig();

			// create cluster testing

			Config config = Config.fromKubeconfig(kube_config);

			KubernetesClient k8s = new DefaultKubernetesClient(config);

			String serviceName = "redis-03";
			String namespace = "redis";
			String clusterClassName = "redis";
			String clusterPlanName = "micro";
			String password = "master77!!";
			String redisPort = "6380";

			String rawJsonCustomResourceObj = Rancher.makeProvisioningRedisWithClusterBrokerJson(serviceName, namespace,clusterClassName,clusterPlanName,password, redisPort);
			k8s.load(new ByteArrayInputStream(rawJsonCustomResourceObj.getBytes())).createOrReplace();

			return "success:provisioning "+serviceName + " in " +namespace +" with password="+password+" and redisPort="+redisPort ;
		}
		catch (Exception e)
		{
			return "fail:"+e.toString();
		}

	}


	@ResponseBody
	@RequestMapping("/bindingredis")
	public String bindingredis(Model model) {
		System.out.println("request /bindingredis");

		try {

			String kube_config = rancherConfig.getRancher().GetKubeConfig();

			// create cluster testing

			Config config = Config.fromKubeconfig(kube_config);

			KubernetesClient k8s = new DefaultKubernetesClient(config);

			String serviceName = "redis-03";
			String namespace = "redis";

			String rawJsonCustomResourceObj = Rancher.makeBindingJson(serviceName, namespace);

			System.out.println(rawJsonCustomResourceObj);

			k8s.load(new ByteArrayInputStream(rawJsonCustomResourceObj.getBytes())).createOrReplace();


			Thread.sleep(3000);

			String secretName = serviceName+"-binding";

			String info = "";

			Secret secret = k8s.secrets().inNamespace(namespace).withName(secretName).get();

			System.out.println(secret);

			Map<String,String> data = secret.getData();

			for(String key : data.keySet() )
			{
				String value = new String(Base64.getDecoder().decode(data.get(key)));

				info = info + key+ "=" + value +";";
			}

			System.out.println(info);

			return "success:binding "+serviceName + " in " +namespace +":" + info;
		}
		catch (Exception e)
		{
			return "fail:"+e.toString();
		}

	}


	@ResponseBody
	@RequestMapping("/unbindingredis")
	public String unbindingredis(Model model) {
		System.out.println("request /unbindingredis");

		try {

			String kube_config = rancherConfig.getRancher().GetKubeConfig();

			// create cluster testing

			Config config = Config.fromKubeconfig(kube_config);

			KubernetesClient k8s = new DefaultKubernetesClient(config);

			String serviceName = "redis-03";
			String namespace = "redis";

			String rawJsonCustomResourceObj = Rancher.makeBindingJson(serviceName, namespace);

			System.out.println(rawJsonCustomResourceObj);


			k8s.load(new ByteArrayInputStream(rawJsonCustomResourceObj.getBytes())).delete();

			return "success:unbinding "+serviceName + " in " +namespace ;
		}
		catch (Exception e)
		{
			return "fail:"+e.toString();
		}

	}


}

