# Semana 5 — Funciones

> Temas (`temas.csv`): *Definición, tipos, parámetros y argumentos*,
> *Hoisting y modelo de ejecución (call stack)*, *Closures y ámbito léxico*.

Cada ejercicio incluye **enunciado** y **solución completa**.

---

## Ejercicio 1 — Definición, tipos, parámetros y argumentos

### Enunciado

Considere el siguiente mini-lenguaje con funciones tipadas:

```bnf
prog     → funcList
funcList → func funcList | func
func     → type id ( params ) block
params   → param | ε
param    → type id , params | type id
type     → int | float | bool
block    → { stmtList }
stmtList → stmt stmtList | ε
stmt     → type id = expr ;          (* declaración con inicialización *)
         | id = expr ;               (* asignación *)
         | return expr ;
         | block
expr     → expr + term | expr - term | term
term     → term * factor | factor
factor   → num | id | id ( args ) | ( expr )
args     → expr , args | expr | ε
```

1. Escriba una definición válida de función con **3 parámetros de tipos
   distintos** y su correspondiente **llamada**.
2. Construya una **tabla de símbolos** (una por función) para esa definición.
3. Para la llamada, verifique: **aridad** (número de argumentos) y **compatibilidad
   de tipos** (argumento ↔ parámetro). Escriba las reglas de verificación.
4. Muestre qué ocurre con estas llamadas y justifique el error:
   - `suma(1, 2)`
   - `suma(1, 2.0, 3)`
   - `suma(1.5, 2, 3)`  (si `suma(int, int, int)`)

### Solución

#### 1. Definición y llamada válidas

```bnf
float promedio ( int n , float suma , bool redondear ) {
    float r = suma / n ;
    return r ;
}
```

Llamada válida:

```bnf
p = promedio ( 10 , 95.5 , true ) ;
```

#### 2. Tabla de símbolos de `promedio`

| Nombre | Categoría | Tipo | Dirección/offset (ambito) |
|--------|-----------|------|---------------------------|
| `promedio` | función (global) | `(int, float, bool) → float` | — |
| `n` | parámetro | `int` | FP+8 (o arg0) |
| `suma` | parámetro | `float` | FP+16 (arg1) |
| `redondear` | parámetro | `bool` | FP+24 (arg2) |
| `r` | variable local | `float` | FP−8 (local) |

Estructura jerárquica (scope léxico):

```
(global)
 └── promedio (scope de función)
      ├── n, suma, redondear  (scope de parámetros)
      └── r                   (scope del bloque { })
```

#### 3. Reglas de verificación

**Regla de aridad:**

```
Si  args.llamada = k  y  params.definición = m
Entonces  k = m   si no   →  error: "número incorrecto de argumentos"
```

**Regla de tipos (compatibilidad, posición a posición):**

```
Para i = 1..m:
    tipo(args[i])  ≤  tipo(params[i])      (≤ = conversión implícita permitida)
    Si no  →  error: "tipo de argumento incompatible en la posición i"
```

Conversión implícita habitual: `int ≤ float` (promoción), `int ≤ bool` no.

**Regla de retorno:**

```
tipo(expresión return)  ≤  tipo(función)
```

Para `promedio`: `float ≤ float` ✓.

Aplicado a la llamada `promedio(10, 95.5, true)`:

| i | Argumento | Tipo | Parámetro | Tipo | ¿Válido? |
|---|-----------|------|-----------|------|----------|
| 1 | `10` | `int` | `n` | `int` | ✓ (`int ≤ int`) |
| 2 | `95.5` | `float` | `suma` | `float` | ✓ |
| 3 | `true` | `bool` | `redondear` | `bool` | ✓ |

Retorno `r : float` ✓. **Llamada válida.**

#### 4. Llamadas con error

Supongamos `suma(int a, int b, int c)`:

| Llamada | Problema | Error |
|---------|----------|-------|
| `suma(1, 2)` | aridad: 2 argumentos ≠ 3 parámetros | **Error de aridad** |
| `suma(1, 2.0, 3)` | argumento 2 es `float`, parámetro es `int`; `float ≰ int` (no hay conversión implícita hacia abajo) | **Error de tipo en posición 2** |
| `suma(1.5, 2, 3)` | argumento 1 es `float`, parámetro `int` | **Error de tipo en posición 1** |

