// ====================================================
// ARCHIVO 1 - FUNCIONALIDADES BÁSICAS
// Grading: Punto 1 (26 pts)
// Cobertura: Variables, Tipos, Operadores, Comentarios, Constantes
// ====================================================

/*
Archivo de prueba para funcionalidades básicas:
- Comentarios de una línea y multilínea
- Declaración de variables (estática, corta, múltiple)
- Constantes con tipos estáticos
- Tipos: int32, float32, bool, rune, string
- Identificadores case-sensitive
- Operadores: aritméticos, de asignación, relacionales, lógicos
- Corto circuito, manejo de nil, precedencia
*/

func main() {
	fmt.Println("=== PRUEBAS BÁSICAS ===")

	// Comentarios: de una línea (arriba) y multilínea (abajo)
	/* Este es comentario
	   multilínea en bloque */

	// 1. DECLARACIÓN ESTÁTICA CON Y SIN INICIALIZACIÓN
	var enteroConValor int32 = 42
	var flotanteConValor float32 = 3.14
	var booleanoConValor bool = true
	var runaConValor rune = 'G'
	var cadenaConValor string = "Golampi"

	var enteroPorDefecto int32
	var flotantePorDefecto float32
	var booleanoPorDefecto bool

	fmt.Println("Valores inicializados:", enteroConValor, flotanteConValor, booleanoConValor, runaConValor, cadenaConValor)
	fmt.Println("Valores por defecto:", enteroPorDefecto, flotantePorDefecto, booleanoPorDefecto)

	// 2. DECLARACIÓN MÚLTIPLE
	var x, y, z int32 = 10, 20, 30
	var a, b float32 = 1.5, 2.5
	fmt.Println("Múltiples:", x, y, z, a, b)

	// 3. DECLARACIÓN CORTA E INFERENCIA DE TIPO
	numeroInferido := 100
	precioInferido := 99.99
	mensaje := "Inferencia"
	p, q, r := 1, 2, 3

	fmt.Println("Corta:", numeroInferido, precioInferido, mensaje, p, q, r)

	// 4. CONSTANTES
	const PI float32 = 3.14159
	const MAX int32 = 1000
	const NOMBRE string = "Golampi"
	const ACTIVO bool = true

	fmt.Println("Constantes:", PI, MAX, NOMBRE, ACTIVO)

	// 5. TIPOS ESTÁTICOS (todos los tipos)
	var t1 int32 = 2147483647
	var t2 float32 = 3.4e38
	var t3 bool = true
	var t4 rune = 'Z'
	var t5 string = "Texto"

	fmt.Println("Tipos:", t1, t2, t3, t4, t5)

	// 6. CASE SENSITIVE
	var variable int32 = 1
	var Variable int32 = 2
	var VARIABLE int32 = 3

	fmt.Println("Case sensitive:", variable, Variable, VARIABLE)

	// 7. ASIGNACIÓN Y OPERADORES DE ASIGNACIÓN COMPUESTA
	var val int32 = 10
	val = 50

	var n1 int32 = 10
	n1 += 5
	fmt.Println("+=:", n1)

	var n2 int32 = 20
	n2 -= 7
	fmt.Println("-=:", n2)

	var n3 int32 = 6
	n3 *= 4
	fmt.Println("*=:", n3)

	var n4 int32 = 40
	n4 /= 5
	fmt.Println("/=:", n4)

	// 8. OPERADORES ARITMÉTICOS
	fmt.Println("+:", 15+25)
	fmt.Println("-:", 50-18)
	fmt.Println("*:", 7*8)
	fmt.Println("/:", 100/3)
	fmt.Println("%:", 17%5)
	fmt.Println("Unarios:", +42, -42)
	fmt.Println("Precedencia:", 2+3*4, "(2+3)*4 =", (2+3)*4)

	// 9. OPERADORES RELACIONALES
	var v1 int32 = 10
	var v2 int32 = 20
	fmt.Println("==:", v1 == v2)
	fmt.Println("!=:", v1 != v2)
	fmt.Println("<:", v1 < v2)
	fmt.Println(">:", v1 > v2)
	fmt.Println("<=:", v1 <= v2)
	fmt.Println(">=:", v1 >= v2)

	// 10. OPERADORES LÓGICOS
	var verd bool = true
	var falso bool = false
	fmt.Println("true && true:", verd && verd)
	fmt.Println("true && false:", verd && falso)
	fmt.Println("true || false:", verd || falso)
	fmt.Println("false || false:", falso || falso)
	fmt.Println("!true:", !verd)
	fmt.Println("!false:", !falso)
	fmt.Println("Complejas:", (10 > 5) && (20 < 30), (10 > 15) || (20 < 30), !(10 == 10))

	// 11. CORTO CIRCUITO
	var cc1 bool = false && true
	var cc2 bool = true || false
	fmt.Println("Corto circuito AND (false):", cc1)
	fmt.Println("Corto circuito OR (true):", cc2)

	// 12. CASOS EXTREMOS
	var maxInt32 int32 = 2147483647
	var minInt32 int32 = -2147483648
	fmt.Println("Extremos int32:", maxInt32, minInt32)
	fmt.Println("100 % 1:", 100%1)
	fmt.Println("String vacío: ''")

	// 13. MANEJO DE nil (representado como ausencia de valor)
	fmt.Println("nil representa ausencia de valor en Golampi")

	fmt.Println("=== FIN DE PRUEBAS BÁSICAS ===")
}
