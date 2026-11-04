# Semana 4 – Ejemplo mínimo de análisis sintáctico con ANTLR4 y C#

Reconoce (solo **análisis sintáctico**) un lenguaje simplificado (`Mini.g4`) con ANTLR4 y .NET 8.
No incluye intérprete, tabla de símbolos ni validación semántica.

## Requisitos

- .NET 8 SDK
- Java 11+ (lo usa `Antlr4BuildTasks` para generar el lexer y el parser al compilar)

## Cómo ejecutar

```bash
cd MiniLang
dotnet build
dotnet run -- pruebas/valida1.txt
```

Sin argumentos imprime `Uso: dotnet run -- <archivo>`.

## Salida esperada

| Archivo | Salida |
|---|---|
| `pruebas/valida1.txt` | `Entrada VÁLIDA` y el árbol `(programa (sentencia (asignacion x = (expr (expr 10) + (expr 5)) ;)) <EOF>)` |
| `pruebas/valida2.txt` | `Entrada VÁLIDA` y el árbol completo |
| `pruebas/invalida1.txt` | `line 1:4 extraneous input '+' expecting {'(', ID, NUM}`<br>`line 2:0 missing ';' at '<EOF>'`<br>`Entrada INVÁLIDA (2 errores sintácticos)` |
| `pruebas/invalida2.txt` | `line 1:3 missing '(' at 'x'`<br>`line 1:9 missing ')' at '{'`<br>`line 3:0 missing ';' at '}'`<br>`Entrada INVÁLIDA (3 errores sintácticos)` |

Los mensajes `line X:Y ...` los emite el listener de errores por defecto de ANTLR.

> El código generado por ANTLR queda en `obj/` y no se sube al repositorio.
