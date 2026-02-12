// Package pricing implements pricing analysis calculations.
//
// Functions:
//   BVR               - Bundle Value Ratio (sum of standalone prices / bundle price)
//   TierGapAnalysis   - Price/value gap analysis between adjacent tiers
//   CostFloor         - Minimum viable price based on cost structure
//   PriceValueRatio   - Customer-perceived value relative to price
//   PremiumPriceIndex  - Premium price relative to market average
//   BundleDiscount     - Effective discount vs standalone sum
package pricing

import (
	"fmt"
	"math"

	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/domain"
)

// Calculator implements all pricing module functions.
type Calculator struct{}

// New creates a new pricing calculator.
func New() *Calculator {
	return &Calculator{}
}

// BVR calculates Bundle Value Ratio = Sum(standalone prices) / bundle price.
// Interpretation: <1.0 negative, 1.0-1.3 marginal, 1.3-1.5 adequate, 1.5-2.0 strong, >2.0 very strong.
func (c *Calculator) BVR(input *domain.AppraisalInput) (*domain.BVRResult, error) {
	if input.Product == nil {
		return nil, fmt.Errorf("product definition required")
	}
	if len(input.Product.Components) == 0 {
		return nil, fmt.Errorf("product must have components for BVR calculation")
	}
	if input.Product.Price <= 0 {
		return nil, fmt.Errorf("product price must be positive")
	}

	standaloneSum := 0.0
	componentValues := make(map[string]float64)
	for _, comp := range input.Product.Components {
		standaloneSum += comp.StandalonePrice
		componentValues[comp.Name] = comp.StandalonePrice
	}

	bvr := standaloneSum / input.Product.Price

	var interp string
	switch {
	case bvr < 1.0:
		interp = "negative_value_proposition"
	case bvr < 1.3:
		interp = "marginal"
	case bvr < 1.5:
		interp = "adequate"
	case bvr < 2.0:
		interp = "strong"
	default:
		interp = "very_strong"
	}

	return &domain.BVRResult{
		BVR:             bvr,
		StandaloneSum:   standaloneSum,
		BundlePrice:     input.Product.Price,
		Interpretation:  interp,
		ComponentValues: componentValues,
	}, nil
}

// TierGapAnalysis evaluates price and value gaps between adjacent tiers.
// For each pair: PriceGap% = (upper - lower) / lower * 100.
// ValueGap = upper perceived value - lower perceived value.
// V/P ratio = ValueGap / PriceGapAbs. >1 = effective upsell, <1 = broken step.
func (c *Calculator) TierGapAnalysis(input *domain.AppraisalInput) (*domain.TierGapResult, error) {
	if len(input.Tiers) < 2 {
		return nil, fmt.Errorf("at least 2 tiers required for gap analysis")
	}

	// Sort tiers by level (assume they come in order, but validate)
	tiers := input.Tiers
	for i := 1; i < len(tiers); i++ {
		if tiers[i].Level <= tiers[i-1].Level {
			return nil, fmt.Errorf("tiers must be ordered by level (ascending)")
		}
	}

	var gaps []domain.TierGap
	for i := 0; i < len(tiers)-1; i++ {
		lower := tiers[i]
		upper := tiers[i+1]

		if lower.Price <= 0 {
			return nil, fmt.Errorf("tier %q has non-positive price", lower.Name)
		}

		priceGapAbs := upper.Price - lower.Price
		priceGapPct := (priceGapAbs / lower.Price) * 100.0

		gap := domain.TierGap{
			FromTier:    lower.Name,
			ToTier:      upper.Name,
			PriceGapAbs: priceGapAbs,
			PriceGapPct: priceGapPct,
		}

		// Value gap requires perceived value on both tiers
		if lower.PerceivedValue != nil && upper.PerceivedValue != nil {
			vGap := *upper.PerceivedValue - *lower.PerceivedValue
			gap.ValueGap = &vGap
			if priceGapAbs > 0 {
				ratio := vGap / priceGapAbs
				gap.ValueToPriceRatio = &ratio
			}
		}

		// Diagnosis
		switch {
		case gap.ValueToPriceRatio != nil && *gap.ValueToPriceRatio > 1.0:
			gap.Diagnosis = "effective_upsell"
		case gap.ValueToPriceRatio != nil && *gap.ValueToPriceRatio < 1.0:
			gap.Diagnosis = "broken_step"
		case gap.ValueToPriceRatio != nil:
			gap.Diagnosis = "neutral"
		case priceGapPct < 10:
			gap.Diagnosis = "gap_too_small"
		case priceGapPct > 80:
			gap.Diagnosis = "gap_too_large"
		default:
			gap.Diagnosis = "insufficient_data"
		}

		gaps = append(gaps, gap)
	}

	return &domain.TierGapResult{Gaps: gaps}, nil
}

