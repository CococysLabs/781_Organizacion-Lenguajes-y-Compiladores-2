using Antlr4.Runtime;

if (args.Length == 0)
{
    Console.WriteLine("Uso: dotnet run -- <archivo>");
    return;
}

string entrada = File.ReadAllText(args[0]);

var input = new AntlrInputStream(entrada);
var lexer = new MiniLexer(input);
var tokens = new CommonTokenStream(lexer);
var parser = new MiniParser(tokens);

var arbol = parser.programa();

if (parser.NumberOfSyntaxErrors == 0)
{
    Console.WriteLine("Entrada VÁLIDA");
    Console.WriteLine(arbol.ToStringTree(parser));
}
else
{
    Console.WriteLine($"Entrada INVÁLIDA ({parser.NumberOfSyntaxErrors} errores sintácticos)");
}
