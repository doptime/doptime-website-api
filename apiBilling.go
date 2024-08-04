package main

// import (
// 	"errors"
// 	"time"
// )

// type UsageRecord struct {
// 	UserID    string
// 	APIID     string
// 	Timestamp time.Time
// 	Params    map[string]interface{}
// 	Cost      float64
// }

// type BillingManager struct {
// 	usageRecords      []UsageRecord
// 	pricingStrategies map[string]pricing.PricingStrategy
// }

// func (bm *BillingManager) RecordUsage(userID, apiID string, params map[string]interface{}) error {
// 	strategy, exists := bm.pricingStrategies[apiID]
// 	if !exists {
// 		return errors.New("no pricing strategy found for API")
// 	}
// 	cost, err := strategy.CalculateCost(params)
// 	if err != nil {
// 		return err
// 	}
// 	record := UsageRecord{
// 		UserID:    userID,
// 		APIID:     apiID,
// 		Timestamp: time.Now(),
// 		Params:    params,
// 		Cost:      cost,
// 	}
// 	bm.usageRecords = append(bm.usageRecords, record)
// 	return nil
// }

// func (bm *BillingManager) GetUserUsage(userID string) ([]UsageRecord, error) {
// 	var records []UsageRecord
// 	for _, record := range bm.usageRecords {
// 		if record.UserID == userID {
// 			records = append(records, record)
// 		}
// 	}
// 	return records, nil
// }
