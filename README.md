# push-swap

Projet Ynov (Bachelor Cybersécurité) en Go.

## Contexte

Pendant une enquête, les logs Windows arrivent dans le désordre. Chaque événement a un numéro unique, l'EventRecordID. Pour reconstruire la chronologie d'une attaque, il faut remettre ces numéros dans l'ordre.

Un numéro qui n'est pas un nombre, ou qui apparaît deux fois, veut dire que l'export est corrompu ou trafiqué. Dans ce cas, les programmes refusent les données et affichent `Error`.

## Les deux programmes

- **push-swap** : reçoit des nombres et affiche la liste des gestes qui les trient.
- **checker** : reçoit des nombres, lit des gestes, les applique, puis affiche `OK` si c'est trié et `KO` sinon.

On trie la pile `a` avec l'aide d'une pile `b`, et seulement avec ces gestes :

| Geste | Effet |
|---|---|
| `sa` / `sb` / `ss` | échange les 2 éléments du haut |
| `pa` / `pb` | pousse le haut d'une pile sur l'autre |
| `ra` / `rb` / `rr` | le haut part en bas |
| `rra` / `rrb` / `rrr` | le bas remonte en haut |

## Installation

```bash
git clone https://github.com/8r1c3-333/push_swap.git
cd push_swap
go build -o push-swap ./cmd/push-swap
go build -o checker ./cmd/checker
```

## Utilisation

```bash
$ ./push-swap "2 1 3 6 5 8"
$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
OK
$ ./push-swap "0 one 2 3"
Error
$ ./push-swap "48213 48209 48215 48209 48212"
Error
```

Sans argument, les deux programmes n'affichent rien.

## Structure

```
push_swap/
├── cmd/
│   ├── checker/main.go     # programme checker
│   └── push-swap/main.go   # programme push-swap
├── internal/
│   ├── parser/             # lecture et vérification des nombres
│   ├── stack/              # piles a et b + les 11 gestes
│   └── solver/             # algorithme de tri
└── go.mod
```

## Algorithme

1. Si la pile est déjà triée, on ne fait rien.
2. Tant qu'il reste plus de 3 nombres dans `a` :
   - on cherche le plus petit ;
   - on le fait monter en haut avec `ra` s'il est dans la 1ère moitié, sinon avec `rra` (le chemin le plus court) ;
   - on le pousse dans `b`.
3. Les 3 nombres restants sont triés directement, en 2 gestes maximum.
4. On ramène tout `b` dans `a` avec `pa`. Comme `b` contient les plus petits, tout revient dans l'ordre.

## Tests

```bash
go test ./...
```

Les tests vérifient le parser (nombres valides, erreurs, doublons), chaque geste des piles, et que la solution du solver trie bien les nombres.