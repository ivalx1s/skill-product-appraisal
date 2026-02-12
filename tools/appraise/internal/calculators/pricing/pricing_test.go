package pricing

import (
	"math"
	"testing"

	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/domain"
)

// ptr returns a pointer to the given float64 value.
func ptr(v float64) *float64 {
	return &v
}

// boolPtr returns a pointer to the given bool value.
func boolPtr(v bool) *bool {
	return &v
}

// strPtr returns a pointer to the given string value.
func strPtr(v string) *string {
	return &v
}

// almostEqual checks float equality within a small epsilon.
func almostEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

const epsilon = 0.0001

// ---------------------------------------------------------------------------
// BVR tests
// ---------------------------------------------------------------------------

func TestBVR(t *testing.T) {
	calc := New()

	tests := []struct {
		name           string
		input          *domain.AppraisalInput
		wantBVR        float64
		wantInterp     string
		wantErr        bool
		errContains    string
	}{
		{
			name: "very_strong_2x_bundle_price",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 100.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 80.0},
						{Name: "B", StandalonePrice: 70.0},
						{Name: "C", StandalonePrice: 50.0},
					},
				},
			},
			wantBVR:    2.0,
			wantInterp: "very_strong",
		},
		{
			name: "strong_1.7x",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 100.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 100.0},
						{Name: "B", StandalonePrice: 70.0},
					},
				},
			},
			wantBVR:    1.7,
			wantInterp: "strong",
		},
		{
			name: "adequate_1.4x",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 100.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 80.0},
						{Name: "B", StandalonePrice: 60.0},
					},
				},
			},
			wantBVR:    1.4,
			wantInterp: "adequate",
		},
		{
			name: "marginal_1.1x",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 100.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 60.0},
						{Name: "B", StandalonePrice: 50.0},
					},
				},
			},
			wantBVR:    1.1,
			wantInterp: "marginal",
		},
		{
			name: "exactly_1.0_break_even",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 100.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 60.0},
						{Name: "B", StandalonePrice: 40.0},
					},
				},
			},
			wantBVR:    1.0,
			wantInterp: "marginal",
		},
		{
			name: "negative_below_1.0",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 100.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 30.0},
						{Name: "B", StandalonePrice: 40.0},
					},
				},
			},
			wantBVR:    0.7,
			wantInterp: "negative_value_proposition",
		},
		{
			name: "single_component",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 50.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 80.0},
					},
				},
			},
			wantBVR:    1.6,
			wantInterp: "strong",
		},
		{
			name: "many_components",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 100.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 30.0},
						{Name: "B", StandalonePrice: 25.0},
						{Name: "C", StandalonePrice: 20.0},
						{Name: "D", StandalonePrice: 15.0},
						{Name: "E", StandalonePrice: 10.0},
						{Name: "F", StandalonePrice: 50.0},
						{Name: "G", StandalonePrice: 40.0},
						{Name: "H", StandalonePrice: 35.0},
					},
				},
			},
			wantBVR:    2.25,
			wantInterp: "very_strong",
		},
		{
			name: "boundary_1.3_exact",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 100.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 130.0},
					},
				},
			},
			wantBVR:    1.3,
			wantInterp: "adequate",
		},
		{
			name: "boundary_1.5_exact",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 100.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 150.0},
					},
				},
			},
			wantBVR:    1.5,
			wantInterp: "strong",
		},
		// Error cases
		{
			name:        "nil_product",
			input:       &domain.AppraisalInput{},
			wantErr:     true,
			errContains: "product definition required",
		},
		{
			name: "no_components",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 100.0,
				},
			},
			wantErr:     true,
			errContains: "product must have components",
		},
		{
			name: "zero_price",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 50.0},
					},
				},
			},
			wantErr:     true,
			errContains: "product price must be positive",
		},
		{
			name: "negative_price",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: -10.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 50.0},
					},
				},
			},
			wantErr:     true,
			errContains: "product price must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.BVR(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.errContains)
				}
				if tt.errContains != "" && !containsStr(err.Error(), tt.errContains) {
					t.Fatalf("expected error containing %q, got %q", tt.errContains, err.Error())
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !almostEqual(result.BVR, tt.wantBVR, epsilon) {
				t.Errorf("BVR = %v, want %v", result.BVR, tt.wantBVR)
			}
			if result.Interpretation != tt.wantInterp {
				t.Errorf("Interpretation = %q, want %q", result.Interpretation, tt.wantInterp)
			}
			if result.BundlePrice != tt.input.Product.Price {
				t.Errorf("BundlePrice = %v, want %v", result.BundlePrice, tt.input.Product.Price)
			}
			// Verify component values map
			for _, comp := range tt.input.Product.Components {
				val, ok := result.ComponentValues[comp.Name]
				if !ok {
					t.Errorf("missing component %q in ComponentValues", comp.Name)
				} else if !almostEqual(val, comp.StandalonePrice, epsilon) {
					t.Errorf("ComponentValues[%q] = %v, want %v", comp.Name, val, comp.StandalonePrice)
				}
			}
		})
	}
}

