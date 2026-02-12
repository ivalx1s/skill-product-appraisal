package product

import (
	"math"
	"testing"

	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/domain"
)

// ptr creates a pointer to a float64 value.
func ptr(v float64) *float64 {
	return &v
}

func approxEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

// ---------------------------------------------------------------------------
// PenetrationRate
// ---------------------------------------------------------------------------

func TestPenetrationRate(t *testing.T) {
	calc := New()

	tests := []struct {
		name           string
		input          *domain.AppraisalInput
		wantValue      float64
		wantInterp     string
		wantErr        bool
		errContains    string
	}{
		{
			name: "typical 10% penetration",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PremiumCustomers: ptr(10000),
					TotalCustomers:   ptr(100000),
				},
			},
			wantValue:  0.10,
			wantInterp: "moderate_penetration",
		},
		{
			name: "very low 1% penetration",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PremiumCustomers: ptr(1000),
					TotalCustomers:   ptr(100000),
				},
			},
			wantValue:  0.01,
			wantInterp: "low_penetration",
		},
		{
			name: "high 25% penetration",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PremiumCustomers: ptr(25000),
					TotalCustomers:   ptr(100000),
				},
			},
			wantValue:  0.25,
			wantInterp: "high_penetration",
		},
		{
			name: "growing 5% boundary",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PremiumCustomers: ptr(5000),
					TotalCustomers:   ptr(100000),
				},
			},
			wantValue:  0.05,
			wantInterp: "growing_penetration",
		},
		{
			name: "just below growing threshold 4.9%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PremiumCustomers: ptr(4900),
					TotalCustomers:   ptr(100000),
				},
			},
			wantValue:  0.049,
			wantInterp: "low_penetration",
		},
		{
			name: "above high threshold 30%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PremiumCustomers: ptr(30000),
					TotalCustomers:   ptr(100000),
				},
			},
			wantValue:  0.30,
			wantInterp: "high_penetration",
		},
		{
			name:        "nil customers",
			input:       &domain.AppraisalInput{},
			wantErr:     true,
			errContains: "customer metrics required",
		},
		{
			name: "nil premium customers",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					TotalCustomers: ptr(100000),
				},
			},
			wantErr:     true,
			errContains: "premium_customers and total_customers required",
		},
		{
			name: "nil total customers",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PremiumCustomers: ptr(10000),
				},
			},
			wantErr:     true,
			errContains: "premium_customers and total_customers required",
		},
		{
			name: "zero total customers",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PremiumCustomers: ptr(100),
					TotalCustomers:   ptr(0),
				},
			},
			wantErr:     true,
			errContains: "total_customers must be positive",
		},
		{
			name: "negative total customers",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PremiumCustomers: ptr(100),
					TotalCustomers:   ptr(-1000),
				},
			},
			wantErr:     true,
			errContains: "total_customers must be positive",
		},
		{
			name: "zero premium customers",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PremiumCustomers: ptr(0),
					TotalCustomers:   ptr(100000),
				},
			},
			wantValue:  0.0,
			wantInterp: "low_penetration",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.PenetrationRate(tt.input)
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
			if !approxEqual(result.Value, tt.wantValue, 1e-9) {
				t.Errorf("value = %f, want %f", result.Value, tt.wantValue)
			}
			if result.Interpretation != tt.wantInterp {
				t.Errorf("interpretation = %q, want %q", result.Interpretation, tt.wantInterp)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// MigrationRate
// ---------------------------------------------------------------------------

func TestMigrationRate(t *testing.T) {
	calc := New()

	tests := []struct {
		name        string
		input       *domain.AppraisalInput
		wantValue   float64
		wantErr     bool
		errContains string
	}{
		{
			name: "typical monthly 3%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					UpgradedCustomers: ptr(3000),
					EligibleBase:      ptr(100000),
				},
			},
			wantValue: 0.03,
		},
		{
			name: "high 10%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					UpgradedCustomers: ptr(10000),
					EligibleBase:      ptr(100000),
				},
			},
			wantValue: 0.10,
		},
		{
			name: "zero upgrades",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					UpgradedCustomers: ptr(0),
					EligibleBase:      ptr(100000),
				},
			},
			wantValue: 0.0,
		},
		{
			name:        "nil customers",
			input:       &domain.AppraisalInput{},
			wantErr:     true,
			errContains: "customer metrics required",
		},
		{
			name: "missing upgraded customers",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					EligibleBase: ptr(100000),
				},
			},
			wantErr:     true,
			errContains: "upgraded_customers and eligible_base required",
		},
		{
			name: "missing eligible base",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					UpgradedCustomers: ptr(3000),
				},
			},
			wantErr:     true,
			errContains: "upgraded_customers and eligible_base required",
		},
		{
			name: "zero eligible base",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					UpgradedCustomers: ptr(3000),
					EligibleBase:      ptr(0),
				},
			},
			wantErr:     true,
			errContains: "eligible_base must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.MigrationRate(tt.input)
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
			if !approxEqual(result.Value, tt.wantValue, 1e-9) {
				t.Errorf("value = %f, want %f", result.Value, tt.wantValue)
			}
			// Check interpretation format
			if result.Interpretation == "" {
				t.Error("interpretation should not be empty")
			}
		})
	}
}

