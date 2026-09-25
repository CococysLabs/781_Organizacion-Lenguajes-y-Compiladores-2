// ====================================================
// ARCHIVO 2 - FUNCIONALIDADES INTERMEDIAS
// Grading: Punto 2 (20 pts)
// Cobertura: Scopes, if/else, switch, for, break, continue
// ====================================================

/*
Archivo de prueba para estructuras de control intermedias:
- Manejo de ámbitos (scopes) con bloques anidados
- Sentencia if simple, if-else, if-else if-else
- Sentencia switch con múltiples casos
- For clásico, for condicional, for infinito
- Sentencias break y continue
- Ámbitos complejos con redeclaración
*/

func main() {
	fmt.Println("=== PRUEBAS INTERMEDIAS ===")

	// 1. ÁMBITOS (SCOPES)
	var variableMain int32 = 100
	fmt.Println("Ámbito main:", variableMain)

	{
		var variableBloque int32 = 50
		fmt.Println("Ámbito bloque:", variableBloque, "acceso a main:", variableMain)
	}

	// 2. SENTENCIA IF - SIMPLE, IF-ELSE, IF-ELSE IF-ELSE
	var edad int32 = 20

	if edad >= 18 {
		fmt.Println("Es mayor de edad")
	}

	var temperatura int32 = 30
	if temperatura > 25 {
		fmt.Println("Hace calor")
	} else {
		fmt.Println("Hace frío")
	}

	var puntuacion int32 = 75
	if puntuacion >= 90 {
		fmt.Println("Calificación: A")
	} else {
		if puntuacion >= 80 {
			fmt.Println("Calificación: B")
		} else {
			if puntuacion >= 70 {
				fmt.Println("Calificación: C")
			} else {
				fmt.Println("Calificación: F")
			}
		}
	}

	// Variables locales dentro de if
	if true {
		var localIF int32 = 999
		fmt.Println("Variable local en if:", localIF)
	}

	// 3. SENTENCIA SWITCH/CASE
	var diaSemana int32 = 3
	fmt.Println("Día de la semana:")
	switch diaSemana {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miércoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6, 7:
		fmt.Println("Fin de semana")
	default:
		fmt.Println("Día inválido")
	}

	var mes int32 = 7
	fmt.Println("Estaciones:")
	switch mes {
	case 12, 1, 2:
		fmt.Println("Verano")
	case 3, 4, 5:
		fmt.Println("Otoño")
	case 6, 7, 8:
		fmt.Println("Invierno")
	case 9, 10, 11:
		fmt.Println("Primavera")
	default:
		fmt.Println("Mes inválido")
	}

	// 4. FOR CLÁSICO
	fmt.Println("For clásico - 1 al 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteración:", i)
	}

	fmt.Println("For con decremento:")
	for j := 10; j >= 6; j-- {
		fmt.Println("Valor:", j)
	}

	fmt.Println("For con saltos de 2:")
	for k := 0; k <= 10; k += 2 {
		fmt.Println(k)
	}

	// For anidado
	fmt.Println("For anidado - Tabla de multiplicar:")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			var producto int32 = i * j
			fmt.Println(i, "x", j, "=", producto)
		}
	}

	// 5. FOR CONDICIONAL (equivalente a while)
	fmt.Println("For condicional:")
	var contador int32 = 1
	for contador <= 5 {
		fmt.Println("Contador:", contador)
		contador++
	}

	// 6. FOR INFINITO CON BREAK
	fmt.Println("For infinito con break:")
	var contInf int32 = 0
	for {
		contInf++
		fmt.Println("Iteración:", contInf)
		if contInf >= 3 {
			break
		}
	}

	// 7. BREAK EN BÚSQUEDA
	fmt.Println("Break en búsqueda:")
	for i := 1; i <= 20; i++ {
		if i == 7 {
			fmt.Println("Encontrado 7, saliendo")
			break
		}
		fmt.Println(i)
	}

	// 8. CONTINUE - SALTAR ELEMENTOS
	fmt.Println("Continue - Solo impares:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("Continue - Saltar múltiplos de 3:")
	for i := 1; i <= 15; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// 9. ÁMBITOS ANIDADOS PROFUNDOS
	fmt.Println("Ámbitos anidados profundos:")
	var nivel1 int32 = 1
	{
		var nivel2 int32 = 2
		fmt.Println("Nivel 2 - acceso:", nivel1)
		{
			var nivel3 int32 = 3
			fmt.Println("Nivel 3 - acceso:", nivel1, nivel2)
			{
				var nivel4 int32 = 4
				fmt.Println("Nivel 4 - acceso:", nivel1, nivel2, nivel3, nivel4)
			}
		}
	}

	// 10. REDECLARACIÓN EN ÁMBITOS DIFERENTES
	fmt.Println("Redeclaración en ámbitos:")
	var nombre string = "Global"
	fmt.Println("Ámbito main:", nombre)
	{
		var nombre string = "Bloque1"
		fmt.Println("Ámbito bloque1:", nombre)
		{
			var nombre string = "Bloque2"
			fmt.Println("Ámbito bloque2:", nombre)
		}
	}

	// 11. IF-FOR-SWITCH COMBINADOS
	fmt.Println("If dentro de for:")
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "es par")
		} else {
			fmt.Println(i, "es impar")
		}
	}

	fmt.Println("Switch dentro de for:")
	for dia := 1; dia <= 7; dia++ {
		switch dia {
		case 1, 7:
			fmt.Println("Día", dia, "- Fin de semana")
		case 2, 3, 4, 5, 6:
			fmt.Println("Día", dia, "- Laboral")
		}
	}

	fmt.Println("=== FIN DE PRUEBAS INTERMEDIAS ===")
}
