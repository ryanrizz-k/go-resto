package resto_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/rynr00/go-resto/internal/mocks"
	"github.com/rynr00/go-resto/internal/model"
	"github.com/rynr00/go-resto/internal/model/constant"
	"github.com/rynr00/go-resto/internal/usecase/resto"
)

var _ = Describe("GinkgoResto", func() {
	var usecase resto.Usecase
	var menuRepoMock *mocks.MockMenuRepository
	var orderRepoMock *mocks.MockOrderRepository
	var userRepoMock *mocks.MockUserRepository
	var mockCtrl *gomock.Controller

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		menuRepoMock = mocks.NewMockMenuRepository(mockCtrl)
		orderRepoMock = mocks.NewMockOrderRepository(mockCtrl)
		userRepoMock = mocks.NewMockUserRepository(mockCtrl)

		usecase = resto.GetUsecase(menuRepoMock, orderRepoMock, userRepoMock)
	})

	Describe("request order info", func() {
		Context("it gave the correct input", func() {
			inputs := model.GetOrderInfoRequest{
				OrderID: "valid_order_id",
				UserID:  "valid_user_id",
			}

			When("the requested orderID is not the user's", func() {
				BeforeEach(func() {
					orderRepoMock.EXPECT().GetOrderInfo(gomock.Any(), inputs.OrderID).
						Times(1).
						Return(model.Order{
							ID:            "valid_order_id",
							UserID:        "valid_user_id_2",
							Status:        constant.OrderStatusFinished,
							ProductOrders: []model.ProductOrder{},
							ReferenceID:   "ref_id",
						}, nil)
				})
				It("returns unauthorized error", func() {
					res, err := usecase.GetOrderInfo(context.Background(), inputs)
					Expect(err).Should(HaveOccurred())
					Expect(err.Error()).To(BeEquivalentTo("unauthorized"))
					Expect(res).To(BeEquivalentTo(model.Order{}))
				})
			})

			When("the requested orderID is the user's", func() {
				BeforeEach(func() {
					orderRepoMock.EXPECT().GetOrderInfo(gomock.Any(), inputs.OrderID).
						Times(1).
						Return(model.Order{
							ID:            "valid_order_id",
							UserID:        "valid_user_id",
							Status:        constant.OrderStatusFinished,
							ProductOrders: []model.ProductOrder{},
							ReferenceID:   "ref_id",
						}, nil)
				})
				It("returns unauthorized error", func() {
					res, err := usecase.GetOrderInfo(context.Background(), inputs)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(res).To(BeEquivalentTo(model.Order{
						ID:            "valid_order_id",
						UserID:        "valid_user_id",
						Status:        constant.OrderStatusFinished,
						ProductOrders: []model.ProductOrder{},
						ReferenceID:   "ref_id",
					}))
				})
			})
		})
	})
})