// ---------------------------------------------------------------------------
// CannibalizationRate
// ---------------------------------------------------------------------------

func TestCannibalizationRate(t *testing.T) {
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
			name: "high cannibalization 80%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					MigratedFromStandalone: ptr(8000),
					PremiumCustomers:       ptr(10000),
				},
			},
			wantValue:  0.80,
			wantInterp: "high_cannibalization_risk",
		},
		{
			name: "low cannibalization 20%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					MigratedFromStandalone: ptr(2000),
					PremiumCustomers:       ptr(10000),
				},
			},
			wantValue:  0.20,
			wantInterp: "acceptable_cannibalization",
		},
		{
			name: "zero cannibalization",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					MigratedFromStandalone: ptr(0),
					PremiumCustomers:       ptr(10000),
				},
			},
			wantValue:  0.0,
			wantInterp: "acceptable_cannibalization",
		},
		{
			name: "100% cannibalization",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					MigratedFromStandalone: ptr(10000),
					PremiumCustomers:       ptr(10000),
				},
			},
			wantValue:  1.0,
			wantInterp: "high_cannibalization_risk",
		},
		{
			name: "moderate cannibalization boundary 31%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					MigratedFromStandalone: ptr(3100),
					PremiumCustomers:       ptr(10000),
				},
			},
			wantValue:  0.31,
			wantInterp: "moderate_cannibalization",
		},
		{
			name: "moderate boundary exact 50%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					MigratedFromStandalone: ptr(5000),
					PremiumCustomers:       ptr(10000),
				},
			},
			wantValue:  0.50,
			wantInterp: "moderate_cannibalization",
		},
		{
			name: "boundary exactly 30% is acceptable",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					MigratedFromStandalone: ptr(3000),
					PremiumCustomers:       ptr(10000),
				},
			},
			wantValue:  0.30,
			wantInterp: "acceptable_cannibalization",
		},
		{
			name: "just above 50% threshold",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					MigratedFromStandalone: ptr(5001),
					PremiumCustomers:       ptr(10000),
				},
			},
			wantValue:  0.5001,
			wantInterp: "high_cannibalization_risk",
		},
		{
			name:        "nil customers",
			input:       &domain.AppraisalInput{},
			wantErr:     true,
			errContains: "customer metrics required",
		},
		{
			name: "missing migrated",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PremiumCustomers: ptr(10000),
				},
			},
			wantErr:     true,
			errContains: "migrated_from_standalone and premium_customers required",
		},
		{
			name: "zero premium customers",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					MigratedFromStandalone: ptr(1000),
					PremiumCustomers:       ptr(0),
				},
			},
			wantErr:     true,
			errContains: "premium_customers must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.CannibalizationRate(tt.input)
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
			if !approxEqual(result.Value, tt.wantValue, 1e-9) {
				t.Errorf("value = %f, want %f", result.Value, tt.wantValue)
			}
			if result.Interpretation != tt.wantInterp {
				t.Errorf("interpretation = %q, want %q", result.Interpretation, tt.wantInterp)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// CrossSellRate
// ---------------------------------------------------------------------------

func TestCrossSellRate(t *testing.T) {
	calc := New()

	tests := []struct {
		name        string
		input       *domain.AppraisalInput
		wantValue   float64
		wantErr     bool
		errContains string
	}{
		{
			name: "typical 30%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PremiumBuyingAddons: ptr(3000),
					PremiumCustomers:    ptr(10000),
				},
			},
			wantValue: 0.30,
		},
		{
			name: "low 5%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PremiumBuyingAddons: ptr(500),
					PremiumCustomers:    ptr(10000),
				},
			},
			wantValue: 0.05,
		},
		{
			name: "high 60%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PremiumBuyingAddons: ptr(6000),
					PremiumCustomers:    ptr(10000),
				},
			},
			wantValue: 0.60,
		},
		{
			name: "zero cross-sell",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PremiumBuyingAddons: ptr(0),
					PremiumCustomers:    ptr(10000),
				},
			},
			wantValue: 0.0,
		},
		{
			name:        "nil customers",
			input:       &domain.AppraisalInput{},
			wantErr:     true,
			errContains: "customer metrics required",
		},
		{
			name: "missing addons",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PremiumCustomers: ptr(10000),
				},
			},
			wantErr:     true,
			errContains: "premium_buying_addons and premium_customers required",
		},
		{
			name: "zero premium customers",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PremiumBuyingAddons: ptr(100),
					PremiumCustomers:    ptr(0),
				},
			},
			wantErr:     true,
			errContains: "premium_customers must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.CrossSellRate(tt.input)
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
			if !approxEqual(result.Value, tt.wantValue, 1e-9) {
				t.Errorf("value = %f, want %f", result.Value, tt.wantValue)
			}
			if result.Interpretation == "" {
				t.Error("interpretation should not be empty")
			}
		})
	}
}

