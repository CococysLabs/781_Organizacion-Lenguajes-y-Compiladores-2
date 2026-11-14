# Semana 8 — Entornos y memoria en un compilador (tabla de símbolos, stack/heap, C3D)

> Temas (`temas.csv`): *Tabla de símbolos y resolución de Variables*,
> *Organización en memoria: Stack, Heap y Punteros*,
> *Traducción de estructuras de control a C3D (etiquetas y saltos)*.

Cada ejercicio incluye **enunciado** y **solución completa**.

---

## Ejercicio 1 — Tabla de símbolos y resolución de variables

### Enunciado

Considere el siguiente programa con bloques anidados:

```bnf
programa → bloque
bloque   → { sentencias }
sentencias → sentencia sentencias | ε
sentencia  → tipo id = expr ;
           | id = expr ;
tipo       → int | float
expr       → expr + term | term
term       → id | num | term * factor | factor
factor     → id | num | ( expr )
```

Programa de ejemplo:

```c
{ int x = 10;
  { float x = 3.5;
    { int x = 7;
      x = x + 1;      // ¿cuál x?
    }
    x = x * 2;        // ¿cuál x?
  }
  x = x + 1;          // ¿cuál x?
}
```

1. Construya la **tabla de símbolos anidada** (un scope por bloque),
   mostrando el estado tras **cada bloque**.
2. Para cada `x = ...`, indique **qué `x` resuelve** el compilador y su tipo.
3. Escriba el algoritmo de **resolución de identificadores** (scope chain).
4. ¿Qué pasaría si escribiéramos `y = 5;` sin declarar `y`?

### Solución

#### 1. Tablas de símbolos por scope

Notación: `S0` = scope raíz, `S1` = bloque interior 1, `S2` = interior 2.

**Tras `int x = 10;` (scope S0):**

| Scope | Nombre | Tipo | Valor inicial |
|-------|--------|------|---------------|
| S0 | `x` | `int` | 10 |

**Tras entrar a `{ float x = 3.5;` (scope S1, hijo de S0):**

| Scope | Nombre | Tipo | Valor inicial |
|-------|--------|------|---------------|
| S0 | `x` | `int` | 10 |
| S1 | `x` | `float` | 3.5 |

*S1 "sombrea" a S0: la búsqueda se detiene en el scope más interno.*

**Tras entrar a `{ int x = 7;` (scope S2, hijo de S1):**

| Scope | Nombre | Tipo | Valor inicial |
|-------|--------|------|---------------|
| S0 | `x` | `int` | 10 |
| S1 | `x` | `float` | 3.5 |
| S2 | `x` | `int` | 7 |

**Tras `x = x + 1;` dentro de S2:**

| Scope | Nombre | Tipo | Valor |
|-------|--------|------|-------|
| S2 | `x` | `int` | **8**  (7 + 1) |

*S2 cierra: se elimina S2 de la pila de scopes.*

**Tras `x = x * 2;` dentro de S1:**

| Scope | Nombre | Tipo | Valor |
|-------|--------|------|-------|
| S1 | `x` | `float` | **7.0**  (3.5 * 2) |

*S1 cierra: se elimina S1.*

**Tras `x = x + 1;` dentro de S0:**

| Scope | Nombre | Tipo | Valor |
|-------|--------|------|-------|
| S0 | `x` | `int` | **11**  (10 + 1) |

#### 2. Resolución de cada `x = ...`

| Sentencia | Scope donde está | Scope resuelto | `x` usada | Tipo | Operación |
|-----------|------------------|----------------|-----------|------|-----------|
| `x = x + 1;` (interno) | S2 | **S2** | `x` de S2 | `int` | `7 + 1 = 8` |
| `x = x * 2;` (medio) | S1 | **S1** | `x` de S1 | `float` | `3.5 * 2 = 7.0` |
| `x = x + 1;` (externo) | S0 | **S0** | `x` de S0 | `int` | `10 + 1 = 11` |

**Shadowing (sombreado):** cada bloque interno define un `x` distinto; la
regla de resolución elige **el más cercano** en la cadena de scopes.

#### 3. Algoritmo de resolución de identificadores

```
resolver(nombre, scope_actual):
    s ← scope_actual
    mientras s ≠ nulo:
        si s.tabla contiene nombre:
            devolver s.tabla[nombre]        // encontrado
        s ← s.padre                         // subir en la cadena léxica
    devolver error "identificador no declarado"
```

La **cadena de scopes** es la del **ámbito léxico** (estructura estática de
bloques), **no** la dinámica de ejecución.

Aplicado a `x` en S2:

```
buscar x en S2 → encontrado (int, 7)      ← se detiene aquí
```

Aplicado a `x` en S1 (después de cerrar S2):

```
buscar x en S1 → encontrado (float, 3.5)  ← nunca llega a S0
```

