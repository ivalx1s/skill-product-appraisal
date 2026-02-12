package bundle

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

// almostEqual checks float equality within a small epsilon.
func almostEqual(a, b, eps float64) bool {
	return math.Abs(a-b) < eps
}

const epsilon = 0.0001

// ---------------------------------------------------------------------------
// ClassifyComponents tests
// ---------------------------------------------------------------------------

func TestClassifyComponents(t *testing.T) {
	calc := New()

	tests := []struct {
		name        string
		input       *domain.AppraisalInput
		wantLeaders int
		wantFillers int
		wantKillers int
		wantClasses map[string]string // component name -> classification
		wantErr     bool
		errContains string
	}{
		{
			name: "leader_high_value_drives_purchase",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{
						Name:           "Premium Music",
						PerceivedValue: ptr(4.5),
						MarginalCost:   ptr(2.0),
						DrivesPurchase: boolPtr(true),
					},
				},
			},
			wantLeaders: 1,
			wantFillers: 0,
			wantKillers: 0,
			wantClasses: map[string]string{"Premium Music": "leader"},
		},
		{
			name: "leader_high_value_no_drives_purchase",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{
						Name:           "Premium Video",
						PerceivedValue: ptr(4.2),
						MarginalCost:   ptr(3.0),
					},
				},
			},
			wantLeaders: 1,
			wantFillers: 0,
			wantKillers: 0,
			wantClasses: map[string]string{"Premium Video": "leader"},
		},
		{
			name: "filler_moderate_value_low_cost",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{
						Name:           "Cloud Storage",
						PerceivedValue: ptr(3.0),
						MarginalCost:   ptr(1.0),
					},
				},
			},
			wantLeaders: 0,
			wantFillers: 1,
			wantKillers: 0,
			wantClasses: map[string]string{"Cloud Storage": "filler"},
		},
		{
			name: "killer_low_value_high_cost",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{
						Name:           "Legacy App",
						PerceivedValue: ptr(1.5),
						MarginalCost:   ptr(5.0),
					},
				},
			},
			wantLeaders: 0,
			wantFillers: 0,
			wantKillers: 1,
			wantClasses: map[string]string{"Legacy App": "killer"},
		},
		{
			name: "killer_removal_increases_wtp",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{
						Name:            "Bloatware",
						PerceivedValue:  ptr(3.5),
						MarginalCost:    ptr(1.0),
						RemovalWTPDelta: ptr(2.0), // positive => killer
					},
				},
			},
			wantLeaders: 0,
			wantFillers: 0,
			wantKillers: 1,
			wantClasses: map[string]string{"Bloatware": "killer"},
		},
		{
			name: "filler_low_value_low_cost_with_option_value",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{
						Name:            "Loyalty Points",
						PerceivedValue:  ptr(2.0),
						MarginalCost:    ptr(0.5),
						RemovalWTPDelta: ptr(-1.0), // negative => has option value
					},
				},
			},
			wantLeaders: 0,
			wantFillers: 1,
			wantKillers: 0,
			wantClasses: map[string]string{"Loyalty Points": "filler"},
		},
		{
			name: "filler_low_value_low_cost_no_removal_delta",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{
						Name:           "Newsletter",
						PerceivedValue: ptr(2.0),
						MarginalCost:   ptr(0.5),
					},
				},
			},
			wantLeaders: 0,
			wantFillers: 1,
			wantKillers: 0,
			wantClasses: map[string]string{"Newsletter": "filler"},
		},
		{
			name: "mixed_classifications",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{
						Name:           "Core Service",
						PerceivedValue: ptr(4.8),
						MarginalCost:   ptr(3.0),
						DrivesPurchase: boolPtr(true),
					},
					{
						Name:           "Bonus Feature",
						PerceivedValue: ptr(3.0),
						MarginalCost:   ptr(1.0),
					},
					{
						Name:           "Unwanted Add-on",
						PerceivedValue: ptr(1.0),
						MarginalCost:   ptr(4.0),
					},
					{
						Name:            "Diluter",
						PerceivedValue:  ptr(2.0),
						RemovalWTPDelta: ptr(0.5), // positive => killer
					},
				},
			},
			wantLeaders: 1,
			wantFillers: 1,
			wantKillers: 2,
			wantClasses: map[string]string{
				"Core Service":   "leader",
				"Bonus Feature":  "filler",
				"Unwanted Add-on": "killer",
				"Diluter":        "killer",
			},
		},
		{
			name: "nil_perceived_value_and_cost_defaults_to_zero",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "Empty Component"},
				},
			},
			// pv=0, mc=0, pv<2.5 and mc is not > pv => falls to pv<2.5 default
			wantLeaders: 0,
			wantFillers: 1,
			wantKillers: 0,
			wantClasses: map[string]string{"Empty Component": "filler"},
		},
		{
			name: "boundary_perceived_value_4.0",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{
						Name:           "Boundary Leader",
						PerceivedValue: ptr(4.0),
						MarginalCost:   ptr(2.0),
					},
				},
			},
			wantLeaders: 1,
			wantFillers: 0,
			wantKillers: 0,
			wantClasses: map[string]string{"Boundary Leader": "leader"},
		},
		{
			name: "boundary_perceived_value_2.5",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{
						Name:           "Boundary Filler",
						PerceivedValue: ptr(2.5),
						MarginalCost:   ptr(1.0),
					},
				},
			},
			wantLeaders: 0,
			wantFillers: 1,
			wantKillers: 0,
			wantClasses: map[string]string{"Boundary Filler": "filler"},
		},
		// Error cases
		{
			name:        "no_components",
			input:       &domain.AppraisalInput{},
			wantErr:     true,
			errContains: "component data required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.ClassifyComponents(tt.input)
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
			if result.Leaders != tt.wantLeaders {
				t.Errorf("Leaders = %d, want %d", result.Leaders, tt.wantLeaders)
			}
			if result.Fillers != tt.wantFillers {
				t.Errorf("Fillers = %d, want %d", result.Fillers, tt.wantFillers)
			}
			if result.Killers != tt.wantKillers {
				t.Errorf("Killers = %d, want %d", result.Killers, tt.wantKillers)
			}
			if len(result.Classifications) != len(tt.wantClasses) {
				t.Errorf("got %d classifications, want %d", len(result.Classifications), len(tt.wantClasses))
			}
			for _, cls := range result.Classifications {
				expected, ok := tt.wantClasses[cls.Name]
				if !ok {
					t.Errorf("unexpected component %q in results", cls.Name)
					continue
				}
				if cls.Classification != expected {
					t.Errorf("component %q: classification = %q, want %q", cls.Name, cls.Classification, expected)
				}
				if cls.Rationale == "" {
					t.Errorf("component %q: rationale should not be empty", cls.Name)
				}
			}
		})
	}
}

