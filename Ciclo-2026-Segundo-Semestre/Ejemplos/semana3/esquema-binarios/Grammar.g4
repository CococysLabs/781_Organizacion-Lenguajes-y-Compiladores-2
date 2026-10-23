grammar Grammar;

@header {
import java.util.*;
}

s: l    { System.out.println($l.value); }
;

l returns [int value]
 : l1=l b   { $value = $l1.value * 2 + $b.value; }
 | b        { $value = $b.value; }
;

b returns [int value] 
 : '0'      { $value = 0; }
 | '1'      { $value = 1; }      
;

WS: [ \t\r\n]+ -> skip;