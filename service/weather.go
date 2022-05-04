package service

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/nickczj/web1/global"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"time"
)

type Weather struct {
	PSI any
	UV  any
}

func Now() (*Weather, error) {
	global.Client = resty.New()

	psi := make(chan any)
	uv := make(chan any)

	ctx, _ := context.WithCancel(context.Background())
	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		defer close(psi)
		resp, err := datagovget("/environment/psi")
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case psi <- resp:
		}
		return nil
	})

	g.Go(func() error {
		defer close(uv)
		resp, err := datagovget("/environment/uv-index")
		if err != nil {
			return err
		}
		log.Info("hmm ")
		select {
		case <-ctx.Done():
			return ctx.Err()
		case uv <- resp:
		}
		return nil
	})

	a, b := <-psi, <-uv

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	return &Weather{a, b}, nil
}

func datagovget(api string) (any, error) {
	now := time.Now().Format("2006-01-02T15:04:05")
	url := "https://api.data.gov.sg/v1" + api

	result := new(any)

	resp, err := global.Client.R().
		SetResult(result).
		SetQueryParams(map[string]string{
			"date_time": now,
		}).
		EnableTrace().
		Get(url)

	log.Info(resp.Request.TraceInfo())

	if err != nil {
		return "", err
	}

	return result, nil
}
