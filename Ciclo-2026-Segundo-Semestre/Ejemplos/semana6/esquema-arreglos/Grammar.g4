grammar Grammar;

@parser::header {
import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
}

@parser::members {
    private static final Map<String, Symbol> SYMBOL_TABLE =
        new LinkedHashMap<String, Symbol>();

    private static class Symbol {
        private final String name;
        private final String type;
        private final List<Integer> dimensions;

        private Symbol(String name, String type, List<Integer> dimensions) {
            this.name = name;
            this.type = type;
            this.dimensions = new ArrayList<Integer>(dimensions);
        }

        private String shape() {
            StringBuilder result = new StringBuilder(type);
            for (Integer dimension : dimensions) {
                result.append("[").append(dimension).append("]");
            }
            return result.toString();
        }

        private int elements() {
            int result = 1;
            for (Integer dimension : dimensions) {
                result *= dimension;
            }
            return result;
        }
    }

    private static void saveSymbol(String name, String type,
                                   List<Integer> dimensions) {
        SYMBOL_TABLE.put(name, new Symbol(name, type, dimensions));
    }

    private static void printSymbolTable() {
        String separator = "+----------------------+----------------------+----------+";
        System.out.println();
        System.out.println("TABLA DE SIMBOLOS");
        System.out.println(separator);
        System.out.printf("| %-20s | %-20s | %8s |%n",
                          "Identificador", "Tipo", "Elementos");
        System.out.println(separator);
        for (Symbol symbol : SYMBOL_TABLE.values()) {
            System.out.printf("| %-20s | %-20s | %8d |%n",
                              symbol.name, symbol.shape(), symbol.elements());
        }
        System.out.println(separator);
    }
}

// S -> L S | epsilon
s
    : l
      {
          saveSymbol($l.name, $l.type, $l.dimensions);
          printSymbolTable();
      }
    ;

// L -> L [ NUM ] | T ID [ NUM ]
l returns [String name, String type, List<Integer> dimensions]
    : l1=l '[' NUM ']'
      {
          $name = $l1.name;
          $type = $l1.type;
          $dimensions = new ArrayList<Integer>($l1.dimensions);
          $dimensions.add(Integer.parseInt($NUM.text));
      }
    | t ID '[' NUM ']'
      {
          $name = $ID.text;
          $type = $t.value;
          $dimensions = new ArrayList<Integer>();
          $dimensions.add(Integer.parseInt($NUM.text));
      }
    ;

// T -> int | float
t returns [String value]
    : 'int'   { $value = "int"; }
    | 'float' { $value = "float"; }
    ;

ID         : [a-zA-Z_] [a-zA-Z_0-9]* ;
NUM        : [0-9]+ ;
WS         : [ \t\r\n]+ -> skip ;
