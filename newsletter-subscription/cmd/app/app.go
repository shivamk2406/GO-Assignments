package app

import (
	"log"
	"net"

	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func Start() error {

	//ctx := context.Background()
	conf, err := InitializeConfig()
	if err != nil {
		log.Println(err)
	}
	db, cleanup, err := initializeDB(conf)
	if err != nil {
		log.Println(err)
	}
	repo := initializeRepo(db)
	serv := initializeUserManagementServer(repo)

	func() {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Println(err)
		}
		s := grpc.NewServer()
		pb.RegisterUserManagementServer(s, serv)
		reflection.Register(s)
		log.Printf("server listening at %v", lis.Addr().String())

		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to Serve: %v", err)

		}

	}()
	//grpcserver.RunServer(ctx, repo)
	defer cleanup()

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
