# Semana 10 – Heap, arreglos multidimensionales (row-major) y punteros en la pila (C3D con ANTLR4 y C#)

Continuación de la semana 7: mismo estilo y mismas expresiones, ahora con arreglos. El visitor genera C3D que:

- reserva memoria en el **heap** (`H = H + total`),
- guarda el **puntero base** de cada arreglo en la **pila** (`stack[SP + slot] = base`),
- calcula la dirección de `a[i][j]...` en **row-major** con el método de Horner (`(i * d1 + j) * d2 + k ...`).

No incluye strings, optimización, funciones ni control de flujo.

## Estructura

- `base/` – punto de partida que se **entrega al estudiante en clase**. Compila, pero **no** tiene `C3DVisitor.cs`;
  `Program.cs` solo imprime `Entrada VÁLIDA` y deja un `// TODO`.
- `solucion/` – versión completa con `C3DVisitor.cs`. **La conserva el auxiliar**, no se entrega.

Ambas carpetas contienen el proyecto `MiniHeap`, la gramática `Heap.g4`, `EnvoltorioC.cs` y la carpeta `pruebas/`.

## Requisitos

- .NET 8 SDK
- Java 11+ (lo usa `Antlr4BuildTasks` para generar el lexer, el parser y el visitor al compilar)
- gcc (para compilar y ejecutar el `salida.c` generado)

## Cómo ejecutar

```bash
cd solucion          # o base
dotnet build
dotnet run -- pruebas/valida1.txt     # imprime el C3D y escribe salida.c
gcc -O0 -o salida salida.c && ./salida
```

Sin argumentos imprime `Uso: dotnet run -- <archivo>`.

## Salida esperada (`solucion/`)

| Archivo | Resultado |
|---|---|
| `valida1.txt` | C3D de 17 líneas (empieza con `t1 = H`, `H = H + 12`, ...) y el programa C imprime `7` |
| `valida2.txt` | el programa C imprime `8` |
| `valida3.txt` | el programa C imprime `9` y `9` |
| `invalida1.txt` | `Error semántico: 'b' no declarado` |
| `invalida2.txt` | `Error semántico: 'a' requiere 2 índices` |
| `invalida3.txt` | `line 2:0 extraneous input 'a' expecting {'[', ';'}`<br>`Entrada INVÁLIDA (1 errores sintácticos)` |

En `base/` las entradas válidas solo imprimen `Entrada VÁLIDA`.

> El código generado por ANTLR queda en `obj/` y `salida`/`salida.c` se ignoran; nada de eso se sube al repositorio.
