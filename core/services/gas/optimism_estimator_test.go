package gas_test

import (
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink/core/services/gas"
	"github.com/smartcontractkit/chainlink/core/services/gas/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_OptimismEstimator(t *testing.T) {
	t.Parallel()

	config := new(mocks.Config)
	client := new(mocks.OptimismRPCClient)
	o := gas.NewOptimismEstimator(config, client)

	calldata := []byte{0x00, 0x00, 0x01, 0x02, 0x03}
	var gasLimit uint64 = 80000

	t.Run("calling EstimateGas on unstarted estimator returns error", func(t *testing.T) {
		_, _, err := o.EstimateGas(calldata, gasLimit)
		assert.EqualError(t, err, "estimator is not started")
	})

	t.Run("calling EstimateGas on started estimator returns prices", func(t *testing.T) {
		client.On("Call", mock.Anything, "rollup_gasPrices").Return(nil).Run(func(args mock.Arguments) {
			res := args.Get(0).(*gas.OptimismGasPricesResponse)
			res.L1GasPrice = big.NewInt(42)
			res.L2GasPrice = big.NewInt(142)
		})

		require.NoError(t, o.Start())
		gasPrice, chainSpecificGasLimit, err := o.EstimateGas(calldata, gasLimit)
		require.NoError(t, err)
		assert.Equal(t, big.NewInt(15000000000), gasPrice)
		assert.Equal(t, 10008, int(chainSpecificGasLimit))
	})

	t.Run("calling BumpGas always returns error", func(t *testing.T) {
		_, _, err := o.BumpGas(big.NewInt(42), gasLimit)
		assert.EqualError(t, err, "bump gas is not supported for optimism")
	})
}