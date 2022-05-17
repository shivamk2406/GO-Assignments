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
	"github.com/shivamk2406/newsletter-subscriptions/internal/pkg/kafka"
	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/user"
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
	repo := initializeRepo(db)
	serv := user.UserManagementService(repo, logger)

	kaf := kafka.NewProducerConsumerService(serv, "my-topic", "news ", []string{"localhost:9092"})
	kaf.NewsProducer(ctx, "my-topic")
	kaf.ConsumeNews(ctx, "my-topic", "news")

	endpoints := user.MakeEndpoint(serv)
	grpcServer := transport.NewServer(endpoints)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterUserManagementServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
	defer cleanup()

	// func() {
	// 	lis, err := net.Listen("tcp", port)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	s := grpc.NewServer()
	// 	pb.RegisterUserManagementServer(s, serv)
	// 	reflection.Register(s)
	// 	log.Printf("server listening at %v", lis.Addr().String())

	// 	if err := s.Serve(lis); err != nil {
	// 		log.Fatalf("Failed to Serve: %v", err)

	// 	}

	// }()
	//grpcserver.RunServer(ctx, repo)

	// conf, err := config.LoadDatabaseConfig()
	// if err != nil {
	// 	log.Println(err)
	// }

	// db, cleanup, err := database.Open(conf)
	// if err != nil {
	// 	log.Println(err)
	// }
	// genres, err := repo.GetGenres()
	// if err != nil {
	// 	log.Println(err)
	// }
	//fmt.Println(genres)
	//subs, err := repo.GetAllSubscriptions()
	//if err != nil {
	//log.Println(err)
	//}
	//for _, val := range subs {
	//fmt.Printf("%d %s %d %d \n", val.ID, val.Name, val.Price, val.Renewal)
	//}
	// users, err := repo.AuthenticateUser("test@test.com")
	// if err != nil {
	// 	log.Println(err)
	// }

	// fmt.Printf("User Info:%s %s %d %d  ", users.Email, users.Name, users.ID, users.SubsID)
	//e2e.InsertGenreData(db)
	//defer cleanup()
	return nil
}