// ---------------------------------------------------------------------------
// TierGapAnalysis tests
// ---------------------------------------------------------------------------

func TestTierGapAnalysis(t *testing.T) {
	calc := New()

	tests := []struct {
		name        string
		input       *domain.AppraisalInput
		wantGaps    int
		checkGaps   []struct {
			fromTier  string
			toTier    string
			diagnosis string
			pricePct  float64
		}
		wantErr     bool
		errContains string
	}{
		{
			name: "two_tiers_with_effective_upsell",
			input: &domain.AppraisalInput{
				Tiers: []domain.TierDefinition{
					// V/P ratio = ValueGap / PriceGapAbs = 150/10 = 15.0 > 1.0
					{Name: "basic", Level: 1, Price: 100.0, PerceivedValue: ptr(50.0)},
					{Name: "premium", Level: 2, Price: 110.0, PerceivedValue: ptr(200.0)},
				},
			},
			wantGaps: 1,
			checkGaps: []struct {
				fromTier  string
				toTier    string
				diagnosis string
				pricePct  float64
			}{
				{fromTier: "basic", toTier: "premium", diagnosis: "effective_upsell", pricePct: 10.0},
			},
		},
		{
			name: "three_tiers_mixed_diagnosis",
			input: &domain.AppraisalInput{
				Tiers: []domain.TierDefinition{
					{Name: "entry", Level: 1, Price: 100.0, PerceivedValue: ptr(2.0)},
					{Name: "middle", Level: 2, Price: 200.0, PerceivedValue: ptr(2.5)},
					// V/P ratio entry->middle = 0.5/100 = 0.005 < 1.0 => broken_step
					// V/P ratio middle->premium = 200/100 = 2.0 > 1.0 => effective_upsell
					{Name: "premium", Level: 3, Price: 300.0, PerceivedValue: ptr(202.5)},
				},
			},
			wantGaps: 2,
			checkGaps: []struct {
				fromTier  string
				toTier    string
				diagnosis string
				pricePct  float64
			}{
				{fromTier: "entry", toTier: "middle", diagnosis: "broken_step", pricePct: 100.0},
				{fromTier: "middle", toTier: "premium", diagnosis: "effective_upsell", pricePct: 50.0},
			},
		},
		{
			name: "without_perceived_values_gap_too_small",
			input: &domain.AppraisalInput{
				Tiers: []domain.TierDefinition{
					{Name: "basic", Level: 1, Price: 100.0},
					{Name: "plus", Level: 2, Price: 105.0},
				},
			},
			wantGaps: 1,
			checkGaps: []struct {
				fromTier  string
				toTier    string
				diagnosis string
				pricePct  float64
			}{
				{fromTier: "basic", toTier: "plus", diagnosis: "gap_too_small", pricePct: 5.0},
			},
		},
		{
			name: "without_perceived_values_gap_too_large",
			input: &domain.AppraisalInput{
				Tiers: []domain.TierDefinition{
					{Name: "basic", Level: 1, Price: 100.0},
					{Name: "ultra", Level: 2, Price: 200.0},
				},
			},
			wantGaps: 1,
			checkGaps: []struct {
				fromTier  string
				toTier    string
				diagnosis string
				pricePct  float64
			}{
				{fromTier: "basic", toTier: "ultra", diagnosis: "gap_too_large", pricePct: 100.0},
			},
		},
		{
			name: "without_perceived_values_moderate_gap",
			input: &domain.AppraisalInput{
				Tiers: []domain.TierDefinition{
					{Name: "basic", Level: 1, Price: 100.0},
					{Name: "plus", Level: 2, Price: 150.0},
				},
			},
			wantGaps: 1,
			checkGaps: []struct {
				fromTier  string
				toTier    string
				diagnosis string
				pricePct  float64
			}{
				{fromTier: "basic", toTier: "plus", diagnosis: "insufficient_data", pricePct: 50.0},
			},
		},
		{
			name: "value_to_price_ratio_exactly_1_neutral",
			input: &domain.AppraisalInput{
				Tiers: []domain.TierDefinition{
					{Name: "basic", Level: 1, Price: 100.0, PerceivedValue: ptr(2.0)},
					{Name: "premium", Level: 2, Price: 200.0, PerceivedValue: ptr(102.0)},
				},
			},
			wantGaps: 1,
			checkGaps: []struct {
				fromTier  string
				toTier    string
				diagnosis string
				pricePct  float64
			}{
				{fromTier: "basic", toTier: "premium", diagnosis: "neutral", pricePct: 100.0},
			},
		},
		// Error cases
		{
			name: "fewer_than_2_tiers",
			input: &domain.AppraisalInput{
				Tiers: []domain.TierDefinition{
					{Name: "only", Level: 1, Price: 100.0},
				},
			},
			wantErr:     true,
			errContains: "at least 2 tiers required",
		},
		{
			name: "no_tiers",
			input: &domain.AppraisalInput{
				Tiers: nil,
			},
			wantErr:     true,
			errContains: "at least 2 tiers required",
		},
		{
			name: "tiers_not_ordered",
			input: &domain.AppraisalInput{
				Tiers: []domain.TierDefinition{
					{Name: "premium", Level: 3, Price: 300.0},
					{Name: "basic", Level: 1, Price: 100.0},
				},
			},
			wantErr:     true,
			errContains: "tiers must be ordered by level",
		},
		{
			name: "tier_with_zero_price",
			input: &domain.AppraisalInput{
				Tiers: []domain.TierDefinition{
					{Name: "free", Level: 1, Price: 0.0},
					{Name: "paid", Level: 2, Price: 100.0},
				},
			},
			wantErr:     true,
			errContains: "non-positive price",
		},
		{
			name: "tier_with_negative_price",
			input: &domain.AppraisalInput{
				Tiers: []domain.TierDefinition{
					{Name: "bad", Level: 1, Price: -10.0},
					{Name: "good", Level: 2, Price: 100.0},
				},
			},
			wantErr:     true,
			errContains: "non-positive price",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.TierGapAnalysis(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.errContains)
				}
				if tt.errContains != "" && !containsStr(err.Error(), tt.errContains) {
					t.Fatalf("expected error containing %q, got %q", tt.errContains, err.Error())
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(result.Gaps) != tt.wantGaps {
				t.Fatalf("got %d gaps, want %d", len(result.Gaps), tt.wantGaps)
			}
			for i, check := range tt.checkGaps {
				gap := result.Gaps[i]
				if gap.FromTier != check.fromTier {
					t.Errorf("gap[%d].FromTier = %q, want %q", i, gap.FromTier, check.fromTier)
				}
				if gap.ToTier != check.toTier {
					t.Errorf("gap[%d].ToTier = %q, want %q", i, gap.ToTier, check.toTier)
				}
				if gap.Diagnosis != check.diagnosis {
					t.Errorf("gap[%d].Diagnosis = %q, want %q", i, gap.Diagnosis, check.diagnosis)
				}
				if !almostEqual(gap.PriceGapPct, check.pricePct, epsilon) {
					t.Errorf("gap[%d].PriceGapPct = %v, want %v", i, gap.PriceGapPct, check.pricePct)
				}
			}
		})
	}
}

