using System.Text;

// Envuelve el C3D generado en un programa C que se puede compilar con gcc para comprobar el resultado.
public static class EnvoltorioC
{
    public static string Generar(List<string> codigo, int temporales)
    {
        var sb = new StringBuilder();
        sb.AppendLine("#include <stdio.h>");
        sb.AppendLine("int heap[1024], stack[256], H = 0, SP = 0;");
        sb.AppendLine("int main(void) {");

        if (temporales > 0)
        {
            var nombres = Enumerable.Range(1, temporales).Select(i => $"t{i}");
            sb.AppendLine($"    int {string.Join(", ", nombres)};");
        }

        foreach (var linea in codigo)
        {
            if (linea.StartsWith("print "))
                sb.AppendLine($"    printf(\"%d\\n\", {linea.Substring("print ".Length)});");
            else
                sb.AppendLine($"    {linea};");
        }

        sb.AppendLine("    return 0;");
        sb.AppendLine("}");
        return sb.ToString();
    }
}
