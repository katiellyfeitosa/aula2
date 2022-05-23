## Parte 1

### Exercício 1 - Impostos de salário
Uma empresa de chocolates precisa calcular o imposto de seus funcionários no momento de
depositar o salário, para cumprir seu objetivo será necessário criar uma função que retorne o
imposto de um salário.
Temos a informação que um dos funcionários ganha um salário de R$50.000 e será
descontado 17% do salário. Um outro funcionário ganha um salário de $150.000, e nesse
caso será descontado mais 10%.

### Exercício 2 - Calcular média
Um colégio precisa calcular a média das notas (por aluno). Precisamos criar uma função na
qual possamos passar N quantidade de números inteiros e devolva a média.
Obs: Caso um dos números digitados seja negativo, a aplicação deve retornar um erro

### Exercício 3 - Calcular salário
Uma empresa marítima precisa calcular o salário de seus funcionários com base no número
de horas trabalhadas por mês e na categoria do funcionário.
Se a categoria for C, seu salário é de R$1.000 por hora
Se a categoria for B, seu salário é de $1.500 por hora mais 20% caso tenha passado de 160h
mensais
Se a categoria for A, seu salário é de $3.000 por hora mais 50% caso tenha passado de 160h
mensais

Calcule o salário dos funcionários conforme as informações abaixo:
- Funcionário de categoria C: 162h mensais
- Funcionário de categoria B: 176h mensais
- Funcionário de categoria A: 172h mensais

### Exercício 4 - Cálculo de estatísticas
Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de
notas dos alunos de um curso, sendo necessário calcular os valores mínimo, máximo e médio
de suas notas.
Será necessário criar uma função que indique que tipo de cálculo deve ser realizado (mínimo,
máximo ou média) e que retorna outra função (e uma mensagem caso o cálculo não esteja
definido) que pode ser passado uma quantidade N de inteiros e retorne o cálculo que foi
indicado na função anterior

## Parte 2

### Exercício 1 - Registro de estudantes
Uma universidade precisa cadastrar os alunos e gerar uma funcionalidade para imprimir os
detalhes dos dados de cada um deles, conforme o exemplo abaixo:
Nome: [Nome do aluno]
Sobrenome: [Sobrenome do aluno]
RG: [RG do aluno]
Data de admissão: [Data de admissão do aluno]

Os valores que estão entre parênteses devem ser substituídos pelos dados fornecidos pelos
alunos.
Para isso é necessário gerar uma estrutura Alunos com as variáveis Nome, Sobrenome, RG,
Data e que tenha um método de detalhamento

### Exercício 2 - Produtos de e-commerce
Diversas lojas de e-commerce precisam realizar funcionalidades no Go para gerenciar
produtos e devolver o valor do preço total.
As empresas têm 3 tipos de produtos:
- Pequeno, Médio e Grande.
Existem custos adicionais para manter o produto no armazém da loja e custos de envio.

Lista de custos adicionais:
- Pequeno: O custo do produto (sem custo adicional)
- Médio: O custo do produto + 3% pela disponibilidade no estoque
- Grande: O custo do produto + 6% pela disponibilidade no estoque + um custo
adicional pelo envio de $2500.

Requisitos:
- Criar uma estrutura “loja” que guarde uma lista de produtos.
- Criar uma estrutura “produto” que guarde o tipo de produto, nome e preço
- Criar uma interface “Produto” que possua o método “CalcularCusto”
- Criar uma interface “Ecommerce” que possua os métodos “Total” e “Adicionar”.
- Será necessário uma função “novoProduto” que receba o tipo de produto, seu nome
e preço, e devolva um Produto.
- Será necessário uma função “novaLoja” que retorne um Ecommerce.
- Interface Produto:
- Deve possuir o método “CalcularCusto”, onde o mesmo deverá calcular o
custo adicional segundo o tipo de produto.

- Interface Ecommerce:
- Deve possuir o método “Total”, onde o mesmo deverá retornar o preço total com
base no custo total dos produtos + o adicional citado anteriormente (caso a categoria
tenha)
- Deve possuir o método “Adicionar”, onde o mesmo deve receber um novo produto
e adicioná-lo a lista da loja