#### 4. Variable no declarada

`y = 5;` con `y` ausente de S_actual…S0 ⇒ la búsqueda falla en toda la cadena:

> **Error en tiempo de compilación:** *"identificador `y` no declarado"*
> (o en lenguajes con *hoisting*/tipado dinámico, podría ser `ReferenceError`
> en tiempo de ejecución).

La tabla de símbolos permite detectarlo **antes** de generar código. ∎

---

## Ejercicio 2 — Organización en memoria: Stack, Heap y punteros

### Enunciado

Considere:

```c
int main() {
    int a = 5;                 // (1)
    int arr[3] = {1, 2, 3};    // (2) arreglo estático en stack
    int *p = &a;               // (3) puntero en stack
    int *h = new int[10];      // (4) memoria dinámica en heap
    *p = *p + 1;               // (5)
    h[0] = 42;                 // (6)
    free(h);                   // (7)
    return a;                  // (8)
}
```

1. Para cada línea **(1)–(8)**, indique en qué **segmento** (stack / heap /
   estáticos) se aloja cada objeto y por qué.
2. Dibuje el **diagrama de memoria** tras la línea (4).
3. Muestre el efecto de `*p = *p + 1` sobre `a` (aliasing).
4. Explique por qué `free(h)` **no** libera `a` ni `arr`, y qué pasa si se
   libera dos veces (*double free*) o no se libera (*memory leak*).

### Solución

#### 1. Segmentos por línea

| Línea | Objeto | Segmento | Motivo |
|-------|--------|----------|--------|
| (1) `int a = 5` | `a` | **Stack** | variable local de `main`, vida ligada a la activación |
| (2) `int arr[3]` | `arr` | **Stack** | arreglo de tamaño fijo y conocido en tiempo de compilación, local |
| (3) `int *p = &a` | `p` (la dirección) | **Stack** | el puntero es una variable local; el *apuntado* (`a`) ya está en stack |
| (4) `int *h = new int[10]` | bloque de 10 `int` | **Heap** | tamaño dinámico, vida independiente de la activación |
| (4) | `h` (la dirección) | **Stack** | la variable puntero es local |
| (5) `*p = *p + 1` | escribe en `a` | **Stack** (via puntero) | solo modifica `a` |
| (6) `h[0] = 42` | `h[0]` | **Heap** | offset dentro del bloque dinámico |
| (7) `free(h)` | libera bloque | **Heap** | devuelve el bloque al *allocator* |
| (8) `return a` | lee `a` | **Stack** | |

#### 2. Diagrama de memoria tras la línea (4)

```
Dirección (ejemplo)     Contenido              Segmento
───────────────────────────────────────────────────────────────
  0x0040 (estáticos)    (código, literales)    DATA/TEXT
        │
  ↑ crece (stack)       ┌─────────────────┐
  0x7FFF...  fp/sp →   │ return addr (8) │  ← marco de main
                        │ a = 5           │  (1)
                        │ arr[0]=1        │  (2)
                        │ arr[1]=2        │
                        │ arr[2]=3        │
                        │ p = 0x7FFF_a    │  (3) → apunta a 'a'
                        │ h = 0x5555_100  │  (4) → apunta al heap
                        └────────┬────────┘
                                 │
                                 ▼
                        ┌─────────────────────────┐
  0x5555_100 (heap) →   │ [0]=? [1]=? … [9]=?    │  (4) new int[10]
                        │ 10 × int (40 bytes)     │
                        └─────────────────────────┘
  ↓ crece (heap)
```

Después de (5) y (6):

```
  a = 6                     (por *p = *p + 1, aliasing)
  h[0] = 42
```

#### 3. Aliasing en `*p = *p + 1`

```
p = &a   ⇒  *p equivale a a
*p = *p + 1
⇒ a = a + 1
⇒ a = 6
```

Ambos nombres (`a` y `*p`) **aliasan** la misma celda. El compilador debe
asumir que escribir por puntero puede modificar cualquier variable visible
(es la razón por la que las optimizaciones con punteros son difíciles: se
necesita *alias analysis*).

#### 4. `free(h)` y errores de memoria

| Situación | ¿Qué se libera? | Consecuencia |
|-----------|------------------|--------------|
| `free(h)` correcto | solo el bloque del **heap** (`0x5555_100`) | `a`, `arr`, `p`, `h` viven en **stack** y se liberan solos al hacer `return` de `main` |
| **Double free** (`free(h)` dos veces) | intenta liberar ya liberado | *undefined behavior*: corrupción del heap, crash o vulnerabilidad |
| **Memory leak** (no llamar `free`) | nada | el bloque queda huérfano en el heap; se acumula y agota memoria |
| Usar `h` tras `free(h)` | — | *use after free* (UB) |

Regla de vida (*lifetime*):

