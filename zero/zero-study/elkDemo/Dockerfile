FROM docker.elastic.co/elasticsearch/elasticsearch:8.4.2

# 创建用户和组
RUN groupadd -g 1000 elasticgroup && \
    useradd -u 1000 -g elasticgroup -s /bin/bash -m elasticuser \

RUN chown -R 1000:1000 /usr/share/elasticsearch

# 切换到 elasticuser
USER elasticuser

# 继续其他配置...
