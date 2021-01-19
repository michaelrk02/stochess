# stochess

Stochastic process calculator (uses markov chain)

## Sample program

```
[michael@arch-vivobook build] $ ./stochess
Number of states: 3
State #1 = A
State #2 = B
State #3 = C
P(A -> A) = 0.8
P(A -> B) = 0.1
P(A -> C) = 0.1
P(B -> A) = 0.1
P(B -> B) = 0.7
P(B -> C) = 0.2
P(C -> A) = 0.2
P(C -> B) = 0.2
P(C -> C) = 0.6
Commands:
 - calc <N> <P(A)> <P(B)> <P(C)>
 - fixed
 - exit
$ calc 2 0.3 0.5 0.2
Probabilities after 2 transitions:
 - P(A) = 0.356000
 - P(B) = 0.377000
 - P(C) = 0.267000
$ fixed
Fixed point:
 - P(A) = 0.421053
 - P(B) = 0.315789
 - P(C) = 0.263158
$ exit
```
