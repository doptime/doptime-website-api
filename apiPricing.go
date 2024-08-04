package main

// type PricingStrategy interface {
// 	CalculateCost(params map[string]interface{}) (float64, error)
// }

// type CallBasedPricing struct {
// 	RatePerCall float64
// }

// func (cbp *CallBasedPricing) CalculateCost(params map[string]interface{}) (float64, error) {
// 	calls := params["calls"].(int)
// 	return float64(calls) * cbp.RatePerCall, nil
// }

// type TrafficBasedPricing struct {
// 	RatePerRequestMB  float64
// 	RatePerResponseMB float64
// }

// func (dbp *TrafficBasedPricing) CalculateCost(params map[string]interface{}) (float64, error) {
// 	dataMB := params["dataMB"].(float64)
// 	return dataMB * dbp.RatePerKBIn, nil
// }

// type TokenBasedPricing struct {
// 	RatePerRequestToken  float64
// 	RatePerResponseToken float64
// }

// func (tbp *TokenBasedPricing) CalculateCost(params map[string]interface{}) (float64, error) {
// 	tokens := params["tokens"].(int)
// 	return float64(tokens) * tbp.RatePerTokenIn, nil
// }
