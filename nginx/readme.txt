对于nginx，可能有些不解

为什么告诉前端接口的ip port url他就可以访问到的。为什么还要加层nginx给他调用？

因为会有这些好处:

1. 统一入口和路由
NGINX 作为反向代理，提供一个统一的入口点。
这意味着所有请求都通过一个公共的端口（在这个例子中是 8081）进入，然后根据请求的路径路由到不同的后端服务。
这种方式使得客户端只需要知道一个 URL，而不需要分别知道各个服务的具体地址和端口。
比如下面的配置，客户端只要知道8081端口，对于服务的具体ip port 它不用担心，并且提高了安全性
server{
  listen 8081;
  access_log /var/log/nginx/looklook.com_access.log;
  error_log /var/log/nginx/looklook.com_error.log;
  location ~ /order/ {
       proxy_set_header Host $http_host;
       proxy_set_header X-Real-IP $remote_addr;
       proxy_set_header REMOTE-HOST $remote_addr;
       proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
       proxy_pass http://looklook:1001;
  }
  location ~ /payment/ {
      proxy_set_header Host $http_host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header REMOTE-HOST $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_pass http://looklook:1002;
  }
}

2. 负载均衡
虽然您的配置中没有展示负载均衡的功能，但 NGINX 可以配置为在多个后端服务器之间分发请求，从而提高应用的可用性和响应速度。
比如
http {
    upstream backend {
        server backend1.example.com;
        server backend2.example.com;
        server backend3.example.com;
    }
    server {
        listen 8081;
        server_name example.com;

        location / {
            proxy_pass http://backend;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
        }
    }
}

proxy_pass http://backend 将请求转发到上游服务器组 backend，NGINX 会自动均匀地分发请求到三个后端服务器。
默认情况下，NGINX 使用轮询算法（Round Robin）进行负载均衡
加权轮询：
upstream backend {
    server backend1.example.com weight=3;
    server backend2.example.com weight=2;
    server backend3.example.com weight=1;
}
最少链接：
upstream backend {
    least_conn;
    server backend1.example.com;
    server backend2.example.com;
    server backend3.example.com;
}

################ 如果我有个服务，是docker swarm部署的，设置的此服务有三副本，现在配置NGINX我得咋配置
比如 docker service create --name my-service --replicas 3 -p 8080:80 my-service-image
这时在nginx里需要配置请求分发到my-service服务，这样配置:
http {
    upstream my_service {
        # Docker Swarm 服务发现
        server my-service:80;
    }
    server {
        listen 8081;
        server_name localhost;
        location / {
            proxy_pass http://my_service;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
        }
    }
}
upstream my_service 使用 Docker Swarm 内置的 DNS 服务发现机制，将请求负载均衡到 my-service 服务的任意副本上。
部署nginx: docker run --name my-nginx -v $(pwd)/nginx.conf:/etc/nginx/nginx.conf:ro -p 8081:8081 --network my-overlay-network nginx
确保nginx容器和swarm在同一个网络里
确保您的 Docker Swarm 网络配置正确，所有服务和 NGINX 容器都在同一个网络中：
docker network create -d overlay my-overlay-network
将您的服务和 NGINX 容器连接到同一个网络：
docker service update --network-add my-overlay-network my-service
docker network connect my-overlay-network my-nginx

################



3. SSL/TLS 终止
NGINX 可以处理 SSL/TLS 加密通信，减轻后端服务的负担。
这样，您可以在 NGINX 和客户端之间使用 HTTPS，在 NGINX 和后端服务之间使用 HTTP，从而简化后端的配置和提高性能。
nginx配置 SSL :  https://blog.csdn.net/m0_63684495/article/details/128748310




















