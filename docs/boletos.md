# Boletos

## Compra aprovada em 1 minuto

```json
{
  "ticket_no": "789 123456789101112 34"
}
```

## Compra expirada em 1 minuto

Esta compra retorna o status cancelada para simular o não pagamento do boleto.
*Simula a transação expirada em 3 dias

```json
{
  "ticket_no": "789 123456789101112 10"
}
```

## Compra recusada em 1 minuto

Qualquer payload diferente dos acima.