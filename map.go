package mfcg

// ID names can be found in a map's JSON data.
const (
	IDEarth     string = "earth"
	IDPlanks    string = "planks"
	IDRivers    string = "rivers"
	IDRoads     string = "roads"
	IDBuildings string = "buildings"
	IDFields    string = "fields"
	IDGreens    string = "greens"
	IDPrisms    string = "prisms"
	IDSquares   string = "squares"
	IDWalls     string = "walls"
	IDWater     string = "water"
	IDValues    string = "values"
)

// Map represents a cartographical map. It contains a collection of features
// commonly used to represent a medieval fantasy city.
type Map struct {
	MetaData
	Earth     Polygon      `json:"earth,omitempty"`
	Planks    []LineString `json:"planks,omitempty"`
	Rivers    []LineString `json:"rivers,omitempty"`
	Roads     []LineString `json:"roads,omitempty"`
	Buildings []Polygon    `json:"buildings,omitempty"`
	Fields    []Polygon    `json:"fields,omitempty"`
	Greens    []Polygon    `json:"greens,omitempty"`
	Prisms    []Polygon    `json:"prisms,omitempty"`
	Squares   []Polygon    `json:"squares,omitempty"`
	Walls     []Polygon    `json:"walls,omitempty"`
	Water     []Polygon    `json:"water,omitempty"`
}

// MetaData contains various map feature parameters and generator details.
type MetaData struct {
	RoadWidth     int     `json:"roadWidth,omitempty"`
	RiverWidth    float64 `json:"riverWidth,omitempty"`
	TowerRadius   float64 `json:"towerRadius,omitempty"`
	WallThickness float64 `json:"wallThickness,omitempty"`
	Generator     string  `json:"generator,omitempty"`
	Version       string  `json:"version,omitempty"`
}

// toMap transforms the provided feature map into a cartographical Map. The
// feature map's keys must match each features' respective ID.
func toMap(feats map[string]feature) (*Map, error) {
	var result Map

	if ft, ok := feats[IDEarth]; ok {
		p, err := coordsToPolygon(ft.Coordinates)
		if err != nil {
			return nil, err
		}
		result.Earth = *p
	}

	if ft, ok := feats[IDPlanks]; ok {
		ls, err := geosToLineStrings(ft.Geometries)
		if err != nil {
			return nil, err
		}
		result.Planks = ls
	}

	if ft, ok := feats[IDRivers]; ok {
		ls, err := geosToLineStrings(ft.Geometries)
		if err != nil {
			return nil, err
		}
		result.Rivers = ls
	}

	if ft, ok := feats[IDRoads]; ok {
		ls, err := geosToLineStrings(ft.Geometries)
		if err != nil {
			return nil, err
		}
		result.Roads = ls
	}

	if ft, ok := feats[IDBuildings]; ok {
		p, err := coordsToPolygons(ft.Coordinates)
		if err != nil {
			return nil, err
		}
		result.Buildings = p
	}

	if ft, ok := feats[IDFields]; ok {
		p, err := coordsToPolygons(ft.Coordinates)
		if err != nil {
			return nil, err
		}
		result.Fields = p
	}

	if ft, ok := feats[IDGreens]; ok {
		p, err := coordsToPolygons(ft.Coordinates)
		if err != nil {
			return nil, err
		}
		result.Greens = p
	}

	if ft, ok := feats[IDPrisms]; ok {
		p, err := coordsToPolygons(ft.Coordinates)
		if err != nil {
			return nil, err
		}
		result.Prisms = p
	}

	if ft, ok := feats[IDSquares]; ok {
		p, err := coordsToPolygons(ft.Coordinates)
		if err != nil {
			return nil, err
		}
		result.Squares = p
	}

	if ft, ok := feats[IDWalls]; ok {
		p, err := geosToPolygons(ft.Geometries)
		if err != nil {
			return nil, err
		}
		result.Walls = p
	}

	if ft, ok := feats[IDWater]; ok {
		p, err := coordsToPolygons(ft.Coordinates)
		if err != nil {
			return nil, err
		}
		result.Water = p
	}

	if ft, ok := feats[IDValues]; ok {
		result.RoadWidth = ft.RoadWidth
		result.RiverWidth = ft.RiverWidth
		result.WallThickness = ft.WallThickness
		result.TowerRadius = ft.TowerRadius
		result.Generator = ft.Generator
		result.Version = ft.Version
	}

	return &result, nil
}
