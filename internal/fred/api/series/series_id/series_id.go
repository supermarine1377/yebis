// package_id provides series_id defined by FRED API.
package series_id

const (
	// FEDFUNDS RATE
	// https://fred.stlouisfed.org/graph/?graph_id=1123034
	FEDFUNDS = "FEDFUNDS"
	// 10-Year Treasury Constant Maturity Minus 2-Year Treasury Constant Maturity (T10Y2Y)
	// https://fred.stlouisfed.org/graph/?graph_id=1123037
	T10Y2Y = "T10Y2Y"
	// Market Yield on U.S. Treasury Securities at 10-Year Constant Maturity
	// FED call it DGS10, but here we call it US10Y because it is more common name.
	// https://fred.stlouisfed.org/graph/?graph_id=1123035
	US10Y = "DGS10"
	// Moody's Seasoned Baa Corporate Bond Yield Relative to Yield on 10-Year Treasury Constant Maturity
	// https://fred.stlouisfed.org/graph/?graph_id=1123036
	BAA10Y = "BAA10Y"
	// Nominal Broad U.S. Dollar Index
	// https://fred.stlouisfed.org/series/DTWEXBGS
	USDINDEX = "DTWEXBGS"
)
