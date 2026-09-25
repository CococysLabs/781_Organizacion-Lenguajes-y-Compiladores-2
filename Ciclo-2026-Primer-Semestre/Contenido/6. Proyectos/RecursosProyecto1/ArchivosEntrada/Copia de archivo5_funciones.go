// ====================================================
// ARCHIVO 5 - FUNCIONES
// Grading: Punto 3 (17 pts) + Punto 4 - Built-in functions (5 pts)
// Cobertura: Todas las variantes de funciones y funciones embebidas
// ====================================================

/*
Archivo de prueba para funciones:
- Funciones sin parámetros
- Funciones con parámetros
- Funciones con retorno simple y múltiples retornos
- Funciones con parámetros por referencia (punteros)
- Funciones recursivas
- 9 funciones requeridas
- Funciones embebidas: fmt.Println, len, now, substr, typeOf
- Hoisting de funciones
*/

// ========================================
// FUNCIONES SIN PARÁMETROS
// ========================================

func mostrarBienvenida() {
	fmt.Println("=== Bienvenido a Pruebas de Funciones ===")
}

func obtenerFechaActual() {
	var fecha string = now()
	fmt.Println("Fecha actual:", fecha)
}

// ========================================
// FUNCIONES CON PARÁMETROS
// ========================================

func saludarPersona(nombre string) {
	fmt.Println("Hola,", nombre)
}

func sumarDosNumeros(a int32, b int32) {
	var resultado int32 = a + b
	fmt.Println("Suma:", resultado)
}

// ========================================
// FUNCIONES CON RETORNO SIMPLE
// ========================================

func multiplicar(x int32, y int32) int32 {
	return x * y
}

func esMayorDeEdad(edad int32) bool {
	if edad >= 18 {
		return true
	}
	return false
}

func concatenar(str1 string, str2 string) string {
	return str1 + str2
}

// ========================================
// FUNCIONES CON MÚLTIPLES RETORNOS
// ========================================

func dividirConValidacion(a int32, b int32) (int32, int32, bool) {
	if b == 0 {
		return 0, 0, false
	}
	return a / b, a % b, true
}

func obtenerMinMax(a int32, b int32, c int32) (int32, int32) {
	var minimo int32 = a
	var maximo int32 = a

	if b < minimo {
		minimo = b
	}
	if c < minimo {
		minimo = c
	}
	if b > maximo {
		maximo = b
	}
	if c > maximo {
		maximo = c
	}

	return minimo, maximo
}

// ========================================
// FUNCIONES CON PARÁMETROS POR REFERENCIA
// ========================================

func duplicarValor(num *int32) {
	*num = *num * 2
}

func intercambiarValores(a *int32, b *int32) {
	var temp int32 = *a
	*a = *b
	*b = temp
}

func incrementarArreglo(arr *[5]int32) {
	for i := 0; i < 5; i++ {
		arr[i] = arr[i] + 1
	}
}

// ========================================
// FUNCIONES RECURSIVAS
// ========================================

