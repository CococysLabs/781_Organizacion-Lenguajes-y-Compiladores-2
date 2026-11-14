using Antlr4.Runtime;

if (args.Length == 0)
{
    Console.WriteLine("Uso: dotnet run -- <archivo>");
    return;
}

string entrada = File.ReadAllText(args[0]);

var input = new AntlrInputStream(entrada);
var lexer = new CalcLexer(input);
var tokens = new CommonTokenStream(lexer);
var parser = new CalcParser(tokens);

var arbol = parser.programa();

if (parser.NumberOfSyntaxErrors > 0)
{
    Console.WriteLine($"Entrada INVÁLIDA ({parser.NumberOfSyntaxErrors} errores sintácticos)");
    return;
}

Console.WriteLine("Entrada VÁLIDA");

// TODO: crear C3DVisitor (hereda de CalcBaseVisitor<string>),
//       visitar el árbol con visitor.Visit(arbol) e imprimir cada instrucción de visitor.Codigo.
