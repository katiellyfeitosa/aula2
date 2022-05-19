# Bootcamp Golang - GO Bases I - aula2
## Parte 1
### Exercício 1 - Imprimindo o nome na tela
1. Crie uma aplicação que tenha uma variável “nome” e outra “idade”.
2. Imprima no terminal o valor de cada variável.

### Exercício 2 - Clima
Uma empresa de meteorologia quer ter um sistema onde possa ter a temperatura, umidade e
pressão atmosférica de diferentes lugares.
1. Declare 3 variáveis especificando o tipo de dado delas, como valor elas devem
possuir a temperatura, umidade e pressão atmosférica de onde você se encontra.
2. Imprima os valores no console.
3. Quais tipos de dados serão atribuídos a essas variáveis?

### Exercício 3 - Declaração de variáveis

Um professor de programação está corrigindo as avaliações de seus alunos na disciplina de
Programação I para fornecer-lhes o feedback correspondente. Um dos pontos do exame é
declarar diferentes variáveis.
Ajude o professor com as seguintes questões:
1. Verifique quais dessas variáveis declaradas pelo aluno estão corretas.
2. Corrigir as incorrectas.

    var 1nome string <br />
    var sobrenome string <br />
    var int idade <br />
    1sobrenome := 6 <br />
    var licenca_para_dirigir = true <br />
    var estatura da pessoa int <br />
    quantidadeDeFilhos := 2 <br />
    

### Exercício 4 - Tipos de dados

Um estudante de programação tentou fazer declarações de variáveis com seus respectivos
tipos em Go mas teve várias dúvidas ao fazê-lo. A partir disso, ele nos deu seu código e
pediu a ajuda de um desenvolvedor experiente que pode:
● Revisar o código e realizar as devidas correções.    
    
    var sobrenome string = "Silva" 
    var idade int = "25" 
    boolean := "false"; 
    var salario string = 4585.90 
    var nome string = "Fellipe" 

## Parte 2

### Exercício 1 - Letras de uma palavra

A Real Academia Brasileira quer saber quantas letras tem uma palavra e então ter cada uma
das letras separadamente para soletrá-la. Para isso terão que:
1. Crie uma aplicação que tenha uma variável com a palavra e imprima o número de
letras que ela contém.
2. Em seguida, imprimi cada uma das letras.

### Exercício 2 - Empréstimo

Um banco deseja conceder empréstimos a seus clientes, mas nem todos podem acessá-los.
Para isso, possui certas regras para saber a qual cliente pode ser concedido. Apenas
concede empréstimos a clientes com mais de 22 anos, empregados e com mais de um ano
de atividade. Dentro dos empréstimos que concede, não cobra juros para quem tem um
salário superior a US$ 100.000.
É necessário fazer uma aplicação que possua essas variáveis e que imprima uma mensagem
de acordo com cada caso.
Dica: seu código deve ser capaz de imprimir pelo menos 3 mensagens diferentes.

### Exercício 3 - A que mês corresponde?

Faça uma aplicação que contenha uma variável com o número do mês.
1. Dependendo do número, imprima o mês correspondente em texto.
2. Ocorre a você que isso pode ser resolvido de maneiras diferentes? Qual você
escolheria e porquê?
Ex: 7 de julho.
    

### Exercício 4 - Exercício 4 - Quantos anos tem...
Um funcionário de uma empresa deseja saber o nome e a idade de um de seus funcionários.
De acordo com o mapa abaixo, ajude a imprimir a idade de Benjamin.

var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

Por outro lado, você também precisa:
- Saiba quantos de seus funcionários têm mais de 21 anos.
- Adicione um novo funcionário à lista, chamado Federico, que tem 25 anos.
- Excluir Pedro do mapa.