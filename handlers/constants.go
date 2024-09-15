package handlers

type GetTransactionSolanaResponse struct {
	TransactionHash string `json:"transactionHash"`
	BlockNumber     int    `json:"blockNumber"`
	Accounts        []struct {
		Account struct {
			Name    string `json:"name"`
			Address string `json:"address"`
		} `json:"account"`
		Writable bool `json:"writable"`
		Signer   bool `json:"signer"`
		Program  bool `json:"program"`
	} `json:"accounts"`
	Header struct {
		NumReadonlySignedAccounts   int `json:"numReadonlySignedAccounts"`
		NumReadonlyUnsignedAccounts int `json:"numReadonlyUnsignedAccounts"`
		NumRequiredSignatures       int `json:"numRequiredSignatures"`
	} `json:"header"`
	Instructions []struct {
		Raw struct {
			Accounts       []any  `json:"accounts"`
			Data           string `json:"data"`
			ProgramIDIndex int    `json:"programIdIndex"`
			StackHeight    any    `json:"stackHeight"`
		} `json:"raw"`
		ProgramID struct {
			Address string `json:"address"`
		} `json:"programId"`
		Parsed struct {
			Transfer struct {
				Lamports int `json:"lamports"`
				Account  struct {
					Address string `json:"address"`
				} `json:"account"`
				Recipient struct {
					Address string `json:"address"`
				} `json:"recipient"`
				Signers []struct {
					Address string `json:"address"`
				} `json:"signers"`
				Writable []struct {
					Address string `json:"address"`
				} `json:"writable"`
			} `json:"Transfer"`
		} `json:"parsed"`
		ProgramID0 struct {
			Name    string `json:"name"`
			Address string `json:"address"`
		} `json:"programId"`
	} `json:"instructions"`
	RecentBlockhash string   `json:"recentBlockhash"`
	Signatures      []string `json:"signatures"`
	Meta            struct {
		ComputeUnitsConsumed int `json:"computeUnitsConsumed"`
		Err                  any `json:"err"`
		Fee                  int `json:"fee"`
		LoadedAddresses      struct {
			Readonly []any `json:"readonly"`
			Writable []any `json:"writable"`
		} `json:"loadedAddresses"`
		LogMessages       []string `json:"logMessages"`
		PostBalances      []any    `json:"postBalances"`
		PostTokenBalances []any    `json:"postTokenBalances"`
		PreBalances       []any    `json:"preBalances"`
		PreTokenBalances  []any    `json:"preTokenBalances"`
		Rewards           []any    `json:"rewards"`
		Status            struct {
			Ok any `json:"Ok"`
		} `json:"status"`
	} `json:"meta"`
	Valid     bool `json:"valid"`
	Blocktime struct {
		Absolute int `json:"absolute"`
		Relative int `json:"relative"`
	} `json:"blocktime"`
	MostImportantInstruction struct {
		Name   string  `json:"name"`
		Weight float64 `json:"weight"`
		Index  int     `json:"index"`
	} `json:"mostImportantInstruction"`
}

