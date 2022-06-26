package model

type Users struct {
	ID       int    `db:"id"; json:"id"`
	Username string `db:"username"; json:"username"`
	Password string `db:"password";json:"password"`
	Email    string `db:"email";json:"email"`
}

type Books struct {
	IDbuku    int    `db:"idbuku";json:"idbuku"`
	ISBN      string `db:"ISBN";json:"ISBN"`
	Judul     string `db:"judul";json:"judul"`
	Penerbit  string `db:"penerbit";json:"penerbit"`
	Pengarang string `db:"pengarang";json:"pengarang"`
	Tahun     string `db:"tahun";json:"tahun"`
	Gambar    string `db:"gambar";json:"gambar"`
	Deskripsi string `db:"deskripsi";json:"deskripsi"`
}

type Admin struct {
	ID       int    `db:"idadmin";json:"idadmin"`
	Username string `db:"username";json:"username"`
	Password string `db:"password";json:"password"`
	Role     string `db:"role";json:"role"`
}

type HistoryBuku struct {
	ID        int    `json:id`
	ID_user   string `idUser`
	ID_books  string `idBooks`
	Last_page string `last_page`
	Favorite  bool
}

type SearchBukuNavBar struct {
	SearchBuku string `json:search`
}

type TempGoogleBooks struct {
	Kind       string `json:"kind"`
	ID         string `json:"id"`
	Etag       string `json:"etag"`
	SelfLink   string `json:"selfLink"`
	VolumeInfo struct {
		Title               string   `json:"title"`
		Subtitle            string   `json:"subtitle"`
		Authors             []string `json:"authors"`
		Publisher           string   `json:"publisher"`
		PublishedDate       string   `json:"publishedDate"`
		Description         string   `json:"description"`
		IndustryIdentifiers []struct {
			Type       string `json:"type"`
			Identifier string `json:"identifier"`
		} `json:"industryIdentifiers"`
		ReadingModes struct {
			Text  bool `json:"text"`
			Image bool `json:"image"`
		} `json:"readingModes"`
		PageCount        int `json:"pageCount"`
		PrintedPageCount int `json:"printedPageCount"`
		Dimensions       struct {
			Height string `json:"height"`
		} `json:"dimensions"`
		PrintType           string   `json:"printType"`
		Categories          []string `json:"categories"`
		AverageRating       float64  `json:"averageRating"`
		RatingsCount        int      `json:"ratingsCount"`
		MaturityRating      string   `json:"maturityRating"`
		AllowAnonLogging    bool     `json:"allowAnonLogging"`
		ContentVersion      string   `json:"contentVersion"`
		PanelizationSummary struct {
			ContainsEpubBubbles  bool `json:"containsEpubBubbles"`
			ContainsImageBubbles bool `json:"containsImageBubbles"`
		} `json:"panelizationSummary"`
		ImageLinks struct {
			SmallThumbnail string `json:"smallThumbnail"`
			Thumbnail      string `json:"thumbnail"`
			Small          string `json:"small"`
			Medium         string `json:"medium"`
			Large          string `json:"large"`
		} `json:"imageLinks"`
		Language            string `json:"language"`
		PreviewLink         string `json:"previewLink"`
		InfoLink            string `json:"infoLink"`
		CanonicalVolumeLink string `json:"canonicalVolumeLink"`
	} `json:"volumeInfo"`
	LayerInfo struct {
		Layers []struct {
			LayerID                  string `json:"layerId"`
			VolumeAnnotationsVersion string `json:"volumeAnnotationsVersion"`
		} `json:"layers"`
	} `json:"layerInfo"`
	SaleInfo struct {
		Country     string `json:"country"`
		Saleability string `json:"saleability"`
		IsEbook     bool   `json:"isEbook"`
		ListPrice   struct {
			Amount       int    `json:"amount"`
			CurrencyCode string `json:"currencyCode"`
		} `json:"listPrice"`
		RetailPrice struct {
			Amount       int    `json:"amount"`
			CurrencyCode string `json:"currencyCode"`
		} `json:"retailPrice"`
		BuyLink string `json:"buyLink"`
		Offers  []struct {
			FinskyOfferType int `json:"finskyOfferType"`
			ListPrice       struct {
				AmountInMicros int64  `json:"amountInMicros"`
				CurrencyCode   string `json:"currencyCode"`
			} `json:"listPrice"`
			RetailPrice struct {
				AmountInMicros int64  `json:"amountInMicros"`
				CurrencyCode   string `json:"currencyCode"`
			} `json:"retailPrice"`
		} `json:"offers"`
	} `json:"saleInfo"`
	AccessInfo struct {
		Country                string `json:"country"`
		Viewability            string `json:"viewability"`
		Embeddable             bool   `json:"embeddable"`
		PublicDomain           bool   `json:"publicDomain"`
		TextToSpeechPermission string `json:"textToSpeechPermission"`
		Epub                   struct {
			IsAvailable  bool   `json:"isAvailable"`
			AcsTokenLink string `json:"acsTokenLink"`
		} `json:"epub"`
		Pdf struct {
			IsAvailable bool `json:"isAvailable"`
		} `json:"pdf"`
		WebReaderLink       string `json:"webReaderLink"`
		AccessViewStatus    string `json:"accessViewStatus"`
		QuoteSharingAllowed bool   `json:"quoteSharingAllowed"`
	} `json:"accessInfo"`
}
