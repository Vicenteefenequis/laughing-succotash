
export type CreditCard = {
	"card_number": string,
	"card_cvc": string,
	"card_expire_at": string,
	"card_owner": string,
}

export type Ticket = {
	"ticket_no": string,
}

export type Pix = {
	"payer_code": string,
}

type PaymentInfoDefaults = {
	value: number
}

type CheckPaymentType<T> = T extends CreditCard
	? CreditCard
	: T extends Ticket
		? Ticket
		: Pix;

export type DefaultTransaction<T> = {
	"transaction-id": string,
	"payment-method": 1 | 2 | 3,
	"payment-info": CheckPaymentType<T> & PaymentInfoDefaults,
	"user-id": string,
	"transaction-cart": string,
	"created-at": string
}