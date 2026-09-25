// ====================================================
// ARCHIVO 3 - ARREGLOS UNIDIMENSIONALES (1D)
// Grading: Punto 5 (5.1-5.5: 10 pts)
// Cobertura: Declaración, acceso, modificación, iteración de arreglos 1D
// ====================================================

/*
Archivo de prueba para arreglos unidimensionales:
- Declaración no inicializada (valores por defecto)
- Declaración inicializada
- Arreglos de tipos: int32, float32, bool, rune, string
- Acceso y modificación de elementos
- Función len()
- Iteración sobre arreglos
- Operaciones: suma, búsqueda, máximo, mínimo
*/

func main() {
	fmt.Println("=== PRUEBAS DE ARREGLOS 1D ===")

	// 1. DECLARACIÓN NO INICIALIZADA
	var numerosVacios [5]int32
	fmt.Println("Arreglo int32 no inicializado [0,0,0,0,0]:")
	fmt.Println(numerosVacios[0], numerosVacios[1], numerosVacios[2], numerosVacios[3], numerosVacios[4])

	var flotantesVacios [3]float32
	fmt.Println("Arreglo float32 no inicializado [0.0,0.0,0.0]:")
	fmt.Println(flotantesVacios[0], flotantesVacios[1], flotantesVacios[2])

	var booleanosVacios [3]bool
	fmt.Println("Arreglo bool no inicializado [false,false,false]:")
	fmt.Println(booleanosVacios[0], booleanosVacios[1], booleanosVacios[2])

	// 2. DECLARACIÓN INICIALIZADA
	var numeros [5]int32 = [5]int32{10, 20, 30, 40, 50}
	fmt.Println("Arreglo int32 inicializado:")
	fmt.Println(numeros[0], numeros[1], numeros[2], numeros[3], numeros[4])

	var precios [4]float32 = [4]float32{19.99, 29.99, 39.99, 49.99}
	fmt.Println("Arreglo float32 inicializado:")
	fmt.Println(precios[0], precios[1], precios[2], precios[3])

	var flags [3]bool = [3]bool{true, false, true}
	fmt.Println("Arreglo bool inicializado:")
	fmt.Println(flags[0], flags[1], flags[2])

	var letras [5]rune = [5]rune{'A', 'B', 'C', 'D', 'E'}
	fmt.Println("Arreglo rune inicializado:")
	fmt.Println(letras[0], letras[1], letras[2], letras[3], letras[4])

	var nombres [4]string = [4]string{"Ana", "Luis", "María", "Carlos"}
	fmt.Println("Arreglo string inicializado:")
	fmt.Println(nombres[0], nombres[1], nombres[2], nombres[3])

	// 3. ACCESO A ELEMENTOS
	fmt.Println("Primer elemento:", numeros[0])
	fmt.Println("Último elemento:", numeros[4])
	fmt.Println("Elemento central:", numeros[2])

	var suma int32 = numeros[0] + numeros[1]
	fmt.Println("Suma de primeros dos:", suma)

	// 4. MODIFICACIÓN DE ELEMENTOS
	numeros[0] = 100
	fmt.Println("Después de modificar numeros[0]:", numeros[0])

	numeros[4] = 500
	fmt.Println("Después de modificar numeros[4]:", numeros[4])

	numeros[1] = numeros[0] + 50
	fmt.Println("Modificado con expresión:", numeros[1])

	// 5. FUNCIÓN len()
	var longitudNumeros int32 = len(numeros)
	fmt.Println("Longitud del arreglo:", longitudNumeros)

	// 6. ITERACIÓN CON FOR
	fmt.Println("Iteración sobre arreglo:")
	for i := 0; i < len(numeros); i++ {
		fmt.Println("Índice", i, ":", numeros[i])
	}

	// 7. OPERACIONES: SUMA
	fmt.Println("Sumar todos los elementos:")
	var sumaTotal int32 = 0
	for i := 0; i < len(numeros); i++ {
		sumaTotal += numeros[i]
	}
	fmt.Println("Suma total:", sumaTotal)

	// 8. OPERACIONES: BUSCAR MÁXIMO
	fmt.Println("Encontrar máximo:")
	var maximo int32 = numeros[0]
	for i := 1; i < len(numeros); i++ {
		if numeros[i] > maximo {
			maximo = numeros[i]
		}
	}
	fmt.Println("Máximo:", maximo)

	// 9. OPERACIONES: BUSCAR MÍNIMO
	fmt.Println("Encontrar mínimo:")
	var minimo int32 = numeros[0]
	for i := 1; i < len(numeros); i++ {
		if numeros[i] < minimo {
			minimo = numeros[i]
		}
	}
	fmt.Println("Mínimo:", minimo)

	// 10. OPERACIONES: CONTAR ELEMENTOS
	fmt.Println("Contar elementos > 100:")
	var contador int32 = 0
	for i := 0; i < len(numeros); i++ {
		if numeros[i] > 100 {
			contador++
		}
	}
	fmt.Println("Cantidad:", contador)

	// 11. OPERACIONES: COPIAR ARREGLO
	var copia [5]int32
	for i := 0; i < len(numeros); i++ {
		copia[i] = numeros[i]
	}
	fmt.Println("Copia:", copia[0], copia[1], copia[2], copia[3], copia[4])

	// 12. OPERACIONES: INVERTIR
	var invertido [5]int32
	for i := 0; i < len(numeros); i++ {
		invertido[i] = numeros[len(numeros)-1-i]
	}
	fmt.Println("Invertido:", invertido[0], invertido[1], invertido[2], invertido[3], invertido[4])

	// 13. OPERACIONES: ORDENAMIENTO BURBUJA
	var paraOrdenar [5]int32 = [5]int32{50, 20, 40, 10, 30}
	fmt.Println("Antes de ordenar:", paraOrdenar[0], paraOrdenar[1], paraOrdenar[2], paraOrdenar[3], paraOrdenar[4])

	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if paraOrdenar[j] > paraOrdenar[j+1] {
				var temp int32 = paraOrdenar[j]
				paraOrdenar[j] = paraOrdenar[j+1]
				paraOrdenar[j+1] = temp
			}
		}
	}
	fmt.Println("Después de ordenar:", paraOrdenar[0], paraOrdenar[1], paraOrdenar[2], paraOrdenar[3], paraOrdenar[4])

	// 14. PROMEDIO DE ARREGLO
	var sumaFloat float32 = 0.0
	for i := 0; i < len(precios); i++ {
		sumaFloat += precios[i]
	}
	var promedio float32 = sumaFloat / float32(len(precios))
	fmt.Println("Promedio de precios:", promedio)

	fmt.Println("=== FIN DE PRUEBAS DE ARREGLOS 1D ===")
}