// ---------------------------------------------------------------------------
// FeatureUtilizationRate
// ---------------------------------------------------------------------------

func TestFeatureUtilizationRate(t *testing.T) {
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
			name: "above target 70%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					FeaturesUsedPerCustomer: ptr(7),
					TotalAvailableFeatures:  ptr(10),
				},
			},
			wantValue:  0.70,
			wantInterp: "healthy_feature_utilization",
		},
		{
			name: "exactly at target 60%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					FeaturesUsedPerCustomer: ptr(6),
					TotalAvailableFeatures:  ptr(10),
				},
			},
			wantValue:  0.60,
			wantInterp: "healthy_feature_utilization",
		},
		{
			name: "below target 40%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					FeaturesUsedPerCustomer: ptr(4),
					TotalAvailableFeatures:  ptr(10),
				},
			},
			wantValue:  0.40,
			wantInterp: "low_utilization_over_provisioning_risk",
		},
		{
			name: "all features used 100%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					FeaturesUsedPerCustomer: ptr(10),
					TotalAvailableFeatures:  ptr(10),
				},
			},
			wantValue:  1.0,
			wantInterp: "healthy_feature_utilization",
		},
		{
			name: "no features used 0%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					FeaturesUsedPerCustomer: ptr(0),
					TotalAvailableFeatures:  ptr(10),
				},
			},
			wantValue:  0.0,
			wantInterp: "low_utilization_over_provisioning_risk",
		},
		{
			name: "just below target 59%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					FeaturesUsedPerCustomer: ptr(5.9),
					TotalAvailableFeatures:  ptr(10),
				},
			},
			wantValue:  0.59,
			wantInterp: "low_utilization_over_provisioning_risk",
		},
		{
			name:        "nil customers",
			input:       &domain.AppraisalInput{},
			wantErr:     true,
			errContains: "customer metrics required",
		},
		{
			name: "missing features used",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					TotalAvailableFeatures: ptr(10),
				},
			},
			wantErr:     true,
			errContains: "features_used_per_customer and total_available_features required",
		},
		{
			name: "zero total features",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					FeaturesUsedPerCustomer: ptr(5),
					TotalAvailableFeatures:  ptr(0),
				},
			},
			wantErr:     true,
			errContains: "total_available_features must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.FeatureUtilizationRate(tt.input)
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
			if !approxEqual(result.Value, tt.wantValue, 1e-9) {
				t.Errorf("value = %f, want %f", result.Value, tt.wantValue)
			}
			if result.Interpretation != tt.wantInterp {
				t.Errorf("interpretation = %q, want %q", result.Interpretation, tt.wantInterp)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// ComponentActivationRate
// ---------------------------------------------------------------------------

func TestComponentActivationRate(t *testing.T) {
	calc := New()

	tests := []struct {
		name       string
		input      *domain.AppraisalInput
		wantValues []float64
		wantInterps []string
		wantErr    bool
		errContains string
	}{
		{
			name: "leader level >70%",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "streaming", Activation30d: ptr(0.80)},
				},
			},
			wantValues:  []float64{0.80},
			wantInterps: []string{"streaming: leader_level_activation"},
		},
		{
			name: "filler level >40%",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "cloud_storage", Activation30d: ptr(0.55)},
				},
			},
			wantValues:  []float64{0.55},
			wantInterps: []string{"cloud_storage: filler_level_activation"},
		},
		{
			name: "below target <40%",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "vpn", Activation30d: ptr(0.15)},
				},
			},
			wantValues:  []float64{0.15},
			wantInterps: []string{"vpn: below_target"},
		},
		{
			name: "multiple components mixed",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "music", Activation30d: ptr(0.75)},
					{Name: "vpn", Activation30d: ptr(0.45)},
					{Name: "insurance", Activation30d: ptr(0.10)},
				},
			},
			wantValues: []float64{0.75, 0.45, 0.10},
			wantInterps: []string{
				"music: leader_level_activation",
				"vpn: filler_level_activation",
				"insurance: below_target",
			},
		},
		{
			name: "nil activation defaults to 0",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "unknown"},
				},
			},
			wantValues:  []float64{0.0},
			wantInterps: []string{"unknown: below_target"},
		},
		{
			name: "exact 70% boundary is leader",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "comp", Activation30d: ptr(0.70)},
				},
			},
			wantValues:  []float64{0.70},
			wantInterps: []string{"comp: leader_level_activation"},
		},
		{
			name: "exact 40% boundary is filler",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "comp", Activation30d: ptr(0.40)},
				},
			},
			wantValues:  []float64{0.40},
			wantInterps: []string{"comp: filler_level_activation"},
		},
		{
			name:        "no components",
			input:       &domain.AppraisalInput{},
			wantErr:     true,
			errContains: "component data required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, err := calc.ComponentActivationRate(tt.input)
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
			if len(results) != len(tt.wantValues) {
				t.Fatalf("got %d results, want %d", len(results), len(tt.wantValues))
			}
			for i, r := range results {
				if !approxEqual(r.Value, tt.wantValues[i], 1e-9) {
					t.Errorf("results[%d].Value = %f, want %f", i, r.Value, tt.wantValues[i])
				}
				if r.Interpretation != tt.wantInterps[i] {
					t.Errorf("results[%d].Interpretation = %q, want %q", i, r.Interpretation, tt.wantInterps[i])
				}
			}
		})
	}
}

