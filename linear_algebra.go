package main

import (
    "errors"
)

type matrix [][]float64
type vector []float64

func newIdentityMatrix(n int) matrix {
    m := make(matrix, n)
    for i := range m {
        m[i] = make(vector, n)
        for j := range m[i] {
            if i == j {
                m[i][j] = 1.0
            } else {
                m[i][j] = 0.0
            }
        }
    }
    return m
}

func (a matrix) mul(b matrix) matrix {
    ra := len(a)
    ca := len(a[0])
    rb := len(b)
    cb := len(b[0])

    if ca != rb {
        panic(errors.New("incompatible matrix sizes"))
    }

    rc := ra
    cc := cb
    c := make(matrix, rc)
    for i := range c {
        c[i] = make(vector, cc)
        for j := range c[i] {
            c[i][j] = 0.0
            for m := range a[i] {
                c[i][j] += a[i][m] * b[m][j]
            }
        }
    }

    return c
}

func (m matrix) transpose() matrix {
    t := make(matrix, len(m[0]))
    for i := range t {
        t[i] = make(vector, len(m))
        for j := range t[i] {
            t[i][j] = m[j][i]
        }
    }
    return t
}

func (m matrix) postMulv(v vector) vector {
    vm := matrix{v}
    return vm.mul(m)[0]
}

func (a vector) add(b vector) vector {
    if len(a) != len(b) {
        panic(errors.New("incompatible vector sizes"))
    }

    c := make(vector, len(a))
    for i := range c {
        c[i] = a[i] + b[i]
    }
    return c
}

func (v vector) muls(s float64) vector {
    r := make(vector, len(v))
    for i := range r {
        r[i] = s * v[i]
    }
    return r
}

