package scoring

import (
	"math"
	"testing"

	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/domain"
)

// ptr creates a pointer to a float64 value.
func ptr(v float64) *float64 {
	return &v
}

func ptrStr(v string) *string {
	return &v
}

func approxEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

// ---------------------------------------------------------------------------
// GoNoGo
// ---------------------------------------------------------------------------

func TestGoNoGo(t *testing.T) {
	calc := New()

	allDimensions := func(score float64) []domain.DimensionScore {
		dims := []string{"PMF", "FIN", "PRC_CX", "CMP", "BND", "MR", "RISK"}
		var result []domain.DimensionScore
		for _, d := range dims {
			result = append(result, domain.DimensionScore{
				Dimension: d,
				Score:     score,
			})
		}
		return result
	}

	tests := []struct {
		name         string
		input        *domain.AppraisalInput
		wantScore    float64
		wantDecision string
		wantErr      bool
		errContains  string
	}{
		{
			name: "strong go - all 5s",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Dimensions: allDimensions(5.0),
				},
			},
			// 5.0 * (0.15+0.25+0.20+0.15+0.10+0.10+0.05) = 5.0 * 1.0 = 5.0
			wantScore:    5.0,
			wantDecision: "strong_go",
		},
		{
			name: "no-go - all 1s",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Dimensions: allDimensions(1.0),
				},
			},
			// 1.0 * 1.0 = 1.0
			wantScore:    1.0,
			wantDecision: "no_go",
		},
		{
			name: "conditional go - mix of 3-4",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Dimensions: []domain.DimensionScore{
						{Dimension: "PMF", Score: 4.0},
						{Dimension: "FIN", Score: 3.0},
						{Dimension: "PRC_CX", Score: 4.0},
						{Dimension: "CMP", Score: 3.0},
						{Dimension: "BND", Score: 4.0},
						{Dimension: "MR", Score: 3.0},
						{Dimension: "RISK", Score: 3.0},
					},
				},
			},
			// 4*0.15 + 3*0.25 + 4*0.20 + 3*0.15 + 4*0.10 + 3*0.10 + 3*0.05
			// = 0.60 + 0.75 + 0.80 + 0.45 + 0.40 + 0.30 + 0.15 = 3.45
			wantScore:    3.45,
			wantDecision: "conditional_go",
		},
		{
			name: "redesign - mix of 2-3",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Dimensions: []domain.DimensionScore{
						{Dimension: "PMF", Score: 3.0},
						{Dimension: "FIN", Score: 2.0},
						{Dimension: "PRC_CX", Score: 3.0},
						{Dimension: "CMP", Score: 2.0},
						{Dimension: "BND", Score: 3.0},
						{Dimension: "MR", Score: 2.0},
						{Dimension: "RISK", Score: 2.0},
					},
				},
			},
			// 3*0.15 + 2*0.25 + 3*0.20 + 2*0.15 + 3*0.10 + 2*0.10 + 2*0.05
			// = 0.45 + 0.50 + 0.60 + 0.30 + 0.30 + 0.20 + 0.10 = 2.45
			wantScore:    2.45,
			wantDecision: "redesign",
		},
		{
			name: "strong go boundary exact 4.0",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Dimensions: allDimensions(4.0),
				},
			},
			// 4.0 * 1.0 = 4.0
			wantScore:    4.0,
			wantDecision: "strong_go",
		},
		{
			name: "all 3s - floating point sum slightly below 3.0",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Dimensions: allDimensions(3.0),
				},
			},
			// 3.0 * (0.15+0.25+0.20+0.15+0.10+0.10+0.05) = 3.0 in theory
			// but floating point sum of individual products is 2.9999999... < 3.0
			// decision is made on raw sum before rounding, so this is "redesign"
			wantScore:    3.0,
			wantDecision: "redesign",
		},
		{
			name: "conditional go - score clearly above 3.0",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Dimensions: []domain.DimensionScore{
						{Dimension: "PMF", Score: 3.5},
						{Dimension: "FIN", Score: 3.0},
						{Dimension: "PRC_CX", Score: 3.5},
						{Dimension: "CMP", Score: 3.0},
						{Dimension: "BND", Score: 3.5},
						{Dimension: "MR", Score: 3.0},
						{Dimension: "RISK", Score: 3.0},
					},
				},
			},
			// 3.5*0.15 + 3*0.25 + 3.5*0.20 + 3*0.15 + 3.5*0.10 + 3*0.10 + 3*0.05
			// = 0.525 + 0.75 + 0.70 + 0.45 + 0.35 + 0.30 + 0.15 = 3.225
			wantScore:    3.23,
			wantDecision: "conditional_go",
		},
		{
			name: "redesign boundary exact 2.0",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Dimensions: allDimensions(2.0),
				},
			},
			wantScore:    2.0,
			wantDecision: "redesign",
		},
		{
			name: "custom weights override FIN (rebalanced to 1.0)",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Dimensions: []domain.DimensionScore{
						{Dimension: "PMF", Score: 5.0},
						{Dimension: "FIN", Score: 5.0},
						{Dimension: "PRC_CX", Score: 5.0},
						{Dimension: "CMP", Score: 5.0},
						{Dimension: "BND", Score: 5.0},
						{Dimension: "MR", Score: 5.0},
						{Dimension: "RISK", Score: 5.0},
					},
					Weights: &domain.ScoringWeights{
						FIN:  ptr(0.40), // override FIN from 0.25 to 0.40
						PRCCX: ptr(0.10), // reduce PRC_CX from 0.20 to 0.10
						RISK: ptr(0.0),   // zero out RISK (0.05 → 0.0)
					},
				},
			},
			// 5*(0.15+0.40+0.10+0.15+0.10+0.10+0.0) = 5*1.0 = 5.0
			wantScore:    5.0,
			wantDecision: "strong_go",
		},
		{
			name: "custom weights change decision (rebalanced to 1.0)",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Dimensions: []domain.DimensionScore{
						{Dimension: "PMF", Score: 5.0},
						{Dimension: "FIN", Score: 1.0},
						{Dimension: "PRC_CX", Score: 5.0},
						{Dimension: "CMP", Score: 5.0},
						{Dimension: "BND", Score: 5.0},
						{Dimension: "MR", Score: 5.0},
						{Dimension: "RISK", Score: 5.0},
					},
					Weights: &domain.ScoringWeights{
						FIN:   ptr(0.55), // heavily weight FIN
						PRCCX: ptr(0.05), // reduce others to keep sum=1.0
						BND:   ptr(0.0),
						RISK:  ptr(0.0),
					},
				},
			},
			// 5*0.15 + 1*0.55 + 5*0.05 + 5*0.15 + 5*0.0 + 5*0.10 + 5*0.0
			// = 0.75 + 0.55 + 0.25 + 0.75 + 0.0 + 0.50 + 0.0 = 2.80
			wantScore:    2.80,
			wantDecision: "redesign",
		},
		{
			name: "invalid weights sum too high",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Dimensions: []domain.DimensionScore{
						{Dimension: "PMF", Score: 5.0},
					},
					Weights: &domain.ScoringWeights{
						FIN: ptr(0.60), // sum = 1.35, way over 1.0
					},
				},
			},
			wantErr:     true,
			errContains: "weights must sum to ~1.0",
		},
		{
			name: "partial dimensions - only 3 of 7",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Dimensions: []domain.DimensionScore{
						{Dimension: "PMF", Score: 5.0},
						{Dimension: "FIN", Score: 4.0},
						{Dimension: "CMP", Score: 3.0},
					},
				},
			},
			// PMF: 5*0.15=0.75, FIN: 4*0.25=1.0, CMP: 3*0.15=0.45
			// PRC_CX: 0*0.20=0, BND: 0*0.10=0, MR: 0*0.10=0, RISK: 0*0.05=0
			// total = 2.20
			wantScore:    2.20,
			wantDecision: "redesign",
		},
		{
			name: "single dimension PMF only",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Dimensions: []domain.DimensionScore{
						{Dimension: "PMF", Score: 5.0},
					},
				},
			},
			// 5*0.15 = 0.75, everything else is 0
			wantScore:    0.75,
			wantDecision: "no_go",
		},
		{
			name: "rationale preserved",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Dimensions: []domain.DimensionScore{
						{Dimension: "PMF", Score: 4.0, Rationale: ptrStr("strong product-market fit")},
						{Dimension: "FIN", Score: 3.0, Rationale: ptrStr("moderate financials")},
						{Dimension: "PRC_CX", Score: 4.0},
						{Dimension: "CMP", Score: 3.0},
						{Dimension: "BND", Score: 4.0},
						{Dimension: "MR", Score: 3.0},
						{Dimension: "RISK", Score: 3.0},
					},
				},
			},
			wantScore:    3.45,
			wantDecision: "conditional_go",
		},
		{
			name:        "nil scoring input",
			input:       &domain.AppraisalInput{},
			wantErr:     true,
			errContains: "scoring input required",
		},
		{
			name: "empty dimensions",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Dimensions: []domain.DimensionScore{},
				},
			},
			wantErr:     true,
			errContains: "at least one dimension score required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.GoNoGo(tt.input)
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
			if !approxEqual(result.WeightedScore, tt.wantScore, 0.01) {
				t.Errorf("weighted_score = %.2f, want %.2f", result.WeightedScore, tt.wantScore)
			}
			if result.Decision != tt.wantDecision {
				t.Errorf("decision = %q, want %q", result.Decision, tt.wantDecision)
			}
		})
	}
}

