// Package financial implements financial viability calculations.
//
// Functions:
//   UnitEconomics          - Revenue, cost, margin per customer
//   GrossMarginPerCustomer - (Revenue - COGS) / customer count
//   CLV                    - Customer Lifetime Value = RPC * margin% * lifespan
//   CACPayback             - Months to recover customer acquisition cost
//   BreakEven              - Units needed: fixed costs / contribution margin
//   CannibalizationNet     - Net revenue after migration losses
//   StressTest             - Margin under costs+20%, growth-30%
//   IncrementalRevenue     - Bundle revenue - lost standalone revenue per customer
//   RevenueUplift          - (premium RPC - base RPC) / base RPC
package financial

import (
	"fmt"

	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/domain"
)

// Calculator implements all financial module functions.
type Calculator struct{}

// New creates a new financial calculator.
func New() *Calculator {
	return &Calculator{}
}

// UnitEconomics calculates per-customer revenue, cost, and margin.
// Revenue per customer = total product revenue / average customer count.
// Cost per customer = (COGS + partner + shared + service) / average customer count.
// Margin = revenue - cost. Viable if margin > 0.
func (c *Calculator) UnitEconomics(input *domain.AppraisalInput) (*domain.UnitEconomicsResult, error) {
	if input.Financials == nil {
		return nil, fmt.Errorf("financial data required")
	}
	f := input.Financials

	if f.AverageCustomerCount == nil || *f.AverageCustomerCount <= 0 {
		return nil, fmt.Errorf("average_customer_count required and must be positive")
	}

	revenue := 0.0
	if f.RevenuePerCustomer != nil {
		revenue = *f.RevenuePerCustomer
	} else if f.TotalProductRevenue != nil {
		revenue = *f.TotalProductRevenue / *f.AverageCustomerCount
	} else {
		return nil, fmt.Errorf("revenue_per_customer or total_product_revenue required")
	}

	cost := 0.0
	if f.DirectCostPerCustomer != nil {
		cost += *f.DirectCostPerCustomer
	}
	if f.PartnerLicensingCost != nil {
		cost += *f.PartnerLicensingCost
	}
	if f.SharedCostPerCustomer != nil {
		cost += *f.SharedCostPerCustomer
	}
	if f.CustomerServiceCost != nil {
		cost += *f.CustomerServiceCost
	}

	margin := revenue - cost
	marginPct := 0.0
	if revenue > 0 {
		marginPct = margin / revenue
	}

	return &domain.UnitEconomicsResult{
		RevenuePerCustomer: revenue,
		CostPerCustomer:    cost,
		MarginPerCustomer:  margin,
		MarginPct:          marginPct,
		Viable:             margin > 0,
	}, nil
}

// GrossMarginPerCustomer calculates (Revenue - COGS) / average customer count.
func (c *Calculator) GrossMarginPerCustomer(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Financials == nil {
		return nil, fmt.Errorf("financial data required")
	}
	f := input.Financials

	if f.TotalProductRevenue == nil {
		return nil, fmt.Errorf("total_product_revenue required")
	}
	if f.COGS == nil {
		return nil, fmt.Errorf("cogs required")
	}
	if f.AverageCustomerCount == nil || *f.AverageCustomerCount <= 0 {
		return nil, fmt.Errorf("average_customer_count required and must be positive")
	}

	grossMargin := (*f.TotalProductRevenue - *f.COGS) / *f.AverageCustomerCount

	return &domain.SingleValueResult{
		Value:          grossMargin,
		Interpretation: fmt.Sprintf("gross_margin_per_customer=%.2f", grossMargin),
	}, nil
}

// CLV calculates Customer Lifetime Value = RPC * gross_margin% * average_lifespan_months.
func (c *Calculator) CLV(input *domain.AppraisalInput) (*domain.CLVResult, error) {
	if input.Financials == nil {
		return nil, fmt.Errorf("financial data required")
	}
	f := input.Financials

	rpc := 0.0
	if f.RevenuePerCustomer != nil {
		rpc = *f.RevenuePerCustomer
	} else if f.TotalProductRevenue != nil && f.AverageCustomerCount != nil && *f.AverageCustomerCount > 0 {
		rpc = *f.TotalProductRevenue / *f.AverageCustomerCount
	} else {
		return nil, fmt.Errorf("revenue_per_customer or (total_product_revenue + average_customer_count) required")
	}

	if f.GrossMarginPct == nil {
		return nil, fmt.Errorf("gross_margin_pct required")
	}
	if f.AverageLifespanMonths == nil {
		return nil, fmt.Errorf("average_lifespan_months required")
	}

	clv := rpc * *f.GrossMarginPct * *f.AverageLifespanMonths

	return &domain.CLVResult{
		CLV:              clv,
		RevenuePerPeriod: rpc,
		GrossMarginPct:   *f.GrossMarginPct,
		LifespanMonths:   *f.AverageLifespanMonths,
	}, nil
}

