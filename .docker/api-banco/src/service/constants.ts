import { CreditCard, Pix, Ticket } from "@src/service/transaction-types";

type TransactionTypes = 'approved' | 'analysing' | 'expired' | 'canceled';

export const cards = new Map<TransactionTypes, CreditCard>([
	[ 'approved', {
		card_number: "1111 1111 1111 1111",
		card_cvc: "123",
		card_expire_at: "05/2027",
		card_owner: "John Doe"
	} ],
	[ 'analysing', {
		card_number: "1111 1111 1111 1112",
		card_cvc: "123",
		card_expire_at: "05/2027",
		card_owner: "John Doe"
	} ]
])

export const tickets = new Map<TransactionTypes, Ticket>([
	[ 'approved', {
		ticket_no: "789 123456789101112 34"
	} ],
	[ 'expired', {
		ticket_no: "789 123456789101112 10"
	} ]
])

export const pixs = new Map<TransactionTypes, Pix>([
	[ 'approved', {
		payer_code: "7babee88-39e5-4438-8162-974719fec685"
	} ],
	[ 'expired', {
		payer_code: "7babee88-39e5-4438-8162-974719fec000"
	} ],
])