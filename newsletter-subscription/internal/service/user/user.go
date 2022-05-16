package user

import (
	"context"
	"os"

	"github.com/go-kit/log"
	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto"
)

type UserManagement interface {
	CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error)
	ListPlans(ctx context.Context, in *pb.ListPlansRequest) (*pb.Plans, error)
	AuthenticateUser(ctx context.Context, in *pb.AuthenticateUserRequest) (*pb.AuthenticateUserResponse, error)
	GetSubscription(ctx context.Context, in *pb.SubscriptionRequest) (*pb.Plan, error)
	CreateSubscription(ctx context.Context, in *pb.CreateSubscriptionRequest) (*pb.SubscriptionResponse, error)
	ListNews(ctx context.Context, in *pb.ListNewsRequest) (*pb.News, error)
	ListNewsByGenre(ctx context.Context, in *pb.ListNewsByGenreRequest) (*pb.News, error)
}

type UserManagementServer struct {
	log  log.Logger
	repo Repository
}

func UserManagementService(repo *Repository, logger log.Logger) UserManagementServer {
	return UserManagementServer{repo: *repo, log: logger}
}
func (r UserManagementServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	model := CreateUserRequest{Name: in.Name, Email: in.Email}
	user, err := r.repo.CreateUser(ctx, model)
	if err != nil {
		return &pb.User{}, err
	}
	return &pb.User{Name: user.Email, Email: user.Email, Active: false}, nil
}

func (r UserManagementServer) ListPlans(ctx context.Context, in *pb.ListPlansRequest) (*pb.Plans, error) {
	model := GetPlansRequests{}
	plans, err := r.repo.GetPlans(ctx, model)
	if err != nil {
		return &pb.Plans{}, err
	}

	var subs []*pb.Plan
	for _, val := range plans.Subs {
		sub := pb.Plan{Id: val.Id, Name: val.Name, Validity: val.Validity}
		var genres []*pb.Genre
		for _, val1 := range val.Genres {
			genre := pb.Genre{Name: val1.Name}
			genres = append(genres, &genre)
		}
		sub.Genres = genres
		subs = append(subs, &sub)
	}
	return &pb.Plans{Subs: subs}, nil
}

func (r UserManagementServer) AuthenticateUser(ctx context.Context, in *pb.AuthenticateUserRequest) (*pb.AuthenticateUserResponse, error) {
	w := log.NewSyncWriter(os.Stderr)
	logger := log.NewLogfmtLogger(w)
	model := AuthenticateUserRequest{Email: in.Email}
	logger.Log("Inside User Request Generated for db %v", model)

	response, err := r.repo.AuthenticateUser(ctx, model)
	if err != nil {
		return &pb.AuthenticateUserResponse{}, err
	}

	user := pb.User{Name: response.User.Name, Email: response.User.Email, Active: response.User.Active}
	return &pb.AuthenticateUserResponse{IsAuthenticated: response.IsAuthenticated, User: &user}, nil
}

func (r UserManagementServer) GetSubscription(ctx context.Context, in *pb.SubscriptionRequest) (*pb.Plan, error) {
	model := SubscriptionRequest{Email: in.Email}
	response, err := r.repo.GetSubscription(ctx, model)
	if err != nil {
		return &pb.Plan{}, err
	}
	var genres []*pb.Genre
	for _, val := range response.Genres {
		genre := pb.Genre{Name: val.Name}
		genres = append(genres, &genre)
	}
	return &pb.Plan{Id: response.Id, Name: response.Name, Validity: response.Validity, Genres: genres}, nil
}

func (r UserManagementServer) CreateSubscription(ctx context.Context, in *pb.CreateSubscriptionRequest) (*pb.SubscriptionResponse, error) {
	model := SetSubscriptionRequest{Email: in.Email, Subsid: in.Subsid}
	response, err := r.repo.SetSubsciption(ctx, model)
	if err != nil {
		return &pb.SubscriptionResponse{}, err
	}
	return &pb.SubscriptionResponse{Email: in.Email, Active: true, Starttime: &response.Starttime, Validity: response.Validity}, nil
}

func (r UserManagementServer) ListNewsByGenre(ctx context.Context, in *pb.ListNewsByGenreRequest) (*pb.News, error) {
	model := GetNewsByGenreRequest{Genre: in.Genre}
	response, err := r.repo.GetNewsByGenre(ctx, model)
	if err != nil {
		return &pb.News{}, err
	}
	var newsString []string
	for _, val := range response.Newss {
		newsString = append(newsString, val)
	}
	return &pb.News{News: newsString}, nil
}

func (r UserManagementServer) ListNews(ctx context.Context, in *pb.ListNewsRequest) (*pb.News, error) {
	model := GetNewsRequest{Subsid: in.Subsid}
	response, err := r.repo.GetNews(ctx, model)
	if err != nil {
		return &pb.News{}, err
	}
	var newsString []string
	for _, val := range response.Newss {
		newsString = append(newsString, val)
	}
	return &pb.News{News: newsString}, nil
}

// func NewRepo(db *gorm.DB) pb.UserManagementServer {
// 	return &UserManagementServer{db: db}
// }

// func (r *UserManagementServer) GetGenresById(id int) []*pb.Genres {
// 	var subscriptionGenres []models.SubscriptionGenre
// 	r.db.Where("subscriptions_id = ?", id).Find(&subscriptionGenres)

// 	var genreIds []int
// 	for _, val := range subscriptionGenres {
// 		genreIds = append(genreIds, val.GenID)
// 	}

// 	var genres []models.Genre
// 	for _, val := range genreIds {
// 		var genre models.Genre
// 		r.db.Where("ID = ?", val).Find(&genre)
// 		genres = append(genres, genre)
// 	}

