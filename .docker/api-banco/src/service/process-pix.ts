import { ServiceTransaction } from "@src/service/service-transaction";
import { TransactionConsumer } from "@src/service/transaction-consumer";
import { HttpClient, PayloadModel } from "@src/service/http-client";
import { DefaultTransaction, Pix } from "@src/service/transaction-types";
import { sleep, Time } from "@src/util";

export class ProcessPix implements ServiceTransaction {

	private transaction: TransactionConsumer;
	private httpClient: HttpClient;
	private readonly payment: DefaultTransaction<Pix>;

	constructor(transaction: TransactionConsumer, httpClient: HttpClient, payment: DefaultTransaction<Pix>) {
		this.transaction = transaction;
		this.httpClient = httpClient;
		this.payment = payment;
	}

	async resolveProcess() {
		const transactionResult = this.transaction.isValidPix(this.payment);

		const payloadResult: PayloadModel = {
			"transaction-id": this.payment["transaction-id"],
			"updated-at": new Date().toISOString(),
			reason: "Pix recusado. Codigo de pagador nao reconhecido",
			"payment-status": "cancelada"
		}

		let time = Time.Millisecond

		if (transactionResult === 'approved') {
			payloadResult.reason = "Pix aprovado"
			payloadResult["payment-status"] = "aprovada"
			time = Time.Second * 5
		}

		if (transactionResult === 'expired') {
			payloadResult.reason = "Pix nao efetuado no prazo de vencimento"
			payloadResult["payment-status"] = "cancelada"
			time = Time.Second * 5
		}

		await sleep(time)
		await this.httpClient.post(payloadResult)
		return payloadResult
	}
}




