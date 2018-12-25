package testhelper

import (
	"fmt"
	"testing"
)

type ISetUpHook interface {
	SetUp()
}

type ITearDown interface {
	TearDown()
}

type ISetUpTearDown interface {
	ISetUpHook
	ITearDown
}

type Tester = func(t *testing.T)

type NoopSetUpTearDown struct {
}

var _ ISetUpTearDown = &NoopSetUpTearDown{}

func (nst *NoopSetUpTearDown) SetUp() {
	fmt.Println("noop setup...")
}

func (nst *NoopSetUpTearDown) TearDown() {
	fmt.Println("noop teardown...")
}

func MakeTests(iSetUpTearDown ISetUpTearDown, t *testing.T, testers ...Tester) {
	iSetUpTearDown.SetUp()
	defer iSetUpTearDown.TearDown()

	for i, test := range testers {
		fmt.Printf("make test %d : \n", i)
		test(t)
	}
}

func BeforeEachMakeTests(iSetUpTearDown ISetUpTearDown, t *testing.T, testers ...Tester) {
	for i, test := range testers {
		iSetUpTearDown.SetUp()
		fmt.Printf("make test %d : \n", i)
		test(t)
	}
}

func AfterEachMakeTests(iSetUpTearDown ISetUpTearDown, t *testing.T, testers ...Tester) {
	for i, test := range testers {
		fmt.Printf("make test %d : \n", i)
		test(t)
		iSetUpTearDown.TearDown()
	}
}

func EachMakeTests(iSetUpTearDown ISetUpTearDown, t *testing.T, testers ...Tester) {
	for i, test := range testers {
		iSetUpTearDown.SetUp()
		fmt.Printf("make test %d : \n", i)
		test(t)
		iSetUpTearDown.TearDown()
	}
}
