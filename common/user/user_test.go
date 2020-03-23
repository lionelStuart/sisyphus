package user

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"testing"
)

func TestUserReturn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockUserRepository(ctrl)

	repo.EXPECT().Find(1).Return(&User{Name: "jim"}, nil)
	repo.EXPECT().Find(2).Return(&User{Name: "marry"}, nil)

	fmt.Println(repo.Find(1))
	fmt.Println(repo.Find(2))
	//fmt.Println(repo.Find(3))
}

func TestReturnDynamic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockUserRepository(ctrl)
	repo.EXPECT().Find(gomock.Any()).DoAndReturn(func(id int) (*User, error) {
		if id == 0 {
			return nil, errors.New("user-svc not exist")
		}
		if id < 100 {
			return &User{
				Name: "LessUser",
			}, nil
		} else {
			return &User{"LargeUser"}, nil
		}
	})
	//t.Log(repo.Find(10))
	t.Log(repo.Find(110))
	//t.Log(repo.Find(0))
}

func TestTimes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockUserRepository(ctrl)
	// 默认期望调用一次
	repo.EXPECT().Find(1).Return(&User{Name: "张三"}, nil)
	// 期望调用2次
	repo.EXPECT().Find(2).Return(&User{Name: "李四"}, nil).Times(2)
	// 调用多少次可以,包括0次
	repo.EXPECT().Find(3).Return(nil, errors.New("user-svc not found")).AnyTimes()

	// 验证一下结果
	fmt.Println(repo.Find(1)) // 这是张三
	fmt.Println(repo.Find(2)) // 这是李四
	fmt.Println(repo.Find(2)) // FindOne(2) 需调用两次,注释本行代码将导致测试不通过
	fmt.Println(repo.Find(3)) // user-svc not found, 不限调用次数，注释掉本行也能通过测试
}

func TestOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockUserRepository(ctrl)
	o1 := repo.EXPECT().Find(1).Return(&User{Name: "张三"}, nil)
	o2 := repo.EXPECT().Find(2).Return(&User{Name: "李四"}, nil)
	o3 := repo.EXPECT().Find(3).Return(nil, errors.New("user-svc not found"))
	gomock.InOrder(o1, o2, o3) //设置调用顺序
	// 按顺序调用，验证一下结果
	fmt.Println(repo.Find(1)) // 这是张三
	fmt.Println(repo.Find(2)) // 这是李四
	fmt.Println(repo.Find(3)) // user-svc not found

	// 如果我们调整了调用顺序，将导致测试不通过：
	// log.Println(repo.FindOne(2)) // 这是李四
	// log.Println(repo.FindOne(1)) // 这是张三
	// log.Println(repo.FindOne(3)) // user-svc not found
}
