package repository

type SalesPerformance struct {
	Object string `json:"object"`

	Impressions int `json:"impressions"`

	Clicks int `json:"clicks"`

	Orders int `json:"orders"`

	Revenue float64 `json:"revenue"`

	AdCost float64 `json:"ad_cost"`

	CTR float64 `json:"ctr"`

	ConversionRate float64 `json:"conversion_rate"`

	ACOS float64 `json:"acos"`
}

type PerformanceEngine struct{}

func NewPerformanceEngine() *PerformanceEngine {

	return &PerformanceEngine{}
}

func (p *PerformanceEngine) Analyze(
	object string,
	impressions int,
	clicks int,
	orders int,
	revenue float64,
	adCost float64,
) SalesPerformance {

	ctr := float64(0)

	if impressions > 0 {

		ctr =
			float64(clicks) /
				float64(impressions) *
				100
	}

	conversion := float64(0)

	if clicks > 0 {

		conversion =
			float64(orders) /
				float64(clicks) *
				100
	}

	acos := float64(0)

	if revenue > 0 {

		acos =
			adCost /
				revenue *
				100
	}

	return SalesPerformance{

		Object: object,

		Impressions: impressions,

		Clicks: clicks,

		Orders: orders,

		Revenue: revenue,

		AdCost: adCost,

		CTR: ctr,

		ConversionRate: conversion,

		ACOS: acos,
	}
}
