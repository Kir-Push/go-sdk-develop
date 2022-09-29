package sdk

import (
	"context"
	"github.com/absmartly/go-sdk/sdk/future"
	"github.com/absmartly/go-sdk/sdk/jsonmodels"
	"testing"
)

var contextData = jsonmodels.ContextData{Experiments: []jsonmodels.Experiment{{Id: 5}}}

type ClientABSMock struct {
}

func (c ClientABSMock) GetContextData() *future.Future {
	return future.Call(func() (future.Value, error) {
		return contextData, nil
	})
}

func (c ClientABSMock) Publish(event jsonmodels.PublishEvent) *future.Future {
	return future.Call(func() (future.Value, error) {
		return nil, nil
	})
}

func TestCreateContext(t *testing.T) {
	var config = ABSmartlyConfig{Client_: ClientABSMock{}}
	var abs = Create(config)
	var contextConfig = ContextConfig{Units_: map[string]string{"user_id": "1234567"}}
	var temp = abs.CreateContext(contextConfig)
	var result = temp
	assertAny(true, result != nil, t)
}

func TestContextWith(t *testing.T) {

	var config = ABSmartlyConfig{Client_: ClientABSMock{}}
	var abs = Create(config)
	var contextConfig = ContextConfig{Units_: map[string]string{"user_id": "1234567"}}
	var result = abs.CreateContextWith(contextConfig, contextData)
	assertAny(true, result.IsReady(), t)
	assertAny(true, result.ReadyFuture_ == nil, t)
	assertAny(true, result.Cassignments_ != nil, t)
	assertAny(map[string]string{"user_id": "1234567"}, result.Units_, t)
}

func TestGetContext(t *testing.T) {

	var config = ABSmartlyConfig{Client_: ClientABSMock{}, ContextDataProvider_: ClientABSMock{}}
	var abs = Create(config)
	var result, err = abs.GetContextData().Get(context.Background())
	assertAny(nil, err, t)
	assertAny(contextData, result, t)
}
