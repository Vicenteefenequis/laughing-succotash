## Objetivo

API de gestão de loja virtual.

## Requisitos

A api deve ter o [controle de produtos](./controle-produtos.md), [categorias de produtos](./categorias-produtos.md), gestao de estoque.
A api deve permitir a compra com cartao de credito (dados fictícios).
A api deve ter usuarios do tipo cliente e loja.

---

## Descrições

### Do controle de produtos

O controle de produtos inclui preço e quantidades de cada produto.
A API deve fazer o controle de entrada e saída de produtos automaticamente com base somente nas transações que envolvem este produto.

### Do pagamento

A API deve aceitar pagamentos dos tipos:
- Boleto (Deve gerar apenas um código aleatório simulando um código de barras numérico)
- Cartão de Crédito (Ver [critérios de aprovação](./aprovacao.md))
- Pix (Gera um base64 com o QRcode de pagamento)

Cada pagamento tem suas regras de aprovação.

**Boleto**

Para pagamentos em Boleto a transação fica com o status de `pendente`.
Transações com o status de pendente expiram em 3 dias (úteis).

**Cartão de crédito**

Para pagamentos em Cartão de crédito a transação fica com o status de `pre-aprovada`.
Transações com o status de pre aprovada expiram em 2 dias (úteis).

**PIX**

Para pagamentos via PIX a transação fica com o status de `aguardando-confirmacao`
Transações via PIX devem receber um evento e podem ficar em aguardando confirmação por
no máximo 10 segundos.

[Confira as regras de aprovação aqui](./aprovacao.md)