// ---------------------------------------------------------------------------
// AttachRate
// ---------------------------------------------------------------------------

func TestAttachRate(t *testing.T) {
	calc := New()

	tests := []struct {
		name        string
		input       *domain.AppraisalInput
		wantValues  []float64
		wantInterps []string
		wantErr     bool
		errContains string
	}{
		{
			name: "typical monthly active component",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "streaming", MonthlyActiveRate: ptr(0.45)},
				},
			},
			wantValues:  []float64{0.45},
			wantInterps: []string{"streaming: active_component"},
		},
		{
			name: "declining dead weight signal",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "vpn", MonthlyActiveRate: ptr(0.10)},
				},
			},
			wantValues:  []float64{0.10},
			wantInterps: []string{"vpn: declining_attach_dead_weight_signal"},
		},
		{
			name: "zero attach rate",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "unused", MonthlyActiveRate: ptr(0.0)},
				},
			},
			wantValues:  []float64{0.0},
			wantInterps: []string{"unused: declining_attach_dead_weight_signal"},
		},
		{
			name: "exact 20% boundary is active",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "comp", MonthlyActiveRate: ptr(0.20)},
				},
			},
			wantValues:  []float64{0.20},
			wantInterps: []string{"comp: active_component"},
		},
		{
			name: "just below 20% boundary",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "comp", MonthlyActiveRate: ptr(0.199)},
				},
			},
			wantValues:  []float64{0.199},
			wantInterps: []string{"comp: declining_attach_dead_weight_signal"},
		},
		{
			name: "nil monthly active rate defaults to 0",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "comp"},
				},
			},
			wantValues:  []float64{0.0},
			wantInterps: []string{"comp: declining_attach_dead_weight_signal"},
		},
		{
			name: "multiple components mixed",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "music", MonthlyActiveRate: ptr(0.65)},
					{Name: "vpn", MonthlyActiveRate: ptr(0.05)},
				},
			},
			wantValues: []float64{0.65, 0.05},
			wantInterps: []string{
				"music: active_component",
				"vpn: declining_attach_dead_weight_signal",
			},
		},
		{
			name:        "no components",
			input:       &domain.AppraisalInput{},
			wantErr:     true,
			errContains: "component data required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, err := calc.AttachRate(tt.input)
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
			if len(results) != len(tt.wantValues) {
				t.Fatalf("got %d results, want %d", len(results), len(tt.wantValues))
			}
			for i, r := range results {
				if !approxEqual(r.Value, tt.wantValues[i], 1e-9) {
					t.Errorf("results[%d].Value = %f, want %f", i, r.Value, tt.wantValues[i])
				}
				if r.Interpretation != tt.wantInterps[i] {
					t.Errorf("results[%d].Interpretation = %q, want %q", i, r.Interpretation, tt.wantInterps[i])
				}
			}
		})
	}
}

