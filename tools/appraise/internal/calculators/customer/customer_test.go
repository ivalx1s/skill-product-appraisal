package customer

import (
	"math"
	"testing"

	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/domain"
)

func ptr(v float64) *float64 { return &v }

const epsilon = 1e-9

func almostEqual(a, b float64) bool {
	if math.IsInf(a, 0) && math.IsInf(b, 0) {
		return (a > 0) == (b > 0)
	}
	return math.Abs(a-b) < epsilon
}

// ---------------------------------------------------------------------------
// ChurnRate
// ---------------------------------------------------------------------------

func TestChurnRate(t *testing.T) {
	calc := New()

	tests := []struct {
		name      string
		input     *domain.AppraisalInput
		wantErr   bool
		wantValue float64
	}{
		{
			name: "typical_5_percent",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					LostCustomers:        ptr(50.0),
					CustomersStartPeriod: ptr(1000.0),
				},
			},
			wantValue: 0.05,
		},
		{
			name: "high_20_percent",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					LostCustomers:        ptr(200.0),
					CustomersStartPeriod: ptr(1000.0),
				},
			},
			wantValue: 0.20,
		},
		{
			name: "zero_churn",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					LostCustomers:        ptr(0.0),
					CustomersStartPeriod: ptr(1000.0),
				},
			},
			wantValue: 0.0,
		},
		{
			name: "100_percent_churn",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					LostCustomers:        ptr(500.0),
					CustomersStartPeriod: ptr(500.0),
				},
			},
			wantValue: 1.0,
		},
		{
			name: "single_customer_lost",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					LostCustomers:        ptr(1.0),
					CustomersStartPeriod: ptr(10000.0),
				},
			},
			wantValue: 0.0001,
		},
		{
			name:    "nil_customers",
			input:   &domain.AppraisalInput{},
			wantErr: true,
		},
		{
			name: "missing_lost_customers",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					CustomersStartPeriod: ptr(1000.0),
				},
			},
			wantErr: true,
		},
		{
			name: "missing_start_period",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					LostCustomers: ptr(50.0),
				},
			},
			wantErr: true,
		},
		{
			name: "zero_start_period",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					LostCustomers:        ptr(50.0),
					CustomersStartPeriod: ptr(0.0),
				},
			},
			wantErr: true,
		},
		{
			name: "negative_start_period",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					LostCustomers:        ptr(50.0),
					CustomersStartPeriod: ptr(-100.0),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.ChurnRate(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !almostEqual(result.Value, tt.wantValue) {
				t.Errorf("Value = %v, want %v", result.Value, tt.wantValue)
			}
			if result.Interpretation == "" {
				t.Error("Interpretation should not be empty")
			}
		})
	}
}

// ---------------------------------------------------------------------------
// RetentionRate
// ---------------------------------------------------------------------------

func TestRetentionRate(t *testing.T) {
	calc := New()

	tests := []struct {
		name      string
		input     *domain.AppraisalInput
		wantErr   bool
		wantValue float64
	}{
		{
			name: "95_percent_retention",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					LostCustomers:        ptr(50.0),
					CustomersStartPeriod: ptr(1000.0),
				},
			},
			wantValue: 0.95, // 1 - 0.05
		},
		{
			name: "80_percent_retention_high_churn",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					LostCustomers:        ptr(200.0),
					CustomersStartPeriod: ptr(1000.0),
				},
			},
			wantValue: 0.80,
		},
		{
			name: "100_percent_retention_zero_churn",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					LostCustomers:        ptr(0.0),
					CustomersStartPeriod: ptr(1000.0),
				},
			},
			wantValue: 1.0,
		},
		{
			name: "0_percent_retention_all_churned",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					LostCustomers:        ptr(500.0),
					CustomersStartPeriod: ptr(500.0),
				},
			},
			wantValue: 0.0,
		},
		{
			name:    "nil_customers_propagates_error",
			input:   &domain.AppraisalInput{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.RetentionRate(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !almostEqual(result.Value, tt.wantValue) {
				t.Errorf("Value = %v, want %v", result.Value, tt.wantValue)
			}
			if result.Interpretation == "" {
				t.Error("Interpretation should not be empty")
			}
		})
	}
}

// ---------------------------------------------------------------------------
// NPS
// ---------------------------------------------------------------------------

