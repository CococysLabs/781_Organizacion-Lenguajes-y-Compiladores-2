# Semana 2 — Análisis Sintáctico

> Temas (`temas.csv`): *Análisis Descendente*, *Análisis Ascendente*,
> *Resolución de ambigüedades sintácticas*, *Esquemas de Traducción*.

Cada ejercicio incluye **enunciado** y **solución completa**.

---

## Ejercicio 1 — Resolución de ambigüedades sintácticas

### Enunciado

La gramática de expresiones aritméticas es **ambigua**:

```bnf
E → E + E
E → E * E
E → id
```

1. Construya **dos árboles de derivación distintos** para la cadena `id + id * id`.
2. Explique por qué esto es un problema para un compilador.
3. Resuelva la ambigüedad imponiendo: `*` tiene mayor precedencia que `+`,
   y ambos son **asociativos por la izquierda**.
4. Verifique que la nueva gramática **NO** es ambigua para `id + id * id`
   (único árbol de derivación).

### Solución

#### 1. Dos árboles de derivación para `id + id * id`

**Árbol A (la suma se resuelve arriba → `*` domina):**

```
        E
      / | \
     E   +   E
     |      /|\
     id    E  *  E
            |   |
            id  id
```

Derivación izquierda (≈): `E ⇒ E + E ⇒ id + E ⇒ id + E * E ⇒ id + id * id`

**Árbol B (la multiplicación se resuelve arriba → `+` domina):**

```
          E
        / | \
       E   *   E
      /|\    / | \
     E  +  E  id  id
     |   |
     id  id
```

Derivación izquierda (≈): `E ⇒ E * E ⇒ E + E * E ⇒ id + E * E ⇒ id + id * id`

#### 2. ¿Por qué es un problema?

La ambigüedad significa que **una cadena tiene más de un árbol de derivación**,
por lo tanto más de un significado (semántica) asociado:

- Árbol A ⇒ interpretación `(id + id) * id`… no: en el árbol A la raíz es `+`,
  la `+` se evalúa **último**, y `*` primero → `id + (id * id)`.
- Árbol B ⇒ raíz es `*`, la `*` se evalúa último → `(id + id) * id`.

El analizador sintáctico no podría decidir **qué árbol pasar al análisis
semántico**, y por tanto el programa tendría un significado no determinado.
Toda gramática ambigua debe eliminarse (o resolverse por precedencias en el
lexer/parser).

#### 3. Gramática resuelta (LL(1), con niveles de precedencia)

Idea: **un no terminal por nivel de precedencia**. Los niveles bajos se
construyen primero (se evalúan al final).

```bnf
E  → T E'          (* suma: precedencia más baja *)
E' → + T E' | ε
T  → F T'          (* multiplicación *)
T' → * F T' | ε
F  → ( E ) | id    (* factor: paréntesis o identificador *)
```

Propiedades impuestas:

- `*` aparece **más abajo** en la cadena de niveles (`T` se expande dentro de `E`)
  ⇒ mayor precedencia.
- `E' → + T E'` es **recursión por la derecha** en la cadena de dependencia
  pero genera **asociatividad izquierda** en la derivación
  (`E ⇒ + T E' ⇒ + T + T E'`… las `+` se anidan a la izquierda).
  Lo mismo para `*` en `T'`.
- Los paréntesis en `F → ( E )` permiten forzar cualquier asociatividad.

#### 4. Verificación: único árbol para `id + id * id`

Única derivación posible:

```
E ⇒ T E' ⇒ F T' E' ⇒ id T' E' ⇒ id E' ⇒ id + T E'
                                    ⇒ id + F T' E' ⇒ id + id T' E'
                                    ⇒ id + id * F T' E' ⇒ id + id * id T'
                                    ⇒ id + id * id ε ⇒ id + id * id
```

Árbol resultante (raíz `+` ⇒ la suma se aplica **al final**, dentro del término
está la multiplicación):

