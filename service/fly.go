package service

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/nickczj/web1/global"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"sort"
	"sync"
	"time"
)

type Fare struct {
	CabinClass    string `json:"cabinClass"`
	DepartureDate string `json:"departureDate"`
	FareFamily    string `json:"fareFamily"`
	Miles         int    `json:"miles"`
	Status        string `json:"status"`
}

type FlightSearch struct {
	Destination string
	Saver       Saver
	Advantage   Advantage
}

type Saver struct {
	Miles     int      `json:"miles"`
	Immediate []string `json:"immediate"`
	Waitlist  []string `json:"waitlist"`
}
type Advantage struct {
	Miles     int      `json:"miles"`
	Immediate []string `json:"immediate"`
	Waitlist  []string `json:"waitlist"`
}

const url = "https://www.singaporeair.com/redemption/getCalendarSearch.form"

func SearchMulti(ctx context.Context, destinations []string) (interface{}, error) {
	var res []FlightSearch

	for _, destination := range destinations {
		search, err := Search(ctx, "SIN", destination)
		if err != nil {
			return nil, err
		}
		search.Destination = destination
		res = append(res, search)
	}

	return res, nil
}

func SearchMultiParallel(ctx context.Context, destinations []string) (interface{}, error) {
	ctx, _ = context.WithCancel(ctx)
	var wg = &sync.WaitGroup{}

	var res []FlightSearch

	for _, destination := range destinations {
		wg.Add(1)
		go func(destination string) {
			search, err := Search(ctx, "SIN", destination)
			if err != nil {
				return
			}
			res = append(res, search)
			wg.Done()
		}(destination)
	}

	wg.Wait()

	return res, nil
}

func Search(ctx context.Context, origin string, destination string) (FlightSearch, error) {
	ctx, _ = context.WithCancel(ctx)

	departureDate := time.Now()
	weeks := []time.Time{departureDate}

	var res []Fare

	w := 0
	for w < 53 {
		departureDate = departureDate.Add(time.Hour * 24 * 7) // add 1 week
		weeks = append(weeks, departureDate)
		w += 1
	}

	var wg = &sync.WaitGroup{}

	w = 0
	for w < 53 {
		wg.Add(1)
		go func(w int) {
			flights, err := flightsearch(origin, destination, weeks[w])
			if err != nil {
				log.Error("Error searching flights: ", err)
			}

			r, err := processresult(flights)
			res = append(res, r...)
			wg.Done()
		}(w)
		w += 1
	}

	wg.Wait()

	//log.Info("res ", res)

	saver := lo.Filter[Fare](res, func(x Fare, _ int) bool {
		return x.Status != "Waitlist" && x.FareFamily == "Saver"
	})

	saverWaitlist := lo.Filter[Fare](res, func(x Fare, _ int) bool {
		return x.Status == "Waitlist" && x.FareFamily == "Saver"
	})

	advantage := lo.Filter[Fare](res, func(x Fare, _ int) bool {
		return x.Status != "Waitlist" && x.FareFamily == "Advantage"
	})

	advantageWaitlist := lo.Filter[Fare](res, func(x Fare, _ int) bool {
		return x.Status == "Waitlist" && x.FareFamily == "Advantage"
	})

	sortbydate(saver)
	sortbydate(saverWaitlist)
	sortbydate(advantage)
	sortbydate(advantageWaitlist)

	var saverMiles int
	if len(saver) != 0 {
		saverMiles = saver[0].Miles
	} else if len(saverWaitlist) != 0 {
		saverMiles = saverWaitlist[0].Miles
	}

	var advantageMiles int
	if len(advantage) != 0 {
		advantageMiles = advantage[0].Miles
	} else if len(advantageWaitlist) != 0 {
		advantageMiles = advantageWaitlist[0].Miles
	}

	flightsearch := FlightSearch{
		Saver: Saver{
			Miles: saverMiles,
			Immediate: lo.Map[Fare, string](saver, func(x Fare, _ int) string {
				return x.DepartureDate
			}),
			Waitlist: lo.Map[Fare, string](saverWaitlist, func(x Fare, _ int) string {
				return x.DepartureDate
			}),
		},
		Advantage: Advantage{
			Miles: advantageMiles,
			Immediate: lo.Map[Fare, string](advantage, func(x Fare, _ int) string {
				return x.DepartureDate
			}),
			Waitlist: lo.Map[Fare, string](advantageWaitlist, func(x Fare, _ int) string {
				return x.DepartureDate
			}),
		},
	}

	return flightsearch, nil
}

func flightsearch(origin string, destination string, departureDate time.Time) (string, error) {
	result := new(interface{})
	userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/110.0"

	resp, err := global.Client.R().
		SetResult(result).
		SetHeaders(map[string]string{
			"Host":                "www.singaporeair.com",
			"User-Agent":          userAgent,
			"Accept":              "application/json,text/plain,*/*",
			"Accept-Language":     "en-US,en;q=0.5",
			"Accept-Encoding":     "gzip,deflate,br",
			"X-Sec-Clge-Req-Type": "ajax",
			"DNT":                 "1",
			"Sec-Fetch-Dest":      "empty",
			"Sec-Fetch-Mode":      "cors",
			"Sec-Fetch-Site":      "same-origin",
			"Referer":             "https://www.singaporeair.com/redemption/loadFlightSearchPage.form",
			"Connection":          "keep-alive",
			"Cookie":              viper.GetString("sia.cookie"),
			"Pragma":              "no-cache",
			"Cache-Control":       "no-cache",
		}).
		SetQueryParams(map[string]string{
			"origin":        origin,
			"destination":   destination,
			"departureDate": departureDate.Format("02_01_2006"),
			"dateRange":     "3",
		}).
		EnableTrace().
		Get(url)

	log.Debug("Resp time: ", resp.Request.TraceInfo().TotalTime)
	log.Debug("Raw Response: ", *result)

	if err != nil {
		return "", err
	}

	if *result == nil {
		return "", &resty.ResponseError{
			Response: nil,
			Err:      nil,
		}
	}

	res, err := json.Marshal(result)

	if err != nil {
		return "", err
	}

	return string(res), nil
}

func processresult(input string) ([]Fare, error) {
	var res []Fare
	value := gjson.Get(input, "flightSearch.response.data.segments.0.fares")

	err := json.Unmarshal([]byte(value.String()), &res)
	if err != nil {
		log.Error("Error unmarshalling JSON: ", err)
		return nil, err
	}
	return res, nil
}

func sortbydate(fares []Fare) {
	sort.Slice(fares, func(i, j int) bool {
		i1, err := time.Parse("2006-01-02", fares[i].DepartureDate)
		if err != nil {
			log.Error("Error parsing departure time: ", err)
		}
		j1, err := time.Parse("2006-01-02", fares[j].DepartureDate)
		if err != nil {
			log.Error("Error parsing departure time: ", err)
		}
		return i1.Before(j1)
	})
}
