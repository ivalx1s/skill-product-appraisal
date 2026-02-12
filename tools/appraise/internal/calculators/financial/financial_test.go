package financial

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
// UnitEconomics
// ---------------------------------------------------------------------------

func TestUnitEconomics(t *testing.T) {
	calc := New()

	tests := []struct {
		name          string
		input         *domain.AppraisalInput
		wantErr       bool
		wantViable    bool
		wantRevenue   float64
		wantCost      float64
		wantMargin    float64
		wantMarginPct float64
	}{
		{
			name: "profitable_with_revenue_per_customer",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					RevenuePerCustomer:    ptr(100.0),
					AverageCustomerCount:  ptr(1000.0),
					DirectCostPerCustomer: ptr(20.0),
					PartnerLicensingCost:  ptr(10.0),
					SharedCostPerCustomer: ptr(5.0),
					CustomerServiceCost:   ptr(5.0),
				},
			},
			wantErr:       false,
			wantViable:    true,
			wantRevenue:   100.0,
			wantCost:      40.0,
			wantMargin:    60.0,
			wantMarginPct: 0.6,
		},
		{
			name: "profitable_with_total_revenue",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalProductRevenue:   ptr(100000.0),
					AverageCustomerCount:  ptr(1000.0),
					DirectCostPerCustomer: ptr(30.0),
				},
			},
			wantErr:       false,
			wantViable:    true,
			wantRevenue:   100.0,
			wantCost:      30.0,
			wantMargin:    70.0,
			wantMarginPct: 0.7,
		},
		{
			name: "breakeven_zero_margin",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					RevenuePerCustomer:    ptr(50.0),
					AverageCustomerCount:  ptr(100.0),
					DirectCostPerCustomer: ptr(30.0),
					SharedCostPerCustomer: ptr(20.0),
				},
			},
			wantErr:       false,
			wantViable:    false, // margin == 0, not > 0
			wantRevenue:   50.0,
			wantCost:      50.0,
			wantMargin:    0.0,
			wantMarginPct: 0.0,
		},
		{
			name: "loss_making_negative_margin",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					RevenuePerCustomer:    ptr(40.0),
					AverageCustomerCount:  ptr(500.0),
					DirectCostPerCustomer: ptr(30.0),
					PartnerLicensingCost:  ptr(15.0),
					SharedCostPerCustomer: ptr(5.0),
				},
			},
			wantErr:       false,
			wantViable:    false,
			wantRevenue:   40.0,
			wantCost:      50.0,
			wantMargin:    -10.0,
			wantMarginPct: -0.25,
		},
		{
			name: "zero_costs",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					RevenuePerCustomer:   ptr(100.0),
					AverageCustomerCount: ptr(10.0),
				},
			},
			wantErr:       false,
			wantViable:    true,
			wantRevenue:   100.0,
			wantCost:      0.0,
			wantMargin:    100.0,
			wantMarginPct: 1.0,
		},
		{
			name:    "nil_financials",
			input:   &domain.AppraisalInput{},
			wantErr: true,
		},
		{
			name: "nil_customer_count",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					RevenuePerCustomer: ptr(100.0),
				},
			},
			wantErr: true,
		},
		{
			name: "zero_customer_count",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					RevenuePerCustomer:   ptr(100.0),
					AverageCustomerCount: ptr(0.0),
				},
			},
			wantErr: true,
		},
		{
			name: "no_revenue_fields",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					AverageCustomerCount: ptr(100.0),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.UnitEconomics(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !almostEqual(result.RevenuePerCustomer, tt.wantRevenue) {
				t.Errorf("RevenuePerCustomer = %v, want %v", result.RevenuePerCustomer, tt.wantRevenue)
			}
			if !almostEqual(result.CostPerCustomer, tt.wantCost) {
				t.Errorf("CostPerCustomer = %v, want %v", result.CostPerCustomer, tt.wantCost)
			}
			if !almostEqual(result.MarginPerCustomer, tt.wantMargin) {
				t.Errorf("MarginPerCustomer = %v, want %v", result.MarginPerCustomer, tt.wantMargin)
			}
			if !almostEqual(result.MarginPct, tt.wantMarginPct) {
				t.Errorf("MarginPct = %v, want %v", result.MarginPct, tt.wantMarginPct)
			}
			if result.Viable != tt.wantViable {
				t.Errorf("Viable = %v, want %v", result.Viable, tt.wantViable)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// GrossMarginPerCustomer
// ---------------------------------------------------------------------------

func TestGrossMarginPerCustomer(t *testing.T) {
	calc := New()

	tests := []struct {
		name      string
		input     *domain.AppraisalInput
		wantErr   bool
		wantValue float64
	}{
		{
			name: "positive_margin",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalProductRevenue:  ptr(500000.0),
					COGS:                 ptr(200000.0),
					AverageCustomerCount: ptr(1000.0),
				},
			},
			wantValue: 300.0,
		},
		{
			name: "zero_margin",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalProductRevenue:  ptr(100000.0),
					COGS:                 ptr(100000.0),
					AverageCustomerCount: ptr(500.0),
				},
			},
			wantValue: 0.0,
		},
		{
			name: "high_revenue_low_cogs",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalProductRevenue:  ptr(1000000.0),
					COGS:                 ptr(50000.0),
					AverageCustomerCount: ptr(100.0),
				},
			},
			wantValue: 9500.0,
		},
		{
			name: "negative_margin",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalProductRevenue:  ptr(100000.0),
					COGS:                 ptr(150000.0),
					AverageCustomerCount: ptr(1000.0),
				},
			},
			wantValue: -50.0,
		},
		{
			name:    "nil_financials",
			input:   &domain.AppraisalInput{},
			wantErr: true,
		},
		{
			name: "missing_revenue",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					COGS:                 ptr(100.0),
					AverageCustomerCount: ptr(10.0),
				},
			},
			wantErr: true,
		},
		{
			name: "missing_cogs",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalProductRevenue:  ptr(100.0),
					AverageCustomerCount: ptr(10.0),
				},
			},
			wantErr: true,
		},
		{
			name: "zero_customer_count",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalProductRevenue:  ptr(100.0),
					COGS:                 ptr(50.0),
					AverageCustomerCount: ptr(0.0),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.GrossMarginPerCustomer(tt.input)
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
// CLV
// ---------------------------------------------------------------------------

func TestCLV(t *testing.T) {
	calc := New()

	tests := []struct {
		name     string
		input    *domain.AppraisalInput
		wantErr  bool
		wantCLV  float64
		wantRPC  float64
		wantGM   float64
		wantLife float64
	}{
		{
			name: "typical_saas_customer",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					RevenuePerCustomer:    ptr(50.0),
					GrossMarginPct:        ptr(0.70),
					AverageLifespanMonths: ptr(24.0),
				},
			},
			wantCLV:  840.0, // 50 * 0.7 * 24
			wantRPC:  50.0,
			wantGM:   0.70,
			wantLife: 24.0,
		},
		{
			name: "high_margin_long_lived",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					RevenuePerCustomer:    ptr(200.0),
					GrossMarginPct:        ptr(0.90),
					AverageLifespanMonths: ptr(48.0),
				},
			},
			wantCLV:  8640.0, // 200 * 0.9 * 48
			wantRPC:  200.0,
			wantGM:   0.90,
			wantLife: 48.0,
		},
		{
			name: "low_margin_short_lived",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					RevenuePerCustomer:    ptr(10.0),
					GrossMarginPct:        ptr(0.10),
					AverageLifespanMonths: ptr(3.0),
				},
			},
			wantCLV:  3.0, // 10 * 0.1 * 3
			wantRPC:  10.0,
			wantGM:   0.10,
			wantLife: 3.0,
		},
		{
			name: "revenue_from_total_and_count",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalProductRevenue:   ptr(100000.0),
					AverageCustomerCount:  ptr(500.0),
					GrossMarginPct:        ptr(0.60),
					AverageLifespanMonths: ptr(12.0),
				},
			},
			wantCLV:  1440.0, // (100000/500) * 0.6 * 12
			wantRPC:  200.0,
			wantGM:   0.60,
			wantLife: 12.0,
		},
		{
			name: "zero_margin",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					RevenuePerCustomer:    ptr(100.0),
					GrossMarginPct:        ptr(0.0),
					AverageLifespanMonths: ptr(12.0),
				},
			},
			wantCLV:  0.0,
			wantRPC:  100.0,
			wantGM:   0.0,
			wantLife: 12.0,
		},
		{
			name:    "nil_financials",
			input:   &domain.AppraisalInput{},
			wantErr: true,
		},
		{
			name: "missing_margin",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					RevenuePerCustomer:    ptr(100.0),
					AverageLifespanMonths: ptr(12.0),
				},
			},
			wantErr: true,
		},
		{
			name: "missing_lifespan",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					RevenuePerCustomer: ptr(100.0),
					GrossMarginPct:     ptr(0.5),
				},
			},
			wantErr: true,
		},
		{
			name: "no_revenue_source",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					GrossMarginPct:        ptr(0.5),
					AverageLifespanMonths: ptr(12.0),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.CLV(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !almostEqual(result.CLV, tt.wantCLV) {
				t.Errorf("CLV = %v, want %v", result.CLV, tt.wantCLV)
			}
			if !almostEqual(result.RevenuePerPeriod, tt.wantRPC) {
				t.Errorf("RevenuePerPeriod = %v, want %v", result.RevenuePerPeriod, tt.wantRPC)
			}
			if !almostEqual(result.GrossMarginPct, tt.wantGM) {
				t.Errorf("GrossMarginPct = %v, want %v", result.GrossMarginPct, tt.wantGM)
			}
			if !almostEqual(result.LifespanMonths, tt.wantLife) {
				t.Errorf("LifespanMonths = %v, want %v", result.LifespanMonths, tt.wantLife)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// CACPayback
// ---------------------------------------------------------------------------

func TestCACPayback(t *testing.T) {
	calc := New()

	tests := []struct {
		name       string
		input      *domain.AppraisalInput
		wantErr    bool
		wantValue  float64
		wantInterp string
	}{
		{
			name: "quick_payback_2_months",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalAcquisitionSpend: ptr(10000.0),
					NewCustomersAcquired:  ptr(100.0),
					RevenuePerCustomer:    ptr(50.0),
					GrossMarginPct:        ptr(1.0), // 100% margin
				},
			},
			wantValue:  2.0, // CAC=100, contrib=50*1.0=50, payback=100/50=2
			wantInterp: "excellent_payback",
		},
		{
			name: "good_payback_5_months",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalAcquisitionSpend: ptr(5000.0),
					NewCustomersAcquired:  ptr(50.0),
					RevenuePerCustomer:    ptr(20.0),
					GrossMarginPct:        ptr(1.0),
				},
			},
			wantValue:  5.0, // CAC=100, contrib=20, payback=5
			wantInterp: "good_payback",
		},
		{
			name: "acceptable_payback_10_months",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalAcquisitionSpend: ptr(10000.0),
					NewCustomersAcquired:  ptr(100.0),
					RevenuePerCustomer:    ptr(10.0),
					GrossMarginPct:        ptr(1.0),
				},
			},
			wantValue:  10.0, // CAC=100, contrib=10, payback=10
			wantInterp: "acceptable_payback",
		},
		{
			name: "slow_payback_18_months",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalAcquisitionSpend: ptr(18000.0),
					NewCustomersAcquired:  ptr(100.0),
					RevenuePerCustomer:    ptr(10.0),
					GrossMarginPct:        ptr(1.0),
				},
			},
			wantValue:  18.0, // CAC=180, contrib=10, payback=18
			wantInterp: "slow_payback",
		},
		{
			name: "concerning_payback_24_months",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalAcquisitionSpend: ptr(24000.0),
					NewCustomersAcquired:  ptr(100.0),
					RevenuePerCustomer:    ptr(10.0),
					GrossMarginPct:        ptr(1.0),
				},
			},
			wantValue:  24.0,
			wantInterp: "concerning_payback",
		},
		{
			name: "with_margin_factor",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalAcquisitionSpend: ptr(5000.0),
					NewCustomersAcquired:  ptr(100.0),
					RevenuePerCustomer:    ptr(50.0),
					GrossMarginPct:        ptr(0.50),
				},
			},
			wantValue:  2.0, // CAC=50, contrib=50*0.5=25, payback=50/25=2
			wantInterp: "excellent_payback",
		},
		{
			name: "no_margin_specified_defaults_to_1",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalAcquisitionSpend: ptr(3000.0),
					NewCustomersAcquired:  ptr(100.0),
					RevenuePerCustomer:    ptr(10.0),
				},
			},
			wantValue:  3.0, // CAC=30, margin=1.0, contrib=10, payback=30/10=3
			wantInterp: "excellent_payback",
		},
		{
			name: "boundary_exactly_3_months",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalAcquisitionSpend: ptr(300.0),
					NewCustomersAcquired:  ptr(1.0),
					RevenuePerCustomer:    ptr(100.0),
					GrossMarginPct:        ptr(1.0),
				},
			},
			wantValue:  3.0,
			wantInterp: "excellent_payback", // <= 3
		},
		{
			name: "zero_margin_error",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalAcquisitionSpend: ptr(1000.0),
					NewCustomersAcquired:  ptr(10.0),
					RevenuePerCustomer:    ptr(50.0),
					GrossMarginPct:        ptr(0.0),
				},
			},
			wantErr: true, // monthly contrib = 0
		},
		{
			name:    "nil_financials",
			input:   &domain.AppraisalInput{},
			wantErr: true,
		},
		{
			name: "missing_acquisition_spend",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					NewCustomersAcquired: ptr(10.0),
					RevenuePerCustomer:   ptr(50.0),
				},
			},
			wantErr: true,
		},
		{
			name: "zero_new_customers",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalAcquisitionSpend: ptr(1000.0),
					NewCustomersAcquired:  ptr(0.0),
					RevenuePerCustomer:    ptr(50.0),
				},
			},
			wantErr: true,
		},
		{
			name: "missing_revenue_per_customer",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					TotalAcquisitionSpend: ptr(1000.0),
					NewCustomersAcquired:  ptr(10.0),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.CACPayback(tt.input)
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
// BreakEven
// ---------------------------------------------------------------------------

func TestBreakEven(t *testing.T) {
	calc := New()

	tests := []struct {
		name          string
		input         *domain.AppraisalInput
		wantErr       bool
		wantUnits     float64
		wantFixed     float64
		wantContrib   float64
	}{
		{
			name: "typical_breakeven",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
				Financials: &domain.FinancialData{
					FixedCosts:          ptr(50000.0),
					VariableCostPerUnit: ptr(40.0),
				},
			},
			wantUnits:   833.333333333, // 50000/60 = 833.33...
			wantFixed:   50000.0,
			wantContrib: 60.0,
		},
		{
			name: "immediate_breakeven_zero_fixed_costs",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 50.0},
				Financials: &domain.FinancialData{
					FixedCosts:          ptr(0.0),
					VariableCostPerUnit: ptr(20.0),
				},
			},
			wantUnits:   0.0,
			wantFixed:   0.0,
			wantContrib: 30.0,
		},
		{
			name: "high_fixed_costs",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 200.0},
				Financials: &domain.FinancialData{
					FixedCosts:          ptr(1000000.0),
					VariableCostPerUnit: ptr(50.0),
				},
			},
			wantUnits:   6666.666666666, // 1000000/150
			wantFixed:   1000000.0,
			wantContrib: 150.0,
		},
		{
			name: "price_equals_variable_cost_error",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 50.0},
				Financials: &domain.FinancialData{
					FixedCosts:          ptr(10000.0),
					VariableCostPerUnit: ptr(50.0),
				},
			},
			wantErr: true, // contribMargin = 0
		},
		{
			name: "price_below_variable_cost_error",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 30.0},
				Financials: &domain.FinancialData{
					FixedCosts:          ptr(10000.0),
					VariableCostPerUnit: ptr(50.0),
				},
			},
			wantErr: true, // contribMargin < 0
		},
		{
			name:    "nil_financials",
			input:   &domain.AppraisalInput{Product: &domain.ProductDefinition{Price: 100.0}},
			wantErr: true,
		},
		{
			name: "nil_product",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					FixedCosts:          ptr(10000.0),
					VariableCostPerUnit: ptr(50.0),
				},
			},
			wantErr: true,
		},
		{
			name: "missing_fixed_costs",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
				Financials: &domain.FinancialData{
					VariableCostPerUnit: ptr(50.0),
				},
			},
			wantErr: true,
		},
		{
			name: "missing_variable_cost",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
				Financials: &domain.FinancialData{
					FixedCosts: ptr(10000.0),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.BreakEven(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !almostEqual(result.BreakEvenUnits, tt.wantUnits) {
				t.Errorf("BreakEvenUnits = %v, want %v", result.BreakEvenUnits, tt.wantUnits)
			}
			if !almostEqual(result.FixedCosts, tt.wantFixed) {
				t.Errorf("FixedCosts = %v, want %v", result.FixedCosts, tt.wantFixed)
			}
			if !almostEqual(result.ContribMargin, tt.wantContrib) {
				t.Errorf("ContribMargin = %v, want %v", result.ContribMargin, tt.wantContrib)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// CannibalizationNet
// ---------------------------------------------------------------------------

func TestCannibalizationNet(t *testing.T) {
	calc := New()

	tests := []struct {
		name        string
		input       *domain.AppraisalInput
		wantErr     bool
		wantDelta   float64
		wantLoss    float64
		wantGain    float64
		wantPositive bool
	}{
		{
			name: "positive_net_bundle_exceeds_loss",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					MigratedCustomerCount:  ptr(100.0),
					MigratedCustomerOldRev: ptr(50.0),
					MigratedCustomerNewRev: ptr(30.0),
					NewPremiumCustomers:    ptr(200.0),
					NewPremiumRevenue:      ptr(60.0),
				},
			},
			wantDelta:    10000.0, // gain=200*60=12000, loss=100*(50-30)=2000, net=10000
			wantLoss:     2000.0,
			wantGain:     12000.0,
			wantPositive: true,
		},
		{
			name: "negative_loss_exceeds_bundle",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					MigratedCustomerCount:  ptr(500.0),
					MigratedCustomerOldRev: ptr(100.0),
					MigratedCustomerNewRev: ptr(40.0),
					NewPremiumCustomers:    ptr(50.0),
					NewPremiumRevenue:      ptr(80.0),
				},
			},
			wantDelta:    -26000.0, // gain=50*80=4000, loss=500*(100-40)=30000, net=-26000
			wantLoss:     30000.0,
			wantGain:     4000.0,
			wantPositive: false,
		},
		{
			name: "exact_breakeven",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					MigratedCustomerCount:  ptr(100.0),
					MigratedCustomerOldRev: ptr(50.0),
					MigratedCustomerNewRev: ptr(30.0),
					NewPremiumCustomers:    ptr(40.0),
					NewPremiumRevenue:      ptr(50.0),
				},
			},
			wantDelta:    0.0, // gain=40*50=2000, loss=100*20=2000, net=0
			wantLoss:     2000.0,
			wantGain:     2000.0,
			wantPositive: false, // 0 is not > 0
		},
		{
			name: "no_migrated_revenue_change",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					MigratedCustomerCount:  ptr(100.0),
					MigratedCustomerOldRev: ptr(50.0),
					MigratedCustomerNewRev: ptr(50.0), // same
					NewPremiumCustomers:    ptr(10.0),
					NewPremiumRevenue:      ptr(100.0),
				},
			},
			wantDelta:    1000.0, // loss=0, gain=1000
			wantLoss:     0.0,
			wantGain:     1000.0,
			wantPositive: true,
		},
		{
			name:    "nil_financials",
			input:   &domain.AppraisalInput{},
			wantErr: true,
		},
		{
			name: "missing_migrated_data",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					NewPremiumCustomers: ptr(10.0),
					NewPremiumRevenue:   ptr(100.0),
				},
			},
			wantErr: true,
		},
		{
			name: "missing_new_premium_data",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					MigratedCustomerCount:  ptr(100.0),
					MigratedCustomerOldRev: ptr(50.0),
					MigratedCustomerNewRev: ptr(30.0),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.CannibalizationNet(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !almostEqual(result.NetRevenueDelta, tt.wantDelta) {
				t.Errorf("NetRevenueDelta = %v, want %v", result.NetRevenueDelta, tt.wantDelta)
			}
			if !almostEqual(result.MigratedRevenueLoss, tt.wantLoss) {
				t.Errorf("MigratedRevenueLoss = %v, want %v", result.MigratedRevenueLoss, tt.wantLoss)
			}
			if !almostEqual(result.NewPremiumRevenueGain, tt.wantGain) {
				t.Errorf("NewPremiumRevenueGain = %v, want %v", result.NewPremiumRevenueGain, tt.wantGain)
			}
			if result.NetPositive != tt.wantPositive {
				t.Errorf("NetPositive = %v, want %v", result.NetPositive, tt.wantPositive)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// StressTest
// ---------------------------------------------------------------------------

func TestStressTest(t *testing.T) {
	calc := New()

	tests := []struct {
		name           string
		input          *domain.AppraisalInput
		wantErr        bool
		wantBase       float64
		wantStressed   float64
		wantSurvives   bool
		wantCostInc    float64
		wantGrowthDec  float64
	}{
		{
			name: "passes_all_default_stress",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					RevenuePerCustomer:    ptr(100.0),
					DirectCostPerCustomer: ptr(20.0),
					PartnerLicensingCost:  ptr(5.0),
				},
			},
			// base: 100-25=75, stressed: 100*0.7 - 25*1.2 = 70-30 = 40
			wantBase:      75.0,
			wantStressed:  40.0,
			wantSurvives:  true,
			wantCostInc:   0.20,
			wantGrowthDec: 0.30,
		},
		{
			name: "fails_under_stress",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					RevenuePerCustomer:    ptr(50.0),
					DirectCostPerCustomer: ptr(40.0),
				},
			},
			// base: 50-40=10, stressed: 50*0.7 - 40*1.2 = 35-48 = -13
			wantBase:      10.0,
			wantStressed:  -13.0,
			wantSurvives:  false,
			wantCostInc:   0.20,
			wantGrowthDec: 0.30,
		},
		{
			name: "custom_stress_factors",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					RevenuePerCustomer:    ptr(100.0),
					DirectCostPerCustomer: ptr(30.0),
					CostIncreasePct:       ptr(0.50),  // +50%
					GrowthDecreasePct:     ptr(0.10),  // -10%
				},
			},
			// base: 100-30=70, stressed: 100*0.9 - 30*1.5 = 90-45 = 45
			wantBase:      70.0,
			wantStressed:  45.0,
			wantSurvives:  true,
			wantCostInc:   0.50,
			wantGrowthDec: 0.10,
		},
		{
			name: "zero_costs_survives",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					RevenuePerCustomer: ptr(100.0),
				},
			},
			// base: 100-0=100, stressed: 100*0.7 - 0 = 70
			wantBase:      100.0,
			wantStressed:  70.0,
			wantSurvives:  true,
			wantCostInc:   0.20,
			wantGrowthDec: 0.30,
		},
		{
			name: "barely_survives_stressed_margin_positive",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					RevenuePerCustomer:    ptr(100.0),
					DirectCostPerCustomer: ptr(55.0),
					SharedCostPerCustomer: ptr(2.0),
				},
			},
			// base: 100-57=43, stressed: 100*0.7 - 57*1.2 = 70-68.4 = 1.6
			wantBase:      43.0,
			wantStressed:  1.6,
			wantSurvives:  true,
			wantCostInc:   0.20,
			wantGrowthDec: 0.30,
		},
		{
			name:    "nil_financials",
			input:   &domain.AppraisalInput{},
			wantErr: true,
		},
		{
			name: "missing_revenue_per_customer",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					DirectCostPerCustomer: ptr(20.0),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.StressTest(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !almostEqual(result.BaseMargin, tt.wantBase) {
				t.Errorf("BaseMargin = %v, want %v", result.BaseMargin, tt.wantBase)
			}
			if !almostEqual(result.StressedMargin, tt.wantStressed) {
				t.Errorf("StressedMargin = %v, want %v", result.StressedMargin, tt.wantStressed)
			}
			if result.SurvivesStress != tt.wantSurvives {
				t.Errorf("SurvivesStress = %v, want %v", result.SurvivesStress, tt.wantSurvives)
			}
			if !almostEqual(result.CostIncrease, tt.wantCostInc) {
				t.Errorf("CostIncrease = %v, want %v", result.CostIncrease, tt.wantCostInc)
			}
			if !almostEqual(result.GrowthDecrease, tt.wantGrowthDec) {
				t.Errorf("GrowthDecrease = %v, want %v", result.GrowthDecrease, tt.wantGrowthDec)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// IncrementalRevenue
// ---------------------------------------------------------------------------

func TestIncrementalRevenue(t *testing.T) {
	calc := New()

	tests := []struct {
		name       string
		input      *domain.AppraisalInput
		wantErr    bool
		wantValue  float64
		wantInterp string
	}{
		{
			name: "positive_increment",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					BundleRevenuePerCust:  ptr(150.0),
					LostStandaloneRevenue: ptr(50.0),
				},
			},
			wantValue:  100.0,
			wantInterp: "positive_incremental_revenue",
		},
		{
			name: "negative_cannibalization_exceeds",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					BundleRevenuePerCust:  ptr(40.0),
					LostStandaloneRevenue: ptr(80.0),
				},
			},
			wantValue:  -40.0,
			wantInterp: "negative_incremental_cannibalization_exceeds_uplift",
		},
		{
			name: "zero_exactly",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					BundleRevenuePerCust:  ptr(50.0),
					LostStandaloneRevenue: ptr(50.0),
				},
			},
			wantValue:  0.0,
			wantInterp: "negative_incremental_cannibalization_exceeds_uplift", // 0 is not > 0
		},
		{
			name: "large_positive",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					BundleRevenuePerCust:  ptr(500.0),
					LostStandaloneRevenue: ptr(10.0),
				},
			},
			wantValue:  490.0,
			wantInterp: "positive_incremental_revenue",
		},
		{
			name: "zero_lost_revenue",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					BundleRevenuePerCust:  ptr(100.0),
					LostStandaloneRevenue: ptr(0.0),
				},
			},
			wantValue:  100.0,
			wantInterp: "positive_incremental_revenue",
		},
		{
			name:    "nil_financials",
			input:   &domain.AppraisalInput{},
			wantErr: true,
		},
		{
			name: "missing_bundle_revenue",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					LostStandaloneRevenue: ptr(50.0),
				},
			},
			wantErr: true,
		},
		{
			name: "missing_lost_standalone",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					BundleRevenuePerCust: ptr(100.0),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.IncrementalRevenue(tt.input)
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
// RevenueUplift
// ---------------------------------------------------------------------------

func TestRevenueUplift(t *testing.T) {
	calc := New()

	tests := []struct {
		name       string
		input      *domain.AppraisalInput
		wantErr    bool
		wantValue  float64
	}{
		{
			name: "strong_uplift_50_percent",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					PremiumRevenue:       ptr(150000.0),
					BaseRevenue:          ptr(100000.0),
					AverageCustomerCount: ptr(1000.0),
				},
			},
			wantValue: 0.5, // (150-100)/100 = 0.5
		},
		{
			name: "modest_uplift_10_percent",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					PremiumRevenue:       ptr(110000.0),
					BaseRevenue:          ptr(100000.0),
					AverageCustomerCount: ptr(1000.0),
				},
			},
			wantValue: 0.1,
		},
		{
			name: "negative_uplift_premium_lower",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					PremiumRevenue:       ptr(80000.0),
					BaseRevenue:          ptr(100000.0),
					AverageCustomerCount: ptr(1000.0),
				},
			},
			wantValue: -0.2, // (80-100)/100 = -0.2
		},
		{
			name: "zero_uplift",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					PremiumRevenue:       ptr(50000.0),
					BaseRevenue:          ptr(50000.0),
					AverageCustomerCount: ptr(100.0),
				},
			},
			wantValue: 0.0,
		},
		{
			name: "double_uplift_100_percent",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					PremiumRevenue:       ptr(200000.0),
					BaseRevenue:          ptr(100000.0),
					AverageCustomerCount: ptr(500.0),
				},
			},
			wantValue: 1.0,
		},
		{
			name:    "nil_financials",
			input:   &domain.AppraisalInput{},
			wantErr: true,
		},
		{
			name: "missing_premium_revenue",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					BaseRevenue:          ptr(100000.0),
					AverageCustomerCount: ptr(1000.0),
				},
			},
			wantErr: true,
		},
		{
			name: "missing_base_revenue",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					PremiumRevenue:       ptr(150000.0),
					AverageCustomerCount: ptr(1000.0),
				},
			},
			wantErr: true,
		},
		{
			name: "zero_customer_count",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					PremiumRevenue:       ptr(150000.0),
					BaseRevenue:          ptr(100000.0),
					AverageCustomerCount: ptr(0.0),
				},
			},
			wantErr: true,
		},
		{
			name: "zero_base_revenue_zero_rpc",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					PremiumRevenue:       ptr(150000.0),
					BaseRevenue:          ptr(0.0),
					AverageCustomerCount: ptr(1000.0),
				},
			},
			wantErr: true, // baseRPC = 0, division by zero guard
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.RevenueUplift(tt.input)
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
