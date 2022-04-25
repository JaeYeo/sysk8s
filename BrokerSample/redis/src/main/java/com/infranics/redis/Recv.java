package com.infranics.redis;


import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisPool;
import redis.clients.jedis.JedisPoolConfig;

public class Recv {


    public static void main(String[] args) {

        try {

            String redis_host = System.getenv("REDIS_HOST");
            String redis_port = System.getenv("REDIS_PORT");
            String redis_password = System.getenv("REDIS_PASSWORD");

            JedisPoolConfig jedis_pool = new JedisPoolConfig();
            JedisPool pool = new JedisPool(jedis_pool, redis_host, Integer.parseInt(redis_port),1000, redis_password);

            Jedis jedis = pool.getResource();
            jedis.set("hello","world");

            String value = jedis.get("hello");

            System.out.printf("redis[\"hello\"]=%s",value);



        } catch(Exception e) {

        } finally {
        }
    }

}