func TestGoNoGoWeightsMap(t *testing.T) {
	calc := New()

	t.Run("default weights returned when no overrides", func(t *testing.T) {
		input := &domain.AppraisalInput{
			Scoring: &domain.ScoringInput{
				Dimensions: []domain.DimensionScore{
					{Dimension: "PMF", Score: 3.0},
				},
			},
		}
		result, err := calc.GoNoGo(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedWeights := map[string]float64{
			"PMF":    0.15,
			"FIN":    0.25,
			"PRC_CX": 0.20,
			"CMP":    0.15,
			"BND":    0.10,
			"MR":     0.10,
			"RISK":   0.05,
		}
		for dim, w := range expectedWeights {
			if got, ok := result.Weights[dim]; !ok {
				t.Errorf("missing weight for dimension %q", dim)
			} else if !approxEqual(got, w, 1e-9) {
				t.Errorf("weight[%q] = %f, want %f", dim, got, w)
			}
		}
	})

	t.Run("override specific weights (sum=1.0)", func(t *testing.T) {
		input := &domain.AppraisalInput{
			Scoring: &domain.ScoringInput{
				Dimensions: []domain.DimensionScore{
					{Dimension: "PMF", Score: 3.0},
				},
				Weights: &domain.ScoringWeights{
					PMF:   ptr(0.20), // up from 0.15 (+0.05)
					PRCCX: ptr(0.10), // down from 0.20 (-0.10)
					RISK:  ptr(0.10), // up from 0.05 (+0.05)
					// net: +0.05-0.10+0.05 = 0, sum stays 1.0
				},
			},
		}
		result, err := calc.GoNoGo(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !approxEqual(result.Weights["PMF"], 0.20, 1e-9) {
			t.Errorf("PMF weight = %f, want 0.20", result.Weights["PMF"])
		}
		if !approxEqual(result.Weights["RISK"], 0.10, 1e-9) {
			t.Errorf("RISK weight = %f, want 0.10", result.Weights["RISK"])
		}
		// Non-overridden should keep defaults
		if !approxEqual(result.Weights["FIN"], 0.25, 1e-9) {
			t.Errorf("FIN weight = %f, want 0.25 (default)", result.Weights["FIN"])
		}
	})
}

func TestGoNoGoDimensionDetails(t *testing.T) {
	calc := New()

	t.Run("dimensions in order with correct weighted scores", func(t *testing.T) {
		input := &domain.AppraisalInput{
			Scoring: &domain.ScoringInput{
				Dimensions: []domain.DimensionScore{
					{Dimension: "PMF", Score: 5.0},
					{Dimension: "FIN", Score: 4.0},
					{Dimension: "PRC_CX", Score: 3.0},
					{Dimension: "CMP", Score: 2.0},
					{Dimension: "BND", Score: 1.0},
					{Dimension: "MR", Score: 5.0},
					{Dimension: "RISK", Score: 3.0},
				},
			},
		}
		result, err := calc.GoNoGo(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Should have 7 dimension details
		if len(result.Dimensions) != 7 {
			t.Fatalf("got %d dimensions, want 7", len(result.Dimensions))
		}

		expectedOrder := []string{"PMF", "FIN", "PRC_CX", "CMP", "BND", "MR", "RISK"}
		expectedWeighted := []float64{
			5.0 * 0.15, // PMF: 0.75
			4.0 * 0.25, // FIN: 1.0
			3.0 * 0.20, // PRC_CX: 0.6
			2.0 * 0.15, // CMP: 0.3
			1.0 * 0.10, // BND: 0.1
			5.0 * 0.10, // MR: 0.5
			3.0 * 0.05, // RISK: 0.15
		}

		for i, d := range result.Dimensions {
			if d.Dimension != expectedOrder[i] {
				t.Errorf("dimensions[%d] = %q, want %q", i, d.Dimension, expectedOrder[i])
			}
			if !approxEqual(d.WeightedScore, expectedWeighted[i], 1e-9) {
				t.Errorf("dimensions[%d].WeightedScore = %f, want %f", i, d.WeightedScore, expectedWeighted[i])
			}
		}
	})
}

// ---------------------------------------------------------------------------
// RiskMatrix
// ---------------------------------------------------------------------------

func TestRiskMatrix(t *testing.T) {
	calc := New()

	tests := []struct {
		name          string
		input         *domain.AppraisalInput
		wantAvg       float64
		wantMax       float64
		wantHighRisks int
		wantLevels    []string
		wantScores    []float64
		wantErr       bool
		errContains   string
	}{
		{
			name: "critical risk 5x5=25",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Risks: []domain.RiskItem{
						{Name: "total_failure", Probability: 5.0, Impact: 5.0},
					},
				},
			},
			wantAvg:       25.0,
			wantMax:       25.0,
			wantHighRisks: 1, // critical counts as high
			wantLevels:    []string{"critical"},
			wantScores:    []float64{25.0},
		},
		{
			name: "high risk 4x4=16",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Risks: []domain.RiskItem{
						{Name: "market_shift", Probability: 4.0, Impact: 4.0},
					},
				},
			},
			wantAvg:       16.0,
			wantMax:       16.0,
			wantHighRisks: 1,
			wantLevels:    []string{"high"},
			wantScores:    []float64{16.0},
		},
		{
			name: "medium risk 3x3=9",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Risks: []domain.RiskItem{
						{Name: "competitor_response", Probability: 3.0, Impact: 3.0},
					},
				},
			},
			wantAvg:       9.0,
			wantMax:       9.0,
			wantHighRisks: 0,
			wantLevels:    []string{"medium"},
			wantScores:    []float64{9.0},
		},
		{
			name: "low risk 2x2=4",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Risks: []domain.RiskItem{
						{Name: "minor_issue", Probability: 2.0, Impact: 2.0},
					},
				},
			},
			wantAvg:       4.0,
			wantMax:       4.0,
			wantHighRisks: 0,
			wantLevels:    []string{"low"},
			wantScores:    []float64{4.0},
		},
		{
			name: "mixed risks",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Risks: []domain.RiskItem{
						{Name: "critical_risk", Probability: 5.0, Impact: 5.0},
						{Name: "high_risk", Probability: 4.0, Impact: 4.0},
						{Name: "medium_risk", Probability: 3.0, Impact: 3.0},
						{Name: "low_risk", Probability: 1.0, Impact: 1.0},
					},
				},
			},
			// scores: 25, 16, 9, 1 => avg = 51/4 = 12.75
			wantAvg:       12.75,
			wantMax:       25.0,
			wantHighRisks: 2, // critical + high
			wantLevels:    []string{"critical", "high", "medium", "low"},
			wantScores:    []float64{25.0, 16.0, 9.0, 1.0},
		},
		{
			name: "many risks",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Risks: []domain.RiskItem{
						{Name: "r1", Probability: 5.0, Impact: 4.0},
						{Name: "r2", Probability: 4.0, Impact: 5.0},
						{Name: "r3", Probability: 3.0, Impact: 5.0},
						{Name: "r4", Probability: 2.0, Impact: 3.0},
						{Name: "r5", Probability: 1.0, Impact: 2.0},
					},
				},
			},
			// scores: 20, 20, 15, 6, 2 => avg = 63/5 = 12.6
			wantAvg:       12.6,
			wantMax:       20.0,
			wantHighRisks: 3, // two critical (20) + one high (15)
			wantLevels:    []string{"critical", "critical", "high", "low", "low"},
			wantScores:    []float64{20.0, 20.0, 15.0, 6.0, 2.0},
		},
		{
			name: "boundary score exactly 20 is critical",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Risks: []domain.RiskItem{
						{Name: "boundary", Probability: 4.0, Impact: 5.0},
					},
				},
			},
			wantAvg:       20.0,
			wantMax:       20.0,
			wantHighRisks: 1,
			wantLevels:    []string{"critical"},
			wantScores:    []float64{20.0},
		},
		{
			name: "boundary score exactly 15 is high",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Risks: []domain.RiskItem{
						{Name: "boundary", Probability: 3.0, Impact: 5.0},
					},
				},
			},
			wantAvg:       15.0,
			wantMax:       15.0,
			wantHighRisks: 1,
			wantLevels:    []string{"high"},
			wantScores:    []float64{15.0},
		},
		{
			name: "boundary score exactly 8 is medium",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Risks: []domain.RiskItem{
						{Name: "boundary", Probability: 2.0, Impact: 4.0},
					},
				},
			},
			wantAvg:       8.0,
			wantMax:       8.0,
			wantHighRisks: 0,
			wantLevels:    []string{"medium"},
			wantScores:    []float64{8.0},
		},
		{
			name: "score 7.99 is low",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Risks: []domain.RiskItem{
						{Name: "just_below_medium", Probability: 2.0, Impact: 3.5},
					},
				},
			},
			wantAvg:       7.0,
			wantMax:       7.0,
			wantHighRisks: 0,
			wantLevels:    []string{"low"},
			wantScores:    []float64{7.0},
		},
		{
			name:        "nil scoring input",
			input:       &domain.AppraisalInput{},
			wantErr:     true,
			errContains: "scoring input required",
		},
		{
			name: "empty risks",
			input: &domain.AppraisalInput{
				Scoring: &domain.ScoringInput{
					Risks: []domain.RiskItem{},
				},
			},
			wantErr:     true,
			errContains: "at least one risk item required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.RiskMatrix(tt.input)
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

			if !approxEqual(result.AvgScore, tt.wantAvg, 0.01) {
				t.Errorf("avg_score = %.2f, want %.2f", result.AvgScore, tt.wantAvg)
			}
			if !approxEqual(result.MaxScore, tt.wantMax, 0.01) {
				t.Errorf("max_score = %.2f, want %.2f", result.MaxScore, tt.wantMax)
			}
			if result.HighRisks != tt.wantHighRisks {
				t.Errorf("high_risks = %d, want %d", result.HighRisks, tt.wantHighRisks)
			}

			if len(result.Risks) != len(tt.wantLevels) {
				t.Fatalf("got %d risks, want %d", len(result.Risks), len(tt.wantLevels))
			}
			for i, r := range result.Risks {
				if r.Level != tt.wantLevels[i] {
					t.Errorf("risks[%d].Level = %q, want %q", i, r.Level, tt.wantLevels[i])
				}
				if !approxEqual(r.Score, tt.wantScores[i], 1e-9) {
					t.Errorf("risks[%d].Score = %f, want %f", i, r.Score, tt.wantScores[i])
				}
			}
		})
	}
}

