grammar Calc;

programa    : asignacion+ EOF ;
asignacion  : ID '=' expr ';' ;
expr        : expr op=('*' | '/') expr   # MulDiv
            | expr op=('+' | '-') expr   # SumaResta
            | '(' expr ')'               # Parentesis
            | NUM                        # Numero
            | ID                         # Variable
            ;
ID          : [a-zA-Z_] [a-zA-Z_0-9]* ;
NUM         : [0-9]+ ;
WS          : [ \t\r\n]+ -> skip ;
COMENTARIO  : '//' ~[\r\n]* -> skip ;
