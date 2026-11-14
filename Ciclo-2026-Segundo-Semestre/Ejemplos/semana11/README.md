# Semana 11 — Puntero *link*, retorno en funciones recursivas y optimización

> Temas (`temas.csv`): *Puntero link y manejo de retorno en funciones recursivas*,
> *Métodos y fases de optimización en la generación de código ensamblador*.

Cada ejercicio incluye **enunciado** y **solución completa**.

---

## Ejercicio 1 — Registro de activación y puntero *link* (FP) en recursión

### Enunciado

Considere la función recursiva:

```c
int factorial(int n) {
    if (n <= 1)
        return 1;
    else
        return n * factorial(n - 1);
}
```

1. Describa el **registro de activación (RA)** de `factorial`: campos y
   función de cada uno.
2. Escriba el **preámbulo** y el **epílogo** de la función en términos de los
   punteros `SP` (stack pointer) y `FP` (frame pointer = puntero *link*).
3. Muestre el contenido del **stack** durante la llamada `factorial(4)`,
   indicando el valor de `FP` y la **instrucción de retorno** guardada.
4. Explique el papel del **puntero de retorno (PC de retorno)** y del
   **puntero de enlace (FP)** y por qué se guardan **por cada activación**.

### Solución

#### 1. Registro de activación (RA) de `factorial`

```
Dirección creciente (hacia arriba del stack)
┌─────────────────────────────┐
│  argumento: n               │  ← pasado por valor (copia)
├─────────────────────────────┤
│  valor de retorno           │  ← donde se escribe el resultado
├─────────────────────────────┤
│  FP antiguo (puntero link)  │  ← enlace al RA del llamador
├─────────────────────────────┤
│  dirección de retorno (PC)  │  ← instrucción siguiente a la llamada
├─────────────────────────────┤
│  registros salvados         │  ← (opcional) callee-saved
├─────────────────────────────┤
│  variable local: n (copia)  │  ← aquí `n` se usa como local
├─────────────────────────────┤
│  temporales / expresión     │  ← p. ej. para `n * fact(n-1)`
└─────────────────────────────┘
         │
         ▼  SP crece hacia abajo
```

| Campo | Propósito |
|-------|-----------|
| argumento `n` | valor pasado por el llamador |
| valor de retorno | espacio donde la llamada lee el resultado |
| **FP antiguo (link)** | **enlace al RA del llamador** — permite restaurar el frame anterior |
| dirección de retorno (PC) | **instrucción a ejecutar tras el `return`** |
| locales/`n` | copia local del parámetro (aquí se modifica) |
| temporales | resultados intermedios de `n * factorial(n-1)` |

#### 2. Preámbulo y epílogo

**Preámbulo** (al inicio de la función):

```asm
factorial:
    # preámbulo: crear el nuevo frame
    addi  sp, sp, -16      # reservar 16 bytes en el stack
    sw    ra, 12(sp)       # guardar dirección de retorno (PC)
    sw    fp, 8(sp)        # guardar el FP antiguo (puntero link)
    addi  fp, sp, 16       # FP apunta al inicio del frame del llamador
    sw    a0, 0(sp)        # guardar el argumento n (a0) en el frame
```

**Epílogo** (antes de cada `return`):

```asm
    # epílogo: restaurar y volver
    lw    ra, 12(sp)       # restaurar dirección de retorno
    lw    fp, 8(sp)        # restaurar FP del llamador (desenlazar)
    addi  sp, sp, 16       # liberar el frame
    ret                    # jalr zero, ra, 0   → salta a la instrucción de retorno
```

#### 3. Stack durante `factorial(4)`

Llamada: `r = factorial(4)`. Supongamos que la llamada a `factorial` está en
la dirección `L_ret` de la función `main`; el valor retornado por la llamada
debe ir a `r`.

**Activación A0 (de `main`, que llama a `factorial(4)`):**

```
FP_main ──► ┌───────────────────────┐
            │ ra_main (retorno)     │
            │ FP_main (link)        │
            │ n = 4  (arg)          │  a0 = 4
            └───────────┬───────────┘
                        │  SP (baja)
```

