package main

func mois() {
	mois := map[string]int{

		"janvier":   31,
		"fevrier":   28,
		"mars":      31,
		"avril":     30,
		"mai":       31,
		"juin":      30,
		"juillet":   31,
		"aout":      31,
		"septembre": 30,
		"octobre":   31,
		"novembre":  30,
		"decembre":  31,
	}
	for key, value := range mois {
		println(key, value)
	}
}
