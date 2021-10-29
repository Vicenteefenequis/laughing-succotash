## Categorias de produtos

Todo produto deve ter uma categoria relacionada para que seja possível fazer uma busca a partir de categorias.

Exemplo:

- Vestuário
  - Camiseta
  - Jaqueta
  - Bermuda
- Cama, mesa e banho
  - Conjunto de banho com 3 toalhas
  - Conjunto cama de casal
  - Lençol
- Eletrodomesticos
  - Aparelho de televisão
  - Smartphone
  - Notebook

Você deve ter um endpoint de busca por categorias (mais de uma categoria).

Exemplo:

```http request
GET {{ $baseURL }}/categories?names=vestuario,cama-mesa-e-banho
```