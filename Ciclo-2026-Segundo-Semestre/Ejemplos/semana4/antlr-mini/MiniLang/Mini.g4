grammar Mini;

programa     : sentencia* EOF ;
sentencia    : asignacion | condicional | impresion ;
asignacion   : ID '=' expr ';' ;
condicional  : 'if' '(' condicion ')' '{' sentencia* '}' ;
impresion    : 'print' '(' expr ')' ';' ;
condicion    : expr OP_REL expr ;
expr         : expr op=('*' | '/') expr   # MulDiv
             | expr op=('+' | '-') expr   # SumaResta
             | '(' expr ')'               # Parentesis
             | NUM                        # Numero
             | ID                         # Variable
             ;
OP_REL       : '>' | '<' | '>=' | '<=' | '==' | '!=' ;
ID           : [a-zA-Z_] [a-zA-Z_0-9]* ;
NUM          : [0-9]+ ;
WS           : [ \t\r\n]+ -> skip ;
COMENTARIO   : '//' ~[\r\n]* -> skip ;
