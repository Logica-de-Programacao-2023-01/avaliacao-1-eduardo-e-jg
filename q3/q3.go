package q3

import "errors"

//Você recebe um tabuleiro retangular de M x N quadrados. Além disso, você tem um número ilimitado de peças de dominó
//padrão de 2 x 1 quadrados. Você pode girar as peças. Você deve colocar o máximo de peças de dominó possível no
//tabuleiro, seguindo as seguintes condições:
//
//1. Cada peça de dominó cobre completamente dois quadrados.
//2. Nenhuma duas peças de dominó se sobrepõem.
//3. Cada peça de dominó está completamente dentro do tabuleiro. É permitido que a peça toque as bordas do tabuleiro.
//
//Encontre o número máximo de peças de dominó que podem ser colocadas sob essas restrições.
//
//Se M ou N forem iguais ou menores que 0, a função deve retornar um erro.

func DominoPieces(m, n int) (int, error) {

	var tamanhoTab int

	var peça int
	peça = 2

	var resultado int
	tamanhoTab = m * n
	if m <= 0 {

		return 0, errors.New("error")
	}

	if n <= 0 {

		return 0, errors.New("error")

	}

	tamanhoTab = m * n

	resultado = tamanhoTab / peça

	if tamanhoTab%2 == 0 {

		return resultado, nil

	} else {

		tamanhoTab -= 1

		return resultado, nil

	}

}
