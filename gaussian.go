package main

import (
    "errors"
    "math"
)

func computeLinearEq(m matrix, v vector) vector {
    rowEchelon := make(matrix, len(v))
    for i := range rowEchelon {
        rowEchelon[i] = make(vector, len(v) + 1)
        copy(rowEchelon[i], m[i])
        rowEchelon[i][len(v)] = v[i]
    }

    gaussianElim(rowEchelon)
    if rowEchelon[len(v) - 1][len(v) - 1] == 0.0 {
        panic(errors.New("non-unique solution detected"))
    }

    rem := make(matrix, len(v))
    rev := make(vector, len(v))
    for i := range v {
        rem[i] = make(vector, len(v))
        copy(rem[i], rowEchelon[i])
        rev[i] = rowEchelon[i][len(v)]
    }
    return backwardSub(rem, rev)
}

func gaussianElim(m matrix) {
    for i := range m {
        r := i + 1
        p := pivot(m, i)
        if i != p {
            tmp := m[p]
            m[p] = m[i]
            m[i] = tmp
        }

        for j := r; j < len(m); j++ {
            factor := m[j][i] / m[i][i]

            v := vector(m[i]).muls(-factor)
            m[j] = vector(m[j]).add(v)
        }
    }
}

func pivot(m matrix, r int) int {
    p := r
    for i := r; i < len(m); i++ {
        if math.Abs(m[i][r]) > math.Abs(m[p][r]) {
            p = i
        }
    }
    return p
}

func backwardSub(m matrix, v vector) vector {
    r := make(vector, len(v))
    for i := len(v) - 1; i >= 0; i-- {
        sigma := 0.0
        for j := i + 1; j < len(m); j++ {
            sigma = sigma + m[i][j] * r[j]
        }
        r[i] = (v[i] - sigma) / m[i][i]
    }
    return r
}

