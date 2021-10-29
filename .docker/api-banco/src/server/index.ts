import express from "express";
import { RabbitMQ } from "@src/queue/rabbit";
import { TransactionConsumer } from "@src/service/transaction-consumer";
import {
  CreditCard,
  DefaultTransaction,
  Pix,
  Ticket,
} from "@src/service/transaction-types";
import { ProcessCreditCard } from "@src/service/process-credit-card";
import { HttpClient } from "@src/service/http-client";
import { ProcessTicket } from "@src/service/process-ticket";
import { Channel, ConsumeMessage } from "amqplib";
import { ProcessPix } from "@src/service/process-pix";

const rabbitMQ = new RabbitMQ();

const app = express();
app.use(express.json());

function ackMessage(message: ConsumeMessage, channel: Channel) {
  return function () {
    channel.ack(message);
  };
}

rabbitMQ.consume(async (message, channel) => {
  const payload = message.content.toString();
  const transaction = new TransactionConsumer();
  const httpClient = new HttpClient();

  const paymentObj = transaction.parseToJSON(payload);

  const isValidTransactionObj = transaction.checkObjectKeys(paymentObj);

  if (!isValidTransactionObj) {
    console.log("Payload de pagamento invalido");
    return;
  }

  const paymentType = transaction.identifyPaymentMethod(
    paymentObj as DefaultTransaction<any>
  );
  const ack = ackMessage(message, channel);

  if (paymentType === "credit_card") {
    const creditCardProcess = new ProcessCreditCard(
      transaction,
      httpClient,
      paymentObj as DefaultTransaction<CreditCard>
    );
    creditCardProcess.resolveProcess().then(console.log).then(ack);
  }

  if (paymentType === "ticket") {
    const ticketProcess = new ProcessTicket(
      transaction,
      httpClient,
      paymentObj as DefaultTransaction<Ticket>
    );

    ticketProcess.resolveProcess().then(console.log).then(ack);
  }

  if (paymentType === "pix") {
    const pixProcess = new ProcessPix(
      transaction,
      httpClient,
      paymentObj as DefaultTransaction<Pix>
    );

    pixProcess.resolveProcess().then(console.log).then(ack);
  }

  return new Promise((resolve) => {
    resolve();
  });
});

app.listen(3000, () => {
  console.log(`RUNNING PORT 3000`);
});
