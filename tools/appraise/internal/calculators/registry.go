// Package calculators provides the top-level registry mapping module.function names
// to calculator implementations.
package calculators

import (
	"encoding/json"
	"fmt"

	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/calculators/bundle"
	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/calculators/customer"
	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/calculators/financial"
	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/calculators/pricing"
	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/calculators/product"
	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/calculators/scoring"
	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/domain"
)

// Registry holds all calculator instances and dispatches calls by module.function.
type Registry struct {
	pricing   *pricing.Calculator
	bundle    *bundle.Calculator
	financial *financial.Calculator
	customer  *customer.Calculator
	product   *product.Calculator
	scoring   *scoring.Calculator
}

// NewRegistry creates a registry with all calculator modules initialized.
func NewRegistry() *Registry {
	return &Registry{
		pricing:   pricing.New(),
		bundle:    bundle.New(),
		financial: financial.New(),
		customer:  customer.New(),
		product:   product.New(),
		scoring:   scoring.New(),
	}
}

// Modules returns the list of available module names.
func (r *Registry) Modules() []string {
	return []string{"pricing", "bundle", "financial", "customer", "product", "scoring"}
}

// Functions returns the list of available function names for a module.
func (r *Registry) Functions(module string) ([]string, error) {
	fns, ok := moduleFunctions[module]
	if !ok {
		return nil, fmt.Errorf("unknown module %q", module)
	}
	return fns, nil
}

// moduleFunctions lists all available functions per module.
var moduleFunctions = map[string][]string{
	"pricing":   {"bvr", "tier_gap", "cost_floor", "price_value_ratio", "premium_price_index", "bundle_discount"},
	"bundle":    {"classify", "dead_weight", "cross_subsidy", "component_activation", "multi_component_usage"},
	"financial": {"unit_economics", "gross_margin", "clv", "cac_payback", "break_even", "cannibalization", "stress_test", "incremental_revenue", "revenue_uplift"},
	"customer":  {"churn_rate", "retention_rate", "nps", "csat", "churn_reduction", "revenue_growth", "service_revenue_share"},
	"product":   {"penetration_rate", "migration_rate", "cannibalization_rate", "cross_sell_rate", "feature_utilization", "component_activation_rate", "attach_rate", "trial_conversion"},
	"scoring":   {"go_no_go", "risk_matrix", "dimension_score"},
}

// Execute runs a calculation by module and function name.
// Returns the result as a JSON-serializable interface{}.
func (r *Registry) Execute(module, function string, input *domain.AppraisalInput) (interface{}, error) {
	key := module + "." + function

	switch key {
	// Pricing module
	case "pricing.bvr":
		return r.pricing.BVR(input)
	case "pricing.tier_gap":
		return r.pricing.TierGapAnalysis(input)
	case "pricing.cost_floor":
		return r.pricing.CostFloor(input)
	case "pricing.price_value_ratio":
		return r.pricing.PriceValueRatio(input)
	case "pricing.premium_price_index":
		return r.pricing.PremiumPriceIndex(input)
	case "pricing.bundle_discount":
		return r.pricing.BundleDiscount(input)

	// Bundle module
	case "bundle.classify":
		return r.bundle.ClassifyComponents(input)
	case "bundle.dead_weight":
		return r.bundle.DeadWeightRatio(input)
	case "bundle.cross_subsidy":
		return r.bundle.CrossSubsidyAnalysis(input)
	case "bundle.component_activation":
		return r.bundle.ComponentActivation(input)
	case "bundle.multi_component_usage":
		return r.bundle.MultiComponentUsage(input)

	// Financial module
	case "financial.unit_economics":
		return r.financial.UnitEconomics(input)
	case "financial.gross_margin":
		return r.financial.GrossMarginPerCustomer(input)
	case "financial.clv":
		return r.financial.CLV(input)
	case "financial.cac_payback":
		return r.financial.CACPayback(input)
	case "financial.break_even":
		return r.financial.BreakEven(input)
	case "financial.cannibalization":
		return r.financial.CannibalizationNet(input)
	case "financial.stress_test":
		return r.financial.StressTest(input)
	case "financial.incremental_revenue":
		return r.financial.IncrementalRevenue(input)
	case "financial.revenue_uplift":
		return r.financial.RevenueUplift(input)

	// Customer module
	case "customer.churn_rate":
		return r.customer.ChurnRate(input)
	case "customer.retention_rate":
		return r.customer.RetentionRate(input)
	case "customer.nps":
		return r.customer.NPS(input)
	case "customer.csat":
		return r.customer.CSAT(input)
	case "customer.churn_reduction":
		return r.customer.ChurnReductionImpact(input)
	case "customer.revenue_growth":
		return r.customer.RevenueGrowthRate(input)
	case "customer.service_revenue_share":
		return r.customer.ServiceRevenueShare(input)

	// Product module
	case "product.penetration_rate":
		return r.product.PenetrationRate(input)
	case "product.migration_rate":
		return r.product.MigrationRate(input)
	case "product.cannibalization_rate":
		return r.product.CannibalizationRate(input)
	case "product.cross_sell_rate":
		return r.product.CrossSellRate(input)
	case "product.feature_utilization":
		return r.product.FeatureUtilizationRate(input)
	case "product.component_activation_rate":
		return r.product.ComponentActivationRate(input)
	case "product.attach_rate":
		return r.product.AttachRate(input)
	case "product.trial_conversion":
		return r.product.TrialConversion(input)

	// Scoring module
	case "scoring.go_no_go":
		return r.scoring.GoNoGo(input)
	case "scoring.risk_matrix":
		return r.scoring.RiskMatrix(input)
	case "scoring.dimension_score":
		// Dimension score requires special handling — extract from scoring input
		if input.Scoring == nil || len(input.Scoring.Dimensions) == 0 {
			return nil, fmt.Errorf("scoring.dimensions required for dimension_score")
		}
		ds := input.Scoring.Dimensions[0]
		return r.scoring.DimensionScoreCalc(ds.Dimension, ds.Score, ds.Rationale)

	default:
		return nil, fmt.Errorf("unknown function %q in module %q", function, module)
	}
}

// ExecuteJSON runs Execute and marshals the result to JSON bytes.
func (r *Registry) ExecuteJSON(module, function string, input *domain.AppraisalInput) ([]byte, error) {
	result, err := r.Execute(module, function, input)
	if err != nil {
		return nil, err
	}
	return json.MarshalIndent(result, "", "  ")
}
