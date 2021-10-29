import { CreditCard, DefaultTransaction } from "@src/service/transaction-types";
import { TransactionConsumer } from "@src/service/transaction-consumer";
import { HttpClient, PayloadModel } from "@src/service/http-client";
import { ServiceTransaction } from "@src/service/service-transaction";
import { sleep, Time } from "@src/util";

export class ProcessCreditCard implements ServiceTransaction{
	private transaction: TransactionConsumer;
	private httpClient: HttpClient;
	private readonly payment: DefaultTransaction<CreditCard>;

	constructor(transaction: TransactionConsumer, httpClient: HttpClient, payment: DefaultTransaction<CreditCard>) {
		this.transaction = transaction
		this.httpClient = httpClient
		this.payment = payment
	}

	async resolveProcess() {
		const transactionResult = this.transaction.transactionResult(this.payment)

		const payloadResult: PayloadModel = {
			"transaction-id": this.payment["transaction-id"],
			"updated-at": new Date().toISOString(),
			reason: "Pagamento recusado. Saldo insuficiente ou dados invalids",
			"payment-status": "cancelada"
		}

		let time = Time.Minute;

		if (transactionResult === 'approved'){
			payloadResult["payment-status"] = "aprovada"
			payloadResult.reason = "OK"
			time = Time.Millisecond
		}

		if (transactionResult === 'analysing') {
			payloadResult["payment-status"] = "pre-aprovada"
			payloadResult.reason = "Em analise contra fraudes"
			time = Time.Millisecond
		}

		await sleep(time)
		await this.httpClient.post(payloadResult)
		return payloadResult
	}

}