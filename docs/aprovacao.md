## Critérios de aprovação

Teremos os seguintes status de transação:

- pendente
- pre-aprovada
- aguardando-confirmacao
- cancelada
- aprovada

Tipos de pagamento:

- 1 Boleto
- 2 Cartão de crédito
- 3 PIX

### Da transação

A transação por padrão é pendente.
Para que a transação seja cancelada ou aprovada é necessária que haja uma
confirmação por parte da instituição bancária.

#### Cartão de crédito

Para transações via cartão de crédito você deverá enviar para o banco um payload semelhante à:

```json
{
  "transaction-id": "d02cdae3-76cc-4d0a-b969-22d4713e5412",
  "payment-method": 2,
  "payment-info": {
    "card_number": "1111 1111 1111 1111",
    "card_cvc": "123",
    "card_expire_at": "05/2027",
    "card_owner": "John Doe",
    "value": 15980 // R$ 159,80 - VALOR SEMPRE EM CENTAVOS
  },
  "user-id": "48ce48e4-fcc3-47f6-9977-604d93871da5",
  "transaction-cart": "5bec5d55-76ea-4df5-9531-9b1134a3ca64",
  "created-at": "2021-10-20T18:09:45.000Z"
}
```

#### Boleto

Para transações via boleto você deverá enviar um payload semelhante:

```json
{
  "transaction-id": "d02cdae3-76cc-4d0a-b969-22d4713e5412",
  "payment-method": 1,
  "payment-info": {
    "ticket_no": "789 123456789101112 34",
    "value": 4990 // R$ 49,90 - VALOR SEMPRE EM CENTAVOS
  },
  "user-id": "48ce48e4-fcc3-47f6-9977-604d93871da5",
  "transaction-cart": "5bec5d55-76ea-4df5-9531-9b1134a3ca64",
  "created-at": "2021-10-20T18:09:45.000Z"
}
```

#### PIX

Para transações via PIX você deverá enviar um payload semelhante:

```json
{
  "transaction-id": "d02cdae3-76cc-4d0a-b969-22d4713e5412",
  "payment-method": 3,
  "payment-info": {
    "payer_code": "7babee88-39e5-4438-8162-974719fec685",
    "value": 215980 // R$ 2.159,80 - VALOR SEMPRE EM CENTAVOS
  },
  "user-id": "48ce48e4-fcc3-47f6-9977-604d93871da5",
  "transaction-cart": "5bec5d55-76ea-4df5-9531-9b1134a3ca64",
  "created-at": "2021-10-20T18:09:45.000Z"
}
```

### Receber respostas de pagamento

Após enviar o payload você **NÃO** terá retorno!
Você deve ter o seguinte endpoint para **RECEBER** o retorno do banco:

```http request
POST http://localhost:3000/transactions/callback
Content-Type: application/json
Accept: application/json

{
"transaction-id": {{ $transaction-id }},
"payment-status": {{ $status }},
"reason": {{ $reason }}
"updated-at": {{ $timestamp }}
}
```

### Simulação de transações

- [Cartões](./cartoes.md)
- [Boleto](./boletos.md)
- [PIX](./pix.md)