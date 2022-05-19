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
	"github.com/shivamk2406/newsletter-subscriptions/cmd/transport"
	"github.com/shivamk2406/newsletter-subscriptions/internal/config"
	newspb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/news"
	subspb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/subscriptions"
	userpb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/user"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/news"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/subscriptions"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/users"
	"github.com/shivamk2406/newsletter-subscriptions/pkg/kafka/consumer"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func Start() error {

	ctx := context.Background()
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	conf, err := InitializeConfig()
	if err != nil {
		fmt.Println(err)
	}

	db, cleanup, err := initializeDB(conf)
	if err != nil {
		fmt.Println(err)
	}

	serv := initializeRegistry(db, logger)

	run(ctx, serv, logger)

	defer cleanup()
	return nil
}

func run(ctx context.Context, serv *service.Registry, logger log.Logger) {
	userEndpoint := users.MakeEndpoint(serv.UsersService)
	newsEndpoint := news.MakeEndpoint(serv.NewsService)
	subsEndpoint := subscriptions.MakeEndpoint(serv.SubscriptionService)

	userServer := transport.NewUserGrpcServer(userEndpoint)
	newsServer := transport.NewNewsServer(newsEndpoint)
	subsServer := transport.NewSubscriptionServer(subsEndpoint)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	port, network, err := config.LoadGrpcConfig()
	if err != nil {
		fmt.Println(err)
	}
	cfg, err := config.LoadConsumerConfig()
	if err != nil {
		fmt.Println(err)
	}
	newConsumer, err := consumer.NewConsumer(ctx, cfg)
	if err != nil {
		fmt.Println(err)
	}
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
		level.Info(logger).Log("msg", "Server started successfully!!")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