// CACPayback calculates months to recover customer acquisition cost.
// CAC = total acquisition spend / new customers acquired.
// Payback = CAC / (monthly revenue per customer * margin%).
func (c *Calculator) CACPayback(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Financials == nil {
		return nil, fmt.Errorf("financial data required")
	}
	f := input.Financials

	if f.TotalAcquisitionSpend == nil || f.NewCustomersAcquired == nil {
		return nil, fmt.Errorf("total_acquisition_spend and new_customers_acquired required")
	}
	if *f.NewCustomersAcquired <= 0 {
		return nil, fmt.Errorf("new_customers_acquired must be positive")
	}

	cac := *f.TotalAcquisitionSpend / *f.NewCustomersAcquired

	rpc := 0.0
	if f.RevenuePerCustomer != nil {
		rpc = *f.RevenuePerCustomer
	} else {
		return nil, fmt.Errorf("revenue_per_customer required")
	}

	margin := 1.0
	if f.GrossMarginPct != nil {
		margin = *f.GrossMarginPct
	}

	monthlyContrib := rpc * margin
	if monthlyContrib <= 0 {
		return nil, fmt.Errorf("monthly contribution must be positive for payback calculation")
	}

	paybackMonths := cac / monthlyContrib

	var interp string
	switch {
	case paybackMonths <= 3:
		interp = "excellent_payback"
	case paybackMonths <= 6:
		interp = "good_payback"
	case paybackMonths <= 12:
		interp = "acceptable_payback"
	case paybackMonths <= 18:
		interp = "slow_payback"
	default:
		interp = "concerning_payback"
	}

	return &domain.SingleValueResult{
		Value:          paybackMonths,
		Interpretation: interp,
	}, nil
}

// BreakEven calculates the number of units needed to break even.
// BreakEvenUnits = fixed costs / (price - variable cost per unit).
func (c *Calculator) BreakEven(input *domain.AppraisalInput) (*domain.BreakEvenResult, error) {
	if input.Financials == nil {
		return nil, fmt.Errorf("financial data required")
	}
	if input.Product == nil {
		return nil, fmt.Errorf("product definition required")
	}
	f := input.Financials

	if f.FixedCosts == nil {
		return nil, fmt.Errorf("fixed_costs required")
	}
	if f.VariableCostPerUnit == nil {
		return nil, fmt.Errorf("variable_cost_per_unit required")
	}

	contribMargin := input.Product.Price - *f.VariableCostPerUnit
	if contribMargin <= 0 {
		return nil, fmt.Errorf("contribution margin must be positive (price > variable cost)")
	}

	units := *f.FixedCosts / contribMargin

	return &domain.BreakEvenResult{
		BreakEvenUnits: units,
		FixedCosts:     *f.FixedCosts,
		ContribMargin:  contribMargin,
	}, nil
}

// CannibalizationNet calculates net revenue impact of cannibalization.
// Net = (new premium customers * new premium revenue) - (migrated customers * (old rev - new rev)).
func (c *Calculator) CannibalizationNet(input *domain.AppraisalInput) (*domain.CannibalizationResult, error) {
	if input.Financials == nil {
		return nil, fmt.Errorf("financial data required")
	}
	f := input.Financials

	if f.MigratedCustomerCount == nil || f.MigratedCustomerOldRev == nil || f.MigratedCustomerNewRev == nil {
		return nil, fmt.Errorf("migrated customer data required (count, old_revenue, new_revenue)")
	}
	if f.NewPremiumCustomers == nil || f.NewPremiumRevenue == nil {
		return nil, fmt.Errorf("new premium customer data required (count, revenue)")
	}

	migratedLoss := *f.MigratedCustomerCount * (*f.MigratedCustomerOldRev - *f.MigratedCustomerNewRev)
	newGain := *f.NewPremiumCustomers * *f.NewPremiumRevenue

	netDelta := newGain - migratedLoss

	return &domain.CannibalizationResult{
		NetRevenueDelta:       netDelta,
		MigratedRevenueLoss:   migratedLoss,
		NewPremiumRevenueGain: newGain,
		NetPositive:           netDelta > 0,
	}, nil
}

