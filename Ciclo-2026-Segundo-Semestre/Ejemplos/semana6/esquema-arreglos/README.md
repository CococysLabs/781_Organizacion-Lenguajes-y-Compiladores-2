Necesario tener instalado ANTLR4 y Java para ejecutar el ejemplo. Además, de la herramienta grun para ejecutar el parser. Esto lo puede instalar siguiendo las instrucciones en la página oficial de ANTLR.

Comandos para ejecutar el ejemplo:

# Generar el lexer y parser
antlr4 Grammar.g4 -o output
# Compilar el código generado
javac output/*.java
# Desplazarse a la carpeta output/
cd output/
# Ejecutar el parser con un archivo de entrada
grun Grammar s -gui 
