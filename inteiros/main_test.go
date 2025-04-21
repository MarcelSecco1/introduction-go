package main

import "testing"

func TestAdiciona(t *testing.T){
	result := Adiciona(2, 3)
	if result != 5 {
		t.Errorf("Soma(2, 3) = %d; want 5", result)
	}
}

func TestSubtrai(t *testing.T){
	result := Subtrai(5, 3)
	if result != 2 {
		t.Errorf("Subtracao(5, 3) = %d; want 2", result)
	}
}
