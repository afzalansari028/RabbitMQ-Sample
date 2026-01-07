# RabbitMQ Docker Setup

To run RabbitMQ locally with the management UI enabled:

`$$ docker run -d --hostname my-rabbit --name some-rabbit -p 5672:5672 -p 15672:15672 rabbitmq:3-management`

This command pulls the official RabbitMQ image and starts a container in detached mode.

Ports Used
5672 – AMQP (Message Communication)

This port is used by applications (such as Go producers and consumers) to publish and consume messages.

15672 – Management & Monitoring UI

This port is used to access the RabbitMQ web-based dashboard in a browser. It helps in monitoring queues, exchanges, and messages.

Default login cred:
`Username: guest`
`Password: guest`
