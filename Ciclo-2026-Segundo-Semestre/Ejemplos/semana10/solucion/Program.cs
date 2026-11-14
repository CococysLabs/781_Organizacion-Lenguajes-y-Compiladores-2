using Antlr4.Runtime;

if (args.Length == 0)
{
    Console.WriteLine("Uso: dotnet run -- <archivo>");
    return;
}

string entrada = File.ReadAllText(args[0]);

var input = new AntlrInputStream(entrada);
var lexer = new HeapLexer(input);
var tokens = new CommonTokenStream(lexer);
var parser = new HeapParser(tokens);

var arbol = parser.programa();

if (parser.NumberOfSyntaxErrors > 0)
{
    Console.WriteLine($"Entrada INVÁLIDA ({parser.NumberOfSyntaxErrors} errores sintácticos)");
    return;
}

var visitor = new C3DVisitor();
try
{
    visitor.Visit(arbol);
}
catch (InvalidOperationException ex)
{
    Console.WriteLine($"Error semántico: {ex.Message}");
    return;
}

foreach (var instruccion in visitor.Codigo)
{
    Console.WriteLine(instruccion);
}

File.WriteAllText("salida.c", EnvoltorioC.Generar(visitor.Codigo, visitor.Temporales));
