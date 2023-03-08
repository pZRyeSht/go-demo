# kafka deploy

## 参数说明
```text
    # 广播主机名称，一般用IP指定
    KAFKA_ADVERTISED_HOST_NAME：
    # Zookeeper连接地址，格式：zoo1：port1,zoo2:port2:/path
    KAFKA_ZOOKEEPER_CONNECT：
    # Kafka启动所使用的的协议及端口
    KAFKA_LISTENERS：
    # Kafka广播地址及端口，告诉客户端，使用什么地址和端口能连接到Kafka，不指定，宿主机以外的客户端将无法连接到Kafka
    KAFKA_ADVERTISED_LISTENERS：
    # 指定BrokerId，如果不指定，将会自己生成
    KAFKA_BROKER_ID：
```

## docker-compose 命令
```text
    # 创建并启动容器，可以使用 -f 参数指定docker-compose.yml文件
    sudo docker-compose up -d
    # 查看状态
    sudo docker-compose ps
    # 停止并删除容器，可以使用 -f 参数指定docker-compose.yml文件
    sudo docker-compose down
```