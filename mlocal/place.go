package mlocal

import (
	"fmt"
	"strings"
	"sync"
)

type Place struct {
	Name      string
	Category  string
	Latitude  float64
	Longitude float64
}

func addPs(ps *[]Place, w Place, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*ps = append(*ps, w)
	m.Unlock()
	wg.Done()
}

func GeneratePlaces(ws []Wilayah) (ps []Place) {
	var wg sync.WaitGroup
	var m sync.Mutex
	for _, w := range ws {
		if w.Level == LEVEL_KABKOTA {
			p := Place{}
			p.Name = fmt.Sprintf("Kantor Pemerintahan %v", w.Name)
			p.Name = strings.Title(strings.ToLower(p.Name))
			p.Category = "Kantor Pemerintah Kabupaten/Kota"
			p.Latitude = w.Latitude
			p.Longitude = w.Longitude
			wg.Add(1)
			go addPs(&ps, p, &wg, &m)

			for i := 1; i <= 3; i++ {
				p := Place{}
				p.Name = fmt.Sprintf("Rumah Sakit %v %v", w.Name, i)
				p.Name = strings.Title(strings.ToLower(p.Name))
				p.Category = "Rumah Sakit"
				p.Latitude = w.Latitude
				p.Longitude = w.Longitude
				wg.Add(1)
				go addPs(&ps, p, &wg, &m)
			}
			for i := 1; i <= 20; i++ {
				p := Place{}
				p.Name = fmt.Sprintf("SMA %v %v", strings.Title(strings.ToLower(w.Name)), i)
				p.Category = "Sekolah Menengah Atas"
				p.Latitude = w.Latitude
				p.Longitude = w.Longitude
				wg.Add(1)
				go addPs(&ps, p, &wg, &m)
			}
		}
		if w.Level == LEVEL_KECAMATAN {
			p := Place{}
			p.Name = fmt.Sprintf("Kantor Pemerintahan Kecamatan %v", strings.Title(strings.ToLower(w.Name)))
			p.Category = "Kantor Pemerintah Kecamatan"
			p.Latitude = w.Latitude
			p.Longitude = w.Longitude
			wg.Add(1)
			go addPs(&ps, p, &wg, &m)
			for i := 1; i <= 5; i++ {
				p := Place{}
				p.Name = fmt.Sprintf("Puskesmas %v %v", strings.Title(strings.ToLower(w.Name)), i)
				p.Category = "Puskesmas"
				p.Latitude = w.Latitude
				p.Longitude = w.Longitude
				wg.Add(1)
				go addPs(&ps, p, &wg, &m)
			}
			for i := 1; i <= 3; i++ {
				p := Place{}
				p.Name = fmt.Sprintf("SMP %v %v", strings.Title(strings.ToLower(w.Name)), i)
				p.Category = "Sekolah Menengah Pertama"
				p.Latitude = w.Latitude
				p.Longitude = w.Longitude
				wg.Add(1)
				go addPs(&ps, p, &wg, &m)
			}
		}
		if w.Level == LEVEL_KELURAHANDESA {
			p := Place{}
			p.Name = fmt.Sprintf("Kantor Pemerintahan Desa %v", strings.Title(strings.ToLower(w.Name)))
			p.Category = "Kantor Pemerintah Kelurahan/Desa"
			p.Latitude = w.Latitude
			p.Longitude = w.Longitude
			wg.Add(1)
			go addPs(&ps, p, &wg, &m)

			for i := 1; i <= 5; i++ {
				p := Place{}
				p.Name = fmt.Sprintf("SD %v %v", strings.Title(strings.ToLower(w.Name)), i)
				p.Category = "Sekolah Dasar"
				p.Latitude = w.Latitude
				p.Longitude = w.Longitude
				wg.Add(1)
				go addPs(&ps, p, &wg, &m)
			}
			for i := 1; i <= 20; i++ {
				p := Place{}
				p.Name = fmt.Sprintf("Tempat Ibadah %v %v", strings.Title(strings.ToLower(w.Name)), i)
				p.Category = "Tempat Ibadah"
				p.Latitude = w.Latitude
				p.Longitude = w.Longitude
				wg.Add(1)
				go addPs(&ps, p, &wg, &m)
			}
		}
	}
	wg.Wait()
	ps = uniquePlaces(ps)
	return
}

func uniquePlaces(slice []Place) []Place {
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
