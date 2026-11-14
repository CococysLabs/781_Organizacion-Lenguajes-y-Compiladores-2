grammar Heap;

programa     : sentencia+ EOF ;

sentencia    : declaracion
             | asignacion
             | impresion
             ;

declaracion  : 'int' ID ('[' NUM ']')+ ';' ;
asignacion   : acceso '=' expr ';' ;
impresion    : 'print' '(' expr ')' ';' ;

acceso       : ID ('[' expr ']')+ ;

expr         : expr op=('*' | '/') expr   # MulDiv
             | expr op=('+' | '-') expr   # SumaResta
             | '(' expr ')'               # Parentesis
             | acceso                     # LecturaArreglo
             | NUM                        # Numero
             ;

ID           : [a-zA-Z_] [a-zA-Z_0-9]* ;
NUM          : [0-9]+ ;
WS           : [ \t\r\n]+ -> skip ;
COMENTARIO   : '//' ~[\r\n]* -> skip ;