**Activación A1 — `factorial(4)`:**

```
FP1 ─────►  ┌───────────────────────┐
            │ ra = L_ret            │  ← instrucción tras la llamada
            │ FP = FP_main (link)   │  ← enlace a A0
            │ n = 4                 │
            │ temporales            │  (para 4 * fact(3))
            └───────────┬───────────┘
```

**Activación A2 — `factorial(3)`:**

```
FP2 ─────►  ┌───────────────────────┐
            │ ra = L_ret (de A1)    │  ← vuelve a A1
            │ FP = FP1 (link)       │  ← enlace a A1
            │ n = 3                 │
            └───────────┬───────────┘
```

**Activación A3 — `factorial(2)`:**

```
FP3 ─────►  ┌───────────────────────┐
            │ ra = L_ret (de A2)    │
            │ FP = FP2 (link)       │
            │ n = 2                 │
            └───────────┬───────────┘
```

**Activación A4 — `factorial(1)`** (caso base):

```
FP4 ─────►  ┌───────────────────────┐
            │ ra = L_ret (de A3)    │
            │ FP = FP3 (link)       │
            │ n = 1                 │
            └───────────────────────┘
   n <= 1  →  return 1  (no hay llamada recursiva)
```

**Desenlaces (retornos) — desapilando:**

| Paso | Acción | Resultado | FP tras volver |
|------|--------|-----------|----------------|
| R1 | A4 retorna `1` a A3 (vía `ra` de A4 = L_ret de A3) | `3?` no: A3 computa `2 * 1 = 2` | FP3 |
| R2 | A3 retorna `2` a A2 | A2 computa `3 * 2 = 6` | FP2 |
| R3 | A2 retorna `6` a A1 | A1 computa `4 * 6 = 24` | FP1 |
| R4 | A1 retorna `24` a `main` | `r = 24` | FP_main |

Cada `ret` usa el `ra` **guardado en su propio frame** (por eso se guarda en
el preámbulo): como la recursión reutiliza la misma dirección de código, el
`ra` de cada nivel es distinto y **debe guardarse por activación**.

#### 4. FP (link) vs PC de retorno

| | **PC de retorno (ra)** | **FP / puntero link** |
|---|---|---|
| ¿Qué guarda? | dirección de la **instrucción siguiente** a la llamada | dirección del **frame del llamador** |
| ¿Para qué? | saber **dónde continuar** tras el `return` | **desenlazar** el frame actual y acceder a args/locales del llamador |
| ¿Quién lo usa? | la instrucción `ret` (`jalr`) | el epílogo y el debugging/`backtrace` |
| Se guarda en cada activación | **sí** (porque la recursión reutiliza el mismo código) | **sí** (para reconstruir la cadena de llamadas) |

Por qué por activación: en la recursión, todas las llamadas ejecutan **el
mismo código** (la misma dirección de retorno `L_ret` en el mismo nivel), pero
cada nivel necesita **su propia copia** del frame (sus `n`, sus temporales).
Sin guardar `ra` y `FP` en el preámbulo, la segunda llamada **sobrescribiría**
los del primer nivel y el retorno sería incorrecto. ∎

---

## Ejercicio 2 — Retorno correcto vía puntero de enlace (cadena de activaciones)

### Enunciado

Continuando con `factorial(4)`:

1. Muestre la **cadena de enlaces FP** (`FP4 → FP3 → FP2 → FP1 → FP_main`)
   y el uso de `backtrace`.
2. Explique paso a paso cómo, tras el `return` en el caso base, se **restaura
   cada `ra` y `FP`** hasta llegar a `main`.
3. ¿Qué ocurriría si el epílogo **olvidara** restaurar el `FP` (no desenlazar)?

### Solución

#### 1. Cadena de enlaces FP

```
FP_main ──► A0 (main)
   ▲
   │ link
FP1 ──► A1 (fact 4) ──ra──► L_ret en main
   ▲
   │ link
FP2 ──► A2 (fact 3) ──ra──► L_ret en A1
   ▲
   │ link
FP3 ──► A3 (fact 2) ──ra──► L_ret en A2
   ▲
   │ link
FP4 ──► A4 (fact 1) ──ra──► L_ret en A3
```