```
          E
        / |  \
       T   E'   (vacío)
       |
       F
       |
       id
```

pero desglosado completo:

```
              E
            / | \
           T   E'
         / | \
        F  T'
        |   |
        id  ε
```

Árbol canónico (forma habitual):

```
            E
          / | \
         T   E'
        /|\    \
       F  T'    + T E'
       |  |         |    \
       id ε         F    E'
                    |     |
                   id    T'
                         |
                        * F T'
                        |  |
                        id ε
```

Simplificando visualmente:

```
                E
           ┌────┴────┐
           T         E'
           │     ┌───┴───┐
           F     +       T E'→ε
           │           ┌─┴─┐
          id           F   T'
                       │  ┌┴┐
                       id * T'
                            │
                            F → id
```

Como la gramática es LL(1) y las cadenas alternativas son disjuntas
(`+`/`*`/`(`/`id`/`ε` están en FIRST separados), **no puede existir un segundo
árbol**: la ambigüedad quedó resuelta. ∎

---

## Ejercicio 2 — Análisis descendente (LL(1))

### Enunciado

Para la gramática LL(1):

```bnf
E  → T E'
E' → + T E' | ε
T  → F T'
T' → * F T' | ε
F  → ( E ) | id
```

1. Calcule **FIRST** y **FOLLOW** de cada no terminal.
2. Construya la **tabla de análisis LL(1)**.
3. Trace el **análisis descendente** (pila de símbolos + entrada) para la
   cadena `id + id * id` (separada por `$` al final).
4. ¿Por qué `FOLLOW(E')` contiene `)`?

### Solución

#### 1. FIRST y FOLLOW

**FIRST** (arranque de las cadenas que genera):

| No terminal | FIRST |
|-------------|-------|
| `F` | `( , id` |
| `T'` | `* , ε` |
| `T` | FIRST(F) = `( , id` |
| `E'` | `+ , ε` |
| `E` | FIRST(T) = `( , id` |

**FOLLOW** (qué puede venir *después* de ese no terminal):

| No terminal | Cálculo | FOLLOW |
|-------------|---------|--------|
| `E` | Símbolo inicial ⇒ `$`; y `F → ( E )` ⇒ `)` | `$ , )` |
| `E'` | `E → T E'`, después de `E'` viene lo que sigue a `E`: `$ , )`; además `E' → + T E'` ⇒ FOLLOW(E') | `$ , )` |
| `T` | `E → T E'` ⇒ FIRST(E')\{ε} = `+`; y si `E' ⇒ ε` ⇒ FOLLOW(E) | `$ , ) , +` |
| `T'` | `T → F T'` ⇒ FOLLOW(T); y `T' → * F T'` ⇒ FOLLOW(T') | `$ , ) , + , *` |
| `F` | `T → F T'` ⇒ FIRST(T')\{ε} = `*`; si `T'⇒ε` ⇒ FOLLOW(T) | `( , id , * , $ , ) , +` |

Resumen:

```
FIRST(E)  = FIRST(T) = FIRST(F) = { ( , id }
FIRST(E') = { + , ε }
FIRST(T') = { * , ε }

FOLLOW(E)  = { $ , ) }
FOLLOW(E') = { $ , ) }
FOLLOW(T)  = { $ , ) , + }
FOLLOW(T') = { $ , ) , + , * }
FOLLOW(F)  = ( id , * , $ , ) , + }   ← FIRST(T') ∪ FOLLOW(T) ∪ FIRST(E')…
```

Verificación de `FOLLOW(F)` en detalle:

