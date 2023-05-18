package word

import (
	"fmt"
	"testing"
)

var s = "Nancy Sophie Cornélie Corry Tendeloo foi uma advogada, feminista e política neerlandesa que atuou na Câmara dos Deputados pela Liga do Pensamento Livre Democrático (VDB) entre 1945 a 1946, e depois pelo recém-formado Partido do Trabalho (PvdA) até sua morte em 1956. Nascida nas Índias Orientais Neerlandesas, Tendeloo estudou Direito na Universidade de Utreque, período em que fez contato com pessoas do movimento pelos direitos das mulheres. Ela se tornou politicamente ativa na década de 1930 e foi eleita para o Conselho Municipal de Amsterdã pelo VDB em 1938. Ela ajudou a garantir o sufrágio universal para as colônias neerlandesas Suriname e Curaçau em 1948. No início de 1955, ela defendeu com sucesso a igualdade salarial e, mais tarde, naquele ano, apresentou uma moção parlamentar para abolir a proibição do emprego estatal para mulheres casadas."

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(s)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(s)
	}
}

func ExampleUseCount() {
	fmt.Println(UseCount(s))
	// Output:
	// map[(PvdA):1 (VDB):1 1930:1 1938.:1 1945:1 1946,:1 1948.:1 1955,:1 1956.:1 Amsterdã:1 Conselho:1 Cornélie:1 Corry:1 Curaçau:1 Câmara:1 Democrático:1 Deputados:1 Direito:1 Ela:2 Liga:1 Livre:1 Municipal:1 Nancy:1 Nascida:1 Neerlandesas,:1 No:1 Orientais:1 Partido:1 Pensamento:1 Sophie:1 Suriname:1 Tendeloo:2 Trabalho:1 Universidade:1 Utreque,:1 VDB:1 a:4 abolir:1 advogada,:1 ajudou:1 ano,:1 apresentou:1 as:1 ativa:1 atuou:1 até:1 casadas.:1 colônias:1 com:2 contato:1 das:1 de:4 defendeu:1 depois:1 direitos:1 do:4 dos:1 década:1 e:4 e,:1 ela:1 eleita:1 em:4 emprego:1 entre:1 estatal:1 estudou:1 feminista:1 fez:1 foi:2 garantir:1 igualdade:1 início:1 mais:1 morte:1 movimento:1 moção:1 mulheres:1 mulheres.:1 na:3 naquele:1 nas:1 neerlandesa:1 neerlandesas:1 o:2 para:4 parlamentar:1 pela:1 pelo:2 pelos:1 período:1 pessoas:1 politicamente:1 política:1 proibição:1 que:2 recém-formado:1 salarial:1 se:1 sua:1 sucesso:1 sufrágio:1 tarde,:1 tornou:1 uma:2 universal:1 Índias:1]
}

func ExampleCount() {
	fmt.Println(Count(s))
	// Output:
	// 135
}

func compareMaps(map1, map2 map[string]int) bool {
	// Check if the maps have the same length
	if len(map1) != len(map2) {
		return false
	}

	// Iterate over the key-value pairs of the first map
	for key, value := range map1 {
		// Check if the key exists in the second map
		if val, ok := map2[key]; !ok || val != value {
			return false
		}
	}

	return true
}

func TestUseCount(t *testing.T) {
	result := UseCount(s)
	expected := map[string]int{
		"(PvdA)":        1,
		"(VDB)":         1,
		"1930":          1,
		"1938.":         1,
		"1945":          1,
		"1946,":         1,
		"1948.":         1,
		"1955,":         1,
		"1956.":         1,
		"Amsterdã":      1,
		"Conselho":      1,
		"Cornélie":      1,
		"Corry":         1,
		"Curaçau":       1,
		"Câmara":        1,
		"Democrático":   1,
		"Deputados":     1,
		"Direito":       1,
		"Ela":           2,
		"Liga":          1,
		"Livre":         1,
		"Municipal":     1,
		"Nancy":         1,
		"Nascida":       1,
		"Neerlandesas,": 1,
		"No":            1,
		"Orientais":     1,
		"Partido":       1,
		"Pensamento":    1,
		"Sophie":        1,
		"Suriname":      1,
		"Tendeloo":      2,
		"Trabalho":      1,
		"Universidade":  1,
		"Utreque,":      1,
		"VDB":           1,
		"a":             4,
		"abolir":        1,
		"advogada,":     1,
		"ajudou":        1,
		"ano,":          1,
		"apresentou":    1,
		"as":            1,
		"ativa":         1,
		"atuou":         1,
		"até":           1,
		"casadas.":      1,
		"colônias":      1,
		"com":           2,
		"contato":       1,
		"das":           1,
		"de":            4,
		"defendeu":      1,
		"depois":        1,
		"direitos":      1,
		"do":            4,
		"dos":           1,
		"década":        1,
		"e":             4,
		"e,":            1,
		"ela":           1,
		"eleita":        1,
		"em":            4,
		"emprego":       1,
		"entre":         1,
		"estatal":       1,
		"estudou":       1,
		"feminista":     1,
		"fez":           1,
		"foi":           2,
		"garantir":      1,
		"igualdade":     1,
		"início":        1,
		"mais":          1,
		"morte":         1,
		"movimento":     1,
		"moção":         1,
		"mulheres":      1,
		"mulheres.":     1,
		"na":            3,
		"naquele":       1,
		"nas":           1,
		"neerlandesa":   1,
		"neerlandesas":  1,
		"o":             2,
		"para":          4,
		"parlamentar":   1,
		"pela":          1,
		"pelo":          2,
		"pelos":         1,
		"período":       1,
		"pessoas":       1,
		"politicamente": 1,
		"política":      1,
		"proibição":     1,
		"que":           2,
		"recém-formado": 1,
		"salarial":      1,
		"se":            1,
		"sua":           1,
		"sucesso":       1,
		"sufrágio":      1,
		"tarde,":        1,
		"tornou":        1,
		"uma":           2,
		"universal":     1,
		"Índias":        1,
	}
	if !compareMaps(result, expected) {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestCount(t *testing.T) {
	result := Count(s)
	expected := 135
	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}
