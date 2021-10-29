## Do controle de produtos

O controle de produtos deve conter o seguinte fluxo:

Sua API deve salvar os carrinhos de compras criados pelos clientes que chegaram no checkout.

Ou seja

Voce deve ter um endpoint de cadastro de carrinho com os relacionamentos de carrinho para produtos.
Os carrinhos de produtos podem ser deletados, porem nao devem ser removidos da base 
de dados

Itens em carrinhos de compras nao removem unidades de estoque. Somente a efetivação da compra remove
unidades do estoque.

 Você deve disponibilizar um endpoint para o reabastecimento de estoque.
 Este endpoint deve receber a quantidade de itens que foram adquiridos para 
 repor o estoque o preço de compra da unidade no distribuidor.
 
Exemplo prático:

Ao fazer a chamada neste endpoint você informa a quantidade de camisetas (por exemplo) e o valor de compra de cada uma. Imagine o valor de R$ 17,00.

Este endpoint deve registrar estas informações e adicionar a quantidade de camisetas no estoque.

---

A gestão de estoque deve reduzir a quantidade de estoque a cada transação confirmada.
A transação só deve ser enviada para confirmação caso haja a quantidade do produto em estoque.

Por exemplo:

Para uma compra de 5 camisetas e você ter apenas 4 em estoque, você deve retornar um erro 403 não é possível efetuar compra com estoque abaixo do pedido.
