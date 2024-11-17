package example

// repo

// IPostRepo IPostRepo
type IPostRepo interface{}

// NewPostRepo NewPostRepo
func NewPostRepo() IPostRepo {
	return new(IPostRepo)
}

// usecase

// IPostUsecase IPostUsecase
type IPostUsecase interface{}
type postUsecase struct {
	repo IPostRepo
}

// NewPostUsecase NewPostUsecase
func NewPostUsecase(repo IPostRepo) IPostUsecase {
	return postUsecase{repo: repo}
}

// service service

// PostService PostService
type PostService struct {
	usecase IPostUsecase
}

// NewPostService NewPostService
func NewPostService(u IPostUsecase) *PostService {
	return &PostService{usecase: u}
}

//func main() {
//	wd, err := os.Getwd()
//	if err != nil {
//		fmt.Printf("get wd err: %+v\n", err)
//		return
//	}
//
//	env := os.Environ()
//	var tags string
//
//	ctx := context.Background()
//	cfg := &packages.Config{
//		Context:    ctx,
//		Mode:       packages.LoadAllSyntax,
//		Dir:        wd,
//		Env:        env,
//		BuildFlags: []string{"-tags=wireinject"},
//		// TODO(light): Use ParseFile to skip function bodies and comments in indirect packages.
//	}
//	if len(tags) > 0 {
//		cfg.BuildFlags[0] += " " + tags
//	}
//	pkgs, err := packages.Load(cfg)
//	fmt.Println(pkgs[0].PkgPath, pkgs[0].GoFiles, pkgs[0].Syntax, err)
//}
