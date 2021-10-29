import { ServiceTransaction } from "@src/service/service-transaction";
import { TransactionConsumer } from "@src/service/transaction-consumer";
import { HttpClient, PayloadModel } from "@src/service/http-client";
import { DefaultTransaction, Ticket } from "@src/service/transaction-types";
import { sleep, Time } from "@src/util";

export class ProcessTicket implements ServiceTransaction{
	private transaction: TransactionConsumer;
	private httpClient: HttpClient;
	private readonly payment: DefaultTransaction<Ticket>;

	constructor(transaction: TransactionConsumer, httpClient: HttpClient, payment: DefaultTransaction<Ticket>) {
		this.transaction = transaction
		this.httpClient = httpClient
		this.payment = payment
	}

	async resolveProcess() {
		const transactionResult = this.transaction.isValidTicket(this.payment)

		const payloadResult: PayloadModel = {
			"transaction-id": this.payment["transaction-id"],
			"updated-at": new Date().toISOString(),
			reason: "Pagamento recusado. Boleto cancelado pelo emissor",
			"payment-status": "cancelada"
		}

		if (transactionResult === 'approved'){
			payloadResult["payment-status"] = "aprovada"
			payloadResult.reason = "OK"
		}

		if (transactionResult === 'expired') {
			payloadResult["payment-status"] = "cancelada"
			payloadResult.reason = "Boleto vencido"
		}

		await sleep(Time.Minute)
		await this.httpClient.post(payloadResult)
		return payloadResult
	}

}