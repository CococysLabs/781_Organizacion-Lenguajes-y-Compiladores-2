// ====================================================
// ARCHIVO 6 - PRUEBAS AVANZADAS E INTEGRACIÓN
// Grading: Integración de todas las características
// Cobertura: Combinación compleja de todas las funcionalidades
// ====================================================

/*
Archivo de integración que combina todas las características:
- Funciones con arreglos multidimensionales
- Recursividad combinada con iteración
- Estructuras de control anidadas
- Funciones embebidas
- Casos límite
*/

// ========================================
// FUNCIONES AUXILIARES
// ========================================

func procesarMatriz(m *[3][3]int32, factor int32) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			m[i][j] = m[i][j] * factor
		}
	}
}

func buscarEnMatriz(m [4][4]int32, valor int32) (int32, int32, bool) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if m[i][j] == valor {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func esPrimo(n int32) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func generarFibonacci(arr *[10]int32) {
	arr[0] = 0
	arr[1] = 1
	for i := 2; i < 10; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
}

func contarVocales(texto string) int32 {
	var contador int32 = 0
	for i := 0; i < len(texto); i++ {
		var letra rune = rune(texto[i])
		if letra == 'a' || letra == 'e' || letra == 'i' || letra == 'o' || letra == 'u' {
			contador++
		}
		if letra == 'A' || letra == 'E' || letra == 'I' || letra == 'O' || letra == 'U' {
			contador++
		}
	}
	return contador
}

func sumaRango(inicio int32, fin int32) int32 {
	if inicio > fin {
		return 0
	}
	if inicio == fin {
		return inicio
	}
	return inicio + sumaRango(inicio+1, fin)
}

func encontrarKesimoMenor(arr [7]int32, k int32) int32 {
	var temp [7]int32
	for i := 0; i < 7; i++ {
		temp[i] = arr[i]
	}

	for i := 0; i < 7; i++ {
		for j := 0; j < 6; j++ {
			if temp[j] > temp[j+1] {
				var aux int32 = temp[j]
				temp[j] = temp[j+1]
				temp[j+1] = aux
			}
		}
	}

	if k >= 0 && k < 7 {
		return temp[k]
	}
	return -1
}

func validarRango(arr [5]int32, min int32, max int32) bool {
	for i := 0; i < 5; i++ {
		if arr[i] < min || arr[i] > max {
			return false
		}
	}
	return true
}

// ========================================
// FUNCIÓN MAIN
// ========================================

func main() {
	fmt.Println("=== PRUEBAS AVANZADAS E INTEGRACIÓN ===")

	// Escenario 1: Gestión de datos con estructuras complejas
	fmt.Println("\n--- Gestión de Calificaciones ---")
	var estudiantes [5]string = [5]string{"Ana", "Luis", "María", "Carlos", "Sofia"}
	var calificaciones [5]int32 = [5]int32{85, 92, 78, 95, 88}

	for i := 0; i < 5; i++ {
		fmt.Println(estudiantes[i], "-", calificaciones[i])
		if calificaciones[i] >= 90 {
			fmt.Println("  Excelente")
		} else {
			if calificaciones[i] >= 80 {
				fmt.Println("  Muy bueno")
			} else {
				fmt.Println("  Bueno")
			}
		}
	}

	var mejorCalif int32 = calificaciones[0]
	var indiceMejor int32 = 0
	for i := 1; i < 5; i++ {
		if calificaciones[i] > mejorCalif {
			mejorCalif = calificaciones[i]
			indiceMejor = i
		}
	}
	fmt.Println("Mejor: ", estudiantes[indiceMejor], mejorCalif)

	// Escenario 2: Análisis matricial
	fmt.Println("\n--- Análisis de Matrices ---")
	var datos [3][3]int32 = [3][3]int32{
		{12, 15, 18},
		{22, 25, 28},
		{32, 35, 38},
	}

	fmt.Println("Matriz original:")
	for i := 0; i < 3; i++ {
		fmt.Println(datos[i][0], datos[i][1], datos[i][2])
	}

	procesarMatriz(&datos, 2)
	fmt.Println("Después de multiplicar por 2:")
	for i := 0; i < 3; i++ {
		fmt.Println(datos[i][0], datos[i][1], datos[i][2])
	}

	// Escenario 3: Búsqueda en estructuras 2D
	fmt.Println("\n--- Búsqueda Matricial ---")
	var tabla [4][4]int32 = [4][4]int32{
		{101, 102, 103, 104},
		{201, 202, 203, 204},
		{301, 302, 303, 304},
		{401, 402, 403, 404},
	}

	var fila, col, encon = buscarEnMatriz(tabla, 303)
	if encon {
		fmt.Println("303 encontrado en [", fila, "][", col, "]")
	}

	// Escenario 4: Números primos y Fibonacci
	fmt.Println("\n--- Series Numéricas ---")
	var primosCont int32 = 0
	var primosEnc int32 = 0
	for n := 2; primosCont < 5; n++ {
		if esPrimo(n) {
			fmt.Println("Primo:", n)
			primosEnc++
			primosCont++
		}
	}

	var fib [10]int32
	generarFibonacci(&fib)
	fmt.Println("Fibonacci:", fib[0], fib[1], fib[2], fib[3], fib[4], fib[5])

	// Escenario 5: Procesamiento de texto
	fmt.Println("\n--- Análisis de Texto ---")
	var textos [3]string = [3]string{
		"Organizacion de Lenguajes",
		"Compiladores",
		"Analisis Semantico",
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Texto", i, ":", textos[i])
		fmt.Println("  Longitud:", len(textos[i]))
		fmt.Println("  Vocales:", contarVocales(textos[i]))
		fmt.Println("  Tipo:", typeOf(textos[i]))

		if len(textos[i]) >= 5 {
			var sub string = substr(textos[i], 0, 5)
			fmt.Println("  Primeros 5:", sub)
		}
	}

	// Escenario 6: Operaciones matemáticas recursivas
	fmt.Println("\n--- Matemáticas Recursivas ---")
	var suma1 int32 = sumaRango(1, 100)
	fmt.Println("Suma 1..100:", suma1)

	var suma2 int32 = sumaRango(10, 50)
	fmt.Println("Suma 10..50:", suma2)

	var datos2 [7]int32 = [7]int32{45, 12, 67, 23, 89, 34, 56}
	fmt.Println("Datos:", datos2[0], datos2[1], datos2[2], datos2[3], datos2[4], datos2[5], datos2[6])
	fmt.Println("3er menor:", encontrarKesimoMenor(datos2, 2))

	// Escenario 7: Validación y control de flujo complejo
	fmt.Println("\n--- Validación y Control ---")
	var entrada [5]int32 = [5]int32{15, 22, 35, 48, 56}

	if validarRango(entrada, 10, 60) {
		fmt.Println("Rango [10,60]: OK")
	}
	if !validarRango(entrada, 20, 40) {
		fmt.Println("Rango [20,40]: FUERA")
	}

	// Escenario 8: Ámbitos complejos
	fmt.Println("\n--- Ámbitos Anidados ---")
	var nivel1 int32 = 100
	{
		var nivel2 int32 = 200
		fmt.Println("Nivel 1:", nivel1, "Nivel 2:", nivel2)
		{
			var nivel3 int32 = 300
			fmt.Println("Nivel 1:", nivel1, "Nivel 2:", nivel2, "Nivel 3:", nivel3)
			for i := 0; i < 2; i++ {
				var nivel4 int32 = 400 + i
				fmt.Println("Nivel 4 (i=", i, "):", nivel4)
			}
		}
	}

	// Escenario 9: Funciones embebidas integradas
	fmt.Println("\n--- Funciones Embebidas ---")
	var texto string = "Golampi 2026"
	fmt.Println("Texto:", texto)
	fmt.Println("Longitud:", len(texto))
	fmt.Println("Subcadena:", substr(texto, 0, 7))
	fmt.Println("Hora:", now())
	fmt.Println("Tipo:", typeOf(texto))

	// Escenario 10: Casos límite
	fmt.Println("\n--- Casos Límite ---")
	var unitario [1]int32 = [1]int32{42}
	fmt.Println("Array unitario:", unitario[0])

	var vacio string = ""
	fmt.Println("String vacío longitud:", len(vacio))

	var minInt int32 = -2147483648
	var maxInt int32 = 2147483647
	fmt.Println("Rango int32:", minInt, "a", maxInt)

	if maxInt > minInt {
		fmt.Println("Comparación: OK")
	}

	fmt.Println("\n=== FIN DE INTEGRACIÓN ===")
}
