package main

import "errors"

type Matrix struct {
	QuantColunas int
	QuantLinhas  int
	Dados        [][]Elemento
}

type Elemento interface {
	Soma(e Elemento) Elemento

	Subtracao(e Elemento) Elemento

	Multi(e Elemento) Elemento

	MultiEscalar(k float64) Elemento

	Divisao(e Elemento) Elemento
}

type LinhaVetor []Elemento

func NewLinhaVetor(tam int) *LinhaVetor {
	l := make([]Elemento, tam)
	return (*LinhaVetor)(&l)
}

func (l *LinhaVetor) Set(i int, e Elemento) {
	if i < 0 || i >= len(*l) {
		panic("po mo preguiça")
	}
	(*l)[i] = e
}

func (l *LinhaVetor) Get(i int) Elemento {
	if i < 0 || i >= len(*l) {
		panic("po mo preguiça")
	}

	return (*l)[i]
}

func (l *LinhaVetor) MultiplicacaoEscalar(k float64) *LinhaVetor {
	nl := NewLinhaVetor(len(*l))
	for i, e := range *l {
		nl.Set(i, e.MultiEscalar(k))
	}
	return nl
}

func (l *LinhaVetor) AdicaoDeLinhas(ol *LinhaVetor) (*LinhaVetor, error) {
	if len(*l) != len(*ol) {
		return nil, errors.New("tamanho incompativel")
	}
	nl := NewLinhaVetor(len(*l))
	for i := range *l {
		el := ol.Get(i)
		nl.Set(i, l.Get(i).Soma(el))
	}
	return nl, nil
}

func NewMatrix(linhas, colunas int) *Matrix {
	m := make([][]Elemento, linhas)
	for i := range m {
		m[i] = make([]Elemento, colunas)
	}
	return &Matrix{QuantLinhas: linhas, QuantColunas: colunas, Dados: m}
}

func (m *Matrix) Get(linha, coluna int) (Elemento, error) {
	if linha < 0 || linha >= m.QuantLinhas || coluna < 0 || coluna >= m.QuantColunas {
		return nil, errors.New("out of bounds")
	}
	return m.Dados[linha][coluna], nil
}

func (m *Matrix) Set(linha, coluna int, e Elemento) error {
	if linha < 0 || linha >= m.QuantLinhas || coluna < 0 || coluna >= m.QuantColunas {
		return errors.New("out of bounds")
	}
	m.Dados[linha][coluna] = e
	return nil
}

func (m *Matrix) Add(n *Matrix) (*Matrix, error) {
	if m.QuantLinhas != n.QuantLinhas || m.QuantColunas != n.QuantColunas {
		return nil, errors.New("Dimensões incompativeis")
	}
}
