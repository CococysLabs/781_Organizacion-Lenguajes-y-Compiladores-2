# Semana 7 – Código de Tres Direcciones (C3D) con un visitor de ANTLR4 y C#

Demostración mínima: generar C3D de expresiones aritméticas (`+ - * /`, paréntesis, números y variables)
recorriendo el árbol con un visitor. No incluye tabla de símbolos, control de flujo, otros tipos,
optimización ni ensamblador.

## Estructura

- `base/` – punto de partida que se **entrega al estudiante en clase**. Compila, pero **no** tiene `C3DVisitor.cs`;
  `Program.cs` solo imprime `Entrada VÁLIDA` y deja un `// TODO`.
- `solucion/` – versión completa con `C3DVisitor.cs`. **La guarda el auxiliar**, no se entrega.

Ambas carpetas contienen el proyecto `MiniCalc`, la gramática `Calc.g4` y la carpeta `pruebas/`.

## Requisitos

- .NET 8 SDK
- Java 11+ (lo usa `Antlr4BuildTasks` para generar el lexer, el parser y el visitor al compilar)

## Cómo ejecutar

```bash
cd solucion          # o base
dotnet build
dotnet run -- pruebas/valida1.txt
```

Sin argumentos imprime `Uso: dotnet run -- <archivo>`.

## Salida esperada (`solucion/`)

| Archivo | Entrada | Salida |
|---|---|---|
| `valida1.txt` | `x = (a + b) * c - d / 2;` | `t1 = a + b`<br>`t2 = t1 * c`<br>`t3 = d / 2`<br>`t4 = t2 - t3`<br>`x = t4` |
| `valida2.txt` | `r = a + b * c;` | `t1 = b * c`<br>`t2 = a + t1`<br>`r = t2` |
| `valida3.txt` | `r = (a + b) * c;`<br>`s = r - 1;` | `t1 = a + b`<br>`t2 = t1 * c`<br>`r = t2`<br>`t3 = r - 1`<br>`s = t3` |
| `invalida1.txt` | `x = (a + b * c;` | `line 1:14 missing ')' at ';'`<br>`Entrada INVÁLIDA (1 errores sintácticos)` |

En `base/` las entradas válidas solo imprimen `Entrada VÁLIDA`.

> El código generado por ANTLR queda en `obj/` y no se sube al repositorio.