// ---------------------------------------------------------------------------
// DeadWeightRatio tests
// ---------------------------------------------------------------------------

func TestDeadWeightRatio(t *testing.T) {
	calc := New()

	tests := []struct {
		name           string
		input          *domain.AppraisalInput
		wantRatio      float64
		wantPasses     bool
		wantDeadWeight []string
		wantErr        bool
		errContains    string
	}{
		{
			name: "all_components_active_0_percent",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "A", MonthlyActiveRate: ptr(0.80)},
					{Name: "B", MonthlyActiveRate: ptr(0.60)},
					{Name: "C", MonthlyActiveRate: ptr(0.40)},
				},
			},
			wantRatio:      0.0,
			wantPasses:     true,
			wantDeadWeight: nil,
		},
		{
			name: "all_dead_weight_100_percent",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "A", MonthlyActiveRate: ptr(0.10)},
					{Name: "B", MonthlyActiveRate: ptr(0.05)},
					{Name: "C", MonthlyActiveRate: ptr(0.15)},
				},
			},
			wantRatio:      1.0,
			wantPasses:     false,
			wantDeadWeight: []string{"A", "B", "C"},
		},
		{
			name: "half_dead_weight_50_percent",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "A", MonthlyActiveRate: ptr(0.80)},
					{Name: "B", MonthlyActiveRate: ptr(0.10)},
					{Name: "C", MonthlyActiveRate: ptr(0.60)},
					{Name: "D", MonthlyActiveRate: ptr(0.05)},
				},
			},
			wantRatio:      0.5,
			wantPasses:     false,
			wantDeadWeight: []string{"B", "D"},
		},
		{
			name: "exactly_at_usage_threshold_20_percent",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "A", MonthlyActiveRate: ptr(0.20)},
					{Name: "B", MonthlyActiveRate: ptr(0.50)},
				},
			},
			// 0.20 is not < 0.20, so A is not dead weight
			wantRatio:      0.0,
			wantPasses:     true,
			wantDeadWeight: nil,
		},
		{
			name: "just_below_usage_threshold",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "A", MonthlyActiveRate: ptr(0.19)},
					{Name: "B", MonthlyActiveRate: ptr(0.50)},
				},
			},
			wantRatio:      0.5,
			wantPasses:     false,
			wantDeadWeight: []string{"A"},
		},
		{
			name: "uses_usage_forecast_when_no_monthly_active",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "A", UsageForecast: ptr(0.30)},
					{Name: "B", UsageForecast: ptr(0.10)},
				},
			},
			wantRatio:      0.5,
			wantPasses:     false,
			wantDeadWeight: []string{"B"},
		},
		{
			name: "monthly_active_takes_priority_over_forecast",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "A", MonthlyActiveRate: ptr(0.50), UsageForecast: ptr(0.05)},
					{Name: "B", MonthlyActiveRate: ptr(0.10), UsageForecast: ptr(0.80)},
				},
			},
			// A uses MonthlyActiveRate=0.50 (active), B uses MonthlyActiveRate=0.10 (dead)
			wantRatio:      0.5,
			wantPasses:     false,
			wantDeadWeight: []string{"B"},
		},
		{
			name: "nil_usage_defaults_to_zero",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "A"},
					{Name: "B"},
				},
			},
			// both default to 0.0, which is < 0.20
			wantRatio:      1.0,
			wantPasses:     false,
			wantDeadWeight: []string{"A", "B"},
		},
		{
			name: "ratio_below_threshold_passes",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "A", MonthlyActiveRate: ptr(0.80)},
					{Name: "B", MonthlyActiveRate: ptr(0.60)},
					{Name: "C", MonthlyActiveRate: ptr(0.50)},
					{Name: "D", MonthlyActiveRate: ptr(0.10)}, // dead
				},
			},
			// 1/4 = 0.25 which is < 0.40
			wantRatio:      0.25,
			wantPasses:     true,
			wantDeadWeight: []string{"D"},
		},
		{
			name: "single_component_active",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "Solo", MonthlyActiveRate: ptr(0.90)},
				},
			},
			wantRatio:      0.0,
			wantPasses:     true,
			wantDeadWeight: nil,
		},
		{
			name: "single_component_dead",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "Solo", MonthlyActiveRate: ptr(0.05)},
				},
			},
			wantRatio:      1.0,
			wantPasses:     false,
			wantDeadWeight: []string{"Solo"},
		},
		// Error cases
		{
			name:        "no_components",
			input:       &domain.AppraisalInput{},
			wantErr:     true,
			errContains: "component data required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.DeadWeightRatio(tt.input)
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
			if !almostEqual(result.DeadWeightRatio, tt.wantRatio, epsilon) {
				t.Errorf("DeadWeightRatio = %v, want %v", result.DeadWeightRatio, tt.wantRatio)
			}
			if result.Passes != tt.wantPasses {
				t.Errorf("Passes = %v, want %v", result.Passes, tt.wantPasses)
			}
			if !almostEqual(result.Threshold, 0.40, epsilon) {
				t.Errorf("Threshold = %v, want 0.40", result.Threshold)
			}
			// Verify dead weight list
			if len(result.DeadWeight) != len(tt.wantDeadWeight) {
				t.Errorf("DeadWeight count = %d, want %d; got %v", len(result.DeadWeight), len(tt.wantDeadWeight), result.DeadWeight)
			} else {
				deadSet := make(map[string]bool)
				for _, name := range result.DeadWeight {
					deadSet[name] = true
				}
				for _, expected := range tt.wantDeadWeight {
					if !deadSet[expected] {
						t.Errorf("expected %q in DeadWeight, got %v", expected, result.DeadWeight)
					}
				}
			}
			// Verify component usage map
			for _, comp := range tt.input.Components {
				if _, ok := result.ComponentUsage[comp.Name]; !ok {
					t.Errorf("missing component %q in ComponentUsage map", comp.Name)
				}
			}
		})
	}
}

