## Parte 1

### Exercício 1 - Imposto sobre o salário #1
1. Em sua função “main”, defina uma variável chamada “salario” e atribua um valor do
tipo “int”.
2. Crie um erro personalizado com uma struct que implemente “Error()” com a
mensagem “erro: O salário digitado não está dentro do valor mínimo" em que seja
disparado quando “salário” for menor que 15.000. Caso contrário, imprima no
console a mensagem “Necessário pagamento de imposto”.

### Exercício 2 - Imposto sobre o salário #2
Faça a mesma coisa que no exercício anterior, mas reformule o código para que, em vez de
“Error()”, seja implementado “errors.New()”.

### Exercício 3 - Imposto sobre o salário #3

Repita o processo anterior, mas agora implementando "fmt.Errorf()", para que a mensagem de
erro receba como parâmetro o valor de "salario", indicando que não atinge o mínimo
tributável (a mensagem exibida pelo console deve dizer : "erro: o mínimo tributável é 15.000 e
o salário informado é: [salario]”, onde [salario] é o valor do tipo int passado pelo parâmetro).

### Exercício 4 - Imposto sobre o salário #4
Vamos fazer com que nosso programa seja um pouco mais complexo e útil.
1. Desenvolva as funções necessárias para permitir que a empresa calcule:
a) Salário mensal de um funcionário segundo a quantidade de horas trabalhadas.
- A função receberá as horas trabalhadas no mês e o valor da hora como
parâmetro.
- Esta função deve retornar mais de um valor (salário calculado e erro).
- No caso de o salário mensal ser igual ou superior a R$ 20.000, 10% devem ser
descontados como imposto.
- Se o número de horas mensais inseridas for menor que 80 ou um número
negativo, a função deverá retornar um erro. Deve indicar "erro: o trabalhador
não pode ter trabalhado menos de 80 horas por mês".

2. Desenvolva o código necessário para cumprir as funcionalidades solicitadas, usando
“errors.New()”, “fmt.Errorf()” e “errors.Unwrap()”. Não esqueça de realizar as validações dos
retornos de erro em sua função “main()”.


