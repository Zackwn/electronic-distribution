package core

type Sublevel struct {
	maxElectrons   Electrons
	layer          int
	representation rune
}

var sublevelsSequence = [...]Sublevel{
	{layer: 1, maxElectrons: 2, representation: 's'},
	{layer: 2, maxElectrons: 2, representation: 's'},
	{layer: 2, maxElectrons: 6, representation: 'p'},
	{layer: 3, maxElectrons: 2, representation: 's'},
	{layer: 3, maxElectrons: 6, representation: 'p'},
	{layer: 4, maxElectrons: 2, representation: 's'},
	{layer: 3, maxElectrons: 10, representation: 'd'},
	{layer: 4, maxElectrons: 6, representation: 'p'},
	{layer: 5, maxElectrons: 2, representation: 's'},
	{layer: 4, maxElectrons: 10, representation: 'd'},
	{layer: 5, maxElectrons: 6, representation: 'p'},
	{layer: 6, maxElectrons: 2, representation: 's'},
	{layer: 4, maxElectrons: 14, representation: 'f'},
	{layer: 5, maxElectrons: 10, representation: 'd'},
	{layer: 6, maxElectrons: 6, representation: 'p'},
	{layer: 7, maxElectrons: 2, representation: 's'},
	{layer: 5, maxElectrons: 14, representation: 'f'},
	{layer: 6, maxElectrons: 10, representation: 'd'},
	{layer: 7, maxElectrons: 6, representation: 'p'},
}
