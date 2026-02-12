// Package domain defines the universal input/output data model for product and bundle appraisal.
// All types are industry-agnostic. Pointers indicate optional fields.
// JSON tags are provided on every exported field for serialization.
package domain

// ---------------------------------------------------------------------------
// Top-level input envelope
// ---------------------------------------------------------------------------

// AppraisalInput is the root structure for all evaluation data.
// An agent (or human) populates the relevant sections and passes the JSON
// to any calculator module.
type AppraisalInput struct {
	Product     *ProductDefinition  `json:"product,omitempty"`
	Tiers       []TierDefinition    `json:"tiers,omitempty"`
	Competitors []CompetitorData    `json:"competitors,omitempty"`
	Customers   *CustomerMetrics    `json:"customers,omitempty"`
	Financials  *FinancialData      `json:"financials,omitempty"`
	Market      *MarketContext      `json:"market,omitempty"`
	Components  []ComponentData     `json:"components,omitempty"`
	Scoring     *ScoringInput       `json:"scoring,omitempty"`
}

// ---------------------------------------------------------------------------
// Product definition
// ---------------------------------------------------------------------------

// ProductDefinition describes the product or bundle being evaluated.
type ProductDefinition struct {
	Name        string      `json:"name"`
	Description *string     `json:"description,omitempty"`
	Price       float64     `json:"price"`
	Currency    *string     `json:"currency,omitempty"`
	Components  []Component `json:"components,omitempty"`
	Features    []Feature   `json:"features,omitempty"`
	Category    *string     `json:"category,omitempty"`
}

// Component is a single element within a bundle.
type Component struct {
	Name            string   `json:"name"`
	StandalonePrice float64  `json:"standalone_price"`
	MarginalCost    *float64 `json:"marginal_cost,omitempty"`
	PerceivedValue  *float64 `json:"perceived_value,omitempty"` // monetary: what customer thinks it's worth
	UsageForecast   *float64 `json:"usage_forecast,omitempty"` // expected % monthly active
	Activation30d   *float64 `json:"activation_30d,omitempty"` // % activated within 30 days
	Category        *string  `json:"category,omitempty"`
	IsSwappable     *bool    `json:"is_swappable,omitempty"`
}

// Feature describes a discrete product capability used in tier/competitive comparisons.
type Feature struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Available   bool    `json:"available"`
	Value       *string `json:"value,omitempty"` // freeform: "unlimited", "10GB", "true", etc.
}

// ---------------------------------------------------------------------------
// Tier definitions (Good-Better-Best)
// ---------------------------------------------------------------------------

// TierDefinition represents one tier in a multi-tier product line.
type TierDefinition struct {
	Name           string    `json:"name"`                      // e.g. "entry", "middle", "premium"
	Level          int       `json:"level"`                     // ordinal: 1=entry, 2=middle, 3=premium
	Price          float64   `json:"price"`
	Features       []Feature `json:"features,omitempty"`
	PerceivedValue *float64  `json:"perceived_value,omitempty"` // aggregate perceived value score
	CustomerShare  *float64  `json:"customer_share,omitempty"`  // actual or expected % of customers
}

// ---------------------------------------------------------------------------
// Competitor data
// ---------------------------------------------------------------------------

// CompetitorData captures a competing product for feature-by-feature comparison.
type CompetitorData struct {
	Name       string    `json:"name"`
	Provider   *string   `json:"provider,omitempty"`
	Price      float64   `json:"price"`
	Currency   *string   `json:"currency,omitempty"`
	Features   []Feature `json:"features,omitempty"`
	Components []Component `json:"components,omitempty"`
	BVR        *float64  `json:"bvr,omitempty"` // pre-calculated or to be computed
}

// ---------------------------------------------------------------------------
// Customer metrics
// ---------------------------------------------------------------------------