Un *backtrace* (depurador) recorre `FP` hacia arriba para reconstruir la
pila de llamadas:

```
#0 factorial (n=1)   ← FP4
#1 factorial (n=2)   ← FP3
#2 factorial (n=3)   ← FP2
#3 factorial (n=4)   ← FP1
#4 main ()           ← FP_main
```

#### 2. Retorno paso a paso desde el caso base

`A4` (fact 1) cumple `n <= 1` ⇒ `return 1`:

```asm
# A4:  a0 = 1  (valor de retorno)
    lw    ra, 12(sp)      # ra ← L_ret de A3 (la llamada que originó A4)
    lw    fp, 8(sp)       # fp ← FP3         (desenlazar a A4)
    addi  sp, sp, 16      # liberar frame de A4
    ret                   # salta a L_ret en A3, con a0 = 1
```

`A3` (fact 2) recibe `a0 = 1`, computa `n * result = 2 * 1 = 2`:

```asm
    # a0 ← 2 * 1 = 2
    lw    ra, 12(sp)      # ra ← L_ret de A2
    lw    fp, 8(sp)       # fp ← FP2
    addi  sp, sp, 16
    ret                   # vuelve a A2 con a0 = 2
```

`A2` (fact 3): `a0 ← 3 * 2 = 6`, retorna a A1.
`A1` (fact 4): `a0 ← 4 * 6 = 24`, retorna a `main`.
`main`: `r = 24`.

En cada paso, el `lw fp, 8(sp)` **restaura el enlace** guardado, de modo que
`FP` siempre apunta al frame del llamador **correcto** — la cadena no se rompe
aunque la recursión sea profunda.

#### 3. Si el epílogo olvida restaurar el FP

```
Error: FP queda "apuntando" al frame actual (o a uno ya liberado)
```

Consecuencias:

| Síntoma | Explicación |
|---------|-------------|
| `backtrace` roto | el depurador no puede recorrer la cadena de llamadas |
| acceso a args del llamador incorrecto | el llamador lee su argumento vía `FP + offset`; si `FP` no se restauró, apunta al frame equivocado |
| corrupción de datos tras varios retornos | el `FP` "viejo" sigue apuntando a memoria que ya fue reutilizada por otra activación (*use after free* del frame) |
| en algunos ABIs, crash | si el epílogo hace `ld fp, ...` desde un `FP` incorrecto |

Es decir: **sin desenlazar el FP, la cadena de activaciones se corrompe** y
los accesos al frame del llamador dejan de ser válidos. ∎

---

## Ejercicio 3 — Optimización de la generación de código

### Enunciado

Dado el siguiente programa fuente y su **C3D sin optimizar**:

```c
int a = 2;
int b = 3;
int c = a + b;
int d = a + b;
int e = c * 0;         // ¿se usa luego?
int f = d + (2 * 3);
if (0 > 1)
    g = 100;           // nunca se ejecuta
```

C3D de referencia (sin optimizar):

```
1: a = 2
2: b = 3
3: t0 = a + b
4: c = t0
5: t1 = a + b
6: d = t1
7: t2 = c * 0
8: e = t2
9: t3 = 2 * 3
10: t4 = d + t3
11: f = t4
12: t5 = 0 > 1
13: ifFalse t5 goto L1
14: g = 100
15: L1:
```

Aplique, explicando cada paso, las siguientes optimizaciones:

1. **Plegado de constantes** (*constant folding*)
2. **Reducción de fuerza** (*strength reduction*)
3. **Propagación de constantes**
4. **Eliminación de código muerto** y **común subexpresión (CSE)**

Finalmente, escriba el **C3D optimizado** y compare el número de instrucciones.

### Solución

#### 1. Plegado de constantes

Sustituir operaciones cuyos operandos **son constantes** por su resultado en
tiempo de compilación.

| Instrucción | Operación | Plegado |
|---|---|---|
| `t3 = 2 * 3` | `2 * 3` | `t3 = 6` |
| `t2 = c * 0` | (aún no: `c` no es constante conocida aquí) | se ve en §3 |
| `t5 = 0 > 1` | `0 > 1` | `t5 = 0` (falso) |

