日志系统大致流程

日志是存在日志文件里的，
filebeat从日志文件里读取日志数据，把日志输出到kafka中作为缓冲
kafka后再接一个过滤器 go-stash ，，go-stash获取kafka中日志根据配置过滤字段，
，然后将过滤后的字段输出到elasticsearch中，最后由kibana显示


先用docker swarm根据docker-compose搭建好elk环境

需要注意的是在部署elastic时，elastic不能是root用户启动，所以需要写个dockerfile去先创建一个elastic的用户
部署完可以检查一下用户是谁，
docker inspect --format='{{.Config.User}}' elasticsearch




