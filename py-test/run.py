import pulsar

client = pulsar.Client('pulsar://localhost:6650')

producer = client.create_producer('my-topic')

for i in range(10):
    producer.send(('Hello-%d' % i).encode('utf-8'))
    print("send message to pulsar")
    
consumer = client.subscribe('my-topic', 'my-subscription')

while True:
    msg = consumer.receive()
    try:
        print("Received message '{}' id='{}'".format(msg.data(), msg.message_id()))
        # 确认已经成功处理消息
        consumer.acknowledge(msg)
    except:
        # 消息处理失败
        consumer.negative_acknowledge(msg)

client.close()