// ---------------------------------------------------------------------------
// CrossSubsidyAnalysis tests
// ---------------------------------------------------------------------------

func TestCrossSubsidyAnalysis(t *testing.T) {
	calc := New()

	tests := []struct {
		name           string
		input          *domain.AppraisalInput
		wantMargin     float64
		wantSustain    bool
		wantSources    int
		wantRecipients int
		wantErr        bool
		errContains    string
	}{
		{
			name: "sources_and_recipients_sustainable",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "Core", RevenueContrib: ptr(100.0), DirectCost: ptr(30.0)},
					{Name: "Add-on", RevenueContrib: ptr(20.0), DirectCost: ptr(50.0)},
				},
			},
			// Core margin: 70, Add-on margin: -30, total: 40
			wantMargin:     40.0,
			wantSustain:    true,
			wantSources:    1,
			wantRecipients: 1,
		},
		{
			name: "all_sources_no_recipients",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "A", RevenueContrib: ptr(100.0), DirectCost: ptr(20.0)},
					{Name: "B", RevenueContrib: ptr(80.0), DirectCost: ptr(30.0)},
				},
			},
			// A margin: 80, B margin: 50, total: 130
			wantMargin:     130.0,
			wantSustain:    true,
			wantSources:    2,
			wantRecipients: 0,
		},
		{
			name: "all_recipients_unsustainable",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "A", RevenueContrib: ptr(10.0), DirectCost: ptr(50.0)},
					{Name: "B", RevenueContrib: ptr(5.0), DirectCost: ptr(30.0)},
				},
			},
			// A margin: -40, B margin: -25, total: -65
			wantMargin:     -65.0,
			wantSustain:    false,
			wantSources:    0,
			wantRecipients: 2,
		},
		{
			name: "balanced_exactly_zero",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "A", RevenueContrib: ptr(50.0), DirectCost: ptr(20.0)},
					{Name: "B", RevenueContrib: ptr(10.0), DirectCost: ptr(40.0)},
				},
			},
			// A margin: 30, B margin: -30, total: 0
			wantMargin:     0.0,
			wantSustain:    false, // 0 is not > 0
			wantSources:    1,
			wantRecipients: 1,
		},
		{
			name: "uses_marginal_cost_when_no_direct_cost",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "A", RevenueContrib: ptr(80.0), MarginalCost: ptr(30.0)},
					{Name: "B", RevenueContrib: ptr(20.0), MarginalCost: ptr(50.0)},
				},
			},
			// A margin: 50, B margin: -30, total: 20
			wantMargin:     20.0,
			wantSustain:    true,
			wantSources:    1,
			wantRecipients: 1,
		},
		{
			name: "direct_cost_takes_priority_over_marginal",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "A", RevenueContrib: ptr(100.0), DirectCost: ptr(40.0), MarginalCost: ptr(10.0)},
				},
			},
			// Uses DirectCost: margin = 100 - 40 = 60
			wantMargin:     60.0,
			wantSustain:    true,
			wantSources:    1,
			wantRecipients: 0,
		},
		{
			name: "nil_revenue_and_cost_default_to_zero",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "Empty"},
				},
			},
			// rev=0, cost=0, margin=0
			wantMargin:     0.0,
			wantSustain:    false,
			wantSources:    1, // margin=0 >= 0, so "source"
			wantRecipients: 0,
		},
		{
			name: "many_components_mixed",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "High Margin", RevenueContrib: ptr(200.0), DirectCost: ptr(50.0)},
					{Name: "Medium Margin", RevenueContrib: ptr(80.0), DirectCost: ptr(60.0)},
					{Name: "Low Margin", RevenueContrib: ptr(10.0), DirectCost: ptr(40.0)},
					{Name: "Break-even", RevenueContrib: ptr(30.0), DirectCost: ptr(30.0)},
				},
			},
			// 150 + 20 + (-30) + 0 = 140
			wantMargin:     140.0,
			wantSustain:    true,
			wantSources:    3, // High, Medium, Break-even (margin >= 0)
			wantRecipients: 1, // Low
		},
		// Error cases
		{
			name:        "no_components",
			input:       &domain.AppraisalInput{},
			wantErr:     true,
			errContains: "component data required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.CrossSubsidyAnalysis(tt.input)
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
			if !almostEqual(result.NetMargin, tt.wantMargin, epsilon) {
				t.Errorf("NetMargin = %v, want %v", result.NetMargin, tt.wantMargin)
			}
			if result.Sustainable != tt.wantSustain {
				t.Errorf("Sustainable = %v, want %v", result.Sustainable, tt.wantSustain)
			}
			if len(result.Sources) != tt.wantSources {
				t.Errorf("Sources count = %d, want %d", len(result.Sources), tt.wantSources)
			}
			if len(result.Recipients) != tt.wantRecipients {
				t.Errorf("Recipients count = %d, want %d", len(result.Recipients), tt.wantRecipients)
			}
			// Verify roles
			for _, src := range result.Sources {
				if src.Role != "source" {
					t.Errorf("source %q has role %q, want %q", src.Name, src.Role, "source")
				}
				if src.NetMargin < 0 {
					t.Errorf("source %q has negative margin %v", src.Name, src.NetMargin)
				}
			}
			for _, rec := range result.Recipients {
				if rec.Role != "recipient" {
					t.Errorf("recipient %q has role %q, want %q", rec.Name, rec.Role, "recipient")
				}
				if rec.NetMargin >= 0 {
					t.Errorf("recipient %q has non-negative margin %v", rec.Name, rec.NetMargin)
				}
			}
		})
	}
}