// CustomerMetrics groups all customer-facing KPI inputs.
type CustomerMetrics struct {
	// Base counts
	TotalCustomers        *float64 `json:"total_customers,omitempty"`
	PremiumCustomers      *float64 `json:"premium_customers,omitempty"`
	NewCustomers          *float64 `json:"new_customers,omitempty"`
	LostCustomers         *float64 `json:"lost_customers,omitempty"`
	CustomersStartPeriod  *float64 `json:"customers_start_period,omitempty"`

	// Churn and retention
	BaseChurnRate         *float64 `json:"base_churn_rate,omitempty"`
	PremiumChurnRate      *float64 `json:"premium_churn_rate,omitempty"`
	ChurnBefore           *float64 `json:"churn_before,omitempty"` // pre-launch churn
	ChurnAfter            *float64 `json:"churn_after,omitempty"`  // post-launch churn

	// NPS
	PromotersPct          *float64 `json:"promoters_pct,omitempty"`   // % scoring 9-10
	DetractorsPct         *float64 `json:"detractors_pct,omitempty"`  // % scoring 0-6

	// CSAT
	SatisfiedResponses    *float64 `json:"satisfied_responses,omitempty"`
	TotalResponses        *float64 `json:"total_responses,omitempty"`

	// Conversion and migration
	TrialUsers            *float64 `json:"trial_users,omitempty"`
	PaidConversions       *float64 `json:"paid_conversions,omitempty"`
	EligibleBase          *float64 `json:"eligible_base,omitempty"`
	UpgradedCustomers     *float64 `json:"upgraded_customers,omitempty"`
	MigratedFromStandalone *float64 `json:"migrated_from_standalone,omitempty"`
	PremiumBuyingAddons   *float64 `json:"premium_buying_addons,omitempty"`

	// Usage and engagement
	FeaturesUsedPerCustomer *float64 `json:"features_used_per_customer,omitempty"`
	TotalAvailableFeatures  *float64 `json:"total_available_features,omitempty"`
	CustomersUsing3Plus     *float64 `json:"customers_using_3plus,omitempty"`

	// Revenue context
	RevenueCurrentPeriod  *float64 `json:"revenue_current_period,omitempty"`
	RevenuePriorPeriod    *float64 `json:"revenue_prior_period,omitempty"`
	AddOnRevenue          *float64 `json:"add_on_revenue,omitempty"`
	TotalRevenue          *float64 `json:"total_revenue,omitempty"`
}

// ---------------------------------------------------------------------------
// Financial data
// ---------------------------------------------------------------------------

// FinancialData groups all financial inputs for unit economics and viability.
type FinancialData struct {
	// Revenue
	TotalProductRevenue    *float64 `json:"total_product_revenue,omitempty"`
	PremiumRevenue         *float64 `json:"premium_revenue,omitempty"`
	BaseRevenue            *float64 `json:"base_revenue,omitempty"`
	BundleRevenuePerCust   *float64 `json:"bundle_revenue_per_customer,omitempty"`
	LostStandaloneRevenue  *float64 `json:"lost_standalone_revenue,omitempty"`
	RevenuePreLaunch       *float64 `json:"revenue_pre_launch,omitempty"`
	RevenuePostLaunch      *float64 `json:"revenue_post_launch,omitempty"`

	// Costs
	COGS                   *float64 `json:"cogs,omitempty"`
	DirectCostPerCustomer  *float64 `json:"direct_cost_per_customer,omitempty"`
	PartnerLicensingCost   *float64 `json:"partner_licensing_cost,omitempty"`
	SharedCostPerCustomer  *float64 `json:"shared_cost_per_customer,omitempty"`
	CustomerServiceCost    *float64 `json:"customer_service_cost,omitempty"`
	TotalAcquisitionSpend  *float64 `json:"total_acquisition_spend,omitempty"`
	FixedCosts             *float64 `json:"fixed_costs,omitempty"`
	VariableCostPerUnit    *float64 `json:"variable_cost_per_unit,omitempty"`

	// Margin targets
	TargetMinMargin        *float64 `json:"target_min_margin,omitempty"` // as decimal (0.10 = 10%)
	GrossMarginPct         *float64 `json:"gross_margin_pct,omitempty"`

	// Customer economics
	AverageCustomerCount   *float64 `json:"average_customer_count,omitempty"`
	RevenuePerCustomer     *float64 `json:"revenue_per_customer,omitempty"`
	AverageLifespanMonths  *float64 `json:"average_lifespan_months,omitempty"`
	NewCustomersAcquired   *float64 `json:"new_customers_acquired,omitempty"`

	// Cannibalization
	MigratedCustomerCount  *float64 `json:"migrated_customer_count,omitempty"`
	MigratedCustomerOldRev *float64 `json:"migrated_customer_old_revenue,omitempty"` // per customer
	MigratedCustomerNewRev *float64 `json:"migrated_customer_new_revenue,omitempty"` // per customer
	NewPremiumCustomers    *float64 `json:"new_premium_customers,omitempty"`
	NewPremiumRevenue      *float64 `json:"new_premium_revenue,omitempty"` // per customer

	// Stress test parameters
	CostIncreasePct        *float64 `json:"cost_increase_pct,omitempty"`  // e.g. 0.20 for +20%
	GrowthDecreasePct      *float64 `json:"growth_decrease_pct,omitempty"` // e.g. 0.30 for -30%
}