func factorial(n int32) int32 {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int32) int32 {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func potencia(base int32, exponente int32) int32 {
	if exponente == 0 {
		return 1
	}
	if exponente == 1 {
		return base
	}
	return base * potencia(base, exponente-1)
}

// ========================================
// FUNCIONES REQUERIDAS (1-9)
// ========================================

// 1. imprimirArbol
func imprimirArbol(altura int32) {
	fmt.Println("Árbol de altura:", altura)
	for i := 0; i < altura; i++ {
		for j := 0; j < altura-i-1; j++ {
			fmt.Println(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Println("*")
		}
		fmt.Println("")
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < altura-1; j++ {
			fmt.Println(" ")
		}
		fmt.Println("|")
	}
}

// 2. calcularVolumenPiramide
func calcularVolumenPiramide(base float32, altura float32) float32 {
	return (base * base * altura) / 3.0
}

// 3. ordenamientoSeleccion
func ordenamientoSeleccion(arr *[5]int32) {
	for i := 0; i < 4; i++ {
		var minIdx int32 = i
		for j := i + 1; j < 5; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			var temp int32 = arr[i]
			arr[i] = arr[minIdx]
			arr[minIdx] = temp
		}
	}
}

// 4. intercambioValores
func intercambioValores(x *int32, y *int32) {
	var temporal int32 = *x
	*x = *y
	*y = temporal
}

// 5. GCD - Greatest Common Divisor
func GCD(a int32, b int32) (int32, int32) {
	if b == 0 {
		return a, 1
	}
	var result, steps = GCD(b, a%b)
	return result, steps + 1
}

// 6. RowInstabilityIndex
func RowInstabilityIndex(m [3][4]float32) float32 {
	var totalInstability float32 = 0.0
	for i := 0; i < 3; i++ {
		var rowSum float32 = 0.0
		for j := 1; j < 4; j++ {
			var diff float32 = m[i][j] - m[i][j-1]
			if diff < 0.0 {
				diff = -diff
			}
			rowSum += diff
		}
		totalInstability += rowSum / 3.0
	}
	return totalInstability / 3.0
}

// 7. Solve2x2 - Resuelve sistema linear 2x2
func Solve2x2(A [2][2]float32, B [2]float32) (float32, float32, bool) {
	var det float32 = A[0][0]*A[1][1] - A[0][1]*A[1][0]
	if det == 0.0 {
		return 0.0, 0.0, false
	}

	var x float32 = (B[0]*A[1][1] - A[0][1]*B[1]) / det
	var y float32 = (A[0][0]*B[1] - B[0]*A[1][0]) / det
	return x, y, true
}

// 8. AverageLayerMean
func AverageLayerMean(cube [2][2][2]float32) float32 {
	var totalLayerMean float32 = 0.0
	for z := 0; z < 2; z++ {
		var layerSum float32 = 0.0
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				layerSum += cube[z][i][j]
			}
		}
		totalLayerMean += layerSum / 4.0
	}
	return totalLayerMean / 2.0
}

// 9. SoftmaxRows
func SoftmaxRows(m [3][3]float32) [3][3]float32 {
	var result [3][3]float32
	for i := 0; i < 3; i++ {
		var maxVal float32 = m[i][0]
		for j := 1; j < 3; j++ {
			if m[i][j] > maxVal {
				maxVal = m[i][j]
			}
		}
		var sumExp float32 = 0.0
		var expVals [3]float32
		for j := 0; j < 3; j++ {
			var expVal float32 = 1.0 + (m[i][j] - maxVal)
			if expVal < 0.1 {
				expVal = 0.1
			}
			expVals[j] = expVal
			sumExp += expVal
		}
		for j := 0; j < 3; j++ {
			result[i][j] = expVals[j] / sumExp
		}
	}
	return result
}

// ========================================
// FUNCIÓN MAIN
// ========================================

