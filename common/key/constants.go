package key

var NetworkMapCovalent = map[string]string{
	"ETH":   "1",
	"BSC":   "56",
	"MATIC": "137",
	"GLMR":  "1284",
	"MOVR":  "1285",
	"ARB":   "42161",
	"FTM":   "250",
}

var NetworkAddressMap = map[string]string{
	"ETH":   "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
	"BSC":   "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",
	"MATIC": "0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270",
	"GLMR":  "0xacc15dc74880c9944775448304b263d191c6077f",
	"MOVR":  "0x98878b06940ae243284ca214f92bb71a2b032b8a",
	"ARB":   "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
	"FTM":   "0x21be370d5312f44cb42ce377bc9b8a0cef1a4c83",
}

var BaseTokensMap = map[string]string{
	"ETH": `{
		"contract_decimals": 18,
		"contract_name":     "Ethereum",
		"contract_ticker_symbol":   "ETH",
		"logo_url":  "https://assets.coingecko.com/coins/images/279/large/ethereum.png?1595348880"
	}`,
	"BSC": `{
		"contract_decimals": 18,
		"contract_name":     "Binance BNB",
		"contract_ticker_symbol":   "BNB",
		"logo_url":  "https://assets.coingecko.com/coins/images/825/large/bnb-icon2_2x.png?1644979850"
	}`,
	"MATIC": `{
		"contract_decimals": 18,
		"contract_name":     "Polygon MATIC",
		"contract_ticker_symbol":   "MATIC",
		"logo_url":  "https://assets.coingecko.com/coins/images/4713/large/matic-token-icon.png?1624446912"
	}`,
	"GLMR": `{
		"contract_decimals": 18,
		"contract_name":     "Moonbeam GLMR",
		"contract_ticker_symbol":   "GLMR",
		"logo_url":  "https://assets.coingecko.com/coins/images/22459/large/glmr.png?1641880985"
	}`,
	"MOVR": `{
		"contract_decimals": 18,
		"contract_name":     "Moonriver MOVR",
		"contract_ticker_symbol":   "MOVR",
		"logo_url":  "https://assets.coingecko.com/coins/images/17984/large/9285.png?1630028620"
	}`,
	"ARB": `{
		"contract_decimals": 18,
		"contract_name":     "Arbitrum Ethereum",
		"contract_ticker_symbol":   "ETH",
		"logo_url":  "https://assets.coingecko.com/coins/images/279/large/ethereum.png?1595348880"
	}`,
	"FTM": `{
		"contract_decimals": 18,
		"contract_name":     "Fantom",
		"contract_ticker_symbol":   "FTM",
		"logo_url":  "https://assets.coingecko.com/coins/images/4001/large/Fantom.png?1558015016"
	}`,
}