// ---------------------------------------------------------------------------
// Market context
// ---------------------------------------------------------------------------

// MarketContext provides market-level data for positioning analysis.
type MarketContext struct {
	MarketAveragePrice         *float64 `json:"market_average_price,omitempty"`
	MarketGrowthRate           *float64 `json:"market_growth_rate,omitempty"`     // CAGR as decimal
	TAM                        *float64 `json:"tam,omitempty"`                    // total addressable market size
	SAM                        *float64 `json:"sam,omitempty"`                    // serviceable addressable market
	CompetitorPremiumPenetration *float64 `json:"competitor_premium_penetration,omitempty"`
	CategoryBudgetShare        *float64 `json:"category_budget_share,omitempty"`  // avg % of income
	IndustryValueTrend         *float64 `json:"industry_value_trend,omitempty"`   // YoY change
}

// ---------------------------------------------------------------------------
// Component-level data (for bundle analysis)
// ---------------------------------------------------------------------------

// ComponentData extends component information with operational metrics.
// Used for Leaders/Fillers/Killers classification and dead weight analysis.
type ComponentData struct {
	Name              string   `json:"name"`
	PerceivedValue    *float64 `json:"perceived_value,omitempty"`    // monetary: what customer thinks it's worth
	MarginalCost      *float64 `json:"marginal_cost,omitempty"`
	UsageForecast     *float64 `json:"usage_forecast,omitempty"`     // expected % monthly active
	Activation30d     *float64 `json:"activation_30d,omitempty"`     // % activated within 30 days
	MonthlyActiveRate *float64 `json:"monthly_active_rate,omitempty"` // actual % monthly active
	StandalonePrice   *float64 `json:"standalone_price,omitempty"`
	StandaloneWTP     *float64 `json:"standalone_wtp,omitempty"`
	DrivesPurchase    *bool    `json:"drives_purchase,omitempty"`     // does this drive purchase intent?
	RemovalWTPDelta   *float64 `json:"removal_wtp_delta,omitempty"`   // +/- change in WTP if removed
	RevenueContrib    *float64 `json:"revenue_contribution,omitempty"`
	DirectCost        *float64 `json:"direct_cost,omitempty"`
	Category          *string  `json:"category,omitempty"`
}

// ---------------------------------------------------------------------------
// Scoring / Go-No-Go input
// ---------------------------------------------------------------------------

// ScoringInput provides dimension scores and optional weight overrides
// for the Go/No-Go weighted scoring.
type ScoringInput struct {
	Dimensions []DimensionScore `json:"dimensions,omitempty"`
	Weights    *ScoringWeights  `json:"weights,omitempty"` // nil = use defaults
	Risks      []RiskItem       `json:"risks,omitempty"`
}

// DimensionScore holds a single dimension's score (1-5).
type DimensionScore struct {
	Dimension   string  `json:"dimension"`   // PMF, FIN, PRC_CX, CMP, BND, MR, RISK
	Score       float64 `json:"score"`        // 1.0 - 5.0
	Rationale   *string `json:"rationale,omitempty"`
}

