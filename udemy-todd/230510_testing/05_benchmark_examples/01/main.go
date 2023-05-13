package main

import (
	"fmt"
	"proj/mystr"
	"strings"
)

const s = "O Giganotossauro foi um gênero de dinossauros terópodes que viveu na Argentina durante o período Cretáceo, no estágio Cenomaniano, há aproximadamente 99,6 a 97 milhões de anos, mais especificamente na Formação Candeleros na Patagônia. Seus primeiros fosseis foram encontrados em 1993 no mesmo local, e foi descrito em 1995 como Giganotosaurus carolinii em homenagem ao seu descobridor, Rubén D. Carolini.O Giganotosaurus foi um dos maiores carnívoros terrestres conhecidos, mas o tamanho exato tem sido difícil de determinar devido à incompletude dos restos encontrados até agora. As estimativas para o espécime mais completo variam de um comprimento de 12 a 13 metros (m), um crânio de 1,53 a 1,80 m de comprimento e um peso de 4,2 a 13,8 toneladas (t). O osso dentário que pertencia a um indivíduo supostamente maior foi usado para extrapolar um comprimento de 13,2 m. Alguns pesquisadores consideraram que o animal é maior que o Tiranossauro, que historicamente tem sido considerado o maior terópode, enquanto outros consideram que eles são aproximadamente iguais em tamanho e as maiores estimativas de tamanho para o Giganotosaurus são exageradas. "

func main() {
	xs := strings.Split(s, " ")
	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n", mystr.Join(xs))
}
