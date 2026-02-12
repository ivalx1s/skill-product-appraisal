// Package customer implements customer-facing KPI calculations.
//
// Functions:
//   ChurnRate              - Customers lost / total customers at start
//   RetentionRate          - 1 - churn rate
//   NPS                    - % promoters (9-10) - % detractors (0-6)
//   CSAT                   - Satisfied responses / total responses
//   ChurnReductionImpact   - (churn_before - churn_after) / churn_before
//   RevenueGrowthRate      - (revenue_t - revenue_t-1) / revenue_t-1
//   ServiceRevenueShare    - Add-on revenue / total revenue
package customer

import (
	"fmt"

	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/domain"
)

// Calculator implements all customer module functions.
type Calculator struct{}

// New creates a new customer calculator.
func New() *Calculator {
	return &Calculator{}
}

// ChurnRate calculates customers lost / total customers at start of period.
func (c *Calculator) ChurnRate(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Customers == nil {
		return nil, fmt.Errorf("customer metrics required")
	}
	m := input.Customers

	if m.LostCustomers == nil {
		return nil, fmt.Errorf("lost_customers required")
	}
	if m.CustomersStartPeriod == nil || *m.CustomersStartPeriod <= 0 {
		return nil, fmt.Errorf("customers_start_period required and must be positive")
	}

	rate := *m.LostCustomers / *m.CustomersStartPeriod

	return &domain.SingleValueResult{
		Value:          rate,
		Interpretation: fmt.Sprintf("churn_rate=%.2f%%", rate*100),
	}, nil
}

// RetentionRate calculates 1 - churn rate.
func (c *Calculator) RetentionRate(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	churnResult, err := c.ChurnRate(input)
	if err != nil {
		return nil, err
	}

	retention := 1.0 - churnResult.Value

	return &domain.SingleValueResult{
		Value:          retention,
		Interpretation: fmt.Sprintf("retention_rate=%.2f%%", retention*100),
	}, nil
}

// NPS calculates Net Promoter Score = % promoters (9-10) - % detractors (0-6).
func (c *Calculator) NPS(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Customers == nil {
		return nil, fmt.Errorf("customer metrics required")
	}
	m := input.Customers

	if m.PromotersPct == nil || m.DetractorsPct == nil {
		return nil, fmt.Errorf("promoters_pct and detractors_pct required")
	}

	nps := *m.PromotersPct - *m.DetractorsPct

	var interp string
	switch {
	case nps >= 50:
		interp = "excellent_nps"
	case nps >= 30:
		interp = "good_nps"
	case nps >= 0:
		interp = "moderate_nps"
	default:
		interp = "negative_nps"
	}

	return &domain.SingleValueResult{
		Value:          nps,
		Interpretation: interp,
	}, nil
}

// CSAT calculates satisfied responses / total responses.
func (c *Calculator) CSAT(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Customers == nil {
		return nil, fmt.Errorf("customer metrics required")
	}
	m := input.Customers

	if m.SatisfiedResponses == nil || m.TotalResponses == nil {
		return nil, fmt.Errorf("satisfied_responses and total_responses required")
	}
	if *m.TotalResponses <= 0 {
		return nil, fmt.Errorf("total_responses must be positive")
	}

	csat := *m.SatisfiedResponses / *m.TotalResponses

	var interp string
	if csat >= 0.80 {
		interp = "premium_level_satisfaction"
	} else if csat >= 0.60 {
		interp = "acceptable_satisfaction"
	} else {
		interp = "low_satisfaction"
	}

	return &domain.SingleValueResult{
		Value:          csat,
		Interpretation: interp,
	}, nil
}

// ChurnReductionImpact calculates (churn_before - churn_after) / churn_before.
// Measures the % improvement in churn after premium/bundle launch.
func (c *Calculator) ChurnReductionImpact(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Customers == nil {
		return nil, fmt.Errorf("customer metrics required")
	}
	m := input.Customers

	if m.ChurnBefore == nil || m.ChurnAfter == nil {
		return nil, fmt.Errorf("churn_before and churn_after required")
	}
	if *m.ChurnBefore <= 0 {
		return nil, fmt.Errorf("churn_before must be positive")
	}

	reduction := (*m.ChurnBefore - *m.ChurnAfter) / *m.ChurnBefore

	var interp string
	switch {
	case reduction >= 0.50:
		interp = "strong_churn_reduction"
	case reduction >= 0.25:
		interp = "moderate_churn_reduction"
	case reduction >= 0.05:
		interp = "modest_churn_reduction"
	case reduction > 0:
		interp = "minimal_churn_reduction"
	default:
		interp = "churn_increased"
	}

	return &domain.SingleValueResult{
		Value:          reduction,
		Interpretation: interp,
	}, nil
}

// RevenueGrowthRate calculates (revenue_current - revenue_prior) / revenue_prior.
func (c *Calculator) RevenueGrowthRate(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Customers == nil {
		return nil, fmt.Errorf("customer metrics required")
	}
	m := input.Customers

	if m.RevenueCurrentPeriod == nil || m.RevenuePriorPeriod == nil {
		return nil, fmt.Errorf("revenue_current_period and revenue_prior_period required")
	}
	if *m.RevenuePriorPeriod <= 0 {
		return nil, fmt.Errorf("revenue_prior_period must be positive")
	}

	growth := (*m.RevenueCurrentPeriod - *m.RevenuePriorPeriod) / *m.RevenuePriorPeriod

	return &domain.SingleValueResult{
		Value:          growth,
		Interpretation: fmt.Sprintf("revenue_growth=%.1f%%", growth*100),
	}, nil
}

// ServiceRevenueShare calculates add-on/service revenue / total revenue.
func (c *Calculator) ServiceRevenueShare(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Customers == nil {
		return nil, fmt.Errorf("customer metrics required")
	}
	m := input.Customers

	if m.AddOnRevenue == nil || m.TotalRevenue == nil {
		return nil, fmt.Errorf("add_on_revenue and total_revenue required")
	}
	if *m.TotalRevenue <= 0 {
		return nil, fmt.Errorf("total_revenue must be positive")
	}

	share := *m.AddOnRevenue / *m.TotalRevenue

	return &domain.SingleValueResult{
		Value:          share,
		Interpretation: fmt.Sprintf("service_revenue_share=%.1f%%", share*100),
	}, nil
}
