package models

import (
	"fmt"
	"strings"
	"sync"
)

type Place struct {
	BaseModel
	Name         string   `json:"name" gorm:"unique"`
	CategoryID   uint     `json:"category_id"`
	Category     Category `json:"-"`
	CityName     string   `json:"city_name"`
	DistrictName string   `json:"district_name"`
	Latitude     float64  `json:"latitude"`
	Longitude    float64  `json:"longitude"`
}

func addPs(ps *[]Place, p Place, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*ps = append(*ps, p)
	m.Unlock()
	wg.Done()
}

func formatTitle(s string) string {
	return strings.Title(strings.ToLower(s))
}

// catmap is "category map" catmap[name]catid
func GeneratePlaces(ws []Wilayah, catmap map[string]uint) (ps []Place) {
	var wg sync.WaitGroup
	var m sync.Mutex
	for _, w := range ws {
		if w.Level == LEVEL_KABKOTA {
			p := Place{}
			p.Name = fmt.Sprintf("Kantor Pemerintahan %v", w.Name)
			p.CategoryID = catmap[CATEGORY_KANTOR_PEM_KABKOTA]
			p.Latitude = w.Latitude
			p.Longitude = w.Longitude
			p.CityName = w.CityName
			p.DistrictName = w.DistrictName
			p.Name = formatTitle(p.Name)
			p.CityName = formatTitle(p.CityName)
			p.DistrictName = formatTitle(p.DistrictName)
			wg.Add(1)
			go addPs(&ps, p, &wg, &m)

			for i := 1; i <= 3; i++ {
				p := Place{}
				p.Name = fmt.Sprintf("Rumah Sakit %v %v", w.Name, i)
				p.CategoryID = catmap[CATEGORY_RUMAH_SAKIT]
				p.Latitude = w.Latitude
				p.Longitude = w.Longitude
				p.CityName = w.CityName
				p.DistrictName = w.DistrictName
				p.Name = formatTitle(p.Name)
				p.CityName = formatTitle(p.CityName)
				p.DistrictName = formatTitle(p.DistrictName)
				wg.Add(1)
				go addPs(&ps, p, &wg, &m)
			}
			for i := 1; i <= 20; i++ {
				p := Place{}
				p.Name = fmt.Sprintf("%v %v", w.Name, i)
				p.CategoryID = catmap[CATEGORY_SMA]
				p.Latitude = w.Latitude
				p.Longitude = w.Longitude
				p.CityName = w.CityName
				p.DistrictName = w.DistrictName
				p.Name = formatTitle(p.Name)
				p.CityName = "SMA " + formatTitle(p.CityName)
				p.DistrictName = formatTitle(p.DistrictName)
				wg.Add(1)
				go addPs(&ps, p, &wg, &m)
			}
		}
		if w.Level == LEVEL_KECAMATAN {
			p := Place{}
			p.Name = fmt.Sprintf("Kantor Pemerintahan Kecamatan %v", w.Name)
			p.CategoryID = catmap[CATEGORY_KANTOR_PEM_KECAMATAN]
			p.Latitude = w.Latitude
			p.Longitude = w.Longitude
			p.CityName = w.CityName
			p.DistrictName = w.DistrictName
			p.Name = formatTitle(p.Name)
			p.CityName = formatTitle(p.CityName)
			p.DistrictName = formatTitle(p.DistrictName)
			wg.Add(1)
			go addPs(&ps, p, &wg, &m)
			for i := 1; i <= 5; i++ {
				p := Place{}
				p.Name = fmt.Sprintf("Puskesmas %v %v", w.Name, i)
				p.CategoryID = catmap[CATEGORY_PUSKESMAS]
				p.Latitude = w.Latitude
				p.Longitude = w.Longitude
				p.CityName = w.CityName
				p.DistrictName = w.DistrictName
				p.Name = formatTitle(p.Name)
				p.CityName = formatTitle(p.CityName)
				p.DistrictName = formatTitle(p.DistrictName)
				wg.Add(1)
				go addPs(&ps, p, &wg, &m)
			}
			for i := 1; i <= 3; i++ {
				p := Place{}
				p.Name = fmt.Sprintf("%v %v", w.Name, i)
				p.CategoryID = catmap[CATEGORY_SMP]
				p.Latitude = w.Latitude
				p.Longitude = w.Longitude
				p.CityName = w.CityName
				p.DistrictName = w.DistrictName
				p.Name = "SMP " + formatTitle(p.Name)
				p.CityName = formatTitle(p.CityName)
				p.DistrictName = formatTitle(p.DistrictName)
				wg.Add(1)
				go addPs(&ps, p, &wg, &m)
			}
		}
		if w.Level == LEVEL_KELURAHANDESA {
			p := Place{}
			p.Name = fmt.Sprintf("Kantor Pemerintahan Desa %v", w.Name)
			p.CategoryID = catmap[CATEGORY_KANTOR_PEM_KELURAHANDESA]
			p.Latitude = w.Latitude
			p.Longitude = w.Longitude
			p.CityName = w.CityName
			p.DistrictName = w.DistrictName
			p.Name = formatTitle(p.Name)
			p.CityName = formatTitle(p.CityName)
			p.DistrictName = formatTitle(p.DistrictName)
			wg.Add(1)
			go addPs(&ps, p, &wg, &m)

			for i := 1; i <= 5; i++ {
				p := Place{}
				p.Name = fmt.Sprintf("%v %v", w.Name, i)
				p.CategoryID = catmap[CATEGORY_SD]
				p.Latitude = w.Latitude
				p.Longitude = w.Longitude
				p.CityName = w.CityName
				p.DistrictName = w.DistrictName
				p.Name = "SD " + formatTitle(p.Name)
				p.CityName = formatTitle(p.CityName)
				p.DistrictName = formatTitle(p.DistrictName)
				wg.Add(1)
				go addPs(&ps, p, &wg, &m)
			}
			for i := 1; i <= 20; i++ {
				p := Place{}
				p.Name = fmt.Sprintf("Tempat Ibadah %v %v", w.Name, i)
				p.CategoryID = catmap[CATEGORY_TEMPAT_IBADAH]
				p.Latitude = w.Latitude
				p.Longitude = w.Longitude
				p.CityName = w.CityName
				p.DistrictName = w.DistrictName
				p.Name = formatTitle(p.Name)
				p.CityName = formatTitle(p.CityName)
				p.DistrictName = formatTitle(p.DistrictName)
				wg.Add(1)
				go addPs(&ps, p, &wg, &m)
			}
		}
	}
	wg.Wait()
	ps = uniquePs(ps)
	return
}

func uniquePs(slice []Place) []Place {
	keys := make(map[string]bool)
	list := []Place{}
	for _, entry := range slice {
		if _, value := keys[entry.Name]; !value {
			keys[entry.Name] = true
			list = append(list, entry)
		}
	}
	return list
}