// 	var genresField []*pb.Genres
// 	for _, val := range genres {
// 		genre := pb.Genres{Name: val.Name}
// 		genresField = append(genresField, &genre)
// 	}
// 	return genresField
// }

// func (r *UserManagementServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
// 	log.Printf("Received: %v %v", in.Name, in.Email)
// 	user := models.User{ID: 1, Email: in.Email, Name: in.Name, Active: false, StartDate: time.Now()}
// 	if err := r.db.Create(&user).Error; err != nil {
// 		return &pb.User{}, err
// 	}

// 	log.Println(user)
// 	return &pb.User{Name: in.Name, Email: in.Email, Active: false}, nil
// }

// func (r *UserManagementServer) GetPlans(ctx context.Context, in *pb.GetPlansRequests) (*pb.Plans, error) {
// 	var subs []models.Subscriptions
// 	if err := r.db.Find(&subs).Error; err != nil {
// 		return &pb.Plans{}, err
// 	}

// 	var plans []*pb.Subscription
// 	for _, val := range subs {
// 		genres := r.GetGenresById(val.ID)
// 		plans = append(plans, &pb.Subscription{Name: val.Name, Id: int32(val.ID), Validity: int32(val.Renewal), Genres: genres})
// 	}
// 	return &pb.Plans{Subs: plans}, nil
// }

// func (r *UserManagementServer) GetAllSubscriptions() ([]models.Subscriptions, error) {
// 	var subscriptions []models.Subscriptions
// 	if err := r.db.Find(&subscriptions).Error; err != nil {
// 		log.Println(err)
// 		return []models.Subscriptions{}, err
// 	}

// 	return subscriptions, nil
// }

// func (r *UserManagementServer) AuthenticateUser(ctx context.Context, in *pb.AuthenticateUserRequest) (*pb.AuthenticateUserResponse, error) {
// 	var user models.User
// 	log.Printf("Received : %v", in.Email)
// 	if err := r.db.First(&user, "email = ?", in.Email).Error; err != nil {
// 		return &pb.AuthenticateUserResponse{IsAuthenticated: false}, err
// 	}
// 	return &pb.AuthenticateUserResponse{IsAuthenticated: true, User: &pb.User{Name: user.Name, Email: user.Email, Active: user.Active}}, nil
// }

// func (r *UserManagementServer) GetSubscription(ctx context.Context, in *pb.SubscriptionRequest) (*pb.Subscription, error) {
// 	log.Printf("Received : %v", in.Email)
// 	var user models.User
// 	row := r.db.First(&user, "email = ?", in.Email)
// 	if err := row.Error; err != nil {
// 		log.Println(err)
// 		return &pb.Subscription{}, fmt.Errorf("You have not Subscribed to any plans")
// 	}

// 	var subs models.Subscriptions
// 	if err := r.db.Where("ID = ?", user.SubsID).Find(&subs).Error; err != nil {
// 		return &pb.Subscription{}, err
// 	}

// 	genresField := r.GetGenresById(subs.ID)
// 	return &pb.Subscription{Id: int32(subs.ID), Name: subs.Name, Validity: int32(subs.Renewal), Genres: genresField}, nil
// }

// func (r *UserManagementServer) SetSubsciption(ctx context.Context, in *pb.SetSubscriptionRequest) (*pb.SetSubscriptionResponse, error) {
// 	log.Printf("Recieved Email %s and Subsid %d ", in.Email, in.Subsid)
// 	var subs models.Subscriptions
// 	r.db.Where("ID = ?", in.Subsid).Find(&subs)
// 	// Update with conditions
// 	if err := r.db.Model(&models.User{}).Where("email = ?", in.Email).
// 		Updates(map[string]interface{}{
// 			"active":     true,
// 			"start_time": time.Now(),
// 			"subsid":     in.Subsid,
// 			"validity":   subs.Renewal}).Error; err != nil {
// 		return &pb.SetSubscriptionResponse{}, err
// 	}

// 	return &pb.SetSubscriptionResponse{Email: in.Email, Active: true, Starttime: timestamppb.Now(), Validity: int32(subs.Renewal)}, nil
// }

// func (r *UserManagementServer) GetNews(ctx context.Context, in *pb.GetNewsRequest) (*pb.News, error) {
// 	log.Printf("Received %v", in.Subsid)
// 	var subscriptionGenres []models.SubscriptionGenre
// 	if err := r.db.Where("subscriptions_id = ?", in.Subsid).Find(&subscriptionGenres).Error; err != nil {
// 		return &pb.News{}, err
// 	}

// 	var genreIds []int
// 	for _, val := range subscriptionGenres {
// 		genreIds = append(genreIds, val.GenID)
// 	}

// 	var newsCollection []models.News
// 	if err := r.db.Where("genreid IN ?", genreIds).Find(&newsCollection).Error; err != nil {
// 		return &pb.News{}, nil
// 	}

// 	var newsString []string
// 	for _, val := range newsCollection {
// 		newsString = append(newsString, val.Description)
// 	}

// 	return &pb.News{News: newsString}, nil
// }

// func (r *UserManagementServer) GetNewsByGenre(ctx context.Context, in *pb.GetNewsByGenreRequest) (*pb.News, error) {
// 	log.Printf("Received %v", in.Genre)
// 	var genre models.Genre

// 	if err := r.db.Where("name = ?", in.Genre).Find(&genre).Error; err != nil {
// 		return &pb.News{}, err
// 	}

// 	var news []models.News
// 	if err := r.db.Where("genreid = ?", genre.ID).Find(&news).Error; err != nil {
// 		return &pb.News{}, err
// 	}

// 	var newsString []string
// 	for _, val := range news {
// 		newsString = append(newsString, val.Description)
// 	}

// 	return &pb.News{News: newsString}, nil

// }
