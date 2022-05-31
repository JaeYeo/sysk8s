#!/bin/bash



if [[ -f "/home/ubuntu/helm-broker/kong-0.1.1/work/nginx_updated" ]]
then
    echo "nginx.conf was updated"
else
    if [[ -f "/home/ubuntu/helm-broker/kong-0.1.1/work/nginx-kong.conf" ]]
    then
        sed -i 's/server_name kong_admin;/server_name kong_admin;\n    auth_basic "admin basic auth";\n    auth_basic_user_file \/opt\/bitnami\/kong\/server\/.htpasswd;/g' /home/ubuntu/helm-broker/kong-0.1.1/work/nginx-kong.conf;
        echo "updated" >> /home/ubuntu/helm-broker/kong-0.1.1/work/nginx_updated
        echo "nginx.conf is updated"
    else
        echo "no nginx.conf"
    fi
fi