```
stack  : se crea al entrar al scope, se destruye al salir   (automático)
heap  : se crea con new/malloc, se destruye con free/delete (explícito)
```

`free` **nunca** toca el stack: solo devuelve bloques asignados al heap. ∎

---

## Ejercicio 3 — Traducción de estructuras de control a C3D

### Enunciado

Traduzca a **código de 3 direcciones (C3D)** con etiquetas y saltos:

1. `if` simple:

```c
if (x < y)
    m = x;
else
    m = y;
```

2. `while`:

```c
while (x < n) {
    x = x + 1;
    s = s + x;
}
```

3. `for` equivalente al `while` anterior:

```c
for (i = 0; i < n; i = i + 1)
    s = s + i;
```

Use temporales `t0, t1, …` y etiquetas `L0, L1, …`.
Además, escriba la **regla de traducción** (esquema) para `if` y `while`.

### Solución

#### 1. Traducción de `if` … `else`

**Regla de traducción (esquema):**

```
if ( E ) S1 else S2   ⇒
        E.code
        ifFalse E.addr goto Lelse
        S1.code
        goto Lend
Lelse:  S2.code
Lend:
```

**Aplicación** (`x < y` → `t0 = x < y`):

```
        ( < , x , y , t0 )        # t0 = x < y
        ifFalse t0 goto L1        # si no se cumple, ir a la rama else
        ( = , x , _ , m )         # m = x
        goto L2                   # saltar el else
L1:     ( = , y , _ , m )         # m = y
L2:     ( _ , _ , _ , _ )         # fin (etiqueta vacía)
```

Árbol anotado (con los atributos `code` y `label`):

```
              if-else
         ┌───────┼────────┐
        cond    then      else
      (x < y)   m=x       m=y
         │
      t0 = x < y
```

#### 2. Traducción de `while`

**Regla de traducción:**

```
while ( E ) S   ⇒
Lstart:
        E.code
        ifFalse E.addr goto Lend
        S.code
        goto Lstart
Lend:
```

**Aplicación** (`x < n` → `t0 = x < n`):

```
L0:     ( < , x , n , t0 )        # condición
        ifFalse t0 goto L1
        ( + , x , 1 , t1 )
        ( = , t1 , _ , x )        # x = x + 1
        ( + , s , x , t2 )
        ( = , t2 , _ , s )        # s = s + x
        goto L0                   # vuelve a evaluar la condición
L1:     ( _ , _ , _ , _ )         # fin del while
```

#### 3. Traducción de `for`

**Regla:** `for (init; cond; upd) S`
⇒ `init; Lstart: if(!cond) goto Lend; S; Lupd: upd; goto Lstart; Lend:`

(El `goto Lstart` salta tras ejecutar el *update* para no re-ejecutar el `init`.)

**Aplicación** (`s = s + i`):

```
        ( = , 0 , _ , i )         # init: i = 0
L0:     ( < , i , n , t0 )        # cond
        ifFalse t0 goto L1
        ( + , s , i , t1 )
        ( = , t1 , _ , s )        # cuerpo: s = s + i
L2:     ( + , i , 1 , t2 )
        ( = , t2 , _ , i )        # update: i = i + 1
        goto L0
L1:     ( _ , _ , _ , _ )         # fin del for
```

#### 4. Equivalencia `for` ⇄ `while`

El `for` del enunciado es **semánticamente idéntico** a:

```c
i = 0;
while (i < n) {
    s = s + i;
    i = i + 1;      // ← el update va al FINAL del cuerpo
}
```

Ambos generan la misma secuencia de cuádruplas salvo el orden de `i = i + 1`
(colocarlo al final del cuerpo, antes de saltar a `L0`).

> **Nota de errores:** el C3D generado permite **recuperación de errores**
> en tiempo de ejecución (p. ej. añadir `ifFalse` de división por cero) y es
> el insumo directo para la generación de código ensamblador (RISC-V/ARM)
> con etiquetas `beq`, `bne`, `j`. ∎

---

## Resumen de la semana

| Ejercicio | Concepto clave |
|-----------|----------------|
| 1 | Pila de scopes, shadowing, algoritmo `resolver(nombre, scope)` |
| 2 | Stack (automático) vs Heap (explícito); aliasing por punteros; leaks/double free |
| 3 | Reglas de traducción de `if`/`while`/`for` → C3D con `goto`/`ifFalse` y etiquetas |

### Preguntas de repaso

1. Si la resolución de identificadores fuera **dinámica**, ¿cambiaría la
   respuesta del Ejercicio 1? ¿Por qué los lenguajes modernos usan léxico?
2. ¿Por qué un arreglo de tamaño variable (`int arr[n]`) va al heap y no al
   stack (en C estándar)?
3. En el `while`, ¿por qué hace falta el `goto L0` al final del cuerpo?
4. ¿Cuántas etiquetas se necesitan para un `if` **sin** `else`?
