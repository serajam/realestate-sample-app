package e2e

type PropertyListRequest struct {
	IDs []int `json:"ids" validate:"required,max=10,dive,gte=0,lte=100000000"`
}

type PropertySearchRequest struct {
	City        string `json:"city" validate:"omitempty,max=255,required_without=Polygon" json:"city,omitempty"`
	CountryCode string `json:"countryCode" validate:"omitempty,max=2,required_without_all=Polygon" json:"countryCode,omitempty"`

	Polygon *Polygon `json:"polygon" validate:"required_without=City"`

	Page int `json:"page" validate:"number,max=100000" json:"page,omitempty"`
	Size int `json:"size" validate:"number,max=1000" json:"size,omitempty"`
	Sort int `json:"sort" validate:"number,min=1,max=15" json:"size,omitempty"`

	PriceFrom     float32 `json:"priceFrom" validate:"omitempty,number"`
	PriceTo       float32 `json:"priceTo" validate:"omitempty,number"`
	HomeSizeFrom  float32 `json:"homeSizeFrom" validate:"omitempty,number"`
	HomeSizeTo    float32 `json:"homeSizeTo" validate:"omitempty,number"`
	LotSizeFrom   float32 `json:"lotSizeFrom" validate:"omitempty,number"`
	LotSizeTo     float32 `json:"lotSizeTo" validate:"omitempty,number"`
	YearBuildFrom uint16  `json:"yearBuildFrom" validate:"omitempty,number"`
	YearBuildTo   uint16  `json:"yearBuildTo" validate:"omitempty,number"`

	Bathroom      uint8 `json:"bathroom" validate:"omitempty,number"`
	BathroomExact *bool `json:"bathroomExact" validate:"omitempty,boolean"`

	Bedroom      uint8 `json:"bedroom" validate:"omitempty,number"`
	BedroomExact *bool `json:"bedroomExact" validate:"omitempty,boolean"`

	Condition    []int `json:"condition" validate:"omitempty,dive,number"`
	HomeType     []int `json:"homeType" validate:"omitempty,dive,number"`
	PropertyType uint8 `json:"propertyType" validate:"omitempty,number"`

	HasAC         *bool  `json:"hasAC" validate:"omitempty,boolean"`
	HasGarage     *bool  `json:"hasGarage" validate:"omitempty,boolean"`
	ParkingNumber *uint8 `json:"parkingNumber" validate:"omitempty,number"`
}

type Polygon struct {
	TopLat     string `json:"topLat" validate:"required_without=City,omitempty,latitude"`
	TopLong    string `json:"topLong" validate:"required_without=City,omitempty,latitude"`
	BottomLat  string `json:"botLat" validate:"required_without=City,omitempty,latitude"`
	BottomLong string `json:"botLong" validate:"required_without=City,omitempty,latitude"`
}

type Location struct {
	Lat  float64 `json:"lat" validate:"required,latitude"`
	Long float64 `json:"long" validate:"required,longitude"`
}

type PropertyResponse struct {
	ID         int `json:"ID"`
	ActualDays int `json:"actualDays"`

	Location      Location `json:"location"`
	Price         float32  `json:"price"`
	PriceCurrency string   `json:"priceCurrency"`

	FullAddress string  `json:"fullAddress"`
	Address     Address `json:"address"`

	HomeSize     float32 `json:"homeSize"`
	LotSize      float32 `json:"lotSize"`
	YearBuild    uint16  `json:"yearBuild"`
	Bedroom      uint8   `json:"bedroom"`
	Bathroom     uint8   `json:"bathroom"`
	Floor        uint8   `json:"floor"`
	TotalFloors  uint8   `json:"totalFloors"`
	PropertyType uint8   `json:"propertyType"`
	HomeType     uint8   `json:"homeType"`
	Condition    uint8   `json:"condition"`

	BrokerName  string `json:"brokerName"`
	Description string `json:"description"`

	IsActiveStatus *bool `json:"isActiveStatus"`
	IsFavorite     *bool `json:"isFavorite"`
	HasImages      *bool `json:"hasImages"`
	HasGarage      *bool `json:"hasGarage"`
	HasVideo       *bool `json:"hasVideo"`
	Has3DTour      *bool `json:"has3DTour"`
	TotalParking   uint8 `json:"totalParking"`
	HasAC          *bool `json:"hasAC"`

	Images []string `json:"images"`
}

type Address struct {
	Country      string `json:"country" validate:"max=100"`
	City         string `json:"city" validate:"max=100"`
	State        string `json:"state" validate:"max=100"`
	Street       string `json:"street" validate:"max=200"`
	ZipCode      string `json:"zipCode" validate:"max=10"`
	HouseNumber  string `json:"houseNumber" validate:"max=10"`
	Neighborhood string `json:"neighborhood" validate:"max=100"`
}

type PropertyCreateRequest struct {
	ID int `json:"ID"`

	Location      Location `json:"location" validate:"required"`
	Price         float32  `json:"price" validate:"required,min=0,max=99999999"`
	PriceCurrency string   `json:"priceCurrency" validate:"required,min=0,max=3"`

	FullAddress string  `json:"fullAddress" validate:"required,min=0,max=1000"`
	Address     Address `json:"address"`

	HomeSize     float32 `json:"homeSize" validate:"required,min=0,max=9999"`
	LotSize      float32 `json:"lotSize" validate:"required,min=0,max=9999"`
	LivingSize   float32 `json:"livingSize" validate:"required,min=0,max=9999"`
	YearBuild    uint16  `json:"yearBuild" validate:"required,number,min=0,max=9999"`
	Bedroom      uint8   `json:"bedroom" validate:"required,number,min=0,max=50"`
	Bathroom     uint8   `json:"bathroom" validate:"required,number,min=0,max=50"`
	Floor        uint8   `json:"floor" validate:"required,number,min=0,max=999"`
	TotalFloors  uint8   `json:"totalFloors" validate:"required,number,min=0,max=999"`
	PropertyType uint8   `json:"propertyType" validate:"required,min=0,max=20"`
	HomeType     uint8   `json:"homeType" validate:"required,min=0,max=20"`
	Condition    uint8   `json:"condition" validate:"required,min=0,max=20"`

	BrokerName  string `json:"brokerName" validate:"required,min=0,max=120"`
	Description string `json:"description" validate:"required,min=0,max=10000"`

	IsActiveStatus *bool `json:"isActiveStatus" validate:"boolean"`
	HasImages      *bool `json:"hasImages" validate:"boolean"`
	HasGarage      *bool `json:"hasGarage" validate:"boolean"`
	HasVideo       *bool `json:"hasVideo" validate:"boolean"`
	Has3DTour      *bool `json:"has3DTour" validate:"boolean"`
	TotalParking   uint8 `json:"totalParking" validate:"min=0,max=100"`
	HasAC          *bool `json:"hasAC" validate:"boolean"`
	Appliance      *bool `json:"appliance" validate:"boolean"`
	Heating        uint8 `json:"heating" validate:"min=0,max=3"`
	PetsAllowed    *bool `json:"petsAllowed" validate:"boolean"`
}

type Point [2]float64