> Nota: si la firma fuera `suma(float, float, float)`, entonces
> `suma(1, 2.0, 3)` **sí** sería válida por promoción `int → float`.

---

## Ejercicio 2 — *Hoisting* y modelo de ejecución (*call stack*)

### Enunciado

Dado el siguiente código en un lenguaje con *hoisting* (estilo JavaScript):

```javascript
function principal() {
    var x = 1;
    mostrar();
    console.log(x);      // ¿qué imprime?
}

function mostrar() {
    console.log(x);      // ¿qué imprime?
    var x = 5;
    console.log(x);      // ¿qué imprime?
}

principal();
```

1. Explique la fase de **hoisting** (elevación) y qué declara antes de ejecutar.
2. Muestre el **call stack** paso a paso en cada `console.log`.
3. Indique la **tabla de símbolos** de `principal` y de `mostrar` tras el hoisting.
4. ¿Cuál es la salida final?

### Solución

#### 1. Fase de hoisting

Antes de ejecutar cualquier sentencia, el motor:

1. **Crea los marcos de activación** (scopes) para `principal` y `mostrar`.
2. **Eleva las declaraciones** (`var x` y las firmas de función) al **inicio**
   de su scope, inicializándolas en `undefined`.

El código *efectivo* tras el hoisting es:

```javascript
function principal() {
    var x;                 // elevado, x = undefined
    x = 1;                 // asignación real
    mostrar();
    console.log(x);
}

function mostrar() {
    var x;                 // elevado AQUÍ, en el scope de mostrar
    console.log(x);        // x local = undefined
    x = 5;
    console.log(x);
}
```

Punto clave: **`var x` de `mostrar` es una variable DISTINTA** de `x` de
`principal` (ámbito léxico separado). `mostrar` **no** ve el `x` de `principal`.

#### 2. Call stack paso a paso

```
Paso 0 (arranque):  call stack = [ main ]

Paso 1: se llama principal()
        call stack = [ main, principal ]
        (se crea el scope de principal: x = undefined)

Paso 2: x = 1
        scope principal: x = 1

Paso 3: se llama mostrar()
        call stack = [ main, principal, mostrar ]
        (se crea el scope de mostrar: x = undefined ← hoisted)

Paso 4: console.log(x)   ← PRIMER LOG
        se resuelve x en el scope actual (mostrar) → undefined
        → imprime: undefined
        call stack = [ main, principal, mostrar ]

Paso 5: x = 5
        scope mostrar: x = 5

Paso 6: console.log(x)   ← SEGUNDO LOG
        x en scope mostrar = 5
        → imprime: 5
        retorna de mostrar():
        call stack = [ main, principal ]

Paso 7: console.log(x)   ← TERCER LOG
        se resuelve x en scope de principal = 1
        → imprime: 1
        retorna de principal():
        call stack = [ main ]
```

#### 3. Tablas de símbolos tras el hoisting

**Scope `principal`:**

| Nombre | Tipo | Estado inicial | Línea de asignación |
|--------|------|----------------|---------------------|
| `x` | `var` (número) | `undefined` | `x = 1` |

**Scope `mostrar`:**

| Nombre | Tipo | Estado inicial | Línea de asignación |
|--------|------|----------------|---------------------|
| `x` | `var` (número) | `undefined` | `x = 5` |

Resolución de identificadores (regla de alcance):

```
console.log(x) dentro de mostrar
   → buscar x en scope(mostrar)       ✓ encontrado (NO sube a principal)

console.log(x) dentro de principal
   → buscar x en scope(principal)     ✓ encontrado
```

#### 4. Salida final

```
undefined
5
1
```

Explicación: el *hoisting* hace que `x` local de `mostrar` exista (como
`undefined`) **antes** de su declaración; nunca hereda el `x = 1` del llamador
porque el ámbito es **léxico/estático**, no dinámico. ∎

---

## Ejercicio 3 — Closures y ámbito léxico

### Enunciado

