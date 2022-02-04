// https://www.youtube.com/watch?v=KOhQCm8f8AE&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=74&ab_channel=AprendaGo

package main

import "fmt"

func main() {
	ufslice := make([]string, 26, 26)

	ufslice = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará",
		"Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul",
		"Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí",
		"Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia",
		"Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(len(ufslice))
	fmt.Println(cap(ufslice))

	for i := 0; i < len(ufslice); i++ {
		fmt.Println(ufslice[i])
	}

}