// CostFloor calculates the minimum viable price:
// CostFloor = direct + partner/licensing + shared + CAC(amortized) + service + target margin.
func (c *Calculator) CostFloor(input *domain.AppraisalInput) (*domain.CostFloorResult, error) {
	if input.Financials == nil {
		return nil, fmt.Errorf("financial data required")
	}
	if input.Product == nil {
		return nil, fmt.Errorf("product definition required for price comparison")
	}

	f := input.Financials
	floor := 0.0

	if f.DirectCostPerCustomer != nil {
		floor += *f.DirectCostPerCustomer
	}
	if f.PartnerLicensingCost != nil {
		floor += *f.PartnerLicensingCost
	}
	if f.SharedCostPerCustomer != nil {
		floor += *f.SharedCostPerCustomer
	}
	if f.TotalAcquisitionSpend != nil && f.NewCustomersAcquired != nil && *f.NewCustomersAcquired > 0 {
		floor += *f.TotalAcquisitionSpend / *f.NewCustomersAcquired
	}
	if f.CustomerServiceCost != nil {
		floor += *f.CustomerServiceCost
	}
	if f.TargetMinMargin != nil {
		// Add margin as absolute amount: floor / (1 - margin%) - floor
		if *f.TargetMinMargin < 1.0 {
			floor = floor / (1.0 - *f.TargetMinMargin)
		}
	}

	margin := input.Product.Price - floor

	return &domain.CostFloorResult{
		CostFloor:    floor,
		CurrentPrice: input.Product.Price,
		Margin:       margin,
		ClearsFloor:  input.Product.Price >= floor,
	}, nil
}

// PriceValueRatio calculates perceived value / actual price.
// Requires customers.satisfied_responses (used as proxy for perceived value score) or product perceived value.
// >1.0 means customers perceive more value than they pay.
func (c *Calculator) PriceValueRatio(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Product == nil {
		return nil, fmt.Errorf("product definition required")
	}

	// Use the first tier's perceived value if available, or average of component perceived values
	var perceivedValue float64
	var found bool

	for _, tier := range input.Tiers {
		if tier.PerceivedValue != nil {
			perceivedValue = *tier.PerceivedValue
			found = true
			break
		}
	}

	if !found && len(input.Components) > 0 {
		sum := 0.0
		count := 0
		for _, comp := range input.Components {
			if comp.PerceivedValue != nil {
				sum += *comp.PerceivedValue
				count++
			}
		}
		if count > 0 {
			perceivedValue = sum / float64(count)
			found = true
		}
	}

	if !found {
		return nil, fmt.Errorf("perceived value data required (tier or component level)")
	}

	if input.Product.Price <= 0 {
		return nil, fmt.Errorf("product price must be positive")
	}

	ratio := perceivedValue / input.Product.Price

	var interp string
	if ratio > 1.0 {
		interp = "positive_value_perception"
	} else if ratio == 1.0 {
		interp = "neutral"
	} else {
		interp = "negative_value_perception"
	}

	return &domain.SingleValueResult{
		Value:          ratio,
		Interpretation: interp,
	}, nil
}

// PremiumPriceIndex calculates premium price / market average price.
func (c *Calculator) PremiumPriceIndex(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Product == nil {
		return nil, fmt.Errorf("product definition required")
	}
	if input.Market == nil || input.Market.MarketAveragePrice == nil {
		return nil, fmt.Errorf("market average price required")
	}
	if *input.Market.MarketAveragePrice <= 0 {
		return nil, fmt.Errorf("market average price must be positive")
	}

	index := input.Product.Price / *input.Market.MarketAveragePrice

	return &domain.SingleValueResult{
		Value:          index,
		Interpretation: fmt.Sprintf("premium_price_is_%.1fx_market_average", index),
	}, nil
}

// BundleDiscount calculates the effective discount percentage.
// Discount = 1 - (bundle price / standalone sum).
func (c *Calculator) BundleDiscount(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Product == nil {
		return nil, fmt.Errorf("product definition required")
	}
	if len(input.Product.Components) == 0 {
		return nil, fmt.Errorf("product must have components")
	}

	standaloneSum := 0.0
	for _, comp := range input.Product.Components {
		standaloneSum += comp.StandalonePrice
	}

	if standaloneSum <= 0 {
		return nil, fmt.Errorf("standalone sum must be positive")
	}

	discount := 1.0 - (input.Product.Price / standaloneSum)

	var interp string
	absPct := math.Abs(discount * 100)
	switch {
	case discount < 0:
		interp = "negative_discount_bundle_costs_more"
	case absPct < 10:
		interp = "weak_discount"
	case absPct < 15:
		interp = "noticeable_discount"
	case absPct <= 30:
		interp = "effective_range"
	case absPct <= 50:
		interp = "aggressive_discount"
	default:
		interp = "extreme_discount_risks_devaluation"
	}

	return &domain.SingleValueResult{
		Value:          discount,
		Interpretation: interp,
	}, nil
}