// ---------------------------------------------------------------------------
// ComponentActivation tests
// ---------------------------------------------------------------------------

func TestComponentActivation(t *testing.T) {
	calc := New()

	tests := []struct {
		name        string
		input       *domain.AppraisalInput
		wantCount   int
		wantRates   map[string]float64
		wantInterps map[string]string // substring checks
		wantErr     bool
		errContains string
	}{
		{
			name: "leaders_above_70_percent",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "Core", Activation30d: ptr(0.85)},
					{Name: "Premium", Activation30d: ptr(0.75)},
				},
			},
			wantCount: 2,
			wantRates: map[string]float64{"Core": 0.85, "Premium": 0.75},
			wantInterps: map[string]string{
				"Core":    "strong_activation",
				"Premium": "strong_activation",
			},
		},
		{
			name: "fillers_above_40_percent",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "Bonus", Activation30d: ptr(0.55)},
					{Name: "Extra", Activation30d: ptr(0.40)},
				},
			},
			wantCount: 2,
			wantRates: map[string]float64{"Bonus": 0.55, "Extra": 0.40},
			wantInterps: map[string]string{
				"Bonus": "moderate_activation",
				"Extra": "moderate_activation",
			},
		},
		{
			name: "below_threshold_weak",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "Weak", Activation30d: ptr(0.15)},
					{Name: "Dead", Activation30d: ptr(0.05)},
				},
			},
			wantCount: 2,
			wantRates: map[string]float64{"Weak": 0.15, "Dead": 0.05},
			wantInterps: map[string]string{
				"Weak": "weak_activation",
				"Dead": "weak_activation",
			},
		},
		{
			name: "mixed_activation_levels",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "Leader", Activation30d: ptr(0.90)},
					{Name: "Filler", Activation30d: ptr(0.55)},
					{Name: "Weak", Activation30d: ptr(0.20)},
				},
			},
			wantCount: 3,
			wantRates: map[string]float64{"Leader": 0.90, "Filler": 0.55, "Weak": 0.20},
			wantInterps: map[string]string{
				"Leader": "strong_activation",
				"Filler": "moderate_activation",
				"Weak":   "weak_activation",
			},
		},
		{
			name: "nil_activation_defaults_to_zero",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "NoData"},
				},
			},
			wantCount: 1,
			wantRates: map[string]float64{"NoData": 0.0},
			wantInterps: map[string]string{
				"NoData": "weak_activation",
			},
		},
		{
			name: "boundary_exactly_70_percent",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "Boundary", Activation30d: ptr(0.70)},
				},
			},
			wantCount: 1,
			wantRates: map[string]float64{"Boundary": 0.70},
			wantInterps: map[string]string{
				"Boundary": "strong_activation",
			},
		},
		{
			name: "boundary_just_below_70_percent",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "AlmostLeader", Activation30d: ptr(0.69)},
				},
			},
			wantCount: 1,
			wantRates: map[string]float64{"AlmostLeader": 0.69},
			wantInterps: map[string]string{
				"AlmostLeader": "moderate_activation",
			},
		},
		{
			name: "boundary_just_below_40_percent",
			input: &domain.AppraisalInput{
				Components: []domain.ComponentData{
					{Name: "AlmostFiller", Activation30d: ptr(0.39)},
				},
			},
			wantCount: 1,
			wantRates: map[string]float64{"AlmostFiller": 0.39},
			wantInterps: map[string]string{
				"AlmostFiller": "weak_activation",
			},
		},
		// Error cases
		{
			name:        "no_components",
			input:       &domain.AppraisalInput{},
			wantErr:     true,
			errContains: "component data required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, err := calc.ComponentActivation(tt.input)
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
			if len(results) != tt.wantCount {
				t.Fatalf("got %d results, want %d", len(results), tt.wantCount)
			}
			for i, r := range results {
				compName := tt.input.Components[i].Name
				expectedRate, ok := tt.wantRates[compName]
				if ok && !almostEqual(r.Value, expectedRate, epsilon) {
					t.Errorf("component %q: rate = %v, want %v", compName, r.Value, expectedRate)
				}
				expectedInterp, ok := tt.wantInterps[compName]
				if ok && !containsStr(r.Interpretation, expectedInterp) {
					t.Errorf("component %q: interpretation = %q, want to contain %q", compName, r.Interpretation, expectedInterp)
				}
			}
		})
	}
}