- `T → F T'`: FIRST(T') \ {ε} = `{*}` ⇒ `* ∈ FOLLOW(F)`.
- `T → F T'` con `T' ⇒ ε`: FOLLOW(T) = `{$, ), +}` ⇒ todos en FOLLOW(F).
- En `F → ( E )` no hay `F` anidado directo; `F` no está dentro de `E`.
- `F → ( E )` aporta `)` a FOLLOW(E), no a FOLLOW(F).

⇒ `FOLLOW(F) = { id?, * , $ , ) , + }`.
Ojo: `id` **no** pertenece a FOLLOW(F): `F` nunca es seguido de `id`
(dos factores seguidos sin operador no son generables). Corrección final:

```
FOLLOW(E)  = { $ , ) }
FOLLOW(E') = { $ , ) }
FOLLOW(T)  = { $ , ) , + }
FOLLOW(T') = { $ , ) , + , * }
FOLLOW(F)  = { * , $ , ) , + }
```

#### 2. Tabla LL(1)  `M[A, a]`

Regla: para `A → α`: `FIRST(α)` → producción; si `ε ∈ FIRST(α)`,
`FOLLOW(A)` → producción.

| | `id` | `+` | `*` | `(` | `)` | `$` |
|---|---|---|---|---|---|---|
| **E** | `E → T E'` | | | `E → T E'` | | |
| **E'** | | `E' → + T E'` | | | `E' → ε` | `E' → ε` |
| **T** | `T → F T'` | | | `T → F T'` | | |
| **T'** | `T' → ε` | `T' → ε` | `T' → * F T'` | | `T' → ε` | `T' → ε` |
| **F** | `F → id` | | | `F → ( E )` | | |

No hay celdas con dos producciones ⇒ la gramática **es LL(1)**.

#### 3. Trazado del análisis para `id + id * id`

Pila se escribe con el **tope a la derecha**. Entrada: `id + id * id $`

| # | Pila | Entrada | Acción |
|---|------|---------|--------|
| 1 | `$ E` | `id + id * id $` | `M[E,id] = E → T E'` |
| 2 | `$ E' T` | `id + id * id $` | `M[T,id] = T → F T'` |
| 3 | `$ E' T' F` | `id + id * id $` | `M[F,id] = F → id` |
| 4 | `$ E' T' id` | `id + id * id $` | coincidencia ⇒ pop `id` |
| 5 | `$ E' T'` | `+ id * id $` | `M[T',+] = T' → ε` ⇒ pop `T'` |
| 6 | `$ E'` | `+ id * id $` | `M[E',+] = E' → + T E'` |
| 7 | `$ E' T +` | `+ id * id $` | coincidencia ⇒ pop `+` |
| 8 | `$ E' T` | `id * id $` | `M[T,id] = T → F T'` |
| 9 | `$ E' T' F` | `id * id $` | `M[F,id] = F → id` |
| 10 | `$ E' T' id` | `id * id $` | coincidencia ⇒ pop `id` |
| 11 | `$ E' T'` | `* id $` | `M[T',*] = T' → * F T'` |
| 12 | `$ E' T' F *` | `* id $` | coincidencia ⇒ pop `*` |
| 13 | `$ E' T' F` | `id $` | `M[F,id] = F → id` |
| 14 | `$ E' T' id` | `id $` | coincidencia ⇒ pop `id` |
| 15 | `$ E' T'` | `$` | `M[T',$] = T' → ε` |
| 16 | `$ E'` | `$` | `M[E',$] = E' → ε` |
| 17 | `$` | `$` | **aceptación** ✓ |

#### 4. ¿Por qué `FOLLOW(E')` contiene `)`?

Porque `E` puede aparecer **entre paréntesis**: `F → ( E )`.
Cuando `E` ha sido consumido, el siguiente símbolo de entrada es `)`.
Como `E'` es la **cola** de `E` (`E → T E'`), todo lo que puede seguir a `E`
también puede seguir a `E'`:

```
F → ( E )   y   E → T E'
⇒  ... (  T  E'  ) ...
                 ^^ tras E' viene ')'
```

Por lo tanto `) ∈ FOLLOW(E')`. ∎

---

## Ejercicio 3 — Análisis ascendente (SLR) y el conflicto *dangling else*

### Enunciado

```bnf
S → if E then S
S → if E then S else S
S → a
E → b
```

1. Escriba la **cerradura de ítems LR(0)** y las **transiciones** de `S₀`.
2. Construya las **acciones de la tabla SLR**.
3. Identifique el **conflicto** y resuélvalo (regla de prioridad *shift* sobre
   *reduce* en `else`).
4. Trace el análisis **shift-reduce** para `if b then if b then a else a`.

### Solución

#### 1. Ítems LR(0) y transiciones

Aumentamos la gramática: `S' → S`.

**Estado `S₀`** = clausura de `S' → •S`:

```
S₀:
  S' → • S
  S  → • if E then S
  S  → • if E then S else S
  S  → • a
```

Transiciones desde `S₀`:

- `goto(S₀, if)` → **`S₁`**
- `goto(S₀, a)` → **`S₂`**
- `goto(S₀, S)` → **`S₃`** (ítem `S' → S •` ⇒ **aceptar**)

**`S₁`** (tras `if`):

```
S  → if • E then S
S  → if • E then S else S
E  → • b
```

- `goto(S₁, b)` → **`S₄`**: `E → b •` (**reduce**)
- `goto(S₁, E)` → **`S₅`**: `S → if E • then S` / `S → if E • then S else S`

**`S₅`** (tras `E`):

```
S  → if E • then S
S  → if E • then S else S
```

- `goto(S₅, then)` → **`S₆`**:

```
S  → if E then • S
S  → if E then • S else S
S  → • if E then S
S  → if E then S • else S     ← ojo: cerradura agrega solo los • al inicio
S  → • if E then S else S
S  → • a
```

Nota: `S₆` contiene:

```
S  → if E then • S
S  → if E then • S else S
S  → • if E then S
S  → • if E then S else S
S  → • a
```

**`goto(S₆, if)` → `S₁`** (reutiliza), **`goto(S₆, a)` → `S₂`**,
**`goto(S₆, S)` → `S₇`**:

```
S₇:
  S  → if E then S •
  S  → if E then S • else S
```

**`S₇` es el estado conflictivo**: en la entrada `else` podemos:
- **shift** `else` (para completar `S → if E then S else S`), o
- **reduce** `S → if E then S` (para cerrar el `if` sin `else`).

- `goto(S₇, else)` → **`S₈`**:

```
S₈:
  S → if E then S else • S
  S → • if E then S
  S → • if E then S else S
  S → • a
```

- `goto(S₈, S)` → **`S₉`**: `S → if E then S else S •` (**reduce**)
- `goto(S₈, if)` → `S₁`, `goto(S₈, a)` → `S₂`.

Estados: `S₀…S₉`.

#### 2. Tabla SLR

FIRST/FOLLOW usados:

```
FIRST(S) = { if , a }        FOLLOW(S) = { $ , else }
FIRST(E) = { b }             FOLLOW(E) = { then }
```

| Estado | `if` | `a` | `b` | `then` | `else` | `$` | Acciones reduce |
|---|---|---|---|---|---|---|---|
| `S₀` | s1 | s2 | | | | | |
| `S₁` | | | s4 | | | | |
| `S₂` | | | | | | | **r `S → a`** (FOLLOW(S)) |
| `S₃` | | | | | | **aceptar** | |
| `S₄` | | | | r `E → b` | | | |
| `S₅` | | | | s6 | | | |
| `S₆` | s1 | s2 | | | | | |
| `S₇` | | | | | **¿s8 o r?** | **r `S → if E then S`** | reduce en `$` y `else` |
| `S₈` | s1 | s2 | | | | | |
| `S₉` | | | | | | | **r `S → if E then S else S`** (en `else`… ver abajo) |

Detalle de `S₉`: `reduce` en FOLLOW(S) = `{$ , else}`.

> **Conflicto en `S₇` con lookahead `else`:**
> - **shift** `else` → s8
> - **reduce** `S → if E then S` (porque `else ∈ FOLLOW(S)`)
>
> Estado `S₇` en `else` = **conflicto shift/reduce**.

#### 3. Resolución del conflicto

Regla clásica (como hace yacc/bison): **prioridad al shift**.

Interpretación sintáctica: si vemos `else`, **el `else` pertenece al `if` más
reciente** (el interno), resolviendo el *dangling else* hacia **dentro**.

Acción corregida en `S₇`, columna `else`: **s8**.

Tabla final corregida:

| Estado | `if` | `a` | `b` | `then` | `else` | `$` |
|---|---|---|---|---|---|---|
| `S₇` | | | | | **s8** | r `S → if E then S` |

#### 4. Trazado shift-reduce para `if b then if b then a else a`

Entrada: `if b then if b then a else a $`
(Pila con tope a la derecha; los símbolos se apilan como tokens/estados)

| # | Pila | Entrada | Acción |
|---|------|---------|--------|
| 1 | `0` | `if b then if b then a else a $` | shift `if` |
| 2 | `0 if 1` | `b then if b then a else a $` | shift `b` |
| 3 | `0 if 1 b 4` | `then if b then a else a $` | **reduce `E → b`** (pop 1, goto 5) |
| 4 | `0 if 1 E 5` | `then if b then a else a $` | shift `then` |
| 5 | `0 if 1 E 5 then 6` | `if b then a else a $` | shift `if` |
| 6 | `0 if 1 E 5 then 6 if 1` | `b then a else a $` | shift `b` |
| 7 | `0 if 1 E 5 then 6 if 1 b 4` | `then a else a $` | **reduce `E → b`** → goto 5 |
| 8 | `0 if 1 E 5 then 6 if 1 E 5` | `then a else a $` | shift `then` |
| 9 | `0 if 1 E 5 then 6 if 1 E 5 then 6` | `a else a $` | shift `a` |
| 10 | `… then 6 a 2` | `else a $` | **reduce `S → a`** → goto 7 |
| 11 | `0 if 1 E 5 then 6 S 7` | `else a $` | **conflicto** → resuelto con **shift `else`** |
| 12 | `0 if 1 E 5 then 6 S 7 else 8` | `a $` | shift `a` |
| 13 | `… else 8 a 2` | `$` | **reduce `S → a`** → goto 9 |
| 14 | `0 if 1 E 5 then 6 S 7 else 8 S 9` | `$` | **reduce `S → if E then S else S`** → goto 7 |
| 15 | `0 if 1 E 5 then 6 S 7` | `$` | **reduce `S → if E then S`** → goto 6… |

continuando: `S → if E then S` deja `S` en el estado 6 con lookahead `$`…
`goto(6, S) = 7` y `S₇` en `$` reduce `S → if E then S` → `goto(S₀,S)=S₃`:

| | | | |
|---|---|---|---|
| 16 | `0 S 3` | `$` | **aceptar** ✓ |

**Árbol resultante** (el `else` se anidó al `if` interno):

```
        S
   ┌────┴────┐
  if E then  S
   │         ┌┴──────────┐
   b        if E then S  else S
             │          ┌─┴─┐
             b         a   a
                          ↑
                    (else emparejado con el if INTERNO)
```

Esto confirma la resolución: *el `else` corresponde al `if` más cercano*. ∎

---

## Ejercicio 4 — Esquemas de traducción (atributos sintetizados y heredados)

### Enunciado

Considere la gramática de expresiones con esquema de traducción que genera
**cuádruplas** (código de 3 direcciones):

```bnf
E   → E1 + T   { E.code = gen("+", E1.code, T.code); }
E   → T        { E.code = T.code; }
T   → T1 * F   { E.code = gen("*", T1.code, F.code); }   (* usar T *)
T   → F        { T.code = F.code; }
F   → id       { F.code = id.lexval; }
F   → ( E )    { F.code = E.code; }
```

1. Identifique los **atributos sintetizados** y los **heredados**.
2. Clasifique el esquema: ¿S-heredado, L-heredado o ambos?
3. Aplique el esquema a `a + b * c` y escriba las **cuádruplas** generadas
   (use temporales `t1, t2, …`).
4. Muestre el **árbol de sintaxis anotado** (con los valores de `code`).

### Solución

#### 1. Atributos

| Atributo | Tipo | Definido en |
|----------|------|-------------|
| `E.code`, `T.code`, `F.code` | **sintetizado** (sube desde los hijos) | `E → E1 + T` etc. |
| `id.lexval` | sintetizado (lo aporta el lexer) | `F → id` |

En esta gramática **no hay atributos heredados**: todos los valores de `code`
se calculan desde las producciones hijas hacia arriba.

#### 2. Clasificación del esquema

Como todos los atributos son **sintetizados** y se evalúan en una pasada
postorden (hijos primero, después el padre), el esquema es:

> **S-heredado** (bottom-up). Es directamente implementable por un parser
> ascendente (LR) o por un analizador descendente con *backtracking* postorden.

#### 3. Cuádruplas para `a + b * c`

Aplicando el esquema (nota: se usará `T.code` en `T → T1 * F`):

Derivación:

```
E → E1 + T
E1 → T1
T1 → F1          (a)
T  → T2 * F
T2 → F2          (b)
F  → id          (c)
```

**Reglas aplicadas (postorden):**

1. `F(a)`: `code = a`
2. `T → F`: `T.code = a`
3. `E → T`: `E1.code = a`
4. `F(b)`: `code = b`
5. `T → F`: `T2.code = b`
6. `F(c)`: `code = c`
7. `T → T2 * F`: `T.code = gen("*", b, c)` → `t1 = b * c`
8. `E → E1 + T`: `E.code = gen("+", a, t1)` → `t2 = a + t1`

**Cuádruplas generadas:**

```
( * , b , c , t1 )
( + , a , t1 , t2 )
```

Resultado final en `E.code` = `t2`.

#### 4. Árbol de sintaxis anotado

```
                 E : t2
              ┌────┴────┐
          E1 : a        T : t1
            │       ┌───┴───┐
          T1 : a   T2 : b   F : c
            │         │
          F1 : a     F : b
            │
          id : a
```

Cada nodo muestra `code`:

- `F → id` aporta el nombre (`a`, `b`, `c`).
- `T → T2 * F` sintetiza `t1 = b * c`.
- `E → E1 + T` sintetiza `t2 = a + t1`.

El recorrido de evaluación es **postorden** (los hijos antes que el padre),
lo que confirma que es un esquema **S-heredado**. ∎

---

## Resumen de la semana

| Ejercicio | Concepto clave |
|-----------|----------------|
| 1 | Ambigüedad ⇒ un no terminal por nivel de precedencia |
| 2 | FIRST/FOLLOW + tabla LL(1) + traza de pila |
| 3 | Ítems LR(0), tabla SLR, conflicto shift/reduce (*dangling else*) |
| 4 | Atributos sintetizados ⇒ esquemas S-heredados ⇒ cuádruplas |

### Preguntas de repaso

1. ¿Toda gramática ambigua falla el test de LL(1)? ¿Y LR(1)?
2. Si en el Ejercicio 1 cambiamos a asociatividad **derecha** para `+`,
   ¿cómo cambiaría `E'`?
3. En el Ejercicio 3, ¿qué pasaría si en `S₇` eligiéramos **reduce** en vez de
   shift para `else`?
4. ¿Qué atributo **heredado** necesitaríamos para generar cuádruplas con
   *saltos* (por ejemplo en `while`)?
