import axios from 'axios';

export type Status = 'pendente' |
	'pre-aprovada' |
	'aguardando-confirmacao' |
	'cancelada' |
	'aprovada';

export type PayloadModel = {
	"transaction-id": string;
	"payment-status": Status;
	"reason": string;
	"updated-at": string;
}

export class HttpClient {
	async post(payload: PayloadModel) {
		try {
			await axios.post('http://localhost:3000/transactions/callback', payload)
		}catch (e: any) {
			console.log("Nao foi possivel executar o callback: ", e.message)
		}
	}

}