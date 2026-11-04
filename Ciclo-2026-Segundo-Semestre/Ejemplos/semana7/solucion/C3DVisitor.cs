using Antlr4.Runtime.Misc;

// Genera Código de Tres Direcciones (C3D) recorriendo el árbol en postorden.
// Cada visita devuelve el "lugar" donde queda el resultado: un literal, una variable o un temporal.
public class C3DVisitor : CalcBaseVisitor<string>
{
    private int _contadorTemporales = 0;

    public List<string> Codigo { get; } = new();

    private string NuevoTemp() => $"t{++_contadorTemporales}";

    public override string VisitNumero([NotNull] CalcParser.NumeroContext ctx)
        => ctx.GetText();

    public override string VisitVariable([NotNull] CalcParser.VariableContext ctx)
        => ctx.GetText();

    public override string VisitParentesis([NotNull] CalcParser.ParentesisContext ctx)
        => Visit(ctx.expr());

    public override string VisitMulDiv([NotNull] CalcParser.MulDivContext ctx)
        => Operacion(ctx.expr(0), ctx.expr(1), ctx.op.Text);

    public override string VisitSumaResta([NotNull] CalcParser.SumaRestaContext ctx)
        => Operacion(ctx.expr(0), ctx.expr(1), ctx.op.Text);

    // Primero se visitan los hijos (postorden), luego se crea el temporal y se emite la instrucción.
    private string Operacion(CalcParser.ExprContext izq, CalcParser.ExprContext der, string op)
    {
        string a = Visit(izq);
        string b = Visit(der);
        string t = NuevoTemp();
        Codigo.Add($"{t} = {a} {op} {b}");
        return t;
    }

    public override string VisitAsignacion([NotNull] CalcParser.AsignacionContext ctx)
    {
        string lugar = Visit(ctx.expr());
        Codigo.Add($"{ctx.ID().GetText()} = {lugar}");
        return "";
    }
}
