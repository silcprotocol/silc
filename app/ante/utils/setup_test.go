package utils_test

import (
	"testing"

	"github.com/silcprotocol/silc/app/ante/testutils"
	"github.com/stretchr/testify/suite"
)

type AnteTestSuite struct {
	*testutils.AnteTestSuite
}

func TestAnteTestSuite(t *testing.T) {
	baseSuite := new(testutils.AnteTestSuite)
	baseSuite.WithLondonHardForkEnabled(true)

	suite.Run(t, &AnteTestSuite{
		AnteTestSuite: baseSuite,
	})
}
