// Package bundle implements bundle composition analysis calculations.
//
// Functions:
//   ClassifyComponents   - Leaders/Fillers/Killers classification (Simon-Kucher)
//   DeadWeightRatio      - Share of components with <20% monthly usage
//   CrossSubsidyAnalysis - Net margin flows between high/low margin components
//   ComponentActivation  - Share activating each component within 30 days
//   MultiComponentUsage  - Share of customers using 3+ components
package bundle

import (
	"fmt"

	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/domain"
)

// Calculator implements all bundle module functions.
type Calculator struct{}

// New creates a new bundle calculator.
func New() *Calculator {
	return &Calculator{}
}

// ClassifyComponents classifies each component as Leader, Filler, or Killer.
//
// Rules:
//   High perceived value AND drives purchase intent -> Leader
//   Moderate perceived value AND low marginal cost -> Filler
//   Low perceived value AND high marginal cost -> Killer
//   Low perceived value AND low marginal cost -> Filler (if option value) or Killer (if dilution)
//   Removing increases WTP (removal_wtp_delta > 0) -> Killer
func (c *Calculator) ClassifyComponents(input *domain.AppraisalInput) (*domain.LFKResult, error) {
	if len(input.Components) == 0 {
		return nil, fmt.Errorf("component data required for classification")
	}

	result := &domain.LFKResult{}

	for _, comp := range input.Components {
		cls := domain.LFKClassification{
			Name:           comp.Name,
			PerceivedValue: comp.PerceivedValue,
			MarginalCost:   comp.MarginalCost,
		}

		// If removing increases WTP, it's a Killer regardless of other factors
		if comp.RemovalWTPDelta != nil && *comp.RemovalWTPDelta > 0 {
			cls.Classification = "killer"
			cls.Rationale = "removing increases WTP (dilution effect)"
			result.Killers++
			result.Classifications = append(result.Classifications, cls)
			continue
		}

		pv := 0.0
		if comp.PerceivedValue != nil {
			pv = *comp.PerceivedValue
		}
		mc := 0.0
		if comp.MarginalCost != nil {
			mc = *comp.MarginalCost
		}
		drivesPurchase := comp.DrivesPurchase != nil && *comp.DrivesPurchase

		switch {
		case pv >= 4.0 && drivesPurchase:
			cls.Classification = "leader"
			cls.Rationale = "high perceived value and drives purchase intent"
			result.Leaders++
		case pv >= 4.0:
			cls.Classification = "leader"
			cls.Rationale = "high perceived value"
			result.Leaders++
		case pv >= 2.5 && mc < pv:
			cls.Classification = "filler"
			cls.Rationale = "moderate perceived value at acceptable cost"
			result.Fillers++
		case pv < 2.5 && mc > pv:
			cls.Classification = "killer"
			cls.Rationale = "low perceived value with high marginal cost"
			result.Killers++
		case pv < 2.5:
			// Low value, low cost: check removal WTP delta
			if comp.RemovalWTPDelta != nil && *comp.RemovalWTPDelta < 0 {
				cls.Classification = "filler"
				cls.Rationale = "low value but has option value (removing decreases WTP)"
				result.Fillers++
			} else {
				cls.Classification = "filler"
				cls.Rationale = "low perceived value but low cost"
				result.Fillers++
			}
		default:
			cls.Classification = "filler"
			cls.Rationale = "default classification"
			result.Fillers++
		}

		result.Classifications = append(result.Classifications, cls)
	}

	return result, nil
}

