package mpesa

import (
	"sync"
)

// responseCodes are codes returned by the M-Pesa API
// INS-0	Request processed successfully
// INS-1	Internal Error
// INS-6	Transaction Failed
// INS-9	Request timeout
// INS-13	Invalid Shortcode Used
// INS-20	Not All Parameters Provided. Please try again.
// INS-21	Parameter validations failed. Please try again.
// INS-28	Invalid ThirdPartyConversationID Used
// INS-996	API Being Used Outside Of Usage Time
// NS-997	API Not Enabled
// INS-998	Invalid Market
var responseCodes = map[string]string{ //nolint:gochecknoglobals
	"INS-0":   "Request processed successfully",
	"INS-1":   "Internal Error",
	"INS-6":   "Transaction Failed",
	"INS-9":   "Request timeout",
	"INS-13":  "Invalid Shortcode Used",
	"INS-20":  "Not All Parameters Provided. Please try again.",
	"INS-21":  "Parameter validations failed. Please try again.",
	"INS-28":  "Invalid ThirdPartyConversationID Used",
	"INS-996": "API Being Used Outside Of Usage Time",
	"INS-997": "API Not Enabled",
	"INS-998": "Invalid Market",
}

type codes struct {
	mu       sync.RWMutex
	codesMap map[string]string
}

func (c *codes) get(code string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// check if the code is available and return it
	// if not return "unknown code"
	if val, ok := c.codesMap[code]; ok {
		return val
	}

	return "unknown code"
}

var (
	once           sync.Once //nolint:gochecknoglobals
	singleInstance *codes    //nolint:gochecknoglobals
)

func getInstance() *codes {
	if singleInstance == nil {
		once.Do(
			func() {
				singleInstance = &codes{
					mu:       sync.RWMutex{},
					codesMap: responseCodes,
				}
			})
	}

	return singleInstance
}

// ResponseCode returns the description of a response code
func ResponseCode(code string) string {
	return getInstance().get(code)
}