func TestNPS(t *testing.T) {
	calc := New()

	tests := []struct {
		name       string
		input      *domain.AppraisalInput
		wantErr    bool
		wantValue  float64
		wantInterp string
	}{
		{
			name: "all_promoters_plus_100",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PromotersPct:  ptr(100.0),
					DetractorsPct: ptr(0.0),
				},
			},
			wantValue:  100.0,
			wantInterp: "excellent_nps",
		},
		{
			name: "all_detractors_minus_100",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PromotersPct:  ptr(0.0),
					DetractorsPct: ptr(100.0),
				},
			},
			wantValue:  -100.0,
			wantInterp: "negative_nps",
		},
		{
			name: "typical_good_nps_plus_35",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PromotersPct:  ptr(55.0),
					DetractorsPct: ptr(20.0),
				},
			},
			wantValue:  35.0,
			wantInterp: "good_nps",
		},
		{
			name: "moderate_nps_plus_15",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PromotersPct:  ptr(40.0),
					DetractorsPct: ptr(25.0),
				},
			},
			wantValue:  15.0,
			wantInterp: "moderate_nps",
		},
		{
			name: "boundary_excellent_exactly_50",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PromotersPct:  ptr(60.0),
					DetractorsPct: ptr(10.0),
				},
			},
			wantValue:  50.0,
			wantInterp: "excellent_nps", // >= 50
		},
		{
			name: "boundary_good_exactly_30",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PromotersPct:  ptr(50.0),
					DetractorsPct: ptr(20.0),
				},
			},
			wantValue:  30.0,
			wantInterp: "good_nps", // >= 30
		},
		{
			name: "boundary_moderate_exactly_0",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PromotersPct:  ptr(30.0),
					DetractorsPct: ptr(30.0),
				},
			},
			wantValue:  0.0,
			wantInterp: "moderate_nps", // >= 0
		},
		{
			name: "just_below_zero_negative",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PromotersPct:  ptr(20.0),
					DetractorsPct: ptr(25.0),
				},
			},
			wantValue:  -5.0,
			wantInterp: "negative_nps",
		},
		{
			name:    "nil_customers",
			input:   &domain.AppraisalInput{},
			wantErr: true,
		},
		{
			name: "missing_promoters",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					DetractorsPct: ptr(20.0),
				},
			},
			wantErr: true,
		},
		{
			name: "missing_detractors",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PromotersPct: ptr(50.0),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.NPS(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !almostEqual(result.Value, tt.wantValue) {
				t.Errorf("Value = %v, want %v", result.Value, tt.wantValue)
			}
			if result.Interpretation != tt.wantInterp {
				t.Errorf("Interpretation = %q, want %q", result.Interpretation, tt.wantInterp)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// CSAT
// ---------------------------------------------------------------------------

func TestCSAT(t *testing.T) {
	calc := New()

	tests := []struct {
		name       string
		input      *domain.AppraisalInput
		wantErr    bool
		wantValue  float64
		wantInterp string
	}{
		{
			name: "perfect_100_percent",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					SatisfiedResponses: ptr(500.0),
					TotalResponses:     ptr(500.0),
				},
			},
			wantValue:  1.0,
			wantInterp: "premium_level_satisfaction",
		},
		{
			name: "low_50_percent",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					SatisfiedResponses: ptr(250.0),
					TotalResponses:     ptr(500.0),
				},
			},
			wantValue:  0.50,
			wantInterp: "low_satisfaction",
		},
		{
			name: "boundary_80_percent_premium",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					SatisfiedResponses: ptr(80.0),
					TotalResponses:     ptr(100.0),
				},
			},
			wantValue:  0.80,
			wantInterp: "premium_level_satisfaction", // >= 0.80
		},
		{
			name: "just_below_80_acceptable",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					SatisfiedResponses: ptr(79.0),
					TotalResponses:     ptr(100.0),
				},
			},
			wantValue:  0.79,
			wantInterp: "acceptable_satisfaction",
		},
		{
			name: "boundary_60_percent_acceptable",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					SatisfiedResponses: ptr(60.0),
					TotalResponses:     ptr(100.0),
				},
			},
			wantValue:  0.60,
			wantInterp: "acceptable_satisfaction", // >= 0.60
		},
		{
			name: "just_below_60_low",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					SatisfiedResponses: ptr(59.0),
					TotalResponses:     ptr(100.0),
				},
			},
			wantValue:  0.59,
			wantInterp: "low_satisfaction",
		},
		{
			name: "zero_satisfied",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					SatisfiedResponses: ptr(0.0),
					TotalResponses:     ptr(100.0),
				},
			},
			wantValue:  0.0,
			wantInterp: "low_satisfaction",
		},
		{
			name:    "nil_customers",
			input:   &domain.AppraisalInput{},
			wantErr: true,
		},
		{
			name: "missing_satisfied",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					TotalResponses: ptr(100.0),
				},
			},
			wantErr: true,
		},
		{
			name: "missing_total",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					SatisfiedResponses: ptr(80.0),
				},
			},
			wantErr: true,
		},
		{
			name: "zero_total_responses",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					SatisfiedResponses: ptr(0.0),
					TotalResponses:     ptr(0.0),
				},
			},
			wantErr: true,
		},
		{
			name: "negative_total_responses",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					SatisfiedResponses: ptr(10.0),
					TotalResponses:     ptr(-5.0),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.CSAT(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !almostEqual(result.Value, tt.wantValue) {
				t.Errorf("Value = %v, want %v", result.Value, tt.wantValue)
			}
			if result.Interpretation != tt.wantInterp {
				t.Errorf("Interpretation = %q, want %q", result.Interpretation, tt.wantInterp)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// ChurnReductionImpact
// ---------------------------------------------------------------------------

func TestChurnReductionImpact(t *testing.T) {
	calc := New()

	tests := []struct {
		name       string
		input      *domain.AppraisalInput
		wantErr    bool
		wantValue  float64
		wantInterp string
	}{
		{
			name: "strong_reduction_50_percent",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					ChurnBefore: ptr(0.10),
					ChurnAfter:  ptr(0.05),
				},
			},
			wantValue:  0.50,
			wantInterp: "strong_churn_reduction", // >= 0.50
		},
		{
			name: "moderate_reduction_33_percent",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					ChurnBefore: ptr(0.15),
					ChurnAfter:  ptr(0.10),
				},
			},
			wantValue:  0.333333333,
			wantInterp: "moderate_churn_reduction", // >= 0.25
		},
		{
			name: "modest_reduction_10_percent",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					ChurnBefore: ptr(0.20),
					ChurnAfter:  ptr(0.18),
				},
			},
			wantValue:  0.10,
			wantInterp: "modest_churn_reduction", // >= 0.05
		},
		{
			name: "minimal_reduction_3_percent",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					ChurnBefore: ptr(0.10),
					ChurnAfter:  ptr(0.097),
				},
			},
			wantValue:  0.03,
			wantInterp: "minimal_churn_reduction", // > 0 but < 0.05
		},
		{
			name: "no_reduction_same_churn",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					ChurnBefore: ptr(0.10),
					ChurnAfter:  ptr(0.10),
				},
			},
			wantValue:  0.0,
			wantInterp: "churn_increased", // <= 0
		},
		{
			name: "churn_increased_negative",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					ChurnBefore: ptr(0.10),
					ChurnAfter:  ptr(0.15),
				},
			},
			wantValue:  -0.50, // (0.10-0.15)/0.10 = -0.50
			wantInterp: "churn_increased",
		},
		{
			name: "complete_elimination_100_percent",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					ChurnBefore: ptr(0.10),
					ChurnAfter:  ptr(0.0),
				},
			},
			wantValue:  1.0,
			wantInterp: "strong_churn_reduction",
		},
		{
			name: "boundary_exactly_25_percent_moderate",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					ChurnBefore: ptr(0.20),
					ChurnAfter:  ptr(0.15),
				},
			},
			wantValue:  0.25,
			wantInterp: "moderate_churn_reduction", // >= 0.25
		},
		{
			name: "boundary_exactly_5_percent_modest",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					ChurnBefore: ptr(0.20),
					ChurnAfter:  ptr(0.19),
				},
			},
			wantValue:  0.05,
			wantInterp: "modest_churn_reduction", // >= 0.05
		},
		{
			name:    "nil_customers",
			input:   &domain.AppraisalInput{},
			wantErr: true,
		},
		{
			name: "missing_churn_before",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					ChurnAfter: ptr(0.05),
				},
			},
			wantErr: true,
		},
		{
			name: "missing_churn_after",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					ChurnBefore: ptr(0.10),
				},
			},
			wantErr: true,
		},
		{
			name: "zero_churn_before",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					ChurnBefore: ptr(0.0),
					ChurnAfter:  ptr(0.05),
				},
			},
			wantErr: true, // division by zero guard
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.ChurnReductionImpact(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !almostEqual(result.Value, tt.wantValue) {
				t.Errorf("Value = %v, want %v", result.Value, tt.wantValue)
			}
			if result.Interpretation != tt.wantInterp {
				t.Errorf("Interpretation = %q, want %q", result.Interpretation, tt.wantInterp)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// RevenueGrowthRate
// ---------------------------------------------------------------------------

func TestRevenueGrowthRate(t *testing.T) {
	calc := New()

	tests := []struct {
		name      string
		input     *domain.AppraisalInput
		wantErr   bool
		wantValue float64
	}{
		{
			name: "high_growth_50_percent",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					RevenueCurrentPeriod: ptr(150000.0),
					RevenuePriorPeriod:   ptr(100000.0),
				},
			},
			wantValue: 0.50,
		},
		{
			name: "decline_minus_10_percent",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					RevenueCurrentPeriod: ptr(90000.0),
					RevenuePriorPeriod:   ptr(100000.0),
				},
			},
			wantValue: -0.10,
		},
		{
			name: "flat_zero_growth",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					RevenueCurrentPeriod: ptr(100000.0),
					RevenuePriorPeriod:   ptr(100000.0),
				},
			},
			wantValue: 0.0,
		},
		{
			name: "double_100_percent_growth",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					RevenueCurrentPeriod: ptr(200000.0),
					RevenuePriorPeriod:   ptr(100000.0),
				},
			},
			wantValue: 1.0,
		},
		{
			name: "small_growth",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					RevenueCurrentPeriod: ptr(101000.0),
					RevenuePriorPeriod:   ptr(100000.0),
				},
			},
			wantValue: 0.01,
		},
		{
			name:    "nil_customers",
			input:   &domain.AppraisalInput{},
			wantErr: true,
		},
		{
			name: "missing_current_period",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					RevenuePriorPeriod: ptr(100000.0),
				},
			},
			wantErr: true,
		},
		{
			name: "missing_prior_period",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					RevenueCurrentPeriod: ptr(150000.0),
				},
			},
			wantErr: true,
		},
		{
			name: "zero_prior_period",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					RevenueCurrentPeriod: ptr(100000.0),
					RevenuePriorPeriod:   ptr(0.0),
				},
			},
			wantErr: true,
		},
		{
			name: "negative_prior_period",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					RevenueCurrentPeriod: ptr(100000.0),
					RevenuePriorPeriod:   ptr(-50000.0),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.RevenueGrowthRate(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !almostEqual(result.Value, tt.wantValue) {
				t.Errorf("Value = %v, want %v", result.Value, tt.wantValue)
			}
			if result.Interpretation == "" {
				t.Error("Interpretation should not be empty")
			}
		})
	}
}

