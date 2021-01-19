package main

import (
    "fmt"
)

func main() {
    nstates := 0
    for nstates < 2 {
        fmt.Printf("Number of states (>1): ")
        fmt.Scanf("%d", &nstates)
    }

    states := make([]string, nstates)
    for i := range states {
        fmt.Printf("State #%d = ", i + 1)
        fmt.Scanf("%s", &states[i])
    }

    trans := make(matrix, nstates)
    for i := range trans {
        trans[i] = make(vector, nstates)
        for j := range trans[i] {
            fmt.Printf("P(%s -> %s) = ", states[i], states[j])
            fmt.Scanf("%f", &trans[i][j])
        }
    }

    fixed := computeFixed(trans)

    calcArgs := ""
    for i := range states {
        if i > 0 {
            calcArgs = fmt.Sprintf("%s ", calcArgs)
        }
        calcArgs = fmt.Sprintf("%s<P(%s)>", calcArgs, states[i])
    }

    fmt.Println("Commands:")
    fmt.Printf(" - calc <N> %s\n", calcArgs)
    fmt.Printf(" - fixed\n")
    fmt.Printf(" - exit\n")
    for {
        var cmd string
        fmt.Printf("$ ")
        fmt.Scanf("%s", &cmd)

        if cmd == "calc" {
            var p int
            fmt.Scanf("%d", &p)

            v := make(vector, nstates)
            for i := range v {
                fmt.Scanf("%f", &v[i])
            }

            vp := compute(trans, v, p)

            fmt.Printf("Probabilities after %d transitions:\n", p)
            for i := range states {
                fmt.Printf(" - P(%s) = %.6f\n", states[i], vp[i])
            }
        } else if cmd == "fixed" {
            fmt.Println("Fixed point:")
            for i := range states {
                fmt.Printf(" - P(%s) = %.6f\n", states[i], fixed[i])
            }
        } else if cmd == "exit" {
            break
        }
    }
}

func compute(trans matrix, v vector, n int) vector {
    m := newIdentityMatrix(len(trans))
    for i := 0; i < n; i++ {
        m = m.mul(trans)
    }
    return m.postMulv(v)
}

func computeFixed(trans matrix) vector {
    t := make(vector, len(trans))

    m := make(matrix, len(trans) - 1)
    for i := range m {
        m[i] = make(vector, len(trans) - 1)
        for j := range m[i] {
            m[i][j] = trans[i][j] - trans[len(trans) - 1][i]
            if i == j {
                m[i][j] = m[i][j] - 1.0
            }
        }
    }

    v := make(vector, len(trans) - 1)
    for i := range v {
        v[i] = -trans[len(trans) - 1][i]
    }

    vr := computeLinearEq(m.transpose(), v)

    copy(t, vr)
    t[len(t) - 1] = 1.0
    for i := 0; i < len(t) - 1; i++ {
        t[len(t) - 1] = t[len(t) - 1] - t[i]
    }

    return t
}

