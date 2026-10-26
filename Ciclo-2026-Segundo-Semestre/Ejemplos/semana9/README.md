# Ejemplos de entornos de ejecución de funciones recursivas
---
# Suma

proc suma {  
  // Argumento 1: a  
  t1 \= pos \- 0  
  t2 \= pila\[t1\]  
  // Argumento 2: b  
  t3 \= pos \- 1  
  t4 \= pila\[t3\]  
  t5 \= t2 \+ t4  
  // Retorno  
  t6 \= pos \- 2  
  pila\[t6\] \= t5  
}

pos \= 0  
pila\[pos+3\] \= 5  
pila\[pos+2\] \= 4  
pos \= pos \+ 3  
call suma  
pos \= pos \- 3  
t7 \= pos \+ 1  
t8 \= pila\[t7\]  
print t8

# Fibonacci

proc fib {  
  // Argumento 1: n  
  t1 \= pos \- 0  
  t2 \= pila\[t1\]  
  if t2 \<= 1 goto L1  
  goto L2  
    
  L1:  
    t2 \= pos \- 1  
    pila\[t2\] \= 1  
    goto Salida  
  L2:  
    // Comienza guardado de temp1  
    t3 \= pos \+ 1  
    pila\[t3\] \= t2  
    // Termina guardado de temp1  
    t4 \= t2 \- 1  
    t5 \= pos \+ 3  
    pila\[t5\] \= t4  
    // Llamada recursiva, ojo con crecimiento del registro de act.  
    pos \= pos \+ 3  
    call fib   
    pos \= pos \- 3  
    // Retorno  
    t6 \= pos \+ 2  
    t7 \= pila\[t6\]  
    // Recuperación de temp1  
    t8 \= pos \+ 1  
    t2 \= pila\[t8\]  
      
    // Comienza guardado de temp2 y temp7  
    t9 \= pos \+ 1  
    pila\[t9\] \= t2  
    t10 \= pos \+ 2  
    pila\[t10\] \= t7  
    // Termina guardado de temp2 y temp7  
    t11 \= t2 \- 2  
    t12 \= pos \+ 4  
    pila\[t12\] \= t11  
    // Llamada recursiva, crece aún mas 2da llamada  
    pos \= pos \+ 4  
    call fib   
    pos \= pos \- 4  
    // Retorno  
    t13 \= pos \+ 3  
    t14 \= pila\[t13\]  
    // Recuperación de temps  
    t15 \= pos \+ 1  
    t2 \= pila\[t15\]  
    t16 \= pos \+ 2  
    t7 \= pila\[t16\]  
      
    t17 \= t7 \+ t14  
    t18 \= pos \- 1  
    pila\[t18\] \= t17  
  Salida:  
}

pos \= 0  
pila\[pos \+ 2\] \= 4  
pos \= pos \+ 2  
call fib  
pos \= pos \- 2  
t19 \= pila\[pos+1\]  
print t19

# esPar y esImpar

proc esPar {  
  // Argumento: n  
  t1 \= pos \- 0  
  t2 \= pila\[t1\]

  if t2 \== 0 goto Par\_Base  
  goto Par\_Recursivo

  Par\_Base:  
    t3 \= pos \- 1  
    pila\[t3\] \= 1  
    goto Par\_Salida

  Par\_Recursivo:  
    // Guardar n  
    t4 \= pos \+ 1  
    pila\[t4\] \= t2

    // Argumento para esImpar(n \- 1\)  
    t5 \= t2 \- 1  
    t6 \= pos \+ 3  
    pila\[t6\] \= t5

    // Llamada a esImpar  
    pos \= pos \+ 3  
    call esImpar  
    pos \= pos \- 3

    // Recuperar resultado  
    t7 \= pos \+ 2  
    t8 \= pila\[t7\]

    // Retorno  
    t9 \= pos \- 1  
    pila\[t9\] \= t8

  Par\_Salida:  
}

proc esImpar {  
  // Argumento: n  
  t10 \= pos \- 0  
  t11 \= pila\[t10\]

  if t11 \== 0 goto Impar\_Base  
  goto Impar\_Recursivo

  Impar\_Base:  
    t12 \= pos \- 1  
    pila\[t12\] \= 0  
    goto Impar\_Salida

  Impar\_Recursivo:  
    // Guardar n  
    t13 \= pos \+ 1  
    pila\[t13\] \= t11

    // Argumento para esPar(n \- 1\)  
    t14 \= t11 \- 1  
    t15 \= pos \+ 3  
    pila\[t15\] \= t14

    // Llamada a esPar  
    pos \= pos \+ 3  
    call esPar  
    pos \= pos \- 3

    // Recuperar resultado  
    t16 \= pos \+ 2  
    t17 \= pila\[t16\]

    // Retorno  
    t18 \= pos \- 1  
    pila\[t18\] \= t17

  Impar\_Salida:  
}

// Programa principal  
pos \= 0

// Argumento n \= 2  
t19 \= pos \+ 2  
pila\[t19\] \= 4

// Crear registro de activación  
pos \= pos \+ 2  
call esPar  
pos \= pos \- 2

// Leer resultado  
t20 \= pos \+ 1  
t21 \= pila\[t20\]

print t21  