// DeadWeightRatio calculates the share of components with <20% monthly active usage.
// Dead weight ratio = dead weight components / total components.
// Threshold: <40% is acceptable.
func (c *Calculator) DeadWeightRatio(input *domain.AppraisalInput) (*domain.DeadWeightResult, error) {
	if len(input.Components) == 0 {
		return nil, fmt.Errorf("component data required")
	}

	const usageThreshold = 0.20
	const ratioThreshold = 0.40

	result := &domain.DeadWeightResult{
		Threshold:      ratioThreshold,
		ComponentUsage: make(map[string]float64),
	}

	deadCount := 0
	for _, comp := range input.Components {
		usage := 0.0
		if comp.MonthlyActiveRate != nil {
			usage = *comp.MonthlyActiveRate
		} else if comp.UsageForecast != nil {
			usage = *comp.UsageForecast
		}

		result.ComponentUsage[comp.Name] = usage

		if usage < usageThreshold {
			deadCount++
			result.DeadWeight = append(result.DeadWeight, comp.Name)
		}
	}

	result.DeadWeightRatio = float64(deadCount) / float64(len(input.Components))
	result.Passes = result.DeadWeightRatio < ratioThreshold

	return result, nil
}

// CrossSubsidyAnalysis evaluates net margin contribution across components.
// High margin components subsidize low margin ones.
// CrossSubsidy = High-Margin Revenue - Low-Margin Subsidy Cost.
func (c *Calculator) CrossSubsidyAnalysis(input *domain.AppraisalInput) (*domain.CrossSubsidyResult, error) {
	if len(input.Components) == 0 {
		return nil, fmt.Errorf("component data required")
	}

	result := &domain.CrossSubsidyResult{}
	totalMargin := 0.0

	for _, comp := range input.Components {
		rev := 0.0
		if comp.RevenueContrib != nil {
			rev = *comp.RevenueContrib
		}
		cost := 0.0
		if comp.DirectCost != nil {
			cost = *comp.DirectCost
		} else if comp.MarginalCost != nil {
			cost = *comp.MarginalCost
		}

		margin := rev - cost
		totalMargin += margin

		csc := domain.CrossSubsidyComponent{
			Name:      comp.Name,
			Revenue:   rev,
			Cost:      cost,
			NetMargin: margin,
		}

		if margin >= 0 {
			csc.Role = "source"
			result.Sources = append(result.Sources, csc)
		} else {
			csc.Role = "recipient"
			result.Recipients = append(result.Recipients, csc)
		}
	}

	result.NetMargin = totalMargin
	result.Sustainable = totalMargin > 0

	return result, nil
}

// ComponentActivation calculates the share of customers activating each component
// within 30 days. Returns per-component activation rates.
// Target: >70% for Leaders, >40% for Fillers.
func (c *Calculator) ComponentActivation(input *domain.AppraisalInput) ([]domain.SingleValueResult, error) {
	if len(input.Components) == 0 {
		return nil, fmt.Errorf("component data required")
	}

	var results []domain.SingleValueResult
	for _, comp := range input.Components {
		rate := 0.0
		if comp.Activation30d != nil {
			rate = *comp.Activation30d
		}

		var interp string
		switch {
		case rate >= 0.70:
			interp = fmt.Sprintf("%s: strong_activation (leader-level)", comp.Name)
		case rate >= 0.40:
			interp = fmt.Sprintf("%s: moderate_activation (filler-level)", comp.Name)
		default:
			interp = fmt.Sprintf("%s: weak_activation (dead_weight_risk)", comp.Name)
		}

		results = append(results, domain.SingleValueResult{
			Value:          rate,
			Interpretation: interp,
		})
	}

	return results, nil
}

// MultiComponentUsage calculates the share of customers using 3+ bundle components.
// Rate = customers using 3+ / total bundle customers.
// Target: >60%.
func (c *Calculator) MultiComponentUsage(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Customers == nil {
		return nil, fmt.Errorf("customer metrics required")
	}
	if input.Customers.CustomersUsing3Plus == nil {
		return nil, fmt.Errorf("customers_using_3plus required")
	}
	if input.Customers.PremiumCustomers == nil || *input.Customers.PremiumCustomers <= 0 {
		return nil, fmt.Errorf("premium_customers required and must be positive")
	}

	rate := *input.Customers.CustomersUsing3Plus / *input.Customers.PremiumCustomers

	var interp string
	if rate >= 0.60 {
		interp = "healthy_multi_component_usage"
	} else {
		interp = "low_multi_component_usage_poor_bundle_composition"
	}

	return &domain.SingleValueResult{
		Value:          rate,
		Interpretation: interp,
	}, nil
}
