// Package product implements product performance KPI calculations.
//
// Functions:
//   PenetrationRate        - Premium customers / total customer base
//   MigrationRate          - Customers upgrading / eligible base per period
//   CannibalizationRate    - Migrated existing / total premium customers
//   CrossSellRate          - Premium buying add-ons / total premium customers
//   FeatureUtilizationRate - Features used per customer / total available features
//   ComponentActivationRate - Customers activating component / total bundle customers
//   AttachRate             - Customers using component monthly / total bundle customers
//   TrialConversion        - Paid conversions / trial users
package product

import (
	"fmt"

	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/domain"
)

// Calculator implements all product module functions.
type Calculator struct{}

// New creates a new product calculator.
func New() *Calculator {
	return &Calculator{}
}

// PenetrationRate calculates premium customers / total customer base.
func (c *Calculator) PenetrationRate(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Customers == nil {
		return nil, fmt.Errorf("customer metrics required")
	}
	m := input.Customers

	if m.PremiumCustomers == nil || m.TotalCustomers == nil {
		return nil, fmt.Errorf("premium_customers and total_customers required")
	}
	if *m.TotalCustomers <= 0 {
		return nil, fmt.Errorf("total_customers must be positive")
	}

	rate := *m.PremiumCustomers / *m.TotalCustomers

	var interp string
	switch {
	case rate >= 0.25:
		interp = "high_penetration"
	case rate >= 0.10:
		interp = "moderate_penetration"
	case rate >= 0.05:
		interp = "growing_penetration"
	default:
		interp = "low_penetration"
	}

	return &domain.SingleValueResult{
		Value:          rate,
		Interpretation: interp,
	}, nil
}

// MigrationRate calculates customers upgrading / eligible base per period.
func (c *Calculator) MigrationRate(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Customers == nil {
		return nil, fmt.Errorf("customer metrics required")
	}
	m := input.Customers

	if m.UpgradedCustomers == nil || m.EligibleBase == nil {
		return nil, fmt.Errorf("upgraded_customers and eligible_base required")
	}
	if *m.EligibleBase <= 0 {
		return nil, fmt.Errorf("eligible_base must be positive")
	}

	rate := *m.UpgradedCustomers / *m.EligibleBase

	return &domain.SingleValueResult{
		Value:          rate,
		Interpretation: fmt.Sprintf("migration_rate=%.2f%%", rate*100),
	}, nil
}

// CannibalizationRate calculates migrated existing customers / total premium customers.
// High cannibalization (>50%) without revenue uplift signals pricing failure.
func (c *Calculator) CannibalizationRate(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Customers == nil {
		return nil, fmt.Errorf("customer metrics required")
	}
	m := input.Customers

	if m.MigratedFromStandalone == nil || m.PremiumCustomers == nil {
		return nil, fmt.Errorf("migrated_from_standalone and premium_customers required")
	}
	if *m.PremiumCustomers <= 0 {
		return nil, fmt.Errorf("premium_customers must be positive")
	}

	rate := *m.MigratedFromStandalone / *m.PremiumCustomers

	var interp string
	switch {
	case rate > 0.50:
		interp = "high_cannibalization_risk"
	case rate > 0.30:
		interp = "moderate_cannibalization"
	default:
		interp = "acceptable_cannibalization"
	}

	return &domain.SingleValueResult{
		Value:          rate,
		Interpretation: interp,
	}, nil
}

// CrossSellRate calculates premium customers buying add-ons / total premium customers.
func (c *Calculator) CrossSellRate(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Customers == nil {
		return nil, fmt.Errorf("customer metrics required")
	}
	m := input.Customers

	if m.PremiumBuyingAddons == nil || m.PremiumCustomers == nil {
		return nil, fmt.Errorf("premium_buying_addons and premium_customers required")
	}
	if *m.PremiumCustomers <= 0 {
		return nil, fmt.Errorf("premium_customers must be positive")
	}

	rate := *m.PremiumBuyingAddons / *m.PremiumCustomers

	return &domain.SingleValueResult{
		Value:          rate,
		Interpretation: fmt.Sprintf("cross_sell_rate=%.2f%%", rate*100),
	}, nil
}

// FeatureUtilizationRate calculates features used per customer / total available features.
// Target: >60%.
func (c *Calculator) FeatureUtilizationRate(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Customers == nil {
		return nil, fmt.Errorf("customer metrics required")
	}
	m := input.Customers

	if m.FeaturesUsedPerCustomer == nil || m.TotalAvailableFeatures == nil {
		return nil, fmt.Errorf("features_used_per_customer and total_available_features required")
	}
	if *m.TotalAvailableFeatures <= 0 {
		return nil, fmt.Errorf("total_available_features must be positive")
	}

	rate := *m.FeaturesUsedPerCustomer / *m.TotalAvailableFeatures

	var interp string
	if rate >= 0.60 {
		interp = "healthy_feature_utilization"
	} else {
		interp = "low_utilization_over_provisioning_risk"
	}

	return &domain.SingleValueResult{
		Value:          rate,
		Interpretation: interp,
	}, nil
}

// ComponentActivationRate calculates customers activating a component / total bundle customers.
// Uses component-level activation_30d data. Returns per-component rates.
// Target: >70% for Leaders, >40% for Fillers.
func (c *Calculator) ComponentActivationRate(input *domain.AppraisalInput) ([]domain.SingleValueResult, error) {
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
			interp = fmt.Sprintf("%s: leader_level_activation", comp.Name)
		case rate >= 0.40:
			interp = fmt.Sprintf("%s: filler_level_activation", comp.Name)
		default:
			interp = fmt.Sprintf("%s: below_target", comp.Name)
		}

		results = append(results, domain.SingleValueResult{
			Value:          rate,
			Interpretation: interp,
		})
	}

	return results, nil
}

// AttachRate calculates customers using component monthly / total bundle customers.
// Uses monthly_active_rate from component data.
func (c *Calculator) AttachRate(input *domain.AppraisalInput) ([]domain.SingleValueResult, error) {
	if len(input.Components) == 0 {
		return nil, fmt.Errorf("component data required")
	}

	var results []domain.SingleValueResult
	for _, comp := range input.Components {
		rate := 0.0
		if comp.MonthlyActiveRate != nil {
			rate = *comp.MonthlyActiveRate
		}

		var interp string
		if rate < 0.20 {
			interp = fmt.Sprintf("%s: declining_attach_dead_weight_signal", comp.Name)
		} else {
			interp = fmt.Sprintf("%s: active_component", comp.Name)
		}

		results = append(results, domain.SingleValueResult{
			Value:          rate,
			Interpretation: interp,
		})
	}

	return results, nil
}

// TrialConversion calculates paid conversions / trial users.
// Benchmarks: self-serve 3-5%, sales-assisted 5-7%, top 8-15%.
func (c *Calculator) TrialConversion(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Customers == nil {
		return nil, fmt.Errorf("customer metrics required")
	}
	m := input.Customers

	if m.PaidConversions == nil || m.TrialUsers == nil {
		return nil, fmt.Errorf("paid_conversions and trial_users required")
	}
	if *m.TrialUsers <= 0 {
		return nil, fmt.Errorf("trial_users must be positive")
	}

	rate := *m.PaidConversions / *m.TrialUsers

	var interp string
	switch {
	case rate >= 0.08:
		interp = "top_performer_conversion"
	case rate >= 0.05:
		interp = "sales_assisted_level"
	case rate >= 0.03:
		interp = "self_serve_level"
	default:
		interp = "below_benchmark"
	}

	return &domain.SingleValueResult{
		Value:          rate,
		Interpretation: interp,
	}, nil
}