// ---------------------------------------------------------------------------
// ServiceRevenueShare
// ---------------------------------------------------------------------------

func TestServiceRevenueShare(t *testing.T) {
	calc := New()

	tests := []struct {
		name      string
		input     *domain.AppraisalInput
		wantErr   bool
		wantValue float64
	}{
		{
			name: "high_share_60_percent",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					AddOnRevenue: ptr(60000.0),
					TotalRevenue: ptr(100000.0),
				},
			},
			wantValue: 0.60,
		},
		{
			name: "low_share_5_percent",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					AddOnRevenue: ptr(5000.0),
					TotalRevenue: ptr(100000.0),
				},
			},
			wantValue: 0.05,
		},
		{
			name: "zero_share",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					AddOnRevenue: ptr(0.0),
					TotalRevenue: ptr(100000.0),
				},
			},
			wantValue: 0.0,
		},
		{
			name: "all_addon_revenue",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					AddOnRevenue: ptr(100000.0),
					TotalRevenue: ptr(100000.0),
				},
			},
			wantValue: 1.0,
		},
		{
			name: "small_total",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					AddOnRevenue: ptr(50.0),
					TotalRevenue: ptr(1000.0),
				},
			},
			wantValue: 0.05,
		},
		{
			name:    "nil_customers",
			input:   &domain.AppraisalInput{},
			wantErr: true,
		},
		{
			name: "missing_addon_revenue",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					TotalRevenue: ptr(100000.0),
				},
			},
			wantErr: true,
		},
		{
			name: "missing_total_revenue",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					AddOnRevenue: ptr(50000.0),
				},
			},
			wantErr: true,
		},
		{
			name: "zero_total_revenue",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					AddOnRevenue: ptr(50000.0),
					TotalRevenue: ptr(0.0),
				},
			},
			wantErr: true,
		},
		{
			name: "negative_total_revenue",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					AddOnRevenue: ptr(50000.0),
					TotalRevenue: ptr(-100.0),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.ServiceRevenueShare(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !almostEqual(result.Value, tt.wantValue) {
				t.Errorf("Value = %v, want %v", result.Value, tt.wantValue)
			}
			if result.Interpretation == "" {
				t.Error("Interpretation should not be empty")
			}
		})
	}
}
