package controllers

import (
	"testing"

	"github.com/blackmagiqq/webproxy2/adapters/requests"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func TestMapRequestToDTOSuccess(t *testing.T) {
	var request requests.CalculatorGetServices
	faker.FakeData(&request)

	ctrl := &CalculatorController{}
	actual := ctrl.mapRequestToInputDTO(&request)

	require.Equal(t, request.Sender.CityID, actual.Sender.CityID)
	require.Equal(t, request.Sender.ContractID, actual.Sender.ContractID)

	require.Equal(t, request.Receiver.CityID, actual.Receiver.CityID)
	require.Equal(t, request.Receiver.ContractID, actual.Receiver.ContractID)
}