// StressTest evaluates margin under adverse conditions (costs +X%, growth -Y%).
// Defaults: costs +20%, growth -30% if not specified in input.
func (c *Calculator) StressTest(input *domain.AppraisalInput) (*domain.StressTestResult, error) {
	if input.Financials == nil {
		return nil, fmt.Errorf("financial data required")
	}
	f := input.Financials

	// Calculate base margin
	if f.RevenuePerCustomer == nil {
		return nil, fmt.Errorf("revenue_per_customer required")
	}

	baseCost := 0.0
	if f.DirectCostPerCustomer != nil {
		baseCost += *f.DirectCostPerCustomer
	}
	if f.PartnerLicensingCost != nil {
		baseCost += *f.PartnerLicensingCost
	}
	if f.SharedCostPerCustomer != nil {
		baseCost += *f.SharedCostPerCustomer
	}
	if f.CustomerServiceCost != nil {
		baseCost += *f.CustomerServiceCost
	}

	baseMargin := *f.RevenuePerCustomer - baseCost

	// Apply stress factors
	costIncrease := 0.20
	if f.CostIncreasePct != nil {
		costIncrease = *f.CostIncreasePct
	}
	growthDecrease := 0.30
	if f.GrowthDecreasePct != nil {
		growthDecrease = *f.GrowthDecreasePct
	}

	stressedCost := baseCost * (1.0 + costIncrease)
	stressedRevenue := *f.RevenuePerCustomer * (1.0 - growthDecrease)
	stressedMargin := stressedRevenue - stressedCost

	return &domain.StressTestResult{
		BaseMargin:     baseMargin,
		StressedMargin: stressedMargin,
		CostIncrease:   costIncrease,
		GrowthDecrease: growthDecrease,
		SurvivesStress: stressedMargin > 0,
	}, nil
}

// IncrementalRevenue calculates bundle revenue minus lost standalone revenue per customer.
// Incremental = bundle_revenue_per_customer - lost_standalone_revenue.
func (c *Calculator) IncrementalRevenue(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Financials == nil {
		return nil, fmt.Errorf("financial data required")
	}
	f := input.Financials

	if f.BundleRevenuePerCust == nil {
		return nil, fmt.Errorf("bundle_revenue_per_customer required")
	}
	if f.LostStandaloneRevenue == nil {
		return nil, fmt.Errorf("lost_standalone_revenue required")
	}

	incremental := *f.BundleRevenuePerCust - *f.LostStandaloneRevenue

	var interp string
	if incremental > 0 {
		interp = "positive_incremental_revenue"
	} else {
		interp = "negative_incremental_cannibalization_exceeds_uplift"
	}

	return &domain.SingleValueResult{
		Value:          incremental,
		Interpretation: interp,
	}, nil
}

// RevenueUplift calculates (premium RPC - base RPC) / base RPC.
func (c *Calculator) RevenueUplift(input *domain.AppraisalInput) (*domain.SingleValueResult, error) {
	if input.Financials == nil {
		return nil, fmt.Errorf("financial data required")
	}
	f := input.Financials

	if f.PremiumRevenue == nil || f.BaseRevenue == nil {
		return nil, fmt.Errorf("premium_revenue and base_revenue required")
	}
	if f.AverageCustomerCount == nil || *f.AverageCustomerCount <= 0 {
		return nil, fmt.Errorf("average_customer_count required and must be positive")
	}

	premiumRPC := *f.PremiumRevenue / *f.AverageCustomerCount
	baseRPC := *f.BaseRevenue / *f.AverageCustomerCount

	if baseRPC <= 0 {
		return nil, fmt.Errorf("base RPC must be positive")
	}

	uplift := (premiumRPC - baseRPC) / baseRPC

	return &domain.SingleValueResult{
		Value:          uplift,
		Interpretation: fmt.Sprintf("revenue_uplift=%.1f%%", uplift*100),
	}, nil
}
