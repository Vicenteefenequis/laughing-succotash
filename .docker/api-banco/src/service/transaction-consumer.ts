import _ from 'lodash';
import { CreditCard, DefaultTransaction, Pix, Ticket } from "@src/service/transaction-types";
import { cards, pixs, tickets } from "@src/service/constants";

export class TransactionConsumer {

	parseToJSON(payload: string): object {
		try {
			return JSON.parse(payload)
		} catch (e: any) {
			return {}
		}

	}

	checkObjectKeys(payload: object): boolean {
		const mustContains = [
			'transaction-id',
			'payment-method',
			'payment-info',
			'user-id',
			'transaction-cart',
			'created-at'
		]

		const objKeys = Object.keys(payload)

		if (!objKeys.length) {
			return false
		}

		return mustContains.every(k => objKeys.includes(k))
	}

	identifyPaymentMethod(payload: DefaultTransaction<any>) {
		const paymentInfo = payload['payment-info']

		if (_.has(paymentInfo, 'card_number')) {
			return 'credit_card'
		}

		if (_.has(paymentInfo, 'ticket_no')) {
			return 'ticket'
		}

		return 'pix'
	}

	transactionResult(payload: DefaultTransaction<CreditCard>) {
		_.unset(payload["payment-info"], 'value')

		if (_.isEqual(payload["payment-info"], cards.get('approved'))) {
			return 'approved';
		}

		if (_.isEqual(payload["payment-info"], cards.get('analysing'))) {
			return 'analysing'
		}

		return 'expired'
	}

	isValidTicket(payload: DefaultTransaction<Ticket>) {
		_.unset(payload["payment-info"], 'value')

		if (_.isEqual(payload["payment-info"], tickets.get('approved'))) {
			return 'approved'
		}

		if (_.isEqual(payload["payment-info"], tickets.get('expired'))) {
			return 'expired'
		}

		return 'canceled'
	}

	isValidPix(payload: DefaultTransaction<Pix>) {
		_.unset(payload["payment-info"], 'value')

		if (_.isEqual(payload["payment-info"], pixs.get('approved'))) {
			return 'approved'
		}

		if (_.isEqual(payload["payment-info"], pixs.get('expired'))) {
			return 'expired'
		}

		return 'canceled'
	}

}