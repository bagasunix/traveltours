package village

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

type villageModel struct {
	Village_id             int64
	Village_name           string
	Sub_district_id        int32
	Sub_district_name      string
	Sub_district_latitude  float64
	Sub_district_longitude float64
	City_id                int16
	City_name              string
	City_latitude          float64
	City_longitude         float64
	Province_id            int16
	Province_name          string
	Province_latitude      float64
	Province_longitude     float64
	Country_id             int16
	Country_name           string
	Country_latitude       float64
	Country_longitude      float64
	Village_postal_code    string
	Distance               float64
}

func subQueryByKeywordsByCoordinateByRadius(g *gormProvider, ctx context.Context, keyWord string, latitude float64, longitude float64) *gorm.DB {
	return g.db.WithContext(ctx).Table(`villages v`).Select(`v.id as village_id, v.name as village_name, sub_district_id, s.name as sub_district_name, s.latitude as sub_district_latitude, s.longitude as sub_district_longitude, city_id, c.name as city_name, c.latitude as city_latitude, c.longitude as city_longitude, province_id, p.name as province_name, p.latitude as province_latitude, p.longitude as province_longitude, country_id, co.name as country_name, co.latitude as country_latitude, co.longitude as country_longitude, v.postal_codes as village_postal_code, calculate_distance(s.latitude, s.longitude, @latitude, @longitude, 'K') as distance`).Joins(`
	JOIN sub_districts s on v.sub_district_id = s.id 
	JOIN cities c on s.city_id = c.id 
	JOIN provinces p on c.province_id = p.id 
	JOIN countries co on p.country_id = co.id `).Where(`v.name like @keyWord or s.name like @keyWord or c.name like @keyWord or p.name like @keyWord or v.postal_codes like @keyWord OR s.postal_codes like @keyWord`, sql.Named("keyWord", keyWord), sql.Named("latitude", latitude), sql.Named("longitude", longitude))
}

func subQueryByKeywords(g *gormProvider, ctx context.Context, keyWord string) *gorm.DB {
	return g.db.WithContext(ctx).Table(`villages v`).Select(`v.id as village_id, v.name as village_name, sub_district_id, s.name as sub_district_name, s.latitude as sub_district_latitude, s.longitude as sub_district_longitude,  city_id, c.name as city_name, c.latitude as city_latitude, c.longitude as city_longitude, province_id, p.name as province_name, p.latitude as province_latitude, p.longitude as province_longitude, country_id, co.name as country_name, co.latitude as country_latitude, co.longitude as country_longitude, v.postal_codes as village_postal_code`).Joins(`
	JOIN sub_districts s on v.sub_district_id = s.id 
	JOIN cities c on s.city_id = c.id 
	JOIN provinces p on c.province_id = p.id 
	JOIN countries co on p.country_id = co.id 
	`).Where(`v.name like @keyWord or s.name like @keyWord or c.name like @keyWord or p.name like @keyWord or v.postal_codes like @keyWord OR s.postal_codes like @keyWord`, sql.Named("keyWord", keyWord)).Order("v.name")
}
