package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	log "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/robfig/cron"
	"github.com/shivamk2406/newsletter-subscriptions/cmd/transport"
	"github.com/shivamk2406/newsletter-subscriptions/internal/config"
	newspb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/news"
	subspb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/subscriptions"
	userpb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/user"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service"
	"github.com/shivamk2406/newsletter-subscriptions/pkg/kafka/consumer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Start() error {

	ctx := context.Background()
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	conf, err := InitializeConfig()
	if err != nil {
		logger.Log(err)
		return err
	}

	db, cleanup, err := initializeDB(conf)
	if err != nil {
		logger.Log(err)
		return err
	}

	serv := initializeRegistry(ctx, db, logger)

	run(ctx, serv, logger)

	defer cleanup()
	return nil
}

func run(ctx context.Context, serv *service.Registry, logger log.Logger) {
	userServer := transport.NewUserGrpcServer(ctx, *serv)
	newsServer := transport.NewNewsServer(ctx, *serv)
	subsServer := transport.NewSubscriptionServer(ctx, *serv)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	port, network, err := config.LoadGrpcConfig()
	if err != nil {
		logger.Log(err)
	}
	cfg, err := config.LoadConsumerConfig()
	if err != nil {
		logger.Log(err)
	}
	newConsumer, err := consumer.NewConsumer(ctx, cfg)
	if err != nil {
		logger.Log(err)

	}

	c := cron.New()
	c.AddFunc("@daily", func() { serv.CronService(ctx, *newConsumer) })
	c.Start()
	go serv.CronService(ctx, *newConsumer)
	go func() {
		grpcListener, err := net.Listen(network, ":"+port)
		if err != nil {
			logger.Log("during", "Listen", "err", err)
			os.Exit(1)
		}
		baseServer := grpc.NewServer()
		userpb.RegisterUserManagementServiceServer(baseServer, userServer)
		subspb.RegisterSubscriptionManagementServiceServer(baseServer, subsServer)
		newspb.RegisterNewsServiceServer(baseServer, newsServer)
		reflection.Register(baseServer)
		level.Info(logger).Log("msg", "Server started successfully!!")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