// ---------------------------------------------------------------------------
// MultiComponentUsage tests
// ---------------------------------------------------------------------------

func TestMultiComponentUsage(t *testing.T) {
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
			name: "all_users_3_plus_healthy",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					CustomersUsing3Plus: ptr(1000.0),
					PremiumCustomers:    ptr(1000.0),
				},
			},
			wantValue:  1.0,
			wantInterp: "healthy_multi_component_usage",
		},
		{
			name: "none_using_3_plus",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					CustomersUsing3Plus: ptr(0.0),
					PremiumCustomers:    ptr(1000.0),
				},
			},
			wantValue:  0.0,
			wantInterp: "low_multi_component_usage_poor_bundle_composition",
		},
		{
			name: "above_60_percent_threshold",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					CustomersUsing3Plus: ptr(700.0),
					PremiumCustomers:    ptr(1000.0),
				},
			},
			wantValue:  0.7,
			wantInterp: "healthy_multi_component_usage",
		},
		{
			name: "below_60_percent_threshold",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					CustomersUsing3Plus: ptr(500.0),
					PremiumCustomers:    ptr(1000.0),
				},
			},
			wantValue:  0.5,
			wantInterp: "low_multi_component_usage_poor_bundle_composition",
		},
		{
			name: "exactly_60_percent_threshold",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					CustomersUsing3Plus: ptr(600.0),
					PremiumCustomers:    ptr(1000.0),
				},
			},
			wantValue:  0.6,
			wantInterp: "healthy_multi_component_usage",
		},
		{
			name: "just_below_60_percent",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					CustomersUsing3Plus: ptr(599.0),
					PremiumCustomers:    ptr(1000.0),
				},
			},
			wantValue:  0.599,
			wantInterp: "low_multi_component_usage_poor_bundle_composition",
		},
		// Error cases
		{
			name:        "nil_customers",
			input:       &domain.AppraisalInput{},
			wantErr:     true,
			errContains: "customer metrics required",
		},
		{
			name: "nil_customers_using_3plus",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					PremiumCustomers: ptr(1000.0),
				},
			},
			wantErr:     true,
			errContains: "customers_using_3plus required",
		},
		{
			name: "nil_premium_customers",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					CustomersUsing3Plus: ptr(500.0),
				},
			},
			wantErr:     true,
			errContains: "premium_customers required and must be positive",
		},
		{
			name: "zero_premium_customers",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					CustomersUsing3Plus: ptr(500.0),
					PremiumCustomers:    ptr(0.0),
				},
			},
			wantErr:     true,
			errContains: "premium_customers required and must be positive",
		},
		{
			name: "negative_premium_customers",
			input: &domain.AppraisalInput{
				Customers: &domain.CustomerMetrics{
					CustomersUsing3Plus: ptr(500.0),
					PremiumCustomers:    ptr(-100.0),
				},
			},
			wantErr:     true,
			errContains: "premium_customers required and must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.MultiComponentUsage(tt.input)
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
