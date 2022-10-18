package main

func listItems(e *Envelope) {
	titleId, err := e.getKey("TitleId")
	if err != nil {
		e.Error(9, "Unable to obtain title.", err)
	}

	attrs, err := e.getKeys("AttributeFilters")
	if err != nil {
		e.Error(5, "AttributeFilters key did not exist!", err)
	}

	var licenceStr string
	for _, attr := range attrs {
		name, value := parseNameValue(attr.InnerText())
		if name == "TitleKind" {
			licenceStr = value
		}
	}

	// Now validate
	licenceKind, err := GetLicenceKind(licenceStr)
	if err != nil {
		e.Error(5, "Invalid TitleKind was passed by SOAP", err)
	}

	// TODO(SketchMaster2001): Query database for items
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
				// Not sure about WSC, but must match the price for the title you are purchasing in Wii no Ma.
				Amount:   0,
				Currency: "POINTS",
			},
			Limits:      LimitStruct(PR),
			LicenseKind: *licenceKind,
		},
	})
}