type GetTransactionSPLResponse struct {
	TransactionHash string `json:"transactionHash"`
	BlockNumber     int    `json:"blockNumber"`
	Accounts        []struct {
		Writable bool `json:"writable"`
		Signer   bool `json:"signer"`
		Program  bool `json:"program"`
		Account  struct {
			Name   string `json:"name"`
			Ticker string `json:"ticker"`
			CmcID  string `json:"cmcId"`
			Logo   string `json:"logo"`
			Meta   struct {
				URL string `json:"url"`
			} `json:"meta"`
			Address string `json:"address"`
		} `json:"account"`
	} `json:"accounts"`
	Header struct {
		NumReadonlySignedAccounts   int `json:"numReadonlySignedAccounts"`
		NumReadonlyUnsignedAccounts int `json:"numReadonlyUnsignedAccounts"`
		NumRequiredSignatures       int `json:"numRequiredSignatures"`
	} `json:"header"`
	Instructions []struct {
		Raw struct {
			Accounts       []any  `json:"accounts"`
			Data           string `json:"data"`
			ProgramIDIndex int    `json:"programIdIndex"`
			StackHeight    any    `json:"stackHeight"`
		} `json:"raw"`
		ProgramID struct {
			Address string `json:"address"`
		} `json:"programId"`
		Parsed struct {
			TransferChecked struct {
				Amount   int `json:"amount"`
				Decimals int `json:"decimals"`
				Source   struct {
					Address string `json:"address"`
				} `json:"source"`
				Mint struct {
					Name   string `json:"name"`
					Ticker string `json:"ticker"`
					CmcID  string `json:"cmcId"`
					Logo   string `json:"logo"`
					Meta   struct {
						URL string `json:"url"`
					} `json:"meta"`
					Address string `json:"address"`
				} `json:"mint"`
				Destination struct {
					Address string `json:"address"`
				} `json:"destination"`
				Owner struct {
					Address string `json:"address"`
				} `json:"owner"`
				Signers []struct {
					Address string `json:"address"`
				} `json:"signers"`
				Writable []struct {
					Address string `json:"address"`
				} `json:"writable"`
			} `json:"TransferChecked"`
		} `json:"parsed"`
		ProgramID0 struct {
			Name    string `json:"name"`
			Address string `json:"address"`
		} `json:"programId"`
	} `json:"instructions"`
	RecentBlockhash string   `json:"recentBlockhash"`
	Signatures      []string `json:"signatures"`
	Meta            struct {
		ComputeUnitsConsumed int `json:"computeUnitsConsumed"`
		Err                  any `json:"err"`
		Fee                  int `json:"fee"`
		LoadedAddresses      struct {
			Readonly []any `json:"readonly"`
			Writable []any `json:"writable"`
		} `json:"loadedAddresses"`
		LogMessages       []string `json:"logMessages"`
		PostBalances      []any    `json:"postBalances"`
		PostTokenBalances []struct {
			AccountIndex int `json:"accountIndex"`
			Mint         struct {
				Name   string `json:"name"`
				Ticker string `json:"ticker"`
				CmcID  string `json:"cmcId"`
				Logo   string `json:"logo"`
				Meta   struct {
					URL string `json:"url"`
				} `json:"meta"`
				Address string `json:"address"`
			} `json:"mint"`
			Owner struct {
				Address string `json:"address"`
			} `json:"owner"`
			ProgramID struct {
				Name    string `json:"name"`
				Address string `json:"address"`
			} `json:"programId"`
			UITokenAmount struct {
				Amount         string `json:"amount"`
				Decimals       int    `json:"decimals"`
				UIAmount       any    `json:"uiAmount"`
				UIAmountString string `json:"uiAmountString"`
			} `json:"uiTokenAmount"`
		} `json:"postTokenBalances"`
		PreBalances      []any `json:"preBalances"`
		PreTokenBalances []struct {
			AccountIndex int `json:"accountIndex"`
			Mint         struct {
				Name   string `json:"name"`
				Ticker string `json:"ticker"`
				CmcID  string `json:"cmcId"`
				Logo   string `json:"logo"`
				Meta   struct {
					URL string `json:"url"`
				} `json:"meta"`
				Address string `json:"address"`
			} `json:"mint"`
			Owner struct {
				Address string `json:"address"`
			} `json:"owner"`
			ProgramID struct {
				Name    string `json:"name"`
				Address string `json:"address"`
			} `json:"programId"`
			UITokenAmount struct {
				Amount         string `json:"amount"`
				Decimals       int    `json:"decimals"`
				UIAmount       any    `json:"uiAmount"`
				UIAmountString string `json:"uiAmountString"`
			} `json:"uiTokenAmount"`
		} `json:"preTokenBalances"`
		Rewards []any `json:"rewards"`
		Status  struct {
			Ok any `json:"Ok"`
		} `json:"status"`
	} `json:"meta"`
	Valid     bool `json:"valid"`
	Blocktime struct {
		Absolute int `json:"absolute"`
		Relative int `json:"relative"`
	} `json:"blocktime"`
	MostImportantInstruction struct {
		Name   string  `json:"name"`
		Weight float64 `json:"weight"`
		Index  int     `json:"index"`
	} `json:"mostImportantInstruction"`
	Smart []struct {
		Type  string `json:"type"`
		Value struct {
			Address string `json:"address"`
		} `json:"value"`
		Value0 struct {
			Amount int `json:"amount"`
			Mint   struct {
				Name   string `json:"name"`
				Ticker string `json:"ticker"`
				CmcID  string `json:"cmcId"`
				Logo   string `json:"logo"`
				Meta   struct {
					URL string `json:"url"`
				} `json:"meta"`
				Address string `json:"address"`
			} `json:"mint"`
			Decimals int `json:"decimals"`
		} `json:"value"`
	} `json:"smart"`
}

type GetBlockNumberResponse []struct {
	Blocknumber       int    `json:"blocknumber"`
	Blockhash         string `json:"blockhash"`
	Previousblockhash string `json:"previousblockhash"`
	Parentslot        int    `json:"parentslot"`
	Blocktime         struct {
		Absolute struct {
			Absolute int `json:"absolute"`
			Relative int `json:"relative"`
		} `json:"absolute"`
		Relative string `json:"relative"`
	} `json:"blocktime"`
	Metrics struct {
		Txcount           int   `json:"txcount"`
		Failedtxs         int   `json:"failedtxs"`
		Totalfees         int   `json:"totalfees"`
		Instructions      int   `json:"instructions"`
		Sucessfultxs      int   `json:"sucessfultxs"`
		Totalvaluemoved   int64 `json:"totalvaluemoved"`
		Innerinstructions int   `json:"innerinstructions"`
	} `json:"metrics"`
	Programstats []struct {
		Count     int `json:"count"`
		ProgramID struct {
			Name    string `json:"name"`
			Address struct {
				Name    string `json:"name"`
				Address string `json:"address"`
			} `json:"address"`
		} `json:"programId"`
		Instructions struct {
			Vote     int `json:"Vote"`
			Withdraw int `json:"Withdraw"`
		} `json:"instructions,omitempty"`
		Instructions0 struct {
			CreateAssociatedTokenAccount int `json:"CreateAssociatedTokenAccount"`
		} `json:"instructions,omitempty"`
	} `json:"programstats"`
	Rewards []struct {
		Pubkey      string `json:"pubkey"`
		Lamports    int    `json:"lamports"`
		Commission  any    `json:"commission"`
		RewardType  string `json:"rewardType"`
		PostBalance int64  `json:"postBalance"`
	} `json:"rewards"`
	Proposer     string `json:"proposer"`
	Ondemand     bool   `json:"ondemand"`
	ProposerData struct {
		NodePubkey string `json:"nodePubkey"`
	} `json:"proposerData"`
}
