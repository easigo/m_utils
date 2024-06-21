package test_case

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_count"
)

func Test_m_count() {

	arith()
}

func arith() {
	a := "0.1"
	b := "0.2"
	// result := m_count.Add(a, b)
	// result := m_count.Sub(a, b)
	// result := m_count.Mul(a, b)
	// result := m_count.Div(a, b)
	// result := m_count.Per(a, b)
	// result := m_count.PerCent(a, b)
	// result := m_count.Rose(a, b)
	result := m_count.Rose(a, b)

	fmt.Println("result", result)

}
