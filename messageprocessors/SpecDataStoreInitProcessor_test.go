package messageprocessors

import (
	m "github.com/manuviswam/gauge-go/gauge_messages"
	t "github.com/manuviswam/gauge-go/testsuit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnExecutionStatusResponseWithSameIdForSpecDataStoreInit(tst *testing.T) {
	msgId := int64(12345)
	context := &t.GaugeContext{
		Steps: make([]t.Step, 0),
	}

	msg := &m.Message{
		MessageType: m.Message_SpecDataStoreInit.Enum(),
		MessageId:   &msgId,
	}

	p := SpecDataStoreInitProcessor{}

	result := p.Process(msg, context)

	assert.Equal(tst, result.MessageType, m.Message_ExecutionStatusResponse.Enum())
	assert.Equal(tst, *result.MessageId, msgId)
}

func TestShouldResetSpecDataStore(tst *testing.T) {
	msgId := int64(12345)
	context := &t.GaugeContext{
		SpecStore: make(map[string]interface{}),
	}
	msg := &m.Message{
		MessageType: m.Message_ScenarioDataStoreInit.Enum(),
		MessageId:   &msgId,
	}

	context.SpecStore["foo"] = "bar"

	p := SpecDataStoreInitProcessor{}

	result := p.Process(msg, context)

	assert.Equal(tst, result.MessageType, m.Message_ExecutionStatusResponse.Enum())
	assert.Equal(tst, *result.MessageId, msgId)
	assert.Nil(tst, context.SpecStore["foo"])
}