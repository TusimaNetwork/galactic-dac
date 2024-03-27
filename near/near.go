package near

import (
	"context"
	"errors"
	"fmt"
	"github.com/0xPolygon/cdk-data-availability/config"
	"github.com/0xPolygon/cdk-data-availability/log"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"sort"
	"strings"
	"time"
)

type NearDA struct {
	db  *PostgresStorage
	cfg config.NearConfig
	ch  chan struct{}
}

func New(pg *pgxpool.Pool, cfg config.NearConfig, ch chan struct{}) *NearDA {
	return &NearDA{
		db:  NewPostgresStorage(pg),
		cfg: cfg,
		ch:  ch,
	}
}

func (s *NearDA) Start() {
	ticker := time.NewTicker(s.cfg.WaitPeriod.Duration)
	for {
		select {
		case <-ticker.C:
			if len(s.ch) <= 0 {
				s.ch <- struct{}{}
			}
		case <-s.ch:
			for len(s.ch) > 0 {
				<-s.ch
			}
			log.Info("start checking uploaded data to Near DA...")
			s.start()
		}
	}
}

func (s *NearDA) start() {
	var (
		err error
	)

	ctx := context.Background()

	datas, err := s.checkData(ctx, nil)
	if (err != nil && errors.Is(err, ErrStateNotFound)) || datas == nil || len(datas) <= 0 {
		log.Info("no data needs to be uploaded to Near DA")
		return
	} else if err != nil {
		log.Errorf("check data error occurred: %v", err)
		return
	}

	for _, da := range datas {
		err := s.uploadData(da.Value, da.Key, da.Id)
		if err != nil {
			break
		}
	}
}

type WrapResData struct {
	Id    int
	Key   string
	Value string
}

func (s *NearDA) checkData(ctx context.Context, dbTx pgx.Tx) ([]WrapResData, error) {
	logId, err := s.GetNearStateLog(ctx, dbTx)
	if err != nil && !errors.Is(err, ErrStateNotFound) {
		return nil, err
	}

	log.Info("111111111111111111111")

	res, err := s.GetNearChainLog(ctx, logId, dbTx)
	if err != nil {
		log.Errorf("err: %v", err)

		return nil, err
	}

	log.Info("222222222222222222222")

	keyLogId := make(map[string]int)
	keys := make([]string, 0)
	for _, re := range res {
		keys = append(keys, re.Key)
		keyLogId[re.Key] = re.Id
	}

	datas, err := s.GetOffChainData(ctx, keys, dbTx)
	if err != nil {
		log.Errorf("err: %v", err)

		return nil, err
	}
	log.Info("333333333333333333333")

	wrapRes := make([]WrapResData, 0)
	for _, da := range datas {
		re := WrapResData{
			Id:    keyLogId[da.Key],
			Key:   da.Key,
			Value: da.Value,
		}
		wrapRes = append(wrapRes, re)
	}

	sort.SliceStable(wrapRes, func(i, j int) bool {
		return wrapRes[i].Id < wrapRes[j].Id
	})
	return wrapRes, nil
}

func (s *NearDA) uploadData(data, key string, logId int) error {

	contractValue := fmt.Sprintf(`{"greeting":"%s", "greetingValue":"%s"}`, key, data)
	res, err := DoCommand(s.cfg.Command, "call", s.cfg.Contract, s.cfg.ContractMethod, contractValue,
		"--account-id", s.cfg.Account)
	if err != nil {
		log.Errorf("upload data to Near DA error occurred: %v", err)
		return err
	}

	fmt.Println(res)
	idIndex := strings.Index(res, "Transaction Id ")
	tx := res[idIndex+15 : idIndex+15+44]
	ctx := context.Background()

	if err := s.StoreNearStateLog(ctx, logId, tx, nil); err != nil {
		log.Errorf("upload data to Near DA error occurred: %v", err)
		return err
	}

	return nil
}
