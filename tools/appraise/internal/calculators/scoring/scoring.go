// Package scoring implements Go/No-Go decision scoring.
//
// Functions:
//   GoNoGo         - Weighted scoring across 7 dimensions with configurable weights
//   RiskMatrix     - Risk probability * impact scoring
//   DimensionScore - Score a single dimension (1-5)
package scoring

import (
	"fmt"
	"math"

	"github.com/ivalx1s/skill-product-appraisal/tools/appraise/internal/domain"
)

// Default weights for Go/No-Go scoring.
const (
	DefaultWeightPMF   = 0.15 // Product-Market Fit
	DefaultWeightFIN   = 0.25 // Financial Viability
	DefaultWeightPRCCX = 0.20 // Pricing + Customer Experience
	DefaultWeightCMP   = 0.15 // Competitive Position
	DefaultWeightBND   = 0.10 // Bundle Composition
	DefaultWeightMR    = 0.10 // Market Reach
	DefaultWeightRISK  = 0.05 // Risk Profile
)

// Calculator implements all scoring module functions.
type Calculator struct{}

// New creates a new scoring calculator.
func New() *Calculator {
	return &Calculator{}
}

// GoNoGo computes weighted total score across all 7 dimensions and returns a decision.
//
// Score >= 4.0: Strong Go
// Score 3.0-3.9: Conditional Go
// Score 2.0-2.9: Redesign
// Score < 2.0: No-Go
func (c *Calculator) GoNoGo(input *domain.AppraisalInput) (*domain.GoNoGoResult, error) {
	if input.Scoring == nil {
		return nil, fmt.Errorf("scoring input required")
	}
	if len(input.Scoring.Dimensions) == 0 {
		return nil, fmt.Errorf("at least one dimension score required")
	}

	// Resolve weights (user overrides or defaults)
	weights := resolveWeights(input.Scoring.Weights)

	// Validate weight sum
	weightSum := 0.0
	for _, w := range weights {
		weightSum += w
	}
	if math.Abs(weightSum-1.0) > 0.01 {
		return nil, fmt.Errorf("weights must sum to ~1.0, got %.4f", weightSum)
	}

	// Build dimension lookup
	dimScores := make(map[string]domain.DimensionScore)
	for _, ds := range input.Scoring.Dimensions {
		dimScores[ds.Dimension] = ds
	}

	// Known dimensions in order
	dimOrder := []string{"PMF", "FIN", "PRC_CX", "CMP", "BND", "MR", "RISK"}

	totalWeighted := 0.0
	var details []domain.DimensionDetail
	var missing []string

	for _, dim := range dimOrder {
		w, ok := weights[dim]
		if !ok {
			continue
		}

		ds, hasScore := dimScores[dim]
		score := 0.0
		var rationale *string
		if hasScore {
			score = ds.Score
			rationale = ds.Rationale
		} else {
			missing = append(missing, dim)
		}

		weighted := score * w
		totalWeighted += weighted

		details = append(details, domain.DimensionDetail{
			Dimension:     dim,
			Score:         score,
			Weight:        w,
			WeightedScore: weighted,
			Rationale:     rationale,
		})
	}

	// Decision
	var decision string
	switch {
	case totalWeighted >= 4.0:
		decision = "strong_go"
	case totalWeighted >= 3.0:
		decision = "conditional_go"
	case totalWeighted >= 2.0:
		decision = "redesign"
	default:
		decision = "no_go"
	}

	result := &domain.GoNoGoResult{
		WeightedScore: math.Round(totalWeighted*100) / 100,
		Decision:      decision,
		Dimensions:    details,
		Weights:       weights,
	}
	if len(missing) > 0 {
		warning := fmt.Sprintf("missing dimensions scored as 0: %v", missing)
		result.Warning = &warning
	}

	return result, nil
}

// RiskMatrix scores each risk as probability * impact and classifies severity.
//
// Score thresholds:
//   >= 20: critical
//   >= 15: high
//   >= 8:  medium
//   < 8:   low
func (c *Calculator) RiskMatrix(input *domain.AppraisalInput) (*domain.RiskMatrixResult, error) {
	if input.Scoring == nil {
		return nil, fmt.Errorf("scoring input required")
	}
	if len(input.Scoring.Risks) == 0 {
		return nil, fmt.Errorf("at least one risk item required")
	}

	result := &domain.RiskMatrixResult{}
	totalScore := 0.0
	maxScore := 0.0

	for _, risk := range input.Scoring.Risks {
		score := risk.Probability * risk.Impact

		var level string
		switch {
		case score >= 20:
			level = "critical"
			result.HighRisks++
		case score >= 15:
			level = "high"
			result.HighRisks++
		case score >= 8:
			level = "medium"
		default:
			level = "low"
		}

		totalScore += score
		if score > maxScore {
			maxScore = score
		}

		result.Risks = append(result.Risks, domain.ScoredRisk{
			Name:        risk.Name,
			Probability: risk.Probability,
			Impact:      risk.Impact,
			Score:       score,
			Level:       level,
		})
	}

	result.AvgScore = totalScore / float64(len(input.Scoring.Risks))
	result.MaxScore = maxScore

	return result, nil
}

// DimensionScoreCalc validates and returns a single dimension score.
// Score must be 1.0-5.0.
func (c *Calculator) DimensionScoreCalc(dimension string, score float64, rationale *string) (*domain.DimensionScore, error) {
	if score < 1.0 || score > 5.0 {
		return nil, fmt.Errorf("score must be between 1.0 and 5.0, got %.1f", score)
	}

	validDimensions := map[string]bool{
		"PMF": true, "FIN": true, "PRC_CX": true,
		"CMP": true, "BND": true, "MR": true, "RISK": true,
	}

	if !validDimensions[dimension] {
		return nil, fmt.Errorf("unknown dimension %q; valid: PMF, FIN, PRC_CX, CMP, BND, MR, RISK", dimension)
	}

	return &domain.DimensionScore{
		Dimension: dimension,
		Score:     score,
		Rationale: rationale,
	}, nil
}

// resolveWeights returns a map of dimension -> weight, using defaults where not overridden.
func resolveWeights(overrides *domain.ScoringWeights) map[string]float64 {
	w := map[string]float64{
		"PMF":    DefaultWeightPMF,
		"FIN":    DefaultWeightFIN,
		"PRC_CX": DefaultWeightPRCCX,
		"CMP":    DefaultWeightCMP,
		"BND":    DefaultWeightBND,
		"MR":     DefaultWeightMR,
		"RISK":   DefaultWeightRISK,
	}

	if overrides == nil {
		return w
	}

	if overrides.PMF != nil {
		w["PMF"] = *overrides.PMF
	}
	if overrides.FIN != nil {
		w["FIN"] = *overrides.FIN
	}
	if overrides.PRCCX != nil {
		w["PRC_CX"] = *overrides.PRCCX
	}
	if overrides.CMP != nil {
		w["CMP"] = *overrides.CMP
	}
	if overrides.BND != nil {
		w["BND"] = *overrides.BND
	}
	if overrides.MR != nil {
		w["MR"] = *overrides.MR
	}
	if overrides.RISK != nil {
		w["RISK"] = *overrides.RISK
	}

	return w
}
