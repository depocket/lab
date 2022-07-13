package curve

type VestingContract struct {
	Name    string
	Address string
}

const (
	CRV_TOKEN_ADDRESS = "0xd533a949740bb3306d119cc777fa900ba034cd52"
)

var (
	VESTING_CONTRACTS = map[string]VestingContract{
		"founder": {
			Address: "0xd2D43555134dC575BF7279F4bA18809645dB0F1D",
			Name:    "Founder",
		},
		"investor": {
			Address: "0x2A7d59E327759acd5d11A8fb652Bf4072d28AC04",
			Name:    "Investor",
		},
		"advisor": {
			Address: "0xf7dBC322d72C1788a1E37eEE738e2eA9C7Fa875e",
			Name:    "Advisor",
		},
		"lp": {
			Address: "0x575ccd8e2d300e2377b43478339e364000318e2c",
			Name:    "LP",
		},
		"employee": {
			Address: "0x679FCB9b33Fc4AE10Ff4f96caeF49c1ae3F8fA67",
			Name:    "Employee",
		},
		"factory": {
			Address: "0x81930d767a75269dc0e9b83938884e342c1fa5f6",
			Name:    "Advisor",
		},
		"factory2": {
			Address: "0x629347824016530fcd9a1990a30658ed9a04c834",
			Name:    "Advisor",
		},
	}
)
