# PIX

## Compra aprovada em 5 segundos

```json
{
  "payer_code": "7babee88-39e5-4438-8162-974719fec685"
}
```

## Compra expirada em 5 segundos

Esta compra retorna o status cancelada para simular o PIX invalido.
*Simula a transação expirada em 3 dias

```json
{
  "payer_code": "7babee88-39e5-4438-8162-974719fec000"
}
```

## Compra recusada na hora

Qualquer payload diferente dos acima.