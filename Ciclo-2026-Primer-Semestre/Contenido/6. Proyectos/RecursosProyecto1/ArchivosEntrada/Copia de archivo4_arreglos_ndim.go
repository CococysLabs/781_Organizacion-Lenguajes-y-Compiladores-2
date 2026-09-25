// ====================================================
// ARCHIVO 4 - ARREGLOS MULTIDIMENSIONALES (N-D)
// Grading: Punto 5 (5.6-5.9: 12 pts)
// Cobertura: Matrices 2D, cubos 3D, acceso, modificación, operaciones
// ====================================================

/*
Archivo de prueba para arreglos multidimensionales:
- Matrices 2D: declaración, inicialización
- Acceso y modificación de elementos en 2D
- Iteración con for anidados
- Operaciones matriciales: suma, transposición, escalar
- Cubos 3D: declaración, acceso, modificación
*/

func main() {
	fmt.Println("=== PRUEBAS DE ARREGLOS MULTIDIMENSIONALES ===")

	// 1. MATRICES 2D - DECLARACIÓN NO INICIALIZADA
	var matriz2x3 [2][3]int32
	fmt.Println("Matriz 2x3 no inicializada:")
	fmt.Println("Fila 0:", matriz2x3[0][0], matriz2x3[0][1], matriz2x3[0][2])
	fmt.Println("Fila 1:", matriz2x3[1][0], matriz2x3[1][1], matriz2x3[1][2])

	// 2. MATRICES 2D - DECLARACIÓN INICIALIZADA
	var matrizCuadrada [2][2]int32 = [2][2]int32{
		{1, 2},
		{3, 4},
	}
	fmt.Println("Matriz 2x2 inicializada:")
	fmt.Println(matrizCuadrada[0][0], matrizCuadrada[0][1])
	fmt.Println(matrizCuadrada[1][0], matrizCuadrada[1][1])

	var matrizRectangular [3][4]int32 = [3][4]int32{
		{10, 20, 30, 40},
		{50, 60, 70, 80},
		{90, 100, 110, 120},
	}
	fmt.Println("Matriz 3x4 inicializada:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println("Posición [", i, "][", j, "]:", matrizRectangular[i][j])
		}
	}

	var matrizStrings [2][3]string = [2][3]string{
		{"Ana", "Luis", "María"},
		{"Carlos", "Sofia", "Diego"},
	}
	fmt.Println("Matriz de strings:")
	fmt.Println(matrizStrings[0][0], matrizStrings[0][1], matrizStrings[0][2])
	fmt.Println(matrizStrings[1][0], matrizStrings[1][1], matrizStrings[1][2])

	// 3. ACCESO A ELEMENTOS 2D
	fmt.Println("Acceso a esquinas matriz 3x4:")
	fmt.Println("Superior izquierda:", matrizRectangular[0][0])
	fmt.Println("Superior derecha:", matrizRectangular[0][3])
	fmt.Println("Inferior izquierda:", matrizRectangular[2][0])
	fmt.Println("Inferior derecha:", matrizRectangular[2][3])

	// 4. MODIFICACIÓN DE ELEMENTOS 2D
	matriz2x3[0][0] = 100
	matriz2x3[0][1] = 200
	matriz2x3[1][0] = 300
	fmt.Println("Matriz 2x3 modificada:")
	fmt.Println(matriz2x3[0][0], matriz2x3[0][1], matriz2x3[0][2])
	fmt.Println(matriz2x3[1][0], matriz2x3[1][1], matriz2x3[1][2])

	// 5. ITERACIÓN 2D
	fmt.Println("Iteración completa matriz 3x4:")
	var sumaMatriz int32 = 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println("[", i, "][", j, "]:", matrizRectangular[i][j])
			sumaMatriz += matrizRectangular[i][j]
		}
	}
	fmt.Println("Suma de todos elementos:", sumaMatriz)

	// 6. MÁXIMO EN MATRIZ
	var maximoMatriz int32 = matrizRectangular[0][0]
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			if matrizRectangular[i][j] > maximoMatriz {
				maximoMatriz = matrizRectangular[i][j]
			}
		}
	}
	fmt.Println("Máximo en matriz:", maximoMatriz)

	// 7. OPERACIONES MATRICIALES - LLENADO SECUENCIAL
	var matrizPatron [4][4]int32
	var valor int32 = 1
	fmt.Println("Matriz 4x4 llenada secuencialmente:")
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			matrizPatron[i][j] = valor
			valor++
		}
	}
	for i := 0; i < 4; i++ {
		fmt.Println(matrizPatron[i][0], matrizPatron[i][1], matrizPatron[i][2], matrizPatron[i][3])
	}

	// 8. OPERACIONES MATRICIALES - SUMA DE MATRICES
	var matrizA [2][2]int32 = [2][2]int32{
		{1, 2},
		{3, 4},
	}
	var matrizB [2][2]int32 = [2][2]int32{
		{5, 6},
		{7, 8},
	}
	var matrizSuma [2][2]int32

	fmt.Println("Suma de matrices:")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			matrizSuma[i][j] = matrizA[i][j] + matrizB[i][j]
		}
	}
	fmt.Println(matrizSuma[0][0], matrizSuma[0][1])
	fmt.Println(matrizSuma[1][0], matrizSuma[1][1])

	// 9. OPERACIONES MATRICIALES - MULTIPLICACIÓN POR ESCALAR
	var escalar int32 = 3
	var matrizEscalada [2][2]int32
	fmt.Println("Multiplicación por escalar 3:")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			matrizEscalada[i][j] = matrizA[i][j] * escalar
		}
	}
	fmt.Println(matrizEscalada[0][0], matrizEscalada[0][1])
	fmt.Println(matrizEscalada[1][0], matrizEscalada[1][1])

	// 10. TRANSPOSICIÓN
	var original [2][3]int32 = [2][3]int32{
		{1, 2, 3},
		{4, 5, 6},
	}
	var transpuesta [3][2]int32
	fmt.Println("Matriz original 2x3:")
	fmt.Println(original[0][0], original[0][1], original[0][2])
	fmt.Println(original[1][0], original[1][1], original[1][2])

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			transpuesta[j][i] = original[i][j]
		}
	}
	fmt.Println("Matriz transpuesta 3x2:")
	for i := 0; i < 3; i++ {
		fmt.Println(transpuesta[i][0], transpuesta[i][1])
	}

	// 11. ARREGLOS 3D - CUBO 2x2x2 NO INICIALIZADO
	var cubo [2][2][2]int32
	fmt.Println("Cubo 2x2x2 no inicializado:")
	fmt.Println("Capa 0:")
	fmt.Println(cubo[0][0][0], cubo[0][0][1])
	fmt.Println(cubo[0][1][0], cubo[0][1][1])
	fmt.Println("Capa 1:")
	fmt.Println(cubo[1][0][0], cubo[1][0][1])
	fmt.Println(cubo[1][1][0], cubo[1][1][1])

	// 12. ARREGLOS 3D - CUBO 2x2x2 INICIALIZADO
	var cuboInicializado [2][2][2]int32 = [2][2][2]int32{
		{
			{1, 2},
			{3, 4},
		},
		{
			{5, 6},
			{7, 8},
		},
	}

	fmt.Println("Cubo 2x2x2 inicializado:")
	fmt.Println("Capa 0:")
	fmt.Println(cuboInicializado[0][0][0], cuboInicializado[0][0][1])
	fmt.Println(cuboInicializado[0][1][0], cuboInicializado[0][1][1])
	fmt.Println("Capa 1:")
	fmt.Println(cuboInicializado[1][0][0], cuboInicializado[1][0][1])
	fmt.Println(cuboInicializado[1][1][0], cuboInicializado[1][1][1])

	// 13. ACCESO Y MODIFICACIÓN EN 3D
	fmt.Println("Acceso a elementos del cubo:")
	fmt.Println("Elemento [0][0][0]:", cuboInicializado[0][0][0])
	fmt.Println("Elemento [1][1][1]:", cuboInicializado[1][1][1])

	cubo[0][0][0] = 100
	cubo[0][0][1] = 200
	cubo[0][1][0] = 300
	cubo[0][1][1] = 400
	cubo[1][0][0] = 500
	cubo[1][0][1] = 600
	cubo[1][1][0] = 700
	cubo[1][1][1] = 800

	fmt.Println("Cubo después de modificar:")
	fmt.Println("Capa 0:")
	fmt.Println(cubo[0][0][0], cubo[0][0][1])
	fmt.Println(cubo[0][1][0], cubo[0][1][1])
	fmt.Println("Capa 1:")
	fmt.Println(cubo[1][0][0], cubo[1][0][1])
	fmt.Println(cubo[1][1][0], cubo[1][1][1])

	// 14. ITERACIÓN 3D
	fmt.Println("Iteración completa del cubo:")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				fmt.Println("[", i, "][", j, "][", k, "]:", cubo[i][j][k])
			}
		}
	}

	fmt.Println("=== FIN DE PRUEBAS MULTIDIMENSIONALES ===")
}