Tras plegar:

```
t3 = 6          (en lugar de t3 = 2 * 3)
t5 = 0          (en lugar de t5 = 0 > 1)
```

#### 2. Reducción de fuerza

Sustituir operaciones **caras** por equivalentes **más baratas**.

| Expresión | Equivalente barato |
|-----------|---------------------|
| `x * 0` | `0` |
| `x * 1` | `x` |
| `x * 2` | `x + x` (o `x << 1`) |
| `x / 2` | `x >> 1` (enteros con signo: cuidado) |

Aplicando a `t2 = c * 0`:

```
t2 = 0        (reducción: c * 0 → 0)
```

Aplicando a `t3 = 6` (ya plegado) no hay nada más; si tuviéramos `d * 2`
sería `d << 1`.

#### 3. Propagación de constantes

Si una variable **siempre tiene el mismo valor constante**, sustituir sus
usos por ese valor.

Sabemos:

- `a = 2`, `b = 3` (constantes propagables)
- `t0 = a + b = 2 + 3 = 5` ⇒ tras plegar la suma: `c = 5`
- `t1 = a + b = 5` ⇒ `d = 5`
- `t2 = c * 0 = 0` ⇒ `e = 0`
- `t3 = 6`, `t4 = d + t3 = 5 + 6 = 11` ⇒ `f = 11`

Propagación completa:

```
c = 5
d = 5
e = 0
f = 11
```

#### 4. Eliminación de código muerto y CSE

**CSE (común subexpresión):** `t0 = a + b` y `t1 = a + b` calculan **lo mismo**.
Reutilizamos `t0` (o directamente el constante 5):

```
# t1 se elimina; d usa t0 (o el valor propagado)
```

**Código muerto:** `if (0 > 1) g = 100;` — la condición es **siempre falsa**
(`t5 = 0`), por lo que la rama `g = 100` **nunca** se ejecuta. Se elimina
todo el bloque `if`:

```
# se eliminan: t5, ifFalse, g = 100, etiqueta L1 (y el goto asociado)
```

#### 5. C3D optimizado

Combinando todas las optimizaciones:

```
# a, b, c, d, e, f son ahora constantes conocidas
a = 2
b = 3
c = 5          # propagación: 2 + 3
d = 5          # CSE + propagación: mismo valor que c
e = 0          # reducción de fuerza: c * 0 → 0
f = 11         # propagación: 5 + 6
# el bloque if (0 > 1) g = 100 se eliminó por completo (código muerto)
```

#### Comparación

| Métrica | Sin optimizar | Optimizado |
|---------|---------------|------------|
| Instrucciones C3D | 15 | 5 |
| Temporales | 6 (`t0`–`t5`) | 0 |
| Reducción | — | **≈ 67 % menos** |

> Fases de optimización recordadas en clase:
> **1)** plegado de constantes → **2)** propagación de constantes →
> **3)** reducción de fuerza → **4)** CSE → **5)** eliminación de código muerto**.
> Se aplican iterativamente hasta alcanzar un punto fijo. ∎

---

## Resumen de la semana

| Ejercicio | Concepto clave |
|-----------|----------------|
| 1 | RA con `ra` (PC retorno) + `FP` (link); preámbulo/epílogo; stack de factorial(4) |
| 2 | Cadena FP → backtrace; sin desenlazar el FP se corrompe el frame del llamador |
| 3 | Constant folding, propagación, strength reduction, CSE, dead-code elimination |

### Preguntas de repaso

1. ¿Por qué el `ra` debe guardarse **en el preámbulo** y no solo usarse?
2. En RISC-V, ¿qué instrucción compone el `ret` (`jalr` con qué argumentos)?
3. Si una función **no es recursiva**, ¿podríamos ahorrar el guardado del `FP`?
   ¿Y del `ra`?
4. ¿Por qué `x * 0` se puede eliminar, pero `x * 0` con `x` de tipo `float`
   puede ser problemático (`NaN`, `-0`)?
5. ¿En qué orden aplicarías las 5 optimizaciones del Ejercicio 3 y por qué?
