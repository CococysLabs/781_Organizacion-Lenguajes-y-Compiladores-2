using Antlr4.Runtime;

public record Arreglo(int Slot, List<int> Dims);

public class C3DVisitor : HeapBaseVisitor<string>
{
    private int _temp = 0;
    private int _siguienteSlot = 0;
    private readonly Dictionary<string, Arreglo> _tabla = new();

    public List<string> Codigo { get; } = new();
    public int Temporales => _temp;

    private string NuevoTemp() => $"t{++_temp}";

    public override string VisitDeclaracion(HeapParser.DeclaracionContext ctx)
    {
        string nombre = ctx.ID().GetText();
        if (_tabla.ContainsKey(nombre))
            throw new InvalidOperationException($"'{nombre}' ya fue declarado");

        var dims = ctx.NUM().Select(n => int.Parse(n.GetText())).ToList();
        int total = dims.Aggregate(1, (a, b) => a * b);

        string baseT = NuevoTemp();
        Codigo.Add($"{baseT} = H");            // base del arreglo
        Codigo.Add($"H = H + {total}");        // reservar celdas en el heap

        int slot = _siguienteSlot++;
        _tabla[nombre] = new Arreglo(slot, dims);
        string dir = NuevoTemp();
        Codigo.Add($"{dir} = SP + {slot}");    // dirección de 'nombre' en la pila
        Codigo.Add($"stack[{dir}] = {baseT}"); // guardar el puntero base
        return "";
    }

    private string Direccion(HeapParser.AccesoContext ctx)
    {
        string nombre = ctx.ID().GetText();
        if (!_tabla.TryGetValue(nombre, out var arr))
            throw new InvalidOperationException($"'{nombre}' no declarado");

        var indices = ctx.expr();
        if (indices.Length != arr.Dims.Count)
            throw new InvalidOperationException($"'{nombre}' requiere {arr.Dims.Count} índices");

        // acum = ((i0 * d1 + i1) * d2 + i2) ...   (Horner)
        string acum = Visit(indices[0]);
        for (int k = 1; k < indices.Length; k++)
        {
            string m = NuevoTemp();
            Codigo.Add($"{m} = {acum} * {arr.Dims[k]}");
            string i = Visit(indices[k]);
            string a = NuevoTemp();
            Codigo.Add($"{a} = {m} + {i}");
            acum = a;
        }

        string d = NuevoTemp();
        Codigo.Add($"{d} = SP + {arr.Slot}");
        string b = NuevoTemp();
        Codigo.Add($"{b} = stack[{d}]");      // puntero base desde la pila
        string r = NuevoTemp();
        Codigo.Add($"{r} = {b} + {acum}");    // dirección final en el heap
        return r;
    }

    public override string VisitAsignacion(HeapParser.AsignacionContext ctx)
    {
        string dir = Direccion(ctx.acceso());
        string valor = Visit(ctx.expr());
        Codigo.Add($"heap[{dir}] = {valor}");
        return "";
    }

    public override string VisitLecturaArreglo(HeapParser.LecturaArregloContext ctx)
    {
        string dir = Direccion(ctx.acceso());
        string t = NuevoTemp();
        Codigo.Add($"{t} = heap[{dir}]");
        return t;
    }

    public override string VisitNumero(HeapParser.NumeroContext ctx) => ctx.GetText();
    public override string VisitParentesis(HeapParser.ParentesisContext ctx) => Visit(ctx.expr());
    public override string VisitMulDiv(HeapParser.MulDivContext ctx)
        => Operacion(ctx.expr(0), ctx.op.Text, ctx.expr(1));
    public override string VisitSumaResta(HeapParser.SumaRestaContext ctx)
        => Operacion(ctx.expr(0), ctx.op.Text, ctx.expr(1));

    private string Operacion(HeapParser.ExprContext izq, string op, HeapParser.ExprContext der)
    {
        string a = Visit(izq);
        string b = Visit(der);
        string t = NuevoTemp();
        Codigo.Add($"{t} = {a} {op} {b}");
        return t;
    }

    public override string VisitImpresion(HeapParser.ImpresionContext ctx)
    {
        Codigo.Add($"print {Visit(ctx.expr())}");
        return "";
    }
}
