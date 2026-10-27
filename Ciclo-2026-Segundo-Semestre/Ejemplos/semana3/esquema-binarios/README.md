Necesario tener instalado ANTLR4 y Java para ejecutar el ejemplo. Además, de la herramienta grun para ejecutar el parser. Esto lo puede instalar siguiendo las instrucciones en la página oficial de [ANTLR](https://github.com/antlr/antlr4/blob/master/doc/getting-started.md).

Comandos para ejecutar el ejemplo:

```bash
# Generar el lexer y parser
antlr4 Grammar.g4 -o output
# Compilar el código generado
javac output/*.java
# Ejecutar el parser con un archivo de entrada
grun Grammar s -gui input
```