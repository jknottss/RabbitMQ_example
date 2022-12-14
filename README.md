#### RabbitMQ-fiber-example

1. Run containers with the RabbitMQ, [Fiber](https://github.com/gofiber/fiber) and consumer by this command:

```bash
make run
```

2. Make HTTP request to the API endpoint:

```console
curl \
    --request GET \
    --url 'http://localhost:3000/send?msg=test'
```

3. Go to RabbitMQ awesome dashboard [localhost:15672](http://localhost:15672) and see `QueueService1` queue with sent messages:
![screen](https://github.com/jknottss/rabbitMQ-fiber-example/blob/main/sample.png)
