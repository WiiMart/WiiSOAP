package main

func listItems(e *Envelope) {
	titleId, err := getKey(e.doc, "TitleId")
	if err != nil {
		e.Error(9, "Unable to obtain title.", err)
	}

	e.AddKVNode("ListResultTotalSize", "1")
	e.AddCustomType(Items{
		TitleId: titleId,
		Contents: ContentsMetadata{
			TitleIncluded: false,
			ContentIndex:  0,
		},
		Attributes: []Attributes{
			{
				Name:  "TitleVersion",
				Value: "0",
			},
			{
				Name:  "Prices",
				Value: "1",
			},
		},
		Ratings: Ratings{
			Name:   "E",
			Rating: 1,
			Age:    9,
		},
		Prices: Prices{
			ItemId: 0,
			Price: Price{
				Amount:   100,
				Currency: "POINTS",
			},
			// Literally every Limit except for PR works
			Limits:      LimitStruct(TR),
			LicenseKind: "RENTAL",
		},
	})
}