// ScoringWeights allows overriding default Go/No-Go weights.
// Default weights: PMF=15%, FIN=25%, PRC_CX=20%, CMP=15%, BND=10%, MR=10%, RISK=5%.
type ScoringWeights struct {
	PMF   *float64 `json:"pmf,omitempty"`    // Product-Market Fit
	FIN   *float64 `json:"fin,omitempty"`    // Financial Viability
	PRCCX *float64 `json:"prc_cx,omitempty"` // Pricing + Customer Experience
	CMP   *float64 `json:"cmp,omitempty"`    // Competitive Position
	BND   *float64 `json:"bnd,omitempty"`    // Bundle Composition
	MR    *float64 `json:"mr,omitempty"`     // Market Reach
	RISK  *float64 `json:"risk,omitempty"`   // Risk Profile
}

// RiskItem describes a single risk for the risk matrix.
type RiskItem struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Probability float64 `json:"probability"` // 1-5
	Impact      float64 `json:"impact"`      // 1-5
	Mitigation  *string `json:"mitigation,omitempty"`
}

// ---------------------------------------------------------------------------
// Calculator result types
// ---------------------------------------------------------------------------

// CalcResult is the standard output envelope for any calculation.
type CalcResult struct {
	Module   string                 `json:"module"`
	Function string                 `json:"function"`
	Value    interface{}            `json:"value"`
	Details  map[string]interface{} `json:"details,omitempty"`
	Error    *string                `json:"error,omitempty"`
}

// BVRResult holds Bundle Value Ratio output.
type BVRResult struct {
	BVR               float64            `json:"bvr"`
	StandaloneSum     float64            `json:"standalone_sum"`
	BundlePrice       float64            `json:"bundle_price"`
	Interpretation    string             `json:"interpretation"`
	ComponentValues   map[string]float64 `json:"component_values,omitempty"`
}

// TierGapResult holds tier gap analysis output.
type TierGapResult struct {
	Gaps []TierGap `json:"gaps"`
}

// TierGap describes the gap between two adjacent tiers.
type TierGap struct {
	FromTier         string  `json:"from_tier"`
	ToTier           string  `json:"to_tier"`
	PriceGapAbs      float64 `json:"price_gap_abs"`
	PriceGapPct      float64 `json:"price_gap_pct"`
	ValueGap         *float64 `json:"value_gap,omitempty"`
	ValueToPriceRatio *float64 `json:"value_to_price_ratio,omitempty"`
	Diagnosis        string  `json:"diagnosis"`
}

// CostFloorResult holds cost floor calculation output.
type CostFloorResult struct {
	CostFloor      float64 `json:"cost_floor"`
	CurrentPrice   float64 `json:"current_price"`
	Margin         float64 `json:"margin"`
	ClearsFloor    bool    `json:"clears_floor"`
}

// LFKResult holds Leaders/Fillers/Killers classification output.
type LFKResult struct {
	Classifications []LFKClassification `json:"classifications"`
	Leaders         int                 `json:"leaders_count"`
	Fillers         int                 `json:"fillers_count"`
	Killers         int                 `json:"killers_count"`
}

// LFKClassification is the classification of a single component.
type LFKClassification struct {
	Name           string  `json:"name"`
	Classification string  `json:"classification"` // "leader", "filler", "killer"
	Rationale      string  `json:"rationale"`
	PerceivedValue *float64 `json:"perceived_value,omitempty"`
	MarginalCost   *float64 `json:"marginal_cost,omitempty"`
}

// DeadWeightResult holds dead weight analysis output.
type DeadWeightResult struct {
	DeadWeightRatio float64             `json:"dead_weight_ratio"`
	Threshold       float64             `json:"threshold"`     // 0.40
	Passes          bool                `json:"passes"`
	DeadWeight      []string            `json:"dead_weight"`   // component names below 20%
	ComponentUsage  map[string]float64  `json:"component_usage"`
}