// ---------------------------------------------------------------------------
// TrialConversion
// ---------------------------------------------------------------------------

func TestTrialConversion(t *testing.T) {
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
			name: "self-serve level 3%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PaidConversions: ptr(300),
					TrialUsers:      ptr(10000),
				},
			},
			wantValue:  0.03,
			wantInterp: "self_serve_level",
		},
		{
			name: "self-serve level 5%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PaidConversions: ptr(500),
					TrialUsers:      ptr(10000),
				},
			},
			wantValue:  0.05,
			wantInterp: "sales_assisted_level",
		},
		{
			name: "sales-assisted level 7%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PaidConversions: ptr(700),
					TrialUsers:      ptr(10000),
				},
			},
			wantValue:  0.07,
			wantInterp: "sales_assisted_level",
		},
		{
			name: "top performer 15%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PaidConversions: ptr(1500),
					TrialUsers:      ptr(10000),
				},
			},
			wantValue:  0.15,
			wantInterp: "top_performer_conversion",
		},
		{
			name: "top performer boundary 8%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PaidConversions: ptr(800),
					TrialUsers:      ptr(10000),
				},
			},
			wantValue:  0.08,
			wantInterp: "top_performer_conversion",
		},
		{
			name: "below benchmark 1%",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PaidConversions: ptr(100),
					TrialUsers:      ptr(10000),
				},
			},
			wantValue:  0.01,
			wantInterp: "below_benchmark",
		},
		{
			name: "zero conversions",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PaidConversions: ptr(0),
					TrialUsers:      ptr(10000),
				},
			},
			wantValue:  0.0,
			wantInterp: "below_benchmark",
		},
		{
			name:        "nil customers",
			input:       &domain.AppraisalInput{},
			wantErr:     true,
			errContains: "customer metrics required",
		},
		{
			name: "missing paid conversions",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					TrialUsers: ptr(10000),
				},
			},
			wantErr:     true,
			errContains: "paid_conversions and trial_users required",
		},
		{
			name: "missing trial users",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PaidConversions: ptr(300),
				},
			},
			wantErr:     true,
			errContains: "paid_conversions and trial_users required",
		},
		{
			name: "zero trial users",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PaidConversions: ptr(300),
					TrialUsers:      ptr(0),
				},
			},
			wantErr:     true,
			errContains: "trial_users must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.TrialConversion(tt.input)
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
			if !approxEqual(result.Value, tt.wantValue, 1e-9) {
				t.Errorf("value = %f, want %f", result.Value, tt.wantValue)
			}
			if result.Interpretation != tt.wantInterp {
				t.Errorf("interpretation = %q, want %q", result.Interpretation, tt.wantInterp)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func containsStr(s, substr string) bool {
	return len(s) >= len(substr) && searchStr(s, substr)
}

func searchStr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
