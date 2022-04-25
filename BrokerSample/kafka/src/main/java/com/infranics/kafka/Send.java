package com.infranics.kafka;

import com.sun.deploy.util.StringUtils;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerConfig;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.clients.producer.RecordMetadata;
import org.apache.kafka.common.serialization.StringSerializer;

import java.util.Properties;
import java.util.Scanner;


public class Send {
    private static final String TOPIC_NAME = "test";
    private static final String FIN_MESSAGE = "exit";

    public static void main(String[] args) {
        Properties properties = new Properties();
        //properties.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, "112.175.114.177:30346");
        properties.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, "kafka-03-0.kafka-03-headless.kafka.svc.cluster.local:30346");
        properties.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class.getName());
        properties.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, StringSerializer.class.getName());

        KafkaProducer<String, String> producer = new KafkaProducer<>(properties);

        while(true) {
            Scanner sc = new Scanner(System.in);
            System.out.print("Input > ");
            String message = sc.nextLine();

            ProducerRecord<String, String> record = new ProducerRecord<>(TOPIC_NAME, message);
            try {

                System.out.print("Send :" + message);
                producer.send(record, (metadata, exception) -> {
                    if (exception != null) {
                        // some exception

                        System.out.println(exception);
                    }
                });

            } catch (Exception e) {
                // exception

                System.out.println(e.toString());
            } finally {
                producer.flush();
            }

            System.out.print("Done :" + message);

            if(message.equalsIgnoreCase(FIN_MESSAGE)) {
                producer.close();
                break;
            }
        }
    }
}