// CrossSubsidyResult holds cross-subsidy analysis output.
type CrossSubsidyResult struct {
	NetMargin        float64                     `json:"net_margin"`
	Sustainable      bool                        `json:"sustainable"`
	Sources          []CrossSubsidyComponent     `json:"sources"`
	Recipients       []CrossSubsidyComponent     `json:"recipients"`
}

// CrossSubsidyComponent is margin detail per component.
type CrossSubsidyComponent struct {
	Name       string  `json:"name"`
	Revenue    float64 `json:"revenue"`
	Cost       float64 `json:"cost"`
	NetMargin  float64 `json:"net_margin"`
	Role       string  `json:"role"` // "source" or "recipient"
}

// UnitEconomicsResult holds per-customer unit economics.
type UnitEconomicsResult struct {
	RevenuePerCustomer     float64 `json:"revenue_per_customer"`
	CostPerCustomer        float64 `json:"cost_per_customer"`
	MarginPerCustomer      float64 `json:"margin_per_customer"`
	MarginPct              float64 `json:"margin_pct"`
	Viable                 bool    `json:"viable"`
}

// CLVResult holds Customer Lifetime Value output.
type CLVResult struct {
	CLV              float64 `json:"clv"`
	RevenuePerPeriod float64 `json:"revenue_per_period"`
	GrossMarginPct   float64 `json:"gross_margin_pct"`
	LifespanMonths   float64 `json:"lifespan_months"`
}

// BreakEvenResult holds break-even analysis output.
type BreakEvenResult struct {
	BreakEvenUnits  float64 `json:"break_even_units"`
	FixedCosts      float64 `json:"fixed_costs"`
	ContribMargin   float64 `json:"contribution_margin"`
}

// StressTestResult holds stress test output.
type StressTestResult struct {
	BaseMargin         float64 `json:"base_margin"`
	StressedMargin     float64 `json:"stressed_margin"`
	CostIncrease       float64 `json:"cost_increase_pct"`
	GrowthDecrease     float64 `json:"growth_decrease_pct"`
	SurvivesStress     bool    `json:"survives_stress"`
}

// CannibalizationResult holds net cannibalization analysis.
type CannibalizationResult struct {
	NetRevenueDelta        float64 `json:"net_revenue_delta"`
	MigratedRevenueLoss    float64 `json:"migrated_revenue_loss"`
	NewPremiumRevenueGain  float64 `json:"new_premium_revenue_gain"`
	NetPositive            bool    `json:"net_positive"`
}

// GoNoGoResult holds the final weighted scoring output.
type GoNoGoResult struct {
	WeightedScore float64             `json:"weighted_score"`
	Decision      string              `json:"decision"` // "strong_go", "conditional_go", "redesign", "no_go"
	Dimensions    []DimensionDetail   `json:"dimensions"`
	Weights       map[string]float64  `json:"weights_used"`
	Warning       *string             `json:"warning,omitempty"` // e.g. missing dimensions
}

// DimensionDetail is a scored dimension with weight applied.
type DimensionDetail struct {
	Dimension      string  `json:"dimension"`
	Score          float64 `json:"score"`
	Weight         float64 `json:"weight"`
	WeightedScore  float64 `json:"weighted_score"`
	Rationale      *string `json:"rationale,omitempty"`
}

// RiskMatrixResult holds risk matrix output.
type RiskMatrixResult struct {
	Risks       []ScoredRisk `json:"risks"`
	AvgScore    float64      `json:"avg_score"`
	MaxScore    float64      `json:"max_score"`
	HighRisks   int          `json:"high_risks"` // score >= 15
}

// ScoredRisk is a risk with computed score.
type ScoredRisk struct {
	Name        string  `json:"name"`
	Probability float64 `json:"probability"`
	Impact      float64 `json:"impact"`
	Score       float64 `json:"score"` // probability * impact
	Level       string  `json:"level"` // "low", "medium", "high", "critical"
}

// SingleValueResult is a generic result for simple ratio/rate calculations.
type SingleValueResult struct {
	Value          float64 `json:"value"`
	Interpretation string  `json:"interpretation,omitempty"`
}
