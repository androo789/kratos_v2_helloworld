package biz

import (
	"context"

	v1 "helloworld/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
// 这个相当于之前的model
// 这是一个DO，就是直接数据表的映射，没有更多的逻辑
type Greeter struct {
	Hello string
}

// GreeterRepo is a Greater repo.
// 相当于对一个model的操作定义在这里。实现在data里面
// repo 是啥意思？按照DDD，repo是对DO的操作，
type GreeterRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	Update(context.Context, *Greeter) (*Greeter, error)
	FindByID(context.Context, int64) (*Greeter, error)
	ListByHello(context.Context, string) ([]*Greeter, error)
	ListAll(context.Context) ([]*Greeter, error)
}

// GreeterUsecase is a Greeter usecase.
//不直接使用repo ，而是封装一层更有业务意义的方法
// use case就是一个类，类下面的方法就是具体的业务方法
type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
//不直接使用save，而是封装一层更有业务意义的方法
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	//为什么要用log对象，直接打log，用包名字可以吗
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	//调用save方法的具体实现
	return uc.repo.Save(ctx, g)
}
