import * as Amqp from "amqp-ts";
import { IServicePanel } from "../namespace/initializer";




@IServicePanel.register
export class MessagingService {
    private static connection: Amqp.Connection;

    public init() {
        MessagingService.connection = new Amqp.Connection(process.env["RABBITMQ_URL"]);
    }

    public static listenToQueue(queueName: string, onMessage: (msg: Amqp.Message) => any) {
        var queue = MessagingService.connection.declareQueue(queueName, {
            durable: true
         });
        queue.activateConsumer(onMessage);
    }
}