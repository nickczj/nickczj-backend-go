package service

import (
	"context"
	"github.com/nickczj/web1/global"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"time"
)

type Weather struct {
	PSI interface{}
	UV  interface{}
}

func Now() (*Weather, error) {
	weather := &Weather{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		resp, err := datagovget(ctx, "/environment/psi")
		if err != nil {
			return err
		}
		weather.PSI = resp
		return nil
	})

	g.Go(func() error {
		resp, err := datagovget(ctx, "/environment/uv-index")
		if err != nil {
			return err
		}
		weather.UV = resp
		return nil
	})

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	return weather, nil
}

func datagovget(ctx context.Context, api string) (interface{}, error) {
	now := time.Now().Format("2006-01-02T15:04:05")
	url := "https://api.data.gov.sg/v1" + api

	result := new(interface{})

	resp, err := global.Client.R().
		SetContext(ctx).
		SetResult(result).
		SetQueryParams(map[string]string{
			"date_time": now,
		}).
		EnableTrace().
		Get(url)

	if err != nil {
		return nil, err
	}

	log.Info("Resp time: ", resp.Request.TraceInfo().TotalTime)

	return result, nil
}
