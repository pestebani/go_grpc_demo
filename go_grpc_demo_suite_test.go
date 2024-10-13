package go_grpc_demo_test

import (
	"context"
	"database/sql"
	"fmt"
	pb "go_grpc_demo/pkg/agenda_server/v1"
	"go_grpc_demo/pkg/client"
	"go_grpc_demo/pkg/model"
	"go_grpc_demo/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoGrpcDemo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoGrpcDemo Suite")
}

func checkAgendas(ag1, ag2 *model.Agenda) {
	Expect(ag1.ID).Should(Equal(ag2.ID))
	Expect(ag1.Name).Should(Equal(ag2.Name))
	Expect(ag1.Email).Should(Equal(ag2.Email))
	Expect(ag1.Phone).Should(Equal(ag2.Phone))
}

var _ = Describe("Test functionality with a postgres database", Ordered, func() {
	var (
		db         *sql.DB
		grpcServer *grpc.Server
		cli        client.Client
	)

	BeforeAll(func() {
		connStr := os.Getenv("DATABASE_URL")
		if connStr == "" {
			connStr = "postgres://postgres:postgres@localhost:5432/test_agenda?sslmode=disable"
		}
		// Connect to the database
		var err error
		db, err = sql.Open("postgres", connStr)

		Expect(err).Should(BeNil())

		grpcServer = grpc.NewServer()

		srv, err := service.NewService()

		Expect(err).Should(BeNil())

		pb.RegisterAgendaServiceServer(grpcServer, srv)

		port := 8083
		lis, err := net.Listen("tcp", fmt.Sprintf("[::]:%d", port))

		cli, err = client.NewClient("localhost:8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
		Expect(err).Should(BeNil())

		go func() {
			if err := grpcServer.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
		}()
	})

	AfterEach(func() {
		// Empty the database
		_, err := db.Exec("DELETE FROM agenda")
		Expect(err).Should(BeNil())
	})

	AfterAll(func() {
		db.Exec("DROP TABLE agenda")
		db.Close()
		grpcServer.GracefulStop()
	})

	It("Test Ping", func() {
		err := cli.Ping(context.Background())
		Expect(err).Should(BeNil())
	})

	It("Test retrieving non-existent id", func() {
		_, err := cli.GetAgenda(context.Background(), 123456789)
		code, isErr := status.FromError(err)
		Expect(isErr).Should(BeTrue())
		Expect(code.Code()).Should(Equal(codes.NotFound))
	})

	It("Test deleting non-existent id", func() {
		err := cli.DeleteAgenda(context.Background(), 123456789)
		code, isErr := status.FromError(err)
		Expect(isErr).Should(BeTrue())
		Expect(code.Code()).Should(Equal(codes.NotFound))
	})

	It("Test updating non-existent id", func() {
		ag := model.Agenda{
			Name:  "testUpdateNonExistent",
			Email: "non.existent@somemail.com",
			Phone: "+1 23456789",
		}

		_, err := cli.UpdateAgenda(context.Background(), 123456789, ag)
		code, isErr := status.FromError(err)
		Expect(isErr).Should(BeTrue())
		Expect(code.Code()).Should(Equal(codes.NotFound))
	})

	It("Test insert and retrieve", func() {
		ag := model.Agenda{
			Name:  "test1",
			Email: "test1@somemail.com",
			Phone: "+1 23456789",
		}

		agCre, err := cli.CreateAgenda(context.Background(), ag)
		Expect(err).Should(BeNil())

		agNew, err := cli.GetAgenda(context.Background(), agCre.ID)
		Expect(err).Should(BeNil())

		ag.ID = agCre.ID

		checkAgendas(&ag, &agNew)
	})

	It("Test insert and retrieve list", func() {
		for i := 0; i < 10; i++ {
			ag := model.Agenda{
				Name:  fmt.Sprintf("test%d", i),
				Email: fmt.Sprintf("test%d@somemail.com", i),
				Phone: fmt.Sprintf("+1 2345678%d", i),
			}
			_, err := cli.CreateAgenda(context.Background(), ag)
			Expect(err).Should(BeNil())
		}

		ags, nextPage, total, err := cli.GetAgendas(context.Background(), 2, 3)
		Expect(err).Should(BeNil())
		Expect(len(ags)).Should(Equal(3))
		Expect(nextPage).Should(Equal(3))
		Expect(total).Should(Equal(10))
	})

	It("Test insert and update an element", func() {
		ag := model.Agenda{
			Name:  "testInsertUpdate",
			Email: "test.insert.update@somemail.com",
			Phone: "+1 23456789",
		}

		agCre, err := cli.CreateAgenda(context.Background(), ag)

		Expect(err).Should(BeNil())

		agCre.Name = "testInsertUpdateUpdated"
		agCre.Email = "new.test.insert.update@somemail.com"
		agCre.Phone = "+1 098765432"

		agUpd, err := cli.UpdateAgenda(context.Background(), agCre.ID, agCre)

		Expect(err).Should(BeNil())
		checkAgendas(&agCre, &agUpd)
	})

	It("Test insert and delete an element", func() {
		ag := model.Agenda{
			Name:  "testInsertDelete",
			Email: "test.insert.delete@somemail.com",
			Phone: "+1 23456789",
		}

		agCre, err := cli.CreateAgenda(context.Background(), ag)
		Expect(err).Should(BeNil())
		_, err = cli.GetAgenda(context.Background(), agCre.ID)
		Expect(err).Should(BeNil())

		err = cli.DeleteAgenda(context.Background(), agCre.ID)
		Expect(err).Should(BeNil())

		_, err = cli.GetAgenda(context.Background(), agCre.ID)

		code, isErr := status.FromError(err)
		Expect(isErr).Should(BeTrue())
		Expect(code.Code()).Should(Equal(codes.NotFound))

		err = cli.DeleteAgenda(context.Background(), agCre.ID)

		code, isErr = status.FromError(err)
		Expect(isErr).Should(BeTrue())
		Expect(code.Code()).Should(Equal(codes.NotFound))
	})

	It("Test insert element with repeated name", func() {
		ag1 := model.Agenda{
			Name:  "testInsertRepeated",
			Email: "test.insert.repeated.1@somemail.com",
			Phone: "+1 234567891",
		}

		ag2 := model.Agenda{
			Name:  "testInsertRepeated",
			Email: "test.insert.repeated.2@somemail.com",
			Phone: "+1 234567892",
		}

		_, err := cli.CreateAgenda(context.Background(), ag1)
		Expect(err).Should(BeNil())

		_, err = cli.CreateAgenda(context.Background(), ag2)

		code, isErr := status.FromError(err)
		Expect(isErr).Should(BeTrue())
		Expect(code.Code()).Should(Equal(codes.AlreadyExists))
	})

	It("Test insert 2 elements and update second element name to the first one", func() {
		ag1 := model.Agenda{
			Name:  "testInsertRepeated",
			Email: "test.insert.repeated.1@somemail.com",
			Phone: "+1 234567891",
		}

		ag2 := model.Agenda{
			Name:  "testInsertRepeatedBis",
			Email: "test.insert.repeated.2@somemail.com",
			Phone: "+1 234567892",
		}

		_, err := cli.CreateAgenda(context.Background(), ag1)
		Expect(err).Should(BeNil())

		agCreated2, err := cli.CreateAgenda(context.Background(), ag2)
		Expect(err).Should(BeNil())

		agCreated2.Name = ag1.Name

		_, err = cli.UpdateAgenda(context.Background(), agCreated2.ID, agCreated2)

		code, isErr := status.FromError(err)
		Expect(isErr).Should(BeTrue())
		Expect(code.Code()).Should(Equal(codes.AlreadyExists))
	})
})