func main() {
	fmt.Println("=== PRUEBAS DE FUNCIONES ===")

	// Funciones sin parámetros
	mostrarBienvenida()
	obtenerFechaActual()

	// Funciones con parámetros
	saludarPersona("Ana")
	sumarDosNumeros(10, 20)

	// Funciones con retorno simple
	fmt.Println("Multiplicación:", multiplicar(7, 8))
	fmt.Println("Mayor de edad (20):", esMayorDeEdad(20))
	fmt.Println("Concatenación:", concatenar("Hola ", "Mundo"))

	// Funciones con múltiples retornos
	var coc, res, val = dividirConValidacion(17, 5)
	if val {
		fmt.Println("17 / 5 = cociente:", coc, "residuo:", res)
	}

	var min, max = obtenerMinMax(15, 42, 8)
	fmt.Println("De 15, 42, 8 - mín:", min, "máx:", max)

	// Parámetros por referencia
	var numero int32 = 25
	duplicarValor(&numero)
	fmt.Println("Después de duplicar:", numero)

	var x int32 = 10
	var y int32 = 20
	intercambiarValores(&x, &y)
	fmt.Println("Intercambiados - x:", x, "y:", y)

	var arr [5]int32 = [5]int32{1, 2, 3, 4, 5}
	incrementarArreglo(&arr)
	fmt.Println("Incrementado:", arr[0], arr[1], arr[2], arr[3], arr[4])

	// Funciones recursivas
	fmt.Println("Factorial(5):", factorial(5))
	fmt.Println("Fibonacci(7):", fibonacci(7))
	fmt.Println("2^8:", potencia(2, 8))

	// Funciones requeridas 1-4
	fmt.Println("--- Función 1: imprimirArbol ---")
	imprimirArbol(3)

	fmt.Println("--- Función 2: calcularVolumenPiramide ---")
	fmt.Println("Volumen(base=10, altura=15):", calcularVolumenPiramide(10.0, 15.0))

	fmt.Println("--- Función 3: ordenamientoSeleccion ---")
	var desordenado [5]int32 = [5]int32{64, 25, 12, 22, 11}
	ordenamientoSeleccion(&desordenado)
	fmt.Println("Ordenado:", desordenado[0], desordenado[1], desordenado[2], desordenado[3], desordenado[4])

	fmt.Println("--- Función 4: intercambioValores ---")
	var v1 int32 = 100
	var v2 int32 = 200
	intercambioValores(&v1, &v2)
	fmt.Println("Intercambiados:", v1, v2)

	// Funciones requeridas 5-9
	fmt.Println("--- Función 5: GCD ---")
	var gcd, pasos = GCD(48, 18)
	fmt.Println("GCD(48, 18) =", gcd, "pasos:", pasos)

	fmt.Println("--- Función 6: RowInstabilityIndex ---")
	var m [3][4]float32 = [3][4]float32{
		{1.0, 3.0, 2.0, 5.0},
		{2.0, 2.0, 4.0, 4.0},
		{1.0, 5.0, 3.0, 7.0},
	}
	fmt.Println("Índice inestabilidad:", RowInstabilityIndex(m))

	fmt.Println("--- Función 7: Solve2x2 ---")
	var A [2][2]float32 = [2][2]float32{{2.0, 1.0}, {1.0, 3.0}}
	var B [2]float32 = [2]float32{5.0, 6.0}
	var sx, sy, ok = Solve2x2(A, B)
	if ok {
		fmt.Println("Solución - x:", sx, "y:", sy)
	}

	fmt.Println("--- Función 8: AverageLayerMean ---")
	var cubo [2][2][2]float32 = [2][2][2]float32{
		{{1.0, 2.0}, {3.0, 4.0}},
		{{5.0, 6.0}, {7.0, 8.0}},
	}
	fmt.Println("Promedio capas:", AverageLayerMean(cubo))

	fmt.Println("--- Función 9: SoftmaxRows ---")
	var mat [3][3]float32 = [3][3]float32{
		{1.0, 2.0, 3.0},
		{4.0, 5.0, 6.0},
		{7.0, 8.0, 9.0},
	}
	var resultado [3][3]float32 = SoftmaxRows(mat)
	for i := 0; i < 3; i++ {
		fmt.Println("Fila", i, ":", resultado[i][0], resultado[i][1], resultado[i][2])
	}

	// Funciones embebidas
	fmt.Println("--- Funciones embebidas ---")

	var texto string = "Golampi"
	fmt.Println("len('Golampi'):", len(texto))

	var arr2 [7]int32 = [7]int32{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("len(arreglo):", len(arr2))

	fmt.Println("now():", now())

	var subcadena string = substr("Organizacion de Lenguajes", 0, 12)
	fmt.Println("substr():", subcadena)

	fmt.Println("typeOf(42):", typeOf(int32(42)))
	fmt.Println("typeOf(3.14):", typeOf(float32(3.14)))
	fmt.Println("typeOf(true):", typeOf(true))
	fmt.Println("typeOf('texto'):", typeOf("texto"))

	// Hoisting
	fmt.Println("--- Hoisting ---")
	funcionDefinidaDespues()

	fmt.Println("=== FIN DE PRUEBAS DE FUNCIONES ===")
}

// Esta función se define después de main (hoisting)
func funcionDefinidaDespues() {
	fmt.Println("Ejecutada mediante hoisting")
}