// ---------------------------------------------------------------------------
// CostFloor tests
// ---------------------------------------------------------------------------

func TestCostFloor(t *testing.T) {
	calc := New()

	tests := []struct {
		name         string
		input        *domain.AppraisalInput
		wantFloor    float64
		wantMargin   float64
		wantClears   bool
		wantErr      bool
		errContains  string
	}{
		{
			name: "all_costs_present_price_above_floor",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
				Financials: &domain.FinancialData{
					DirectCostPerCustomer:  ptr(20.0),
					PartnerLicensingCost:   ptr(10.0),
					SharedCostPerCustomer:  ptr(5.0),
					TotalAcquisitionSpend:  ptr(1000.0),
					NewCustomersAcquired:   ptr(100.0),
					CustomerServiceCost:    ptr(3.0),
					TargetMinMargin:        ptr(0.10),
				},
			},
			// floor = 20 + 10 + 5 + (1000/100=10) + 3 = 48, then / (1-0.10) = 48/0.9 = 53.3333
			wantFloor:  53.3333,
			wantMargin: 46.6667,
			wantClears: true,
		},
		{
			name: "price_below_floor",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 30.0},
				Financials: &domain.FinancialData{
					DirectCostPerCustomer:  ptr(20.0),
					PartnerLicensingCost:   ptr(10.0),
					SharedCostPerCustomer:  ptr(5.0),
					CustomerServiceCost:    ptr(3.0),
					TargetMinMargin:        ptr(0.10),
				},
			},
			// floor = 20 + 10 + 5 + 3 = 38, then / 0.9 = 42.2222
			wantFloor:  42.2222,
			wantMargin: -12.2222,
			wantClears: false,
		},
		{
			name: "some_costs_nil",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 50.0},
				Financials: &domain.FinancialData{
					DirectCostPerCustomer: ptr(15.0),
					CustomerServiceCost:   ptr(5.0),
				},
			},
			// floor = 15 + 5 = 20, no margin adjustment
			wantFloor:  20.0,
			wantMargin: 30.0,
			wantClears: true,
		},
		{
			name: "zero_new_customers_skips_cac",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 50.0},
				Financials: &domain.FinancialData{
					DirectCostPerCustomer: ptr(10.0),
					TotalAcquisitionSpend: ptr(5000.0),
					NewCustomersAcquired:  ptr(0.0), // zero => skip CAC
				},
			},
			wantFloor:  10.0,
			wantMargin: 40.0,
			wantClears: true,
		},
		{
			name: "margin_at_boundary",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
				Financials: &domain.FinancialData{
					DirectCostPerCustomer: ptr(90.0),
					TargetMinMargin:       ptr(0.10),
				},
			},
			// floor = 90 / 0.9 = 100, price == floor
			wantFloor:  100.0,
			wantMargin: 0.0,
			wantClears: true,
		},
		// Error cases
		{
			name: "nil_financials",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
			},
			wantErr:     true,
			errContains: "financial data required",
		},
		{
			name: "nil_product",
			input: &domain.AppraisalInput{
				Financials: &domain.FinancialData{
					DirectCostPerCustomer: ptr(10.0),
				},
			},
			wantErr:     true,
			errContains: "product definition required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.CostFloor(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.errContains)
				}
				if tt.errContains != "" && !containsStr(err.Error(), tt.errContains) {
					t.Fatalf("expected error containing %q, got %q", tt.errContains, err.Error())
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !almostEqual(result.CostFloor, tt.wantFloor, epsilon) {
				t.Errorf("CostFloor = %v, want %v", result.CostFloor, tt.wantFloor)
			}
			if !almostEqual(result.Margin, tt.wantMargin, epsilon) {
				t.Errorf("Margin = %v, want %v", result.Margin, tt.wantMargin)
			}
			if result.ClearsFloor != tt.wantClears {
				t.Errorf("ClearsFloor = %v, want %v", result.ClearsFloor, tt.wantClears)
			}
			if result.CurrentPrice != tt.input.Product.Price {
				t.Errorf("CurrentPrice = %v, want %v", result.CurrentPrice, tt.input.Product.Price)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// PriceValueRatio tests
// ---------------------------------------------------------------------------

func TestPriceValueRatio(t *testing.T) {
	calc := New()

	tests := []struct {
		name        string
		input       *domain.AppraisalInput
		wantValue   float64
		wantInterp  string
		wantErr     bool
		errContains string
	}{
		{
			name: "value_greater_than_price_from_tier",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
				Tiers: []domain.TierDefinition{
					{Name: "premium", Level: 1, Price: 100.0, PerceivedValue: ptr(150.0)},
				},
			},
			wantValue:  1.5,
			wantInterp: "positive_value_perception",
		},
		{
			name: "value_less_than_price_from_tier",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
				Tiers: []domain.TierDefinition{
					{Name: "basic", Level: 1, Price: 100.0, PerceivedValue: ptr(50.0)},
				},
			},
			wantValue:  0.5,
			wantInterp: "negative_value_perception",
		},
		{
			name: "value_equals_price_neutral",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
				Tiers: []domain.TierDefinition{
					{Name: "standard", Level: 1, Price: 100.0, PerceivedValue: ptr(100.0)},
				},
			},
			wantValue:  1.0,
			wantInterp: "neutral",
		},
		{
			name: "value_from_component_average",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
				Components: []domain.ComponentData{
					{Name: "A", PerceivedValue: ptr(200.0)},
					{Name: "B", PerceivedValue: ptr(100.0)},
				},
			},
			// average = (200+100)/2 = 150, ratio = 150/100 = 1.5
			wantValue:  1.5,
			wantInterp: "positive_value_perception",
		},
		{
			name: "tier_takes_priority_over_components",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
				Tiers: []domain.TierDefinition{
					{Name: "premium", Level: 1, Price: 100.0, PerceivedValue: ptr(80.0)},
				},
				Components: []domain.ComponentData{
					{Name: "A", PerceivedValue: ptr(200.0)},
				},
			},
			wantValue:  0.8,
			wantInterp: "negative_value_perception",
		},
		{
			name: "components_with_some_nil_perceived_values",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
				Components: []domain.ComponentData{
					{Name: "A", PerceivedValue: ptr(300.0)},
					{Name: "B"}, // nil perceived value
					{Name: "C", PerceivedValue: ptr(100.0)},
				},
			},
			// average = (300+100)/2 = 200, ratio = 200/100 = 2.0
			wantValue:  2.0,
			wantInterp: "positive_value_perception",
		},
		// Error cases
		{
			name:        "nil_product",
			input:       &domain.AppraisalInput{},
			wantErr:     true,
			errContains: "product definition required",
		},
		{
			name: "no_perceived_value_data",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
			},
			wantErr:     true,
			errContains: "perceived value data required",
		},
		{
			name: "zero_price",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 0.0},
				Tiers: []domain.TierDefinition{
					{Name: "tier", Level: 1, Price: 100.0, PerceivedValue: ptr(50.0)},
				},
			},
			wantErr:     true,
			errContains: "product price must be positive",
		},
		{
			name: "tiers_without_perceived_value_and_no_components",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
				Tiers: []domain.TierDefinition{
					{Name: "basic", Level: 1, Price: 100.0},
				},
			},
			wantErr:     true,
			errContains: "perceived value data required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.PriceValueRatio(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.errContains)
				}
				if tt.errContains != "" && !containsStr(err.Error(), tt.errContains) {
					t.Fatalf("expected error containing %q, got %q", tt.errContains, err.Error())
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !almostEqual(result.Value, tt.wantValue, epsilon) {
				t.Errorf("Value = %v, want %v", result.Value, tt.wantValue)
			}
			if result.Interpretation != tt.wantInterp {
				t.Errorf("Interpretation = %q, want %q", result.Interpretation, tt.wantInterp)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// PremiumPriceIndex tests
// ---------------------------------------------------------------------------

func TestPremiumPriceIndex(t *testing.T) {
	calc := New()

	tests := []struct {
		name        string
		input       *domain.AppraisalInput
		wantValue   float64
		wantInterp  string
		wantErr     bool
		errContains string
	}{
		{
			name: "2x_market_average",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 200.0},
				Market:  &domain.MarketContext{MarketAveragePrice: ptr(100.0)},
			},
			wantValue:  2.0,
			wantInterp: "premium_price_is_2.0x_market_average",
		},
		{
			name: "0.5x_market_average",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 50.0},
				Market:  &domain.MarketContext{MarketAveragePrice: ptr(100.0)},
			},
			wantValue:  0.5,
			wantInterp: "premium_price_is_0.5x_market_average",
		},
		{
			name: "exact_market_average",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
				Market:  &domain.MarketContext{MarketAveragePrice: ptr(100.0)},
			},
			wantValue:  1.0,
			wantInterp: "premium_price_is_1.0x_market_average",
		},
		{
			name: "high_premium",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 500.0},
				Market:  &domain.MarketContext{MarketAveragePrice: ptr(100.0)},
			},
			wantValue:  5.0,
			wantInterp: "premium_price_is_5.0x_market_average",
		},
		// Error cases
		{
			name:        "nil_product",
			input:       &domain.AppraisalInput{Market: &domain.MarketContext{MarketAveragePrice: ptr(100.0)}},
			wantErr:     true,
			errContains: "product definition required",
		},
		{
			name: "nil_market",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
			},
			wantErr:     true,
			errContains: "market average price required",
		},
		{
			name: "nil_market_average_price",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
				Market:  &domain.MarketContext{},
			},
			wantErr:     true,
			errContains: "market average price required",
		},
		{
			name: "zero_market_average_price",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
				Market:  &domain.MarketContext{MarketAveragePrice: ptr(0.0)},
			},
			wantErr:     true,
			errContains: "market average price must be positive",
		},
		{
			name: "negative_market_average_price",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
				Market:  &domain.MarketContext{MarketAveragePrice: ptr(-50.0)},
			},
			wantErr:     true,
			errContains: "market average price must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.PremiumPriceIndex(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.errContains)
				}
				if tt.errContains != "" && !containsStr(err.Error(), tt.errContains) {
					t.Fatalf("expected error containing %q, got %q", tt.errContains, err.Error())
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !almostEqual(result.Value, tt.wantValue, epsilon) {
				t.Errorf("Value = %v, want %v", result.Value, tt.wantValue)
			}
			if result.Interpretation != tt.wantInterp {
				t.Errorf("Interpretation = %q, want %q", result.Interpretation, tt.wantInterp)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// BundleDiscount tests
// ---------------------------------------------------------------------------

func TestBundleDiscount(t *testing.T) {
	calc := New()

	tests := []struct {
		name        string
		input       *domain.AppraisalInput
		wantValue   float64
		wantInterp  string
		wantErr     bool
		errContains string
	}{
		{
			name: "50_percent_discount_aggressive",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 100.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 120.0},
						{Name: "B", StandalonePrice: 80.0},
					},
				},
			},
			// standalone sum = 200, discount = 1 - (100/200) = 0.5 = 50%
			wantValue:  0.5,
			wantInterp: "aggressive_discount",
		},
		{
			name: "20_percent_discount_effective",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 80.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 50.0},
						{Name: "B", StandalonePrice: 50.0},
					},
				},
			},
			// standalone sum = 100, discount = 1 - (80/100) = 0.2 = 20%
			wantValue:  0.2,
			wantInterp: "effective_range",
		},
		{
			name: "negative_discount_bundle_costs_more",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 120.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 50.0},
						{Name: "B", StandalonePrice: 50.0},
					},
				},
			},
			// standalone sum = 100, discount = 1 - (120/100) = -0.2
			wantValue:  -0.2,
			wantInterp: "negative_discount_bundle_costs_more",
		},
		{
			name: "weak_discount_5_percent",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 95.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 50.0},
						{Name: "B", StandalonePrice: 50.0},
					},
				},
			},
			// standalone sum = 100, discount = 1 - (95/100) = 0.05 = 5%
			wantValue:  0.05,
			wantInterp: "weak_discount",
		},
		{
			name: "noticeable_discount_12_percent",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 88.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 50.0},
						{Name: "B", StandalonePrice: 50.0},
					},
				},
			},
			// standalone sum = 100, discount = 1 - (88/100) = 0.12 = 12%
			wantValue:  0.12,
			wantInterp: "noticeable_discount",
		},
		{
			name: "extreme_discount_over_50_percent",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 30.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 50.0},
						{Name: "B", StandalonePrice: 50.0},
					},
				},
			},
			// standalone sum = 100, discount = 1 - (30/100) = 0.7 = 70%
			wantValue:  0.7,
			wantInterp: "extreme_discount_risks_devaluation",
		},
		{
			name: "29_percent_effective_range",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 71.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 50.0},
						{Name: "B", StandalonePrice: 50.0},
					},
				},
			},
			// standalone sum = 100, discount = 1 - (71/100) = 0.29 = 29%
			wantValue:  0.29,
			wantInterp: "effective_range",
		},
		{
			name: "exact_50_percent_boundary_aggressive",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 50.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 50.0},
						{Name: "B", StandalonePrice: 50.0},
					},
				},
			},
			// standalone sum = 100, discount = 1 - (50/100) = 0.5 = 50%
			wantValue:  0.5,
			wantInterp: "aggressive_discount",
		},
		{
			name: "zero_discount",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 100.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 50.0},
						{Name: "B", StandalonePrice: 50.0},
					},
				},
			},
			// standalone sum = 100, discount = 1 - (100/100) = 0 = 0%
			wantValue:  0.0,
			wantInterp: "weak_discount",
		},
		// Error cases
		{
			name:        "nil_product",
			input:       &domain.AppraisalInput{},
			wantErr:     true,
			errContains: "product definition required",
		},
		{
			name: "no_components",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{Price: 100.0},
			},
			wantErr:     true,
			errContains: "product must have components",
		},
		{
			name: "zero_standalone_sum",
			input: &domain.AppraisalInput{
				Product: &domain.ProductDefinition{
					Price: 100.0,
					Components: []domain.Component{
						{Name: "A", StandalonePrice: 0.0},
						{Name: "B", StandalonePrice: 0.0},
					},
				},
			},
			wantErr:     true,
			errContains: "standalone sum must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.BundleDiscount(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.errContains)
				}
				if tt.errContains != "" && !containsStr(err.Error(), tt.errContains) {
					t.Fatalf("expected error containing %q, got %q", tt.errContains, err.Error())
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !almostEqual(result.Value, tt.wantValue, epsilon) {
				t.Errorf("Value = %v, want %v", result.Value, tt.wantValue)
			}
			if result.Interpretation != tt.wantInterp {
				t.Errorf("Interpretation = %q, want %q", result.Interpretation, tt.wantInterp)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// TestNew verifies constructor
// ---------------------------------------------------------------------------

func TestNew(t *testing.T) {
	calc := New()
	if calc == nil {
		t.Fatal("New() returned nil")
	}
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func containsStr(s, substr string) bool {
	return len(s) >= len(substr) && searchStr(s, substr)
}

func searchStr(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