```javascript
function crearContador() {
    var cuenta = 0;
    return function() {
        cuenta = cuenta + 1;
        return cuenta;
    };
}

var c1 = crearContador();
var c2 = crearContador();

console.log(c1());    // ?
console.log(c1());    // ?
console.log(c2());    // ?
console.log(c1());    // ?
```

1. Explique qué es un **closure** (función + su entorno léxico capturado).
2. Dibuje los **entornos (scopes) enlazados** en cadena (*scope chain*) después
   de ejecutar `var c1 = ...` y `var c2 = ...`.
3. Trace la ejecución mostrando el valor de `cuenta` en cada entorno.
4. Responda: ¿por qué `c1` y `c2` **no** comparten el contador?

### Solución

#### 1. Definición de *closure*

Un **closure** es:

> Una **función anidada** junto con el **entorno léxico** (scopes) en el que
> fue **creada**, el cual permanece **accesible** incluso después de que la
> función externa haya retornado.

En este caso, la función interna `function() { cuenta = cuenta + 1; ... }`
**captura** el scope de `crearContador` (donde vive `cuenta`).

#### 2. Entornos enlazados (scope chain)

Tras `var c1 = crearContador();`:

```
[ scope global ]
   c1 ──────┐
   c2 (aún undefined)
            │
            ▼
[ scope crearContador #1 ]        ← entorno capturado por c1
   cuenta = 0
   (la función interna retorna y "cerrada" sobre este scope)
            ▲
            │
   closure de c1
```

Tras `var c2 = crearContador();` se crea **un entorno nuevo**:

```
[ scope global ]
   c1 ──────► [ scope crearContador #1 ]   cuenta = 0  ◄── closure c1
   c2 ──────► [ scope crearContador #2 ]   cuenta = 0  ◄── closure c2
```

Cadena de búsqueda (*scope chain*) dentro de `c1()`:

```
scope(c1) → scope(crearContador #1) → scope(global)
```

#### 3. Trazado de ejecución

| Paso | Acción | Entorno #1 (`cuenta`) | Entorno #2 (`cuenta`) | Retorno |
|------|--------|-----------------------|-----------------------|---------|
| 0 | `crearContador()` #1 | `0` | — | `c1` = closure sobre #1 |
| 1 | `crearContador()` #2 | `0` | `0` | `c2` = closure sobre #2 |
| 2 | `c1()` | **1** | 0 | `1` |
| 3 | `c1()` | **2** | 0 | `2` |
| 4 | `c2()` | 2 | **1** | `1` |
| 5 | `c1()` | **3** | 1 | `3` |

#### 4. ¿Por qué no comparten el contador?

Porque **cada invocación de `crearContador` crea un scope nuevo** con su
propia variable `cuenta`. Los closures `c1` y `c2` capturan **entornos
diferentes**:

```
c1 → cuenta₁  (creado en la 1ª llamada)
c2 → cuenta₂  (creado en la 2ª llamada)
```

El alcance es **léxico** (estático): depende de *dónde está escrita* la
función, pero la **instancia del scope** se crea en tiempo de ejecución
**cada vez** que se ejecuta la función externa. ∎

### Ejercicio complementario (para clase)

```javascript
function sumador(n) {
    return function(x) { return n + x; };
}
var doble = sumador(2);
var triple = sumador(3);
console.log(doble(5));    // 7
console.log(triple(5));   // 8
```

Mismo principio: `n` es un parámetro del scope capturado; cada llamada a
`sumador` congela un `n` distinto en su closure.

---

## Resumen de la semana

| Ejercicio | Concepto clave |
|-----------|----------------|
| 1 | Firma de función + aridad + compatibilidad de tipos (≤) |
| 2 | Hoisting: declaraciones elevadas; call stack; scopes por activación |
| 3 | Closure = función + entorno léxico capturado; un scope por invocación |

### Preguntas de repaso

1. Si en el Ejercicio 2 no hubiera *hoisting*, ¿qué error ocurriría en la
   primera línea de `mostrar`?
2. ¿Qué pasaría si en el Ejercicio 3 el closure usara `let cuenta` en lugar de
   `var cuenta` dentro de un bucle `for`?
3. Distinga: **ámbito dinámico** vs **ámbito léxico** — ¿cuál usa el ejemplo 2?
4. ¿Los parámetros formales pertenecen al scope de la función o al del bloque?