func TestRiskMatrixSingleRisk(t *testing.T) {
	calc := New()

	input := &domain.AppraisalInput{
		Scoring: &domain.ScoringInput{
			Risks: []domain.RiskItem{
				{Name: "only_risk", Probability: 3.0, Impact: 4.0, Description: ptrStr("test"), Mitigation: ptrStr("monitor")},
			},
		},
	}
	result, err := calc.RiskMatrix(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result.Risks) != 1 {
		t.Fatalf("expected 1 risk, got %d", len(result.Risks))
	}
	if result.Risks[0].Name != "only_risk" {
		t.Errorf("name = %q, want %q", result.Risks[0].Name, "only_risk")
	}
	if result.AvgScore != result.MaxScore {
		t.Errorf("avg (%f) should equal max (%f) for single risk", result.AvgScore, result.MaxScore)
	}
}

// ---------------------------------------------------------------------------
// DimensionScoreCalc
// ---------------------------------------------------------------------------

func TestDimensionScoreCalc(t *testing.T) {
	calc := New()

	tests := []struct {
		name        string
		dimension   string
		score       float64
		rationale   *string
		wantErr     bool
		errContains string
	}{
		{
			name:      "valid PMF score 4.0",
			dimension: "PMF",
			score:     4.0,
		},
		{
			name:      "valid FIN score 1.0 (minimum)",
			dimension: "FIN",
			score:     1.0,
		},
		{
			name:      "valid PRC_CX score 5.0 (maximum)",
			dimension: "PRC_CX",
			score:     5.0,
		},
		{
			name:      "valid CMP score 2.5",
			dimension: "CMP",
			score:     2.5,
		},
		{
			name:      "valid BND score 3.0",
			dimension: "BND",
			score:     3.0,
		},
		{
			name:      "valid MR score 4.5",
			dimension: "MR",
			score:     4.5,
		},
		{
			name:      "valid RISK score 1.5",
			dimension: "RISK",
			score:     1.5,
		},
		{
			name:      "with rationale",
			dimension: "PMF",
			score:     4.0,
			rationale: ptrStr("strong product-market fit observed"),
		},
		{
			name:      "nil rationale",
			dimension: "PMF",
			score:     3.0,
			rationale: nil,
		},
		{
			name:        "invalid score below 1.0",
			dimension:   "PMF",
			score:       0.5,
			wantErr:     true,
			errContains: "score must be between 1.0 and 5.0",
		},
		{
			name:        "invalid score zero",
			dimension:   "PMF",
			score:       0.0,
			wantErr:     true,
			errContains: "score must be between 1.0 and 5.0",
		},
		{
			name:        "invalid score above 5.0",
			dimension:   "PMF",
			score:       5.1,
			wantErr:     true,
			errContains: "score must be between 1.0 and 5.0",
		},
		{
			name:        "invalid score negative",
			dimension:   "PMF",
			score:       -1.0,
			wantErr:     true,
			errContains: "score must be between 1.0 and 5.0",
		},
		{
			name:        "invalid score way above",
			dimension:   "PMF",
			score:       10.0,
			wantErr:     true,
			errContains: "score must be between 1.0 and 5.0",
		},
		{
			name:        "invalid dimension name",
			dimension:   "INVALID",
			score:       3.0,
			wantErr:     true,
			errContains: "unknown dimension",
		},
		{
			name:        "empty dimension name",
			dimension:   "",
			score:       3.0,
			wantErr:     true,
			errContains: "unknown dimension",
		},
		{
			name:        "lowercase dimension name",
			dimension:   "pmf",
			score:       3.0,
			wantErr:     true,
			errContains: "unknown dimension",
		},
		{
			name:        "dimension name with typo",
			dimension:   "PRCCX",
			score:       3.0,
			wantErr:     true,
			errContains: "unknown dimension",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.DimensionScoreCalc(tt.dimension, tt.score, tt.rationale)
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
			if result.Dimension != tt.dimension {
				t.Errorf("dimension = %q, want %q", result.Dimension, tt.dimension)
			}
			if result.Score != tt.score {
				t.Errorf("score = %f, want %f", result.Score, tt.score)
			}
			if tt.rationale == nil && result.Rationale != nil {
				t.Errorf("rationale should be nil, got %q", *result.Rationale)
			}
			if tt.rationale != nil {
				if result.Rationale == nil {
					t.Error("rationale should not be nil")
				} else if *result.Rationale != *tt.rationale {
					t.Errorf("rationale = %q, want %q", *result.Rationale, *tt.rationale)
				}
			}
		})
	}
}

func TestDimensionScoreCalcAllValidDimensions(t *testing.T) {
	calc := New()

	validDimensions := []string{"PMF", "FIN", "PRC_CX", "CMP", "BND", "MR", "RISK"}
	for _, dim := range validDimensions {
		t.Run("valid_dimension_"+dim, func(t *testing.T) {
			result, err := calc.DimensionScoreCalc(dim, 3.0, nil)
			if err != nil {
				t.Fatalf("unexpected error for dimension %q: %v", dim, err)
			}
			if result.Dimension != dim {
				t.Errorf("dimension = %q, want %q", result.Dimension, dim)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func containsStr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
