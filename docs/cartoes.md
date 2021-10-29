# Cartões

## Compra aprovada na hora

```json
{
  "card_number": "1111 1111 1111 1111",
  "card_cvc": "123",
  "card_expire_at": "05/2027",
  "card_owner": "John Doe"
}
```

## Compra em análise

Esta compra retorna o status pre-aprovada
*Simula uma análise mais aprofundada pela instituição bancária

```json
{
  "card_number": "1111 1111 1111 1112",
  "card_cvc": "123",
  "card_expire_at": "05/2027",
  "card_owner": "John Doe"
}
```

## Compra recusada em 1 minuto

Qualquer payload diferente